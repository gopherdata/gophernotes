/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * global.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/atomic"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type I = interface{}

// ================================= Untyped =================================

type UntypedLit = untyped.Lit

var untypedOne = UntypedLit{Kind: r.Int, Val: constant.MakeInt64(1)}

func MakeUntypedLit(kind r.Kind, val constant.Value, basicTypes *[]xr.Type) UntypedLit {
	return untyped.MakeLit(kind, val, basicTypes)
}

// ================================= Lit =================================

// Lit represents a literal value, i.e. a typed or untyped constant
type Lit struct {

	// Type is nil only for literal nils.
	// for all other literals, it is reflect.TypeOf(Lit.Value)
	//
	// when Lit is embedded in other structs that represent non-constant expressions,
	// Type is the first type returned by the expression (nil if returns no values)
	Type xr.Type

	// Value is one of:
	//   nil, bool, int, int8, int16, int32, int64,
	//   uint, uint8, uint16, uint32, uint64, uintptr,
	//   float32, float64, complex64, complex128, string,
	//   UntypedLit
	//
	// when Lit is embedded in other structs that represent non-constant expressions,
	// Value is usually nil
	Value I
}

// Untyped returns true if Lit is an untyped constant
func (lit *Lit) Untyped() bool {
	_, ok := lit.Value.(UntypedLit)
	return ok
}

// UntypedKind returns the reflect.Kind of untyped constants,
// i.e. their "default type"
func (lit *Lit) UntypedKind() r.Kind {
	if untyp, ok := lit.Value.(UntypedLit); ok {
		return untyp.Kind
	} else {
		return r.Invalid
	}
}

func (lit *Lit) ConstValue() r.Value {
	v := r.ValueOf(lit.Value)
	if lit.Type != nil {
		rtype := lit.Type.ReflectType()
		if !v.IsValid() {
			v = r.Zero(rtype)
		} else if !lit.Untyped() && v.Type() != rtype {
			v = convert(v, rtype)
		}
	}
	return v
}

func (lit Lit) String() string {
	switch val := lit.Value.(type) {
	case string, nil:
		return fmt.Sprintf("%#v", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ================================= EFlags =================================

// EFlags represents the flags of an expression
type EFlags uint32

const (
	EIsNil EFlags = 1 << iota
	EIsTypeAssert
)

func (f EFlags) IsNil() bool {
	return f&EIsNil != 0
}

func MakeEFlag(flag bool, iftrue EFlags) EFlags {
	if flag {
		return iftrue
	}
	return 0
}

func EFlag4Value(value I) EFlags {
	if value == nil {
		return EIsNil
	}
	return 0
}

// ================================= Expr =================================

// Expr represents an expression in the "compiler"
type Expr struct {
	Lit
	Types []xr.Type // in case the expression produces multiple values. if nil, use Lit.Type.
	Fun   I         // function that evaluates the expression at runtime.
	Sym   *Symbol   // in case the expression is a symbol
	EFlags
}

func (e *Expr) Const() bool {
	return e.Value != nil || e.IsNil()
}

// NumOut returns the number of values that an expression will produce when evaluated
func (e *Expr) NumOut() int {
	if e.Types == nil {
		return 1
	}
	return len(e.Types)
}

// Out returns the i-th type that an expression will produce when evaluated
func (e *Expr) Out(i int) xr.Type {
	if i == 0 && e.Types == nil {
		return e.Type
	}
	return e.Types[i]
}

func (e *Expr) String() string {
	if e == nil {
		return "nil"
	}
	var str string
	if e.Const() {
		str = fmt.Sprintf("Expr{Type: %v, Value: %v}", e.Type, e.Lit.String())
	} else if e.NumOut() == 1 {
		str = fmt.Sprintf("Expr{Type: %v, Fun: %#v}", e.Type, e.Fun)
	} else {
		str = fmt.Sprintf("Expr{Types: %v, Fun: %#v}", e.Types, e.Fun)
	}
	return str
}

// ================================= Stmt =================================

// Stmt represents a statement in the fast interpreter
type Stmt func(*Env) (Stmt, *Env)

// ================================= Builtin =================================

// Builtin represents a builtin function in the fast interpreter
type Builtin struct {
	// interpreted code should not access "compile": not exported.
	// compile usually needs to modify Symbol: pass it by value.
	Compile func(c *Comp, sym Symbol, node *ast.CallExpr) *Call
	ArgMin  uint16
	ArgMax  uint16
}

// ================================= Function =================================

// Function represents a function that accesses *Interp in the fast interpreter
type Function struct {
	Fun  interface{}
	Type xr.Type
}

// ================================= Macro =================================

// Macro represents a macro in the fast interpreter
type Macro struct {
	closure func(args []r.Value) (results []r.Value)
	argNum  int
}

// ================================= BindClass =================================

// BindDescriptor uses two bits to store the class.
// use all remaining bits as unsigned => we lose only one bit
// when representing non-negative ints
type BindClass uint

const (
	ConstBind BindClass = iota
	FuncBind
	VarBind
	IntBind
)

func (class BindClass) String() string {
	switch class {
	case ConstBind:
		return "const"
	case FuncBind:
		return "func"
	case VarBind:
		return "var"
	case IntBind:
		return "intvar"
	default:
		return fmt.Sprintf("unknown%d", uint(class))
	}
}

// ================================== BindDescriptor =================================

// the zero value of BindDescriptor is a valid descriptor for all constants,
// and also for functions and variables named "_"
type BindDescriptor BindClass

const (
	bindClassMask  = BindClass(0x3)
	bindIndexShift = 2

	NoIndex             = int(-1)                   // index of constants, functions and variables named "_"
	ConstBindDescriptor = BindDescriptor(ConstBind) // bind descriptor for all constants
)

func (class BindClass) MakeDescriptor(index int) BindDescriptor {
	class &= bindClassMask
	return BindDescriptor((index+1)<<bindIndexShift | int(class))
}

// IntBind returns true if BindIndex refers to a slot in Env.IntBinds (the default is a slot in Env.Binds)
func (desc BindDescriptor) Class() BindClass {
	return BindClass(desc) & bindClassMask
}

// Index returns the slice index to use in Env.Binds or Env.IntBinds to access a variable or function.
// returns NoIndex for variables and functions named "_"
func (desc BindDescriptor) Index() int {
	index := int(desc>>bindIndexShift) - 1
	// debugf("BindDescriptor=%v, class=%v, index=%v", desc, desc.Class(), index)
	return index
}

func (desc BindDescriptor) Settable() bool {
	class := desc.Class()
	return class == IntBind || class == VarBind
}

func (desc BindDescriptor) String() string {
	return fmt.Sprintf("%s index=%d", desc.Class(), desc.Index())
}

// ================================== Bind =================================

// Bind represents a constant, variable, function or builtin in the "compiler"
type Bind struct {
	Lit
	Desc BindDescriptor
	Name string
}

func (bind *Bind) String() string {
	return fmt.Sprintf("{%s name=%q value=%v type=<%v>}", bind.Desc, bind.Name, bind.Lit.Value, bind.Lit.Type)
}

func (bind *Bind) Const() bool {
	return bind.Desc.Class() == ConstBind
}

// return bind value for constant binds.
// if bind is untyped constant, returns UntypedLit wrapped in reflect.Value
func (bind *Bind) ConstValue() r.Value {
	if !bind.Const() {
		return Nil
	}
	return bind.Lit.ConstValue()
}

// return bind value.
// if bind is untyped constant, returns UntypedLit wrapped in reflect.Value
func (bind *Bind) RuntimeValue(env *Env) r.Value {
	var v r.Value
	switch bind.Desc.Class() {
	case ConstBind:
		v = bind.Lit.ConstValue()
	case IntBind:
		expr := bind.intExpr(&env.Run.Stringer)
		// no need for Interp.RunExpr(): expr is a local variable,
		// not a statement or a function call that may be stopped by the debugger
		v = expr.AsX1()(env)
	default:
		v = env.Vals[bind.Desc.Index()]
	}
	return v
}

func (bind *Bind) AsVar(upn int, opt PlaceOption) *Var {
	class := bind.Desc.Class()
	switch class {
	case VarBind, IntBind:
		return &Var{Upn: upn, Desc: bind.Desc, Type: bind.Type, Name: bind.Name}
	default:
		Errorf("%s a %s: %s <%v>", opt, class, bind.Name, bind.Type)
		return nil
	}
}

func (bind *Bind) AsSymbol(upn int) *Symbol {
	return &Symbol{Bind: *bind, Upn: upn}
}

func (c *Comp) BindUntyped(kind r.Kind, value constant.Value) *Bind {
	untypedlit := MakeUntypedLit(kind, value, &c.Universe.BasicTypes)
	return &Bind{Lit: Lit{Type: c.TypeOfUntypedLit(), Value: untypedlit}, Desc: ConstBindDescriptor}
}

// ================================== Symbol, Var, Place =================================

// Symbol represents a resolved constant, function, variable or builtin
type Symbol struct {
	Bind
	Upn int
}

func (sym *Symbol) AsVar(opt PlaceOption) *Var {
	return sym.Bind.AsVar(sym.Upn, opt)
}

// Var represents a settable variable
type Var struct {
	// when Var is embedded in other structs that represent non-identifiers,
	// Upn and Desc are usually the zero values
	Upn  int
	Desc BindDescriptor
	Type xr.Type
	Name string
}

func (va *Var) AsSymbol() *Symbol {
	return &Symbol{
		Bind: Bind{
			Lit:  Lit{Type: va.Type},
			Desc: va.Desc,
			Name: va.Name,
		},
		Upn: va.Upn,
	}
}

func (va *Var) AsPlace() *Place {
	return &Place{Var: *va}
}

// Place represents a settable place or, equivalently, its address
type Place struct {
	Var
	// Fun is nil for variables.
	// For non-variables, returns a settable and addressable reflect.Value: the place itself.
	// For map[key], Fun returns the map itself (which may NOT be settable).
	// Call Fun only once, it may have side effects!
	Fun func(*Env) r.Value
	// Addr is nil for variables.
	// For non-variables, it will return the address of the place.
	// For map[key], it is nil since map[key] is not addressable
	// Call Addr only once, it may have side effects!
	Addr func(*Env) r.Value
	// used only for map[key], returns key. call it only once, it may have side effects!
	MapKey  func(*Env) r.Value
	MapType xr.Type
}

func (place *Place) IsVar() bool {
	return place.Fun == nil
}

type PlaceOption bool // the reason why we want a place: either to write into it, or to take its address

const (
	PlaceSettable PlaceOption = false
	PlaceAddress  PlaceOption = true
)

func (opt PlaceOption) String() string {
	if opt == PlaceAddress {
		return "cannot take the address of"
	} else {
		return "cannot assign to"
	}
}

// ================================== Comp, Env =================================

type CompileOptions int

const (
	COptKeepUntyped CompileOptions = 1 << iota // if set, Compile() on expressions will keep all untyped constants as such (in expressions where Go compiler would compute an untyped constant too)
	COptDefaults    CompileOptions = 0
)

type Code struct {
	List       []Stmt
	DebugPos   []token.Pos // for debugging interpreted code: position of each statement
	WithDefers bool        // true if code contains some defers
}

type LoopInfo struct {
	Break      *int
	Continue   *int
	ThisLabels []string // sorted. for labeled "switch" and "for"
}

func (l *LoopInfo) HasLabel(label string) bool {
	i := sort.SearchStrings(l.ThisLabels, label)
	return i >= 0 && i < len(l.ThisLabels) && l.ThisLabels[i] == label
}

type FuncInfo struct {
	Name         string
	Param        []*Bind
	Result       []*Bind
	NamedResults bool
}

const (
	poolCapacity = 32
)

type ExecFlags uint32

const (
	EFStartDefer ExecFlags = 1 << iota // true next executed function body is a defer
	EFDefer                            // function body being executed is a defer
	EFDebug                            // function body is executed with debugging enabled
)

func (ef ExecFlags) StartDefer() bool {
	return ef&EFStartDefer != 0
}

func (ef ExecFlags) IsDefer() bool {
	return ef&EFDefer != 0
}

func (ef ExecFlags) IsDebug() bool {
	return ef&EFDebug != 0
}

func (ef *ExecFlags) SetDefer(flag bool) {
	if flag {
		(*ef) |= EFDefer
	} else {
		(*ef) &^= EFDefer
	}
}

func (ef *ExecFlags) SetStartDefer(flag bool) {
	if flag {
		(*ef) |= EFStartDefer
	} else {
		(*ef) &^= EFStartDefer
	}
}

func (ef *ExecFlags) SetDebug(flag bool) {
	if flag {
		(*ef) |= EFDebug
	} else {
		(*ef) &^= EFDebug
	}
}

type DebugOp struct {
	// statements at env.CallDepth < Depth will be executed in single-stepping mode,
	// i.e. invoking the debugger after every statement
	Depth int
	// nil = do not panic.
	// otherwise, address of value to panic() in order to terminate execution
	Panic *interface{}
}

var (
	// NEVER modify these!
	DebugOpContinue = DebugOp{0, nil}
	DebugOpStep     = DebugOp{MaxInt, nil}
)

type Debugger interface {
	Breakpoint(ir *Interp, env *Env) DebugOp
	At(ir *Interp, env *Env) DebugOp
}

// IrGlobals contains interpreter configuration
type IrGlobals struct {
	gls  map[uintptr]*Run
	lock atomic.SpinLock
	Globals
}

// Run contains per-goroutine interpreter runtime bookeeping information
type Run struct {
	*IrGlobals
	goid         uintptr // owner goroutine id
	Interrupt    Stmt
	Signals      Signals // set by defer, return, breakpoint, debugger and Run.interrupt(os.Signal)
	ExecFlags    ExecFlags
	CurrEnv      *Env        // caller of current function. used ONLY at function entry to build call stack
	InstallDefer func()      // defer function to be installed
	DeferOfFun   *Env        // function whose defer are running
	PanicFun     *Env        // the currently panicking function
	Panic        interface{} // current panic. needed for recover()
	CmdOpt       CmdOpt
	Debugger     Debugger
	DebugDepth   int // depth of function to debug with single-step
	PoolSize     int
	Pool         [poolCapacity]*Env
}

// CompGlobals contains interpreter compile bookeeping information
type CompGlobals struct {
	*IrGlobals
	Universe     *xr.Universe
	KnownImports map[string]*Import // map[path]*Import cache of known imports
	interf2proxy map[r.Type]r.Type  // interface -> proxy
	proxy2interf map[r.Type]xr.Type // proxy -> interface
	Prompt       string
}

func (cg *CompGlobals) CompileOptions() CompileOptions {
	var opts CompileOptions
	if cg.Options&OptKeepUntyped != 0 {
		opts = COptKeepUntyped
	}
	return opts
}

type CompBinds struct {
	Binds      map[string]*Bind
	BindNum    int // len(Binds) == BindNum + IntBindNum + # of constants
	IntBindNum int
	// if address of some Env.Ints[index] was taken, we must honor it:
	// we can no longer reallocate Env.Ints[], thus we cannot declare IntBind variables
	// beyond Env.Ints[] capacity. In such case, we set IntBindMax to cap(Env.Ints):
	// Comp.NewBind() will allocate IntBind variables only up to IntBindMax,
	// then switch and allocate them as VarBind instead (they are slower and each one allocates memory)
	IntBindMax int
	Types      map[string]xr.Type
	Name       string // set by "package" directive
	Path       string
}

// Comp is a tree-of-closures builder: it transforms ast.Nodes into closures
// for faster execution. Consider it a poor man's compiler (hence the name)
type Comp struct {
	*CompGlobals
	CompBinds
	// UpCost is the number of *Env.Outer hops to perform at runtime to reach the *Env corresponding to *Comp.Outer
	// usually equals one. will be zero if this *Comp defines no local variables/functions.
	UpCost    int
	Depth     int
	Code      Code      // "compiled" code
	Loop      *LoopInfo // != nil when compiling a for or switch
	Func      *FuncInfo // != nil when compiling a function
	Labels    map[string]*int
	Outer     *Comp
	FuncMaker *funcMaker // used by debugger command 'backtrace' to obtain function name, type and binds for arguments and results
}

// ================================= Env =================================

type EnvBinds struct {
	Vals []r.Value
	Ints []uint64
}

// Env is the interpreter's runtime environment
type Env struct {
	EnvBinds
	Outer           *Env
	IP              int
	Code            []Stmt
	Run             *Run
	FileEnv         *Env
	DebugPos        []token.Pos // for debugging interpreted code: position of each statement
	DebugComp       *Comp       // for debugging interpreted code: compiler with Binds, and to rebuild an Interp if needed
	Caller          *Env        // for debugging interpreted code: previous function in call stack. nil for nested *Env
	CallDepth       int         // for debugging interpreted code: depth of call stack
	UsedByClosure   bool        // a bitfield would introduce more races among goroutines
	IntAddressTaken bool        // true if &Env.Ints[index] was executed... then we cannot reuse or reallocate Ints
}

// ================================= Import =================================

// Import represents an imported package.
// we cannot name it "Package" because it conflicts with ast2.Package
type Import struct {
	// model as a combination of CompBinds and EnvBinds, because to support the command 'package PATH'
	// we must convert Comp+Env to Import and vice-versa.
	// This has the added benefit of allowing packages to freely mix
	// interpreted and compiled constants, functions, variables and types.
	CompBinds
	*EnvBinds
	env *Env
}
