package cmd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
)

/*
*  Copyright (c) 2022 by Robert (Bob) Shofner
*  SPDX-License-Identifier: BSD-2-clause
 */
/*
  Description: Set of functions to simulate a console device.
*/
type Console struct {
	Content *fyne.Container
	Buttons []*widget.Button
	Focus   func()
	wait    chan string
	view    *fyne.Container
	row     int // current row
	rowMax  int
	entry   *widget.Entry
	prompt  string
	font    fyne.TextStyle
}

var Prompt = ">> "

func NewConsole(prompt string, buttons []*widget.Button, rowMax int) *Console {
	console := Console{
		Content: nil,
		Buttons: buttons,
		wait:    make(chan string),
		view:    nil,
		row:     0,
		rowMax:  rowMax,
		entry:   widget.NewEntry(),
		prompt:  prompt,
		font:    fyne.TextStyle{Monospace: true},
	}
	var buttonBar *fyne.Container
	if len(buttons) > 0 {
		buttonBar = container.NewHBox(layout.NewSpacer())
		for _, button := range buttons {
			buttonBar.Objects = append(buttonBar.Objects, button)
		}
	}
	console.entry.PlaceHolder = console.prompt
	bottom := container.NewVBox(console.entry)
	if buttonBar != nil {
		bottom.Objects = append(bottom.Objects, buttonBar)
	}
	console.view = container.NewVBox()
	console.Content = container.NewBorder(nil, bottom, nil, nil, console.view)
	console.Focus = func() { settings.Window.Canvas().Focus(console.entry) }
	return &console
}
func (c *Console) Ask(required string) (b string) {
	c.entry.SetText("")
	c.entry.OnSubmitted = func(response string) {
		if response == "" && required != "" {
			c.Speak(c.prompt + required)
		} else {
			c.Speak(c.prompt + response)
			c.wait <- response
		}
	}
	b = <-c.wait
	return

}
func (c *Console) AskYesNo(required string) (b bool) {
	c.entry.SetText("")
	c.entry.OnSubmitted = func(response string) {
		if response != "" {
			l := strings.ToLower(response)[:1]
			if l == "y" || l == "n" {
				c.Speak(Prompt + response)
				c.wait <- strings.ToLower(response)[:1]
				return
			}
		}
		c.entry.SetText("")
		c.Speak(Prompt + required)
	}
	c.Focus()
	ok := <-c.wait
	if ok == "y" {
		b = true
	}
	return
}
func (c *Console) Speak(txt string) {
	lines := strings.Split(txt, "\n")
	if len(lines) > c.rowMax {
		lines = lines[len(lines)-c.rowMax:]
	}
	avail := c.rowMax - c.row
	if avail < len(lines) {
		// remove from top
		c.view.Objects = c.view.Objects[len(lines)-avail:]
		c.row = c.rowMax - len(lines)
	}
	for _, line := range lines {
		t := canvas.NewText(line, theme.ForegroundColor())
		t.TextStyle = c.font
		c.view.Objects = append(c.view.Objects, t)
		c.row++
	}
	c.view.Refresh()
}
