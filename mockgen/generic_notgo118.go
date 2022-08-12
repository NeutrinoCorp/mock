// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !go1.18
// +build !go1.18

package main

import (
	"fmt"
	"go/ast"

	"github.com/golang/mock/mockgen/model"
)

func getTypeSpecTypeParams(ts *ast.TypeSpec) []*ast.Field {
	return nil
}

func (p *fileParser) parseGenericType(pkg string, typ ast.Expr, tps map[string]bool, tm *instTypeMatcher) (model.Type, error) {
	return nil, nil
}

func getIdentTypeParams(decl interface{}) string {
	return ""
}

func (p *fileParser) parseGenericField(field *ast.Field, it *namedInterface, iface *model.Interface, pkg string, tps map[string]bool) error {
	return fmt.Errorf("don't know how to mock method of type %T", field.Type)
}
