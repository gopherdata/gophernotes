/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * bind.go
 *
 *  Created on Apr 27, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"sort"

	"github.com/cosmos72/gomacro/fast"
)

// show local variables
func (d *Debugger) Vars() {
	env := d.env
	var envs []*fast.Env
	for env != nil {
		envs = append(envs, env)
		env = env.Outer
		if env == nil || env.FileEnv == env {
			break // omit global variables
		}
	}
	d.showEnvs(envs)
}

func (d *Debugger) showEnvs(envs []*fast.Env) {
	// show outermost scope first
	for i := len(envs) - 1; i >= 0; i-- {
		d.showEnv(envs[i])
	}
}

func (d *Debugger) showEnv(env *fast.Env) {
	c := env.DebugComp
	if c == nil || (c.BindNum == 0 && c.IntBindNum == 0) {
		return
	}
	o := d.globals.Output
	o.Fprintf(o.Stdout, "// ----------\n")
	binds := make([]*fast.Bind, len(c.Binds))
	i := 0
	for _, bind := range c.Binds {
		binds[i] = bind
		i++
	}
	sort.Slice(binds, func(i, j int) bool {
		return binds[i].Name < binds[j].Name
	})
	g := c.CompGlobals
	for _, bind := range binds {
		value := bind.RuntimeValue(g, env)
		o.Fprintf(o.Stdout, "%s\t= %v\t// %v\n", bind.Name, value, bind.Type)
	}
}

// =============================================================================

func (d *Debugger) showBinds(env *fast.Env, binds []*fast.Bind) {
	g := d.globals
	for i, bind := range binds {
		if i != 0 {
			g.Fprintf(g.Stdout, ", ")
		}
		d.showBind(env, bind)
	}
}

func (d *Debugger) showBind(env *fast.Env, bind *fast.Bind) {
	value := bind.RuntimeValue(d.interp.Comp.CompGlobals, env)
	var ivalue interface{} = value
	if !value.IsValid() {
		ivalue = "nil"
	}

	g := d.globals
	if name := bind.Name; len(name) != 0 {
		g.Fprintf(g.Stdout, "%s=%v <%v>", name, ivalue, bind.Type)
	} else {
		g.Fprintf(g.Stdout, "%v <%v>", ivalue, bind.Type)
	}
}
