// Code generated by "core generate -webcore content"; DO NOT EDIT.

package main

import (
	"errors"
	"fmt"
	"maps"
	"strings"

	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/webcore"
)

func init() {
	maps.Copy(webcore.Examples, WebcoreExamples)
}

// WebcoreExamples are the compiled webcore examples for this app.
var WebcoreExamples = map[string]func(parent gi.Widget){
	"getting-started/hello-world-0": func(parent gi.Widget) {
		b := parent
		gi.NewButton(b).SetText("Hello, World!")
	},
	"basics/widgets-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").SetIcon(icons.Add)
	},
	"basics/widgets-1": func(parent gi.Widget) {
		sw := gi.NewSwitch(parent).SetText("Switch me!")
		// Later...
		gi.MessageSnackbar(parent, sw.Text)
	},
	"basics/events-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Button clicked")
		})
	},
	"basics/events-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprint("Button clicked at ", e.Pos()))
			e.SetHandled() // this event will not be handled by other event handlers now
		})
	},
	"basics/icons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Download").SetIcon(icons.Download)
	},
	"widgets/buttons-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-2": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Share").SetIcon(icons.Share).SetMenu(func(m *gi.Scene) {
			gi.NewButton(m).SetText("Copy link")
			gi.NewButton(m).SetText("Send message")
		})
	},
	"widgets/buttons-3": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonFilled).SetText("Filled")
	},
	"widgets/buttons-4": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonTonal).SetText("Tonal")
	},
	"widgets/buttons-5": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonElevated).SetText("Elevated")
	},
	"widgets/buttons-6": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonOutlined).SetText("Outlined")
	},
	"widgets/buttons-7": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonText).SetText("Text")
	},
	"widgets/buttons-8": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonAction).SetText("Action")
	},
	"widgets/text-fields-0": func(parent gi.Widget) {
		gi.NewTextField(parent)
	},
	"widgets/text-fields-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Name:")
		gi.NewTextField(parent).SetPlaceholder("Jane Doe")
	},
	"widgets/text-fields-2": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("Hello, world!")
	},
	"widgets/text-fields-3": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("This is a really long sentence that demonstrates how text field content can overflow onto multiple lines")
	},
	"widgets/text-fields-4": func(parent gi.Widget) {
		gi.NewTextField(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/text-fields-5": func(parent gi.Widget) {
		gi.NewTextField(parent).SetTypePassword()
	},
	"widgets/text-fields-6": func(parent gi.Widget) {
		gi.NewTextField(parent).AddClearButton()
	},
	"widgets/text-fields-7": func(parent gi.Widget) {
		gi.NewTextField(parent).SetLeadingIcon(icons.Euro).SetTrailingIcon(icons.OpenInNew, func(e events.Event) {
			gi.MessageSnackbar(parent, "Opening shopping cart")
		})
	},
	"widgets/text-fields-8": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.SetValidator(func() error {
			if !strings.Contains(tf.Text(), "Go") {
				return errors.New("Must contain Go")
			}
			return nil
		})
	},
}
