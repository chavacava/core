// Code generated by "goki generate"; DO NOT EDIT.

package main

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gix/filetree"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
)

// FileBrowseType is the [gti.Type] for [FileBrowse]
var FileBrowseType = gti.AddType(&gti.Type{
	Name:       "main.FileBrowse",
	ShortName:  "main.FileBrowse",
	IDName:     "file-browse",
	Doc:        "FileBrowse is a simple file browser / viewer / editor with a file tree and\none or more editor windows.  It is based on an early version of the Gide\nIDE framework, and remains simple to test / demo the file tree component.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"ProjRoot", &gti.Field{Name: "ProjRoot", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename", Directives: gti.Directives{}, Tag: "desc:\"root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename\""}},
		{"ActiveFilename", &gti.Field{Name: "ActiveFilename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "filename of the currently-active texteditor", Directives: gti.Directives{}, Tag: "desc:\"filename of the currently-active texteditor\""}},
		{"Changed", &gti.Field{Name: "Changed", Type: "bool", LocalType: "bool", Doc: "has the root changed?  we receive update signals from root for changes", Directives: gti.Directives{}, Tag: "json:\"-\" desc:\"has the root changed?  we receive update signals from root for changes\""}},
		{"Files", &gti.Field{Name: "Files", Type: "*goki.dev/gix/filetree.Tree", LocalType: "*filetree.Tree", Doc: "all the files in the project directory and subdirectories", Directives: gti.Directives{}, Tag: "desc:\"all the files in the project directory and subdirectories\""}},
		{"NTextEditors", &gti.Field{Name: "NTextEditors", Type: "int", LocalType: "int", Doc: "number of texteditors available for editing files (default 2) -- configurable with n-text-views property", Directives: gti.Directives{}, Tag: "xml:\"n-text-views\" desc:\"number of texteditors available for editing files (default 2) -- configurable with n-text-views property\""}},
		{"ActiveTextEditorIdx", &gti.Field{Name: "ActiveTextEditorIdx", Type: "int", LocalType: "int", Doc: "index of the currently-active texteditor -- new files will be viewed in other views if available", Directives: gti.Directives{}, Tag: "json:\"-\" desc:\"index of the currently-active texteditor -- new files will be viewed in other views if available\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Frame", &gti.Field{Name: "Frame", Type: "goki.dev/gi/v2/gi.Frame", LocalType: "gi.Frame", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"UpdateFiles", &gti.Method{Name: "UpdateFiles", Doc: "UpdateFiles updates the list of files saved in project", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"OpenPath", &gti.Method{Name: "OpenPath", Doc: "OpenPath opens a new browser viewer at given path, which can either be a\nspecific file or a directory containing multiple files of interest -- opens\nin current FileBrowse object if it is empty, or otherwise opens a new\nwindow.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"path", &gti.Field{Name: "path", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveActiveView", &gti.Method{Name: "SaveActiveView", Doc: "SaveActiveView saves the contents of the currently-active texteditor", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveActiveViewAs", &gti.Method{Name: "SaveActiveViewAs", Doc: "SaveActiveViewAs save with specified filename the contents of the\ncurrently-active texteditor", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"ConfigToolbar", &gti.Method{Name: "ConfigToolbar", Doc: "", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"tb", &gti.Field{Name: "tb", Type: "*goki.dev/gi/v2/gi.Toolbar", LocalType: "*gi.Toolbar", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
	}),
	Instance: &FileBrowse{},
})

// NewFileBrowse adds a new [FileBrowse] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewFileBrowse(par ki.Ki, name ...string) *FileBrowse {
	return par.NewChild(FileBrowseType, name...).(*FileBrowse)
}

// KiType returns the [*gti.Type] of [FileBrowse]
func (t *FileBrowse) KiType() *gti.Type {
	return FileBrowseType
}

// New returns a new [*FileBrowse] value
func (t *FileBrowse) New() ki.Ki {
	return &FileBrowse{}
}

// SetProjRoot sets the [FileBrowse.ProjRoot]:
// root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename
func (t *FileBrowse) SetProjRoot(v gi.FileName) *FileBrowse {
	t.ProjRoot = v
	return t
}

// SetActiveFilename sets the [FileBrowse.ActiveFilename]:
// filename of the currently-active texteditor
func (t *FileBrowse) SetActiveFilename(v gi.FileName) *FileBrowse {
	t.ActiveFilename = v
	return t
}

// SetChanged sets the [FileBrowse.Changed]:
// has the root changed?  we receive update signals from root for changes
func (t *FileBrowse) SetChanged(v bool) *FileBrowse {
	t.Changed = v
	return t
}

// SetFiles sets the [FileBrowse.Files]:
// all the files in the project directory and subdirectories
func (t *FileBrowse) SetFiles(v *filetree.Tree) *FileBrowse {
	t.Files = v
	return t
}

// SetNtextEditors sets the [FileBrowse.NTextEditors]:
// number of texteditors available for editing files (default 2) -- configurable with n-text-views property
func (t *FileBrowse) SetNtextEditors(v int) *FileBrowse {
	t.NTextEditors = v
	return t
}

// SetActiveTextEditorIdx sets the [FileBrowse.ActiveTextEditorIdx]:
// index of the currently-active texteditor -- new files will be viewed in other views if available
func (t *FileBrowse) SetActiveTextEditorIdx(v int) *FileBrowse {
	t.ActiveTextEditorIdx = v
	return t
}

// SetTooltip sets the [FileBrowse.Tooltip]
func (t *FileBrowse) SetTooltip(v string) *FileBrowse {
	t.Tooltip = v
	return t
}

// SetCustomContextMenu sets the [FileBrowse.CustomContextMenu]
func (t *FileBrowse) SetCustomContextMenu(v func(m *gi.Scene)) *FileBrowse {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [FileBrowse.StackTop]
func (t *FileBrowse) SetStackTop(v int) *FileBrowse {
	t.StackTop = v
	return t
}

// SetStripes sets the [FileBrowse.Stripes]
func (t *FileBrowse) SetStripes(v gi.Stripes) *FileBrowse {
	t.Stripes = v
	return t
}
