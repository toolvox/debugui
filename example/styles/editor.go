// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025 The Ebitengine Authors

package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/toolvox/debugui"
)

func (g *Game) styleEditorWindow(ctx *debugui.Context) {
	_, height := ebiten.WindowSize()
	x, y := 20, 20
	ctx.Window("Style Editor", image.Rect(x, y, x+350, y+height-40), func(layout debugui.ContainerLayout) {
		ctx.Header("Dimensions", true, func() {
			ctx.SetGridLayout([]int{-2, -1}, nil)
			ctx.Text("Default Width:")
			ctx.NumberField(&g.defaultWidth, 1)
			ctx.Text("Default Height:")
			ctx.NumberField(&g.defaultHeight, 1)
			ctx.Text("Padding:")
			ctx.NumberField(&g.padding, 1)
			ctx.Text("Spacing:")
			ctx.NumberField(&g.spacing, 1)
			ctx.Text("Indent:")
			ctx.NumberField(&g.indent, 1)
			ctx.Text("Title Height:")
			ctx.NumberField(&g.titleHeight, 1)
			ctx.Text("Scrollbar Size:")
			ctx.NumberField(&g.scrollbarSize, 1)
			ctx.Text("Thumb Size:")
			ctx.NumberField(&g.thumbSize, 1)
		})

		ctx.Header("Colors", true, func() {
			g.colorEditor(ctx, "Text", &g.colorText)
			g.colorEditor(ctx, "Border", &g.colorBorder)
			g.colorEditor(ctx, "Window BG", &g.colorWindowBG)
			g.colorEditor(ctx, "Title BG", &g.colorTitleBG)
			g.colorEditor(ctx, "Title BG Trans", &g.colorTitleBGTransparent)
			g.colorEditor(ctx, "Title Text", &g.colorTitleText)
			g.colorEditor(ctx, "Panel BG", &g.colorPanelBG)
			g.colorEditor(ctx, "Button", &g.colorButton)
			g.colorEditor(ctx, "Button Hover", &g.colorButtonHover)
			g.colorEditor(ctx, "Button Focus", &g.colorButtonFocus)
			g.colorEditor(ctx, "Base", &g.colorBase)
			g.colorEditor(ctx, "Base Hover", &g.colorBaseHover)
			g.colorEditor(ctx, "Base Focus", &g.colorBaseFocus)
			g.colorEditor(ctx, "Scroll Base", &g.colorScrollBase)
			g.colorEditor(ctx, "Scroll Thumb", &g.colorScrollThumb)
		})

		ctx.Header("Export", true, func() {
			ctx.Button("Print Style to Stdout").On(func() {
				g.printStyle()
			})
		})
	})
}

func (g *Game) colorEditor(ctx *debugui.Context, label string, colorVals *[4]int) {
	ctx.TreeNode(label, func() {
		ctx.SetGridLayout([]int{-3, -2, 40}, nil)

		ctx.Text("R:")
		ctx.Slider(&colorVals[0], 0, 255, 1)
		ctx.Text(fmt.Sprintf("%d", colorVals[0]))

		ctx.Text("G:")
		ctx.Slider(&colorVals[1], 0, 255, 1)
		ctx.Text(fmt.Sprintf("%d", colorVals[1]))

		ctx.Text("B:")
		ctx.Slider(&colorVals[2], 0, 255, 1)
		ctx.Text(fmt.Sprintf("%d", colorVals[2]))

		ctx.Text("A:")
		ctx.Slider(&colorVals[3], 0, 255, 1)
		ctx.Text(fmt.Sprintf("%d", colorVals[3]))

		ctx.GridCell(func(bounds image.Rectangle) {})
		ctx.GridCell(func(bounds image.Rectangle) {})
		ctx.GridCell(func(bounds image.Rectangle) {
			g.drawColorPreview(ctx, bounds, colorVals)
		})
	})
}

func (g *Game) drawColorPreview(ctx *debugui.Context, bounds image.Rectangle, colorVals *[4]int) {
	ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
		scale := ctx.Scale()
		col := color.RGBA{
			R: byte(colorVals[0]),
			G: byte(colorVals[1]),
			B: byte(colorVals[2]),
			A: byte(colorVals[3]),
		}
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				for sy := 0; sy < scale; sy++ {
					for sx := 0; sx < scale; sx++ {
						screen.Set(x*scale+sx, y*scale+sy, col)
					}
				}
			}
		}
	})
}

func (g *Game) printStyle() {
	colors := [15]color.RGBA{
		g.toRGBA(g.colorText),
		g.toRGBA(g.colorBorder),
		g.toRGBA(g.colorWindowBG),
		g.toRGBA(g.colorTitleBG),
		g.toRGBA(g.colorTitleBGTransparent),
		g.toRGBA(g.colorTitleText),
		g.toRGBA(g.colorPanelBG),
		g.toRGBA(g.colorButton),
		g.toRGBA(g.colorButtonHover),
		g.toRGBA(g.colorButtonFocus),
		g.toRGBA(g.colorBase),
		g.toRGBA(g.colorBaseHover),
		g.toRGBA(g.colorBaseFocus),
		g.toRGBA(g.colorScrollBase),
		g.toRGBA(g.colorScrollThumb),
	}

	fmt.Printf("debugui.NewStyle(\"mystyle\", %d, %d, %d, %d, %d, %d, %d, %d,\n\t[...]color.RGBA{\n",
		g.defaultWidth, g.defaultHeight, g.padding, g.spacing, g.indent, g.titleHeight, g.scrollbarSize, g.thumbSize)

	colorNames := []string{"text", "border", "windowBG", "titleBG", "titleBGTransparent", "titleText",
		"panelBG", "button", "buttonHover", "buttonFocus", "base", "baseHover", "baseFocus", "scrollBase", "scrollThumb"}

	for i, col := range colors {
		fmt.Printf("\t\t{%d, %d, %d, %d}, // %s\n", col.R, col.G, col.B, col.A, colorNames[i])
	}
	fmt.Printf("\t},\n)\n")
}

func (g *Game) toRGBA(vals [4]int) color.RGBA {
	return color.RGBA{
		R: byte(vals[0]),
		G: byte(vals[1]),
		B: byte(vals[2]),
		A: byte(vals[3]),
	}
}

func (g *Game) updateLiveStyle(ctx *debugui.Context) {
	debugui.NewStyle("live",
		g.defaultWidth, g.defaultHeight,
		g.padding, g.spacing,
		g.indent, g.titleHeight,
		g.scrollbarSize, g.thumbSize,
		[15]color.RGBA{
			g.toRGBA(g.colorText),
			g.toRGBA(g.colorBorder),
			g.toRGBA(g.colorWindowBG),
			g.toRGBA(g.colorTitleBG),
			g.toRGBA(g.colorTitleBGTransparent),
			g.toRGBA(g.colorTitleText),
			g.toRGBA(g.colorPanelBG),
			g.toRGBA(g.colorButton),
			g.toRGBA(g.colorButtonHover),
			g.toRGBA(g.colorButtonFocus),
			g.toRGBA(g.colorBase),
			g.toRGBA(g.colorBaseHover),
			g.toRGBA(g.colorBaseFocus),
			g.toRGBA(g.colorScrollBase),
			g.toRGBA(g.colorScrollThumb),
		},
	)
}
