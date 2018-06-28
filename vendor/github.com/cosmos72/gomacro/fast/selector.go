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
 * selector.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// SelectorExpr compiles foo.bar, i.e. read access to methods, struct fields and imported packages
func (c *Comp) SelectorExpr(node *ast.SelectorExpr) *Expr {
	e, t := c.Expr1OrType(node.X)
	if t != nil {
		return c.selectorType(node, t)
	}
	t = e.Type
	eorig := e
	name := node.Sel.Name
	if t.Kind() == r.Ptr && t.ReflectType() == rtypeOfPtrImport && e.Const() {
		// access symbol from imported package, for example fmt.Printf
		imp := e.Value.(*Import)
		return imp.selector(name, &c.Stringer)
	}
	if t.Kind() == r.Ptr && t.Elem().Kind() == r.Struct {
		t = t.Elem()
		fun := e.AsX1()
		e = exprFun(t, func(env *Env) r.Value {
			return fun(env).Elem()
		})
	}
	switch t.Kind() {
	case r.Struct:
		field, fieldok, mtd, mtdok := c.LookupFieldOrMethod(t, name)
		if fieldok {
			return c.compileField(e, field)
		} else if mtdok {
			return c.compileMethod(node, eorig, mtd)
		}
	default:
		// interfaces and non-struct named types can have methods, but no fields
		mtd, mtdn := c.LookupMethod(t, name)
		switch mtdn {
		case 0:
		case 1:
			return c.compileMethod(node, eorig, mtd)
		default:
			c.Errorf("type %s has %d methods %q, expression is ambiguous: %v", t, mtdn, name, node)
		}
	}
	c.Errorf("type %s has no field or method %q: %v", t, name, node)
	return nil
}

// selectorType compiles foo.bar where 'foo' is a type
func (c *Comp) selectorType(node *ast.SelectorExpr, t xr.Type) *Expr {
	mtd, count := c.LookupMethod(t, node.Sel.Name)
	if count == 0 {
		c.Errorf("type <%v> has no method %q: %v", t, node.Sel, node)
	} else if count > 1 {
		c.Errorf("type <%v> has %d wrapper methods %q all at the same depth=%d - expression is ambiguous: %v", t, count, node.Sel, len(mtd.FieldIndex), node)
	}
	return c.compileMethodAsFunc(t, mtd)
}

// lookup fields and methods at the same time... it's and error if both exist at the same depth
func (c *Comp) LookupFieldOrMethod(t xr.Type, name string) (xr.StructField, bool, xr.Method, bool) {
	field, fieldok, mtd, mtdok, err := c.TryLookupFieldOrMethod(t, name)
	if err != nil {
		c.Error(err)
	}
	return field, fieldok, mtd, mtdok
}

// lookup fields and methods at the same time... it's and error if both exist at the same depth
func (c *Comp) TryLookupFieldOrMethod(t xr.Type, name string) (xr.StructField, bool, xr.Method, bool, error) {
	field, fieldn := c.LookupField(t, name)
	mtd, mtdn := c.LookupMethod(t, name)
	if c.Options&OptDebugField != 0 {
		c.Debugf("LookupFieldOrMethod for %v.%v found %d fields:  %#v", t, name, fieldn, field)
	}
	if c.Options&OptDebugMethod != 0 {
		c.Debugf("LookupFieldOrMethod for %v.%v found %d methods: %#v", t, name, mtdn, mtd)
	}
	fielddepth := len(field.Index)
	mtddepth := len(mtd.FieldIndex) + 1
	var err error
	if fieldn != 0 && mtdn != 0 {
		if fielddepth < mtddepth {
			// prefer the field
			mtdn = 0
		} else if fielddepth > mtddepth {
			// prefer the method
			fieldn = 0
		} else {
			err = c.MakeRuntimeError("type %v has %d field(s) and %d method(s) named %q at depth %d",
				t, fieldn, mtdn, name, fielddepth)
		}
	}
	if fieldn > 1 {
		err = c.MakeRuntimeError("type %v has %d fields named %q at depth %d", t, fieldn, name, fielddepth)
	} else if mtdn > 1 {
		err = c.MakeRuntimeError("type %v has %d methods named %q at depth %d", t, mtdn, name, mtddepth)
	}
	if err != nil {
		return field, fieldn == 1, mtd, mtdn == 1, err
	}
	return field, fieldn == 1, mtd, mtdn == 1, nil
}

// list direct and embedded field names that start with prefix,
// and  explicit and wrapper methods that start with prefix
func (c *Comp) listFieldsAndMethods(t xr.Type, prefix string) []string {
	var names []string
	size := len(prefix)

	collectMethods := func(typ xr.Type) {
		if t.Kind() == r.Ptr {
			t = t.Elem()
			if t.Kind() == r.Interface {
				// ignore pointer-to-interface
				return
			}
		}
		for i, n := 0, typ.NumMethod(); i < n; i++ {
			if name := typ.Method(i).Name; len(name) >= size && name[:size] == prefix {
				names = append(names, name)
			}
		}
	}
	if t.Kind() == r.Ptr {
		t = t.Elem()
		if t.Kind() == r.Interface {
			// ignore pointer-to-interface
			return nil
		}
	}
	collectMethods(t)
	if t.Kind() == r.Struct {
		size := len(prefix)
		c.Universe.VisitFields(t, func(field xr.StructField) {
			if name := field.Name; len(name) >= size && name[:size] == prefix {
				names = append(names, name)
			}
			collectMethods(field.Type)
		})
	}
	return names
}

// LookupField performs a breadth-first search for struct field with given name
func (c *Comp) LookupField(t xr.Type, name string) (field xr.StructField, numfound int) {
	return t.FieldByName(name, c.FileComp().Path)
}

// LookupMethod performs a breadth-first search for method with given name
func (c *Comp) LookupMethod(t xr.Type, name string) (mtd xr.Method, numfound int) {
	return t.MethodByName(name, c.FileComp().Path)
}

// field0 is a variant of reflect.Value.Field, also accepts pointer values
// and dereferences pointer ONLY if index < 0 (actually used index will be ^index)
func field0(v r.Value, index int) r.Value {
	switch v.Kind() {
	// also accept interface xr.Forward and extract concrete type from it
	case r.Ptr, r.Interface:
		v = v.Elem()
	}
	return v.Field(index)
}

// fieldByIndex is a variant of reflect.Value.FieldByIndex, also accepts pointer values
// and dereferences pointers ONLY if index[i] < 0 (actually used index will be ^index[i])
func fieldByIndex(v r.Value, index []int) r.Value {
	for _, x := range index {
		switch v.Kind() {
		// also accept interface xr.Forward and extract concrete type from it
		case r.Ptr, r.Interface:
			v = v.Elem()
		}
		v = v.Field(x)
	}
	return v
}

// descend embedded fields, detect any pointer-to-struct that must be dereferenced
func descendEmbeddedFields(t xr.Type, field xr.StructField) []int {
	// currently a no-op
	return field.Index
}

func (c *Comp) compileField(e *Expr, field xr.StructField) *Expr {
	objfun := e.AsX1()
	index := descendEmbeddedFields(e.Type, field)
	t := field.Type
	var fun I

	// c.Debugf("compileField: field=%#v", field)
	if len(index) == 1 {
		index0 := index[0]
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				return field0(obj, index0).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				return int(field0(obj, index0).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				return int8(field0(obj, index0).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				return int16(field0(obj, index0).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				return int32(field0(obj, index0).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				return field0(obj, index0).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				return uint(field0(obj, index0).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				return uint8(field0(obj, index0).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				return uint16(field0(obj, index0).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				return uint32(field0(obj, index0).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				return field0(obj, index0).Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				obj := objfun(env)
				return uintptr(field0(obj, index0).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				return float32(field0(obj, index0).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				return field0(obj, index0).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				return complex64(field0(obj, index0).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				return field0(obj, index0).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				return field0(obj, index0).String()
			}
		default:
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				return field0(obj, index0)
			}
		}
	} else {
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				return fieldByIndex(obj, index).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				return int(fieldByIndex(obj, index).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				return int8(fieldByIndex(obj, index).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				return int16(fieldByIndex(obj, index).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				return int32(fieldByIndex(obj, index).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				return fieldByIndex(obj, index).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				return uint(fieldByIndex(obj, index).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				return uint8(fieldByIndex(obj, index).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				return uint16(fieldByIndex(obj, index).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				return uint32(fieldByIndex(obj, index).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				return fieldByIndex(obj, index).Uint()
			}
		case r.Uintptr:

			fun = func(env *Env) uintptr {
				obj := objfun(env)
				return uintptr(fieldByIndex(obj, index).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				return float32(fieldByIndex(obj, index).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				return fieldByIndex(obj, index).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				return complex64(fieldByIndex(obj, index).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				return fieldByIndex(obj, index).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				return fieldByIndex(obj, index).String()
			}
		default:
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				return fieldByIndex(obj, index)
			}
		}
	}
	return exprFun(t, fun)
}

func (c *Comp) changeFirstParam(tfirstparam, t xr.Type) xr.Type {
	nin := t.NumIn()
	if nin == 0 {
		c.Errorf("compileMethod: inconsistent method type: expecting at least the receiver, found zero input parameters: <%v>", t)
	}
	params := make([]xr.Type, nin)
	params[0] = tfirstparam
	for i := 1; i < nin; i++ {
		params[i] = t.In(i)
	}
	nout := t.NumOut()
	results := make([]xr.Type, nout)
	for i := 0; i < nout; i++ {
		results[i] = t.Out(i)
	}
	return c.Universe.FuncOf(params, results, t.IsVariadic())
}

func (c *Comp) removeFirstParam(t xr.Type) xr.Type {
	nin := t.NumIn()
	if nin == 0 {
		c.Errorf("compileMethod: inconsistent method type: expecting at least the receiver, found zero input parameters: <%v>", t)
	}
	params := make([]xr.Type, nin-1)
	for i := 1; i < nin; i++ {
		params[i-1] = t.In(i)
	}
	nout := t.NumOut()
	results := make([]xr.Type, nout)
	for i := 0; i < nout; i++ {
		results[i] = t.Out(i)
	}
	return c.Universe.FuncOf(params, results, t.IsVariadic())
}

// compileMethod compiles expr.method
// relatively slow, but simple: return a closure with the receiver already bound
func (c *Comp) compileMethod(node *ast.SelectorExpr, e *Expr, mtd xr.Method) *Expr {
	obj2method := c.compileObjGetMethod(e.Type, mtd)
	fun := e.AsX1()
	tclosure := c.removeFirstParam(mtd.Type)

	return exprX1(tclosure, func(env *Env) r.Value {
		return obj2method(fun(env))
	})
}

// create and return a function that, given a reflect.Value, returns its method specified by mtd
func (c *Comp) compileObjGetMethod(t xr.Type, mtd xr.Method) (ret func(r.Value) r.Value) {
	if c.Options&OptDebugMethod != 0 {
		c.Debugf("compileObjGetMethod for %v.%v: method is %#v", t, mtd.Name, mtd)
	}
	index := mtd.Index
	tfunc := mtd.Type
	rtclosure := c.removeFirstParam(tfunc).ReflectType()

	tfield, fieldindex, addressof, deref := c.computeMethodFieldIndex(t, mtd)
	rtfield := tfield.ReflectType()

	rmtd, ok := rtfield.MethodByName(mtd.Name)

	if ok && xr.QName1(tfield) == xr.QName1(rtfield) && c.compatibleMethodType(tfield, mtd, rmtd) {
		// closures for methods declared by compiled code are available
		// simply with reflect.Value.Method(index). Easy.
		index := rmtd.Index

		switch len(fieldindex) {
		case 0:
			if addressof {
				ret = func(obj r.Value) r.Value {
					return obj.Addr().Method(index)
				}
			} else if deref {
				ret = func(obj r.Value) r.Value {
					return obj.Elem().Method(index)
				}
			} else {
				ret = func(obj r.Value) r.Value {
					return obj.Method(index)
				}
			}
		case 1:
			fieldindex := fieldindex[0]
			ret = func(obj r.Value) r.Value {
				return field0(obj, fieldindex)
				return obj.Method(index)
			}
		default:
			ret = func(obj r.Value) r.Value {
				obj = fieldByIndex(obj, fieldindex)
				return obj.Method(index)
			}
		}
	} else {
		tname := t.Name()
		methodname := mtd.Name

		// method declared by interpreted code, manually build the closure.
		//
		// It's not possible to call r.MakeFunc() only once at compile-time,
		// because the closure passed to it needs access to a variable holding the receiver.
		// such variable would be evaluated only once at compile-time,
		// not once per method extraction!
		funs := mtd.Funs
		variadic := tfunc.IsVariadic()

		if funs == nil {
			c.Errorf("method declared but not yet implemented: %s.%s", tname, methodname)
		} else if len(*funs) <= index || (*funs)[index].Kind() != r.Func {
			// c.Warnf("method declared but not yet implemented: %s.%s", tname, methodname)
		} else if c.Options&OptDebugMethod != 0 {
			c.Debugf("compiling method %v.%s <%v>: method declared by interpreted code, manually building the closure reflect.Type <%v>",
				tname, methodname, mtd.Type, rtclosure)
		}
		// Go compiled code crashes when extracting a method from nil interface,
		// NOT later when calling the method.
		//
		// On the other hand, Go compiled code can extract methods from a nil pointer to named type,
		// and it will crash later calling the method ONLY if the method implementation dereferences the receiver.
		//
		// Reproduce the same behaviour
		if t.Kind() == r.Interface {
			ret = compileInterfaceGetMethod(fieldindex, deref, index)
		} else {
			switch len(fieldindex) {
			case 0:
				ret = func(obj r.Value) r.Value {
					if addressof {
						obj = obj.Addr()
					} else if deref {
						obj = obj.Elem()
					}
					fun := (*funs)[index] // retrieve the function as soon as possible (early bind)
					if fun == Nil {
						Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
					}
					return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
						args = append([]r.Value{obj}, args...)
						// Debugf("invoking <%v> with args %v", fun.Type(), fullargs
						if variadic {
							return fun.CallSlice(args)
						} else {
							return fun.Call(args)
						}
					})
				}
			case 1:
				fieldindex := fieldindex[0]
				ret = func(obj r.Value) r.Value {
					obj = field0(obj, fieldindex)
					// Debugf("invoking method <%v> on receiver <%v> (addressof=%t, deref=%t)", (*funs)[index].Type(), obj.Type(), addressof, deref)
					if addressof {
						obj = obj.Addr()
					} else if deref {
						obj = obj.Elem()
					}
					fun := (*funs)[index] // retrieve the function as soon as possible (early bind)
					if fun == Nil {
						Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
					}
					return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
						args = append([]r.Value{obj}, args...)
						// Debugf("invoking <%v> with args %v", fun.Type(), fullargs)
						if variadic {
							return fun.CallSlice(args)
						} else {
							return fun.Call(args)
						}
					})
				}
			default:
				ret = func(obj r.Value) r.Value {
					obj = fieldByIndex(obj, fieldindex)
					if addressof {
						obj = obj.Addr()
					} else if deref {
						obj = obj.Elem()
					}
					fun := (*funs)[index] // retrieve the function as soon as possible (early bind)
					if fun == Nil {
						Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
					}
					return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
						args = append([]r.Value{obj}, args...)
						// Debugf("invoking <%v> with args %v", fun.Type(), fullargs)
						if variadic {
							return fun.CallSlice(args)
						} else {
							return fun.Call(args)
						}
					})
				}
			}
		}
	}
	return ret
}

// return true if t is not an interface and mtd.Type().ReflectType() == rmtd.Type,
// or if t is an interface and rmtd.Type is the same as mtd.Type().ReflectType() _minus_ the receiver
func (c *Comp) compatibleMethodType(t xr.Type, mtd xr.Method, rmtd r.Method) bool {
	rt1 := mtd.Type.ReflectType()
	rt2 := rmtd.Type
	if t.Kind() != r.Interface {
		return rt1 == rt2
	}
	return rt1.NumIn()-1 == rt2.NumIn() && c.removeFirstParam(mtd.Type).ReflectType() == rt2
}

func compileInterfaceGetMethod(fieldindex []int, deref bool, index int) func(r.Value) r.Value {
	switch len(fieldindex) {
	case 0:
		return func(obj r.Value) r.Value {
			if deref {
				obj = obj.Elem()
			}
			return xr.EmulatedInterfaceGetMethod(obj, index)
		}
	case 1:
		fieldindex := fieldindex[0]
		return func(obj r.Value) r.Value {
			obj = field0(obj, fieldindex)
			if deref {
				obj = obj.Elem()
			}
			return xr.EmulatedInterfaceGetMethod(obj, index)
		}
	default:
		return func(obj r.Value) r.Value {
			obj = fieldByIndex(obj, fieldindex)
			if deref {
				obj = obj.Elem()
			}
			return xr.EmulatedInterfaceGetMethod(obj, index)
		}
	}
}

// compute and return the dereferences and addressof to perform while descending
// the embedded fields described by mtd.FieldIndex []int
// also check that addressof will be performed on addressable fields
func (c *Comp) computeMethodFieldIndex(t xr.Type, mtd xr.Method) (fieldtype xr.Type, fieldindex []int, addressof bool, deref bool) {
	fieldindex = mtd.FieldIndex
	var copied, indirect bool

	// descend embedded fields
	for i, x := range mtd.FieldIndex {
		if t.Kind() == r.Ptr {
			// embedded field (or initial value) is a pointer, dereference it.
			t = t.Elem()
			indirect = true
			if !copied {
				copied = true
				fieldindex = make([]int, len(mtd.FieldIndex))
				copy(fieldindex, mtd.FieldIndex)
			}
			fieldindex[i] = ^x // remember we need a pointer dereference at runtime
		}
		t = t.Field(x).Type
	}
	tfunc := mtd.Type
	trecv := tfunc.In(0)

	objPointer := t.Kind() == r.Ptr      // field is pointer?
	recvPointer := trecv.Kind() == r.Ptr // method with pointer receiver?
	addressof = !objPointer && recvPointer
	deref = objPointer && !recvPointer

	debug := c.Options&OptDebugMethod != 0
	if debug {
		c.Debugf("compiling method %v.%v", t.Name(), mtd.Name)
	}
	if t.AssignableTo(trecv) {
		addressof = false
		deref = false
		if debug {
			c.Debugf("compiling method %v.%v: value is assignable to receiver", t.Name(), mtd.Name)
		}
	} else if addressof && c.Universe.PtrTo(t).AssignableTo(trecv) {
		// c.Debugf("method call <%v> will take address of receiver <%v>", tfunc, t)
		// ensure receiver is addressable. maybe it was simply dereferenced by Comp.SelectorExpr
		// or maybe we need to explicitly take its address
		if indirect {
			if len(fieldindex) != 0 {
				// easy, we dereferenced some expression while descending embedded fields
				// so the receiver is addressable
				if debug {
					c.Debugf("compiling method %v.%v: address-of-value is assignable to receiver", t.Name(), mtd.Name)
				}
			} else {
				// even easier, the initial expression already contains the address we want
				addressof = false
				if debug {
					c.Debugf("compiling method %v.%v: value was initially an address", t.Name(), mtd.Name)
				}
			}
		} else {
			// manually compile "& receiver_expression"
			if debug {
				c.Debugf("compiling method %v.%v: compiling address-of-value", t.Name(), mtd.Name)
			}
			// FIXME restore and complete these addressability checks
			/*
				if len(index) != 0 {
					// must execute addressof at runtime, just check that struct is addressable
					c.addressOf(node.X)
				} else {
					e = c.addressOf(node.X)
					addressof = false
				}
			*/
		}
		t = c.Universe.PtrTo(t)
	} else if deref && t.Elem().AssignableTo(trecv) {
		t = t.Elem()
		if debug {
			c.Debugf("method call <%v> will dereference receiver <%v>", tfunc, t)
		}
	} else {
		c.Errorf("cannot use <%v> as <%v> in receiver of method <%v>", t, trecv, tfunc)
	}
	return t, fieldindex, addressof, deref
}

// compileMethodAsFunc compiles a method as a function, for example time.Duration.String.
// The method receiver will be the first argument of returned function.
func (c *Comp) compileMethodAsFunc(t xr.Type, mtd xr.Method) *Expr {
	tsave := t
	fieldindex := mtd.FieldIndex
	var copied bool

	// descend embedded fields
	for i, x := range mtd.FieldIndex {
		if t.Kind() == r.Ptr && t.Elem().Kind() == r.Struct {
			// embedded field (or initial value) is a pointer, dereference it.
			if !copied {
				copied = true
				fieldindex = make([]int, len(mtd.FieldIndex))
				copy(fieldindex, mtd.FieldIndex)
			}
			fieldindex[i] = ^x // remember we neeed a pointer dereference at runtime
			t = t.Elem()
		}
		t = t.Field(x).Type
	}

	index := mtd.Index
	tfunc := mtd.Type
	trecv := tfunc.In(0)

	objPointer := t.Kind() == r.Ptr      // field is pointer?
	recvPointer := trecv.Kind() == r.Ptr // method with pointer receiver?
	addressof := !objPointer && recvPointer
	deref := objPointer && !recvPointer

	// convert a method (i.e. with first param used as receiver) to regular function
	// and, if needed, create wrapper method for embedded field
	if recvPointer {
		// receiver is pointer-to-tsave
		if tsave.Kind() != r.Ptr {
			tsave = c.Universe.PtrTo(tsave)
			if len(fieldindex) != 0 && fieldindex[0] >= 0 {
				// remember we neeed a pointer dereference at runtime
				fieldindex[0] = ^fieldindex[0]
			}
		}
	} else {
		// receiver is tsave
		if tsave.Kind() == r.Ptr {
			tsave = tsave.Elem()
			if len(fieldindex) != 0 && fieldindex[0] < 0 {
				// no pointer dereference at runtime
				fieldindex[0] = ^fieldindex[0]
			}
		}
	}
	tfunc = c.changeFirstParam(tsave, tfunc)

	if len(fieldindex) == 0 {
		// tsave is a named type, while trecv may be an unnamed interface:
		// use tsave for correctness
		t = tsave
	} else {
		t = trecv
	}
	if t.Kind() == r.Ptr {
		t = t.Elem()
	}
	rtype := t.ReflectType()

	var ret r.Value

	// c.Debugf("compileMethodAsFunc: t = <%v> has %d methods, rtype = <%v> has %d methods", t, t.NumMethod(), rtype, rtype.NumMethod())

	if t.NumMethod() == rtype.NumMethod() && t.Named() && xr.QName1(t) == xr.QName1(rtype) {
		// methods declared by compiled code are available
		// simply with reflect.Type.Method(index). Easy.
		rmethod, ok := rtype.MethodByName(mtd.Name)
		if !ok {
			c.Errorf("inconsistent type <%v>: reflect.Type <%v> has no method %q", t, rtype, mtd.Name)
		}
		rfunc := rmethod.Func

		if rfunc.Kind() != r.Func {
			if rtype.Kind() != r.Interface {
				c.Errorf("inconsistent type <%v>: reflect.Type <%v> has method %q with callable function = nil", t, rtype, mtd.Name)
			}
			// invoking interface method... retrieve the function at runtime
			rindex := rmethod.Index // usually == index. may differ if we removed wrapper methods from t
			switch len(fieldindex) {
			case 0:
				ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
					return args[0].Method(rindex).Call(args[1:])
				})
			case 1:
				fieldindex := fieldindex[0]
				ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
					args[0] = field0(args[0], fieldindex)
					return args[0].Method(rindex).Call(args[1:])
				})
			default:
				ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
					args[0] = fieldByIndex(args[0], fieldindex)
					return args[0].Method(rindex).Call(args[1:])
				})
			}
		} else {
			// invoking method of named type
			switch len(fieldindex) {
			case 0:
				ret = rfunc
			case 1:
				fieldindex := fieldindex[0]
				ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
					args[0] = field0(args[0], fieldindex)
					return rfunc.Call(args)
				})
			default:
				ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
					args[0] = fieldByIndex(args[0], fieldindex)
					return rfunc.Call(args)
				})
			}
		}
	} else {
		// method declared by interpreted code, manually retrieve it.
		funs := mtd.Funs

		tname := t.Name()
		methodname := mtd.Name
		if funs == nil {
			c.Errorf("method declared but not yet implemented: %s.%s", tname, methodname)
		} else if len(*funs) <= index || (*funs)[index].Kind() != r.Func {
			// c.Warnf("method declared but not yet implemented: %s.%s", tname, methodname)
		}

		switch len(fieldindex) {
		case 0:
			ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
				if addressof {
					args[0] = args[0].Addr()
				} else if deref {
					args[0] = args[0].Elem()
				}
				fun := (*funs)[index] // retrieve the function as soon as possible (early bind)
				if fun == Nil {
					Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
				}
				return fun.Call(args)
			})
		case 1:
			fieldindex := fieldindex[0]
			ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
				args[0] = field0(args[0], fieldindex)
				// Debugf("invoking method <%v> on receiver <%v> (addressof=%t, deref=%t)", (*funs)[index].Type(), obj.Type(), addressof, deref)
				if addressof {
					args[0] = args[0].Addr()
				} else if deref {
					args[0] = args[0].Elem()
				}
				fun := (*funs)[index]
				if fun == Nil {
					Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
				}
				return fun.Call(args)
			})
		default:
			ret = r.MakeFunc(tfunc.ReflectType(), func(args []r.Value) []r.Value {
				args[0] = fieldByIndex(args[0], fieldindex)
				if addressof {
					args[0] = args[0].Addr()
				} else if deref {
					args[0] = args[0].Elem()
				}
				fun := (*funs)[index] // retrieve the function as soon as possible (early bind)
				if fun == Nil {
					Errorf("method is declared but not yet implemented: %s.%s", tname, methodname)
				}
				return fun.Call(args)
			})
		}
	}
	return c.exprValue(tfunc, ret.Interface())
}

// SelectorPlace compiles a.b returning a settable and/or addressable Place
func (c *Comp) SelectorPlace(node *ast.SelectorExpr, opt PlaceOption) *Place {
	obje := c.Expr1(node.X, nil)
	te := obje.Type
	name := node.Sel.Name
	if te.ReflectType() == rtypeOfPtrImport && obje.Const() {
		// access settable and/or addressable variable from imported package, for example os.Stdout
		imp := obje.Value.(*Import)
		return imp.selectorPlace(c, name, opt)
	}
	ispointer := false
	switch te.Kind() {
	case r.Ptr:
		ispointer = true
		te = te.Elem()
		if te.Kind() != r.Struct {
			break
		}
		objfun := obje.AsX1()
		obje = exprFun(te, func(env *Env) r.Value {
			obj := objfun(env)
			// Debugf("SelectorPlace: obj = %v <%v> (expecting pointer to struct)", obj, obj.Type())
			return obj.Elem()
		})
		fallthrough
	case r.Struct:
		field, fieldn := c.LookupField(te, name)
		if fieldn == 0 {
			break
		} else if fieldn > 1 {
			c.Errorf("type %v has %d fields named %q, all at depth %d", te, fieldn, name, len(field.Index))
			return nil
		}
		// if ispointer, field is automatically settable and addressable
		// because the 'a' in 'a.b' is actually a pointer
		if !ispointer {
			c.checkAddressableField(node)
		}
		return c.compileFieldPlace(obje, field)
	}
	c.Errorf("type %v has no field %q: %v", te, name, node)
	return nil
}

// checkSettableField check that a struct field is settable and addressable.
// by Go specs, this requires the struct itself to be settable and addressable.
func (c *Comp) checkAddressableField(node *ast.SelectorExpr) {
	panicking := true
	defer func() {
		if panicking {
			rec := recover()
			c.Pos = node.Pos()
			c.Errorf("cannot assign to %v\n\t%v", node, rec)
		}
	}()
	c.placeOrAddress(node.X, PlaceAddress, nil)
	panicking = false
}

func (c *Comp) compileFieldPlace(obje *Expr, field xr.StructField) *Place {
	// c.Debugf("compileFieldPlace: field=%#v", field)
	objfun := obje.AsX1()
	index := descendEmbeddedFields(obje.Type, field)
	t := field.Type
	var fun, addr func(*Env) r.Value

	if len(index) == 1 {
		index0 := index[0]
		fun = func(env *Env) r.Value {
			obj := objfun(env)
			return field0(obj, index0)
		}
		addr = func(env *Env) r.Value {
			obj := objfun(env)
			return field0(obj, index0).Addr()
		}
	} else {
		fun = func(env *Env) r.Value {
			obj := objfun(env)
			return fieldByIndex(obj, index)
		}
		addr = func(env *Env) r.Value {
			obj := objfun(env)
			return fieldByIndex(obj, index).Addr()
		}
	}
	return &Place{Var: Var{Type: t, Name: field.Name}, Fun: fun, Addr: addr}
}
