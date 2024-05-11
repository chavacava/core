// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"testing"

	"cogentcore.org/core/core"
)

func TestTreeView(t *testing.T) {
	b := core.NewBody()

	fr := core.NewFrame()
	core.NewButton(fr)
	core.NewText(fr)
	core.NewButton(core.NewLayout(fr))

	NewTreeView(b).SyncTree(fr)
	b.AssertRender(t, "tree-view/basic")
}

func TestTreeViewReadOnly(t *testing.T) {
	b := core.NewBody()

	fr := core.NewFrame()
	core.NewButton(fr)
	core.NewText(fr)
	core.NewButton(core.NewLayout(fr))

	NewTreeView(b).SyncTree(fr).SetReadOnly(true)
	b.AssertRender(t, "tree-view/read-only")
}
