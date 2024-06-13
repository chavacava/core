// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

/* TODO(config)
import (
	"errors"
	"strings"
	"testing"

	"cogentcore.org/core/core"
	"cogentcore.org/core/tree"
)

type validator string

func (v *validator) Validate() error {
	if !strings.Contains(string(*v), "@") {
		return errors.New("must have an @")
	}
	return nil
}

func TestValidatorValid(t *testing.T) {
	b := core.NewBody()
	v := NewValue(b, validator("my@string"))
	b.AssertRender(t, "text/validator-valid", func() {
		v.AsWidgetBase().SendChange() // trigger validation
	})
}

func TestValidatorInvalid(t *testing.T) {
	b := core.NewBody()
	v := NewValue(b, validator("my string"))
	b.AssertRender(t, "text/validator-invalid", func() {
		v.AsWidgetBase().SendChange() // trigger validation
	})
}

type fieldValidator struct {
	Name  string
	Email string
}

func (v *fieldValidator) ValidateField(field string) error {
	switch field {
	case "Name":
		if !strings.Contains(v.Name, " ") {
			return errors.New("need full name")
		}
	case "Email":
		if !strings.Contains(v.Email, "@") || !strings.Contains(v.Email, ".") {
			return errors.New("must have a . and @")
		}
	}
	return nil
}

func TestFieldValidatorValid(t *testing.T) {
	b := core.NewBody()
	v := NewForm(b).SetStruct(&fieldValidator{Name: "Go Gopher", Email: "me@example.com"})
	b.AssertRender(t, "text/field-validator-valid", func() {
		v.WidgetWalkDown(func(kwi core.Widget, kwb *core.WidgetBase) bool {
			kwb.SendChange() // trigger validation
			return tree.Continue
		})
	})
}

func TestFieldValidatorInvalid(t *testing.T) {
	b := core.NewBody()
	v := NewForm(b).SetStruct(&fieldValidator{Name: "Go Gopher", Email: "me@example"})
	b.AssertRender(t, "text/field-validator-invalid", func() {
		v.WidgetWalkDown(func(kwi core.Widget, kwb *core.WidgetBase) bool {
			kwb.SendChange() // trigger validation
			return tree.Continue
		})
	})
}
*/
