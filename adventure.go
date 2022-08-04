package main

/*
 * Adventure.
 *
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * SPDX-License-Identifier: BSD-2-clause
 *
 * Translate to GO by Robert Shofner (c) 2022
 *  compile with: -ldflags -H=windowsgui
 *
 */
import (
	"adventure/cmd"
	"adventure/images"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
	"time"
)

func main() {
	settings := cmd.Settings{AllowSave: true, Logger: false}
	settings.App = app.NewWithID("com.scsi.adventure")
	settings.Window = settings.App.NewWindow("GO Adventure (by Bob)")
	settings.App.Settings().SetTheme(&myTheme{})

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "-log" {
		logger, err := os.OpenFile("adventure.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err == nil {
			defer func() {
				_ = logger.Close()
			}()
			log.SetOutput(logger)
			settings.Logger = true
		}
	}

	cave := canvas.NewImageFromResource(resourceColossalCaveAdventureMapJpg)
	cave.FillMode = canvas.ImageFillContain
	splash := container.NewMax(cave)

	buttons := make([]*widget.Button, 0)
	f1 := widget.NewButton("score", func() {
		log.Println("score")
	})
	f2 := widget.NewButton("inventory", func() {
		log.Println("inventory")
	})
	buttons = append(buttons, f1)
	buttons = append(buttons, f2)

	settings.TTY, settings.CRT = cmd.NewConsole(cmd.Prompt, buttons, 30), display.NewDisplay()
	content := container.NewHSplit(settings.TTY.Content, settings.CRT.Content)
	content.Offset = .7

	settings.Window.SetContent(splash)
	settings.Window.Resize(fyne.NewSize(800, 500))
	settings.Window.SetFixedSize(false)

	settings.CRT.Show(true, "building") // top
	settings.CRT.Show(false, "map")     // bottom
	settings.TTY.Speak("\nGO Adventure (by Bob)\n  Kudos to:\nBSD (c) 1977, 2005\n  Will Crowther \n  Don Woods\nBSD (c) 2017\n  Eric S. Raymond\nColossal Cave images\n  Mari J. Michaelis @www.spitenet.com\nImages (Simplified Pixabay License) from pixabay.com")

	go func() { // show entiree cave map for a bit
		time.Sleep(time.Millisecond * 3000)
		splash.Objects[0] = content
		splash.Refresh()
		settings.TTY.Focus()
	}()

	go cmd.Main(settings)
	settings.Window.ShowAndRun()
	settings.App.Quit()
}

var _ fyne.Theme = (*myTheme)(nil)

type myTheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 { // SizeNamePadding
	switch name {
	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameText:
		return 12
	}
	return theme.DefaultTheme().Size(name)
}
