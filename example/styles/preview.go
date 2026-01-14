// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025 The Ebitengine Authors

package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/toolvox/debugui"
)

func (g *Game) previewWindow(ctx *debugui.Context) {
	// Update the live style with current values
	g.updateLiveStyle(ctx)

	_, height := ebiten.WindowSize()
	x, y := 390, 20
	ctx.SetStyle("live")
	defer ctx.SetStyle("default")

	ctx.Window("Style Preview", image.Rect(x, y, x+500, y+height-40), func(layout debugui.ContainerLayout) {
		ctx.Header("Buttons", true, func() {
			ctx.SetGridLayout([]int{-1, -1, -1}, nil)
			ctx.Button("Button 1")
			ctx.Button("Button 2")
			ctx.Button("Button 3")
		})

		ctx.Header("Checkboxes", true, func() {
			ctx.Checkbox(&g.checkbox1, "Checkbox Option 1")
			ctx.Checkbox(&g.checkbox2, "Checkbox Option 2")
		})

		ctx.Header("Text Input", true, func() {
			ctx.TextField(&g.textField)
		})

		ctx.Header("Number Controls", true, func() {
			ctx.SetGridLayout([]int{-2, -1}, nil)
			ctx.Text("Integer:")
			ctx.NumberField(&g.numberValue, 1)
			ctx.Text("Float:")
			ctx.NumberFieldF(&g.floatValue, 0.1, 2)
			ctx.Text("Slider:")
			ctx.Slider(&g.sliderValue, 0, 100, 1)
		})

		ctx.Header("Dropdown", true, func() {
			ctx.SetGridLayout([]int{-2, -1}, nil)
			ctx.Text("Select Option:")
			ctx.Dropdown(&g.selectedOpt, g.dropdownOpts)
		})

		ctx.Header("Tree Nodes", true, func() {
			ctx.TreeNode("Parent Node 1", func() {
				ctx.Text("Child item 1")
				ctx.Text("Child item 2")
				ctx.TreeNode("Nested Node", func() {
					ctx.Text("Deeply nested item")
					ctx.Button("Nested Button")
				})
			})
			ctx.TreeNode("Parent Node 2", func() {
				ctx.Checkbox(&g.checkbox1, "Nested Checkbox 1")
				ctx.Checkbox(&g.checkbox2, "Nested Checkbox 2")
			})
		})

		ctx.Header("Panel Example", true, func() {
			ctx.Panel(func(layout debugui.ContainerLayout) {
				ctx.Text("This is content inside a panel.")
				ctx.Text("Panels have their own background color.")
				ctx.Button("Panel Button")
			})
		})

		ctx.Header("Mixed Grid Layout", true, func() {
			ctx.SetGridLayout([]int{-1, -1}, nil)
			ctx.GridCell(func(bounds image.Rectangle) {
				ctx.Text("Left Column")
				ctx.Button("Button A")
				ctx.Button("Button B")
			})
			ctx.GridCell(func(bounds image.Rectangle) {
				ctx.Text("Right Column")
				ctx.Button("Button C")
				ctx.Button("Button D")
			})
		})

		ctx.Header("Text Display", false, func() {
			ctx.Text("This is regular text using the current style.")
			ctx.Text(fmt.Sprintf("Selected: %s", g.dropdownOpts[g.selectedOpt]))
			ctx.Text(fmt.Sprintf("Slider value: %d", g.sliderValue))
			ctx.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		})

		ctx.Header("Popup Example", true, func() {
			popupID := ctx.Popup(func(layout debugui.ContainerLayout, id debugui.PopupID) {
				ctx.Text("This is a popup!")
				ctx.Button("Option 1")
				ctx.Button("Option 2")
				ctx.Button("Close").On(func() {
					ctx.ClosePopup(id)
				})
			})
			ctx.Button("Open Popup").On(func() {
				ctx.OpenPopup(popupID)
			})
		})
	})
}
