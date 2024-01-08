// Code generated by "goki generate"; DO NOT EDIT.

package texteditor

import (
	"image"

	"goki.dev/gi/v2/gi"
	"goki.dev/girl/paint"
	"goki.dev/girl/units"
	"goki.dev/gix/texteditor/textbuf"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
)

// DiffViewType is the [gti.Type] for [DiffView]
var DiffViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gix/texteditor.DiffView",
	ShortName:  "texteditor.DiffView",
	IDName:     "diff-view",
	Doc:        "DiffView presents two side-by-side TextEditor windows showing the differences\nbetween two files (represented as lines of strings).",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"FileA", &gti.Field{Name: "FileA", Type: "string", LocalType: "string", Doc: "first file name being compared", Directives: gti.Directives{}, Tag: ""}},
		{"FileB", &gti.Field{Name: "FileB", Type: "string", LocalType: "string", Doc: "second file name being compared", Directives: gti.Directives{}, Tag: ""}},
		{"RevA", &gti.Field{Name: "RevA", Type: "string", LocalType: "string", Doc: "revision for first file, if relevant", Directives: gti.Directives{}, Tag: ""}},
		{"RevB", &gti.Field{Name: "RevB", Type: "string", LocalType: "string", Doc: "revision for second file, if relevant", Directives: gti.Directives{}, Tag: ""}},
		{"BufA", &gti.Field{Name: "BufA", Type: "*goki.dev/gix/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for A showing the aligned edit view", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" set:\"-\""}},
		{"BufB", &gti.Field{Name: "BufB", Type: "*goki.dev/gix/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for B showing the aligned edit view", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" set:\"-\""}},
		{"AlignD", &gti.Field{Name: "AlignD", Type: "goki.dev/gix/texteditor/textbuf.Diffs", LocalType: "textbuf.Diffs", Doc: "aligned diffs records diff for aligned lines", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" set:\"-\""}},
		{"Diffs", &gti.Field{Name: "Diffs", Type: "goki.dev/gix/texteditor/textbuf.DiffSelected", LocalType: "textbuf.DiffSelected", Doc: "Diffs applied", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Frame", &gti.Field{Name: "Frame", Type: "goki.dev/gi/v2/gi.Frame", LocalType: "gi.Frame", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"SaveFileA", &gti.Method{Name: "SaveFileA", Doc: "SaveFileA saves the current state of file A to given filename", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveFileB", &gti.Method{Name: "SaveFileB", Doc: "SaveFileB saves the current state of file B to given filename", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
	}),
	Instance: &DiffView{},
})

// NewDiffView adds a new [DiffView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewDiffView(par ki.Ki, name ...string) *DiffView {
	return par.NewChild(DiffViewType, name...).(*DiffView)
}

// KiType returns the [*gti.Type] of [DiffView]
func (t *DiffView) KiType() *gti.Type {
	return DiffViewType
}

// New returns a new [*DiffView] value
func (t *DiffView) New() ki.Ki {
	return &DiffView{}
}

// SetFileA sets the [DiffView.FileA]:
// first file name being compared
func (t *DiffView) SetFileA(v string) *DiffView {
	t.FileA = v
	return t
}

// SetFileB sets the [DiffView.FileB]:
// second file name being compared
func (t *DiffView) SetFileB(v string) *DiffView {
	t.FileB = v
	return t
}

// SetRevA sets the [DiffView.RevA]:
// revision for first file, if relevant
func (t *DiffView) SetRevA(v string) *DiffView {
	t.RevA = v
	return t
}

// SetRevB sets the [DiffView.RevB]:
// revision for second file, if relevant
func (t *DiffView) SetRevB(v string) *DiffView {
	t.RevB = v
	return t
}

// SetDiffs sets the [DiffView.Diffs]:
// Diffs applied
func (t *DiffView) SetDiffs(v textbuf.DiffSelected) *DiffView {
	t.Diffs = v
	return t
}

// SetTooltip sets the [DiffView.Tooltip]
func (t *DiffView) SetTooltip(v string) *DiffView {
	t.Tooltip = v
	return t
}

// SetCustomContextMenu sets the [DiffView.CustomContextMenu]
func (t *DiffView) SetCustomContextMenu(v func(m *gi.Scene)) *DiffView {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [DiffView.StackTop]
func (t *DiffView) SetStackTop(v int) *DiffView {
	t.StackTop = v
	return t
}

// SetStripes sets the [DiffView.Stripes]
func (t *DiffView) SetStripes(v gi.Stripes) *DiffView {
	t.Stripes = v
	return t
}

// DiffTextEditorType is the [gti.Type] for [DiffTextEditor]
var DiffTextEditorType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gix/texteditor.DiffTextEditor",
	ShortName:  "texteditor.DiffTextEditor",
	IDName:     "diff-text-editor",
	Doc:        "DiffTextEditor supports double-click based application of edits from one\nbuffer to the other.",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Editor", &gti.Field{Name: "Editor", Type: "goki.dev/gix/texteditor.Editor", LocalType: "Editor", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &DiffTextEditor{},
})

// NewDiffTextEditor adds a new [DiffTextEditor] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewDiffTextEditor(par ki.Ki, name ...string) *DiffTextEditor {
	return par.NewChild(DiffTextEditorType, name...).(*DiffTextEditor)
}

// KiType returns the [*gti.Type] of [DiffTextEditor]
func (t *DiffTextEditor) KiType() *gti.Type {
	return DiffTextEditorType
}

// New returns a new [*DiffTextEditor] value
func (t *DiffTextEditor) New() ki.Ki {
	return &DiffTextEditor{}
}

// SetTooltip sets the [DiffTextEditor.Tooltip]
func (t *DiffTextEditor) SetTooltip(v string) *DiffTextEditor {
	t.Tooltip = v
	return t
}

// SetCustomContextMenu sets the [DiffTextEditor.CustomContextMenu]
func (t *DiffTextEditor) SetCustomContextMenu(v func(m *gi.Scene)) *DiffTextEditor {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [DiffTextEditor.StackTop]
func (t *DiffTextEditor) SetStackTop(v int) *DiffTextEditor {
	t.StackTop = v
	return t
}

// SetPlaceholder sets the [DiffTextEditor.Placeholder]
func (t *DiffTextEditor) SetPlaceholder(v string) *DiffTextEditor {
	t.Placeholder = v
	return t
}

// SetCursorWidth sets the [DiffTextEditor.CursorWidth]
func (t *DiffTextEditor) SetCursorWidth(v units.Value) *DiffTextEditor {
	t.CursorWidth = v
	return t
}

// SetLineNumberColor sets the [DiffTextEditor.LineNumberColor]
func (t *DiffTextEditor) SetLineNumberColor(v image.Image) *DiffTextEditor {
	t.LineNumberColor = v
	return t
}

// SetSelectColor sets the [DiffTextEditor.SelectColor]
func (t *DiffTextEditor) SetSelectColor(v image.Image) *DiffTextEditor {
	t.SelectColor = v
	return t
}

// SetHighlightColor sets the [DiffTextEditor.HighlightColor]
func (t *DiffTextEditor) SetHighlightColor(v image.Image) *DiffTextEditor {
	t.HighlightColor = v
	return t
}

// SetCursorColor sets the [DiffTextEditor.CursorColor]
func (t *DiffTextEditor) SetCursorColor(v image.Image) *DiffTextEditor {
	t.CursorColor = v
	return t
}

// SetLinkHandler sets the [DiffTextEditor.LinkHandler]
func (t *DiffTextEditor) SetLinkHandler(v func(tl *paint.TextLink)) *DiffTextEditor {
	t.LinkHandler = v
	return t
}

// EditorType is the [gti.Type] for [Editor]
var EditorType = gti.AddType(&gti.Type{
	Name:      "goki.dev/gix/texteditor.Editor",
	ShortName: "texteditor.Editor",
	IDName:    "editor",
	Doc:       "Editor is a widget for editing multiple lines of text (as compared to\n[gi.TextField] for a single line).  The Editor is driven by a [Buf]\nbuffer which contains all the text, and manages all the edits,\nsending update signals out to the views.\n\nUse SetNeedsRender to drive an render update for any change that does\nnot change the line-level layout of the text.\nUse SetNeedsLayout whenever there are changes across lines that require\nre-layout of the text.  This sets the Widget NeedsRender flag and triggers\nlayout during that render.\n\nMultiple views can be attached to a given buffer.  All updating in the\nEditor should be within a single goroutine, as it would require\nextensive protections throughout code otherwise.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "goki", Directive: "embedder", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Buf", &gti.Field{Name: "Buf", Type: "*goki.dev/gix/texteditor.Buf", LocalType: "*Buf", Doc: "the text buffer that we're editing", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"Placeholder", &gti.Field{Name: "Placeholder", Type: "string", LocalType: "string", Doc: "text that is displayed when the field is empty, in a lower-contrast manner", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"placeholder\""}},
		{"CursorWidth", &gti.Field{Name: "CursorWidth", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "width of cursor -- set from cursor-width property (inherited)", Directives: gti.Directives{}, Tag: "xml:\"cursor-width\""}},
		{"LineNumberColor", &gti.Field{Name: "LineNumberColor", Type: "image.Image", LocalType: "image.Image", Doc: "the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"SelectColor", &gti.Field{Name: "SelectColor", Type: "image.Image", LocalType: "image.Image", Doc: "the color used for the user text selection background color; this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"HighlightColor", &gti.Field{Name: "HighlightColor", Type: "image.Image", LocalType: "image.Image", Doc: "the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"CursorColor", &gti.Field{Name: "CursorColor", Type: "image.Image", LocalType: "image.Image", Doc: "the color used for the text field cursor (caret); this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"NLines", &gti.Field{Name: "NLines", Type: "int", LocalType: "int", Doc: "number of lines in the view -- sync'd with the Buf after edits, but always reflects storage size of Renders etc", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"Renders", &gti.Field{Name: "Renders", Type: "[]goki.dev/girl/paint.Text", LocalType: "[]paint.Text", Doc: "renders of the text lines, with one render per line (each line could visibly wrap-around, so these are logical lines, not display lines)", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"Offs", &gti.Field{Name: "Offs", Type: "[]float32", LocalType: "[]float32", Doc: "starting render offsets for top of each line", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoDigs", &gti.Field{Name: "LineNoDigs", Type: "int", LocalType: "int", Doc: "number of line number digits needed", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoOff", &gti.Field{Name: "LineNoOff", Type: "float32", LocalType: "float32", Doc: "horizontal offset for start of text after line numbers", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoRender", &gti.Field{Name: "LineNoRender", Type: "goki.dev/girl/paint.Text", LocalType: "paint.Text", Doc: "render for line numbers", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorPos", &gti.Field{Name: "CursorPos", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "current cursor position", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorTarg", &gti.Field{Name: "CursorTarg", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "target cursor position for externally-set targets: ensures that it is visible", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorCol", &gti.Field{Name: "CursorCol", Type: "int", LocalType: "int", Doc: "desired cursor column -- where the cursor was last when moved using left / right arrows -- used when doing up / down to not always go to short line columns", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"PosHistIdx", &gti.Field{Name: "PosHistIdx", Type: "int", LocalType: "int", Doc: "current index within PosHistory", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectStart", &gti.Field{Name: "SelectStart", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "starting point for selection -- will either be the start or end of selected region depending on subsequent selection.", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectReg", &gti.Field{Name: "SelectReg", Type: "goki.dev/gix/texteditor/textbuf.Region", LocalType: "textbuf.Region", Doc: "current selection region", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"PrevSelectReg", &gti.Field{Name: "PrevSelectReg", Type: "goki.dev/gix/texteditor/textbuf.Region", LocalType: "textbuf.Region", Doc: "previous selection region, that was actually rendered -- needed to update render", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"Highlights", &gti.Field{Name: "Highlights", Type: "[]goki.dev/gix/texteditor/textbuf.Region", LocalType: "[]textbuf.Region", Doc: "highlighted regions, e.g., for search results", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"Scopelights", &gti.Field{Name: "Scopelights", Type: "[]goki.dev/gix/texteditor/textbuf.Region", LocalType: "[]textbuf.Region", Doc: "highlighted regions, specific to scope markers", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectMode", &gti.Field{Name: "SelectMode", Type: "bool", LocalType: "bool", Doc: "if true, select text as cursor moves", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ForceComplete", &gti.Field{Name: "ForceComplete", Type: "bool", LocalType: "bool", Doc: "if true, complete regardless of any disqualifying reasons", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ISearch", &gti.Field{Name: "ISearch", Type: "goki.dev/gix/texteditor.ISearch", LocalType: "ISearch", Doc: "interactive search data", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"QReplace", &gti.Field{Name: "QReplace", Type: "goki.dev/gix/texteditor.QReplace", LocalType: "QReplace", Doc: "query replace data", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"FontHeight", &gti.Field{Name: "FontHeight", Type: "float32", LocalType: "float32", Doc: "font height, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LineHeight", &gti.Field{Name: "LineHeight", Type: "float32", LocalType: "float32", Doc: "line height, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"FontAscent", &gti.Field{Name: "FontAscent", Type: "float32", LocalType: "float32", Doc: "font ascent, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"FontDescent", &gti.Field{Name: "FontDescent", Type: "float32", LocalType: "float32", Doc: "font descent, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"NLinesChars", &gti.Field{Name: "NLinesChars", Type: "image.Point", LocalType: "image.Point", Doc: "height in lines and width in chars of the visible area", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LinesSize", &gti.Field{Name: "LinesSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "total size of all lines as rendered", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"TotalSize", &gti.Field{Name: "TotalSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "TotalSize = LinesSize plus extra space and line numbers etc", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LineLayoutSize", &gti.Field{Name: "LineLayoutSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "LineLayoutSize is Geom.Size.Actual.Total subtracting\nextra space and line numbers -- this is what\nLayoutStdLR sees for laying out each line", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"BlinkOn", &gti.Field{Name: "BlinkOn", Type: "bool", LocalType: "bool", Doc: "oscillates between on and off for blinking", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorMu", &gti.Field{Name: "CursorMu", Type: "sync.Mutex", LocalType: "sync.Mutex", Doc: "mutex protecting cursor rendering -- shared between blink and main code", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
		{"HasLinks", &gti.Field{Name: "HasLinks", Type: "bool", LocalType: "bool", Doc: "at least one of the renders has links -- determines if we set the cursor for hand movements", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LinkHandler", &gti.Field{Name: "LinkHandler", Type: "func(tl *goki.dev/girl/paint.TextLink)", LocalType: "func(tl *paint.TextLink)", Doc: "handles link clicks -- if nil, they are sent to the standard web URL handler", Directives: gti.Directives{}, Tag: ""}},
		{"lastRecenter", &gti.Field{Name: "lastRecenter", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"lastAutoInsert", &gti.Field{Name: "lastAutoInsert", Type: "rune", LocalType: "rune", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"lastFilename", &gti.Field{Name: "lastFilename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"Lookup", &gti.Method{Name: "Lookup", Doc: "Lookup attempts to lookup symbol at current location, popping up a window\nif something is found", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
	}),
	Instance: &Editor{},
})

// NewEditor adds a new [Editor] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewEditor(par ki.Ki, name ...string) *Editor {
	return par.NewChild(EditorType, name...).(*Editor)
}

// KiType returns the [*gti.Type] of [Editor]
func (t *Editor) KiType() *gti.Type {
	return EditorType
}

// New returns a new [*Editor] value
func (t *Editor) New() ki.Ki {
	return &Editor{}
}

// EditorEmbedder is an interface that all types that embed Editor satisfy
type EditorEmbedder interface {
	AsEditor() *Editor
}

// AsEditor returns the given value as a value of type Editor if the type
// of the given value embeds Editor, or nil otherwise
func AsEditor(k ki.Ki) *Editor {
	if k == nil || k.This() == nil {
		return nil
	}
	if t, ok := k.(EditorEmbedder); ok {
		return t.AsEditor()
	}
	return nil
}

// AsEditor satisfies the [EditorEmbedder] interface
func (t *Editor) AsEditor() *Editor {
	return t
}

// SetPlaceholder sets the [Editor.Placeholder]:
// text that is displayed when the field is empty, in a lower-contrast manner
func (t *Editor) SetPlaceholder(v string) *Editor {
	t.Placeholder = v
	return t
}

// SetCursorWidth sets the [Editor.CursorWidth]:
// width of cursor -- set from cursor-width property (inherited)
func (t *Editor) SetCursorWidth(v units.Value) *Editor {
	t.CursorWidth = v
	return t
}

// SetLineNumberColor sets the [Editor.LineNumberColor]:
// the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties
func (t *Editor) SetLineNumberColor(v image.Image) *Editor {
	t.LineNumberColor = v
	return t
}

// SetSelectColor sets the [Editor.SelectColor]:
// the color used for the user text selection background color; this should be set in Stylers like all other style properties
func (t *Editor) SetSelectColor(v image.Image) *Editor {
	t.SelectColor = v
	return t
}

// SetHighlightColor sets the [Editor.HighlightColor]:
// the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties
func (t *Editor) SetHighlightColor(v image.Image) *Editor {
	t.HighlightColor = v
	return t
}

// SetCursorColor sets the [Editor.CursorColor]:
// the color used for the text field cursor (caret); this should be set in Stylers like all other style properties
func (t *Editor) SetCursorColor(v image.Image) *Editor {
	t.CursorColor = v
	return t
}

// SetLinkHandler sets the [Editor.LinkHandler]:
// handles link clicks -- if nil, they are sent to the standard web URL handler
func (t *Editor) SetLinkHandler(v func(tl *paint.TextLink)) *Editor {
	t.LinkHandler = v
	return t
}

// SetTooltip sets the [Editor.Tooltip]
func (t *Editor) SetTooltip(v string) *Editor {
	t.Tooltip = v
	return t
}

// SetCustomContextMenu sets the [Editor.CustomContextMenu]
func (t *Editor) SetCustomContextMenu(v func(m *gi.Scene)) *Editor {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [Editor.StackTop]
func (t *Editor) SetStackTop(v int) *Editor {
	t.StackTop = v
	return t
}

// TwinEditorsType is the [gti.Type] for [TwinEditors]
var TwinEditorsType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gix/texteditor.TwinEditors",
	ShortName:  "texteditor.TwinEditors",
	IDName:     "twin-editors",
	Doc:        "TwinEditors presents two side-by-side [Editor]s in [gi.Splits]\nthat scroll in sync with each other.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"BufA", &gti.Field{Name: "BufA", Type: "*goki.dev/gix/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for A", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
		{"BufB", &gti.Field{Name: "BufB", Type: "*goki.dev/gix/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for B", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Splits", &gti.Field{Name: "Splits", Type: "goki.dev/gi/v2/gi.Splits", LocalType: "gi.Splits", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TwinEditors{},
})

// NewTwinEditors adds a new [TwinEditors] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTwinEditors(par ki.Ki, name ...string) *TwinEditors {
	return par.NewChild(TwinEditorsType, name...).(*TwinEditors)
}

// KiType returns the [*gti.Type] of [TwinEditors]
func (t *TwinEditors) KiType() *gti.Type {
	return TwinEditorsType
}

// New returns a new [*TwinEditors] value
func (t *TwinEditors) New() ki.Ki {
	return &TwinEditors{}
}

// SetBufA sets the [TwinEditors.BufA]:
// textbuf for A
func (t *TwinEditors) SetBufA(v *Buf) *TwinEditors {
	t.BufA = v
	return t
}

// SetBufB sets the [TwinEditors.BufB]:
// textbuf for B
func (t *TwinEditors) SetBufB(v *Buf) *TwinEditors {
	t.BufB = v
	return t
}

// SetTooltip sets the [TwinEditors.Tooltip]
func (t *TwinEditors) SetTooltip(v string) *TwinEditors {
	t.Tooltip = v
	return t
}

// SetCustomContextMenu sets the [TwinEditors.CustomContextMenu]
func (t *TwinEditors) SetCustomContextMenu(v func(m *gi.Scene)) *TwinEditors {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [TwinEditors.StackTop]
func (t *TwinEditors) SetStackTop(v int) *TwinEditors {
	t.StackTop = v
	return t
}
