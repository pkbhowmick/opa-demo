// Copyright 2020 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

// +build !opa_wasm

package wasm

import (
	"context"
	"errors"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/resolver"
)

// Resolver is a stub implementation of a resolver.Resolver.
type Resolver struct {
}

// Entrypoints unimplemented.
func (r *Resolver) Entrypoints() []ast.Ref {
	panic("unreachable")
}

// Close unimplemented.
func (r *Resolver) Close() {
	panic("unreachable")
}

// Eval unimplemented.
func (r *Resolver) Eval(context.Context, resolver.Input) (resolver.Result, error) {

	panic("unreachable")
}

// SetData unimplemented.
func (r *Resolver) SetData(context.Context, interface{}) error {
	panic("unreachable")
}

// SetDataPath unimplemented.
func (r *Resolver) SetDataPath(context.Context, []string, interface{}) error {
	panic("unreachable")
}

// RemoveDataPath unimplemented.
func (r *Resolver) RemoveDataPath(context.Context, []string) error {
	panic("unreachable")
}

// New unimplemented. Will always return an error.
func New([]ast.Ref, []byte, interface{}) (*Resolver, error) {
	return nil, errors.New("WebAssembly runtime not supported in this build")
}
