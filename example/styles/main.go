// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025 The Ebitengine Authors

package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/toolvox/debugui"
)

type Game struct {
	debugUI             debugui.DebugUI
	inputCapturingState debugui.InputCapturingState

	// Style parameters
	defaultWidth  int
	defaultHeight int
	padding       int
	spacing       int
	indent        int
	titleHeight   int
	scrollbarSize int
	thumbSize     int

	// Colors (RGBA values)
	colorText               [4]int
	colorBorder             [4]int
	colorWindowBG           [4]int
	colorTitleBG            [4]int
	colorTitleBGTransparent [4]int
	colorTitleText          [4]int
	colorPanelBG            [4]int
	colorButton             [4]int
	colorButtonHover        [4]int
	colorButtonFocus        [4]int
	colorBase               [4]int
	colorBaseHover          [4]int
	colorBaseFocus          [4]int
	colorScrollBase         [4]int
	colorScrollThumb        [4]int

	// Preview window state
	checkbox1    bool
	checkbox2    bool
	textField    string
	selectedOpt  int
	sliderValue  int
	numberValue  int
	floatValue   float64
	dropdownOpts []string
}

func NewGame() (*Game, error) {
	g := &Game{
		// Initialize with default style values
		defaultWidth:  60,
		defaultHeight: 18,
		padding:       5,
		spacing:       4,
		indent:        debugui.LineHeight(),
		titleHeight:   24,
		scrollbarSize: 12,
		thumbSize:     8,

		colorText:               [4]int{230, 230, 230, 255},
		colorBorder:             [4]int{60, 60, 60, 255},
		colorWindowBG:           [4]int{45, 45, 45, 230},
		colorTitleBG:            [4]int{30, 30, 30, 255},
		colorTitleBGTransparent: [4]int{20, 20, 20, 204},
		colorTitleText:          [4]int{240, 240, 240, 255},
		colorPanelBG:            [4]int{0, 0, 0, 0},
		colorButton:             [4]int{75, 75, 75, 255},
		colorButtonHover:        [4]int{95, 95, 95, 255},
		colorButtonFocus:        [4]int{115, 115, 115, 255},
		colorBase:               [4]int{30, 30, 30, 255},
		colorBaseHover:          [4]int{35, 35, 35, 255},
		colorBaseFocus:          [4]int{40, 40, 40, 255},
		colorScrollBase:         [4]int{43, 43, 43, 255},
		colorScrollThumb:        [4]int{30, 30, 30, 255},

		checkbox1:    true,
		checkbox2:    false,
		textField:    "Type here...",
		sliderValue:  50,
		numberValue:  42,
		floatValue:   3.14,
		dropdownOpts: []string{"Option A", "Option B", "Option C", "Option D"},
	}

	return g, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	inputCaptured, err := g.debugUI.Update(func(ctx *debugui.Context) error {
		g.styleEditorWindow(ctx)
		g.previewWindow(ctx)
		return nil
	})
	if err != nil {
		return err
	}
	g.inputCapturingState = inputCaptured
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x40, 0x40, 0x80, 0xff})
	g.debugUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("Style Editor Demo")
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	g, err := NewGame()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := ebiten.RunGame(g); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
