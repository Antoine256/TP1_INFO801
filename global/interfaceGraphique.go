package global

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"image/color"
	"os"
)

func InitInterface(etatPompe *string, etatVentilateur *string) {
	func() {
		w := app.NewWindow()
		err := Run(w, etatPompe, etatVentilateur)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
}

func Run(w *app.Window, etatPompe *string, etatVentilateur *string) error {
	th := material.NewTheme()
	var ops op.Ops
	var ctx context.Context
	fmt.Println(ctx)
	for {
		switch e := w.NextEvent().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H4(th, "Etat de la pompe : "+*etatPompe+"\nEtat du ventilateur : "+*etatVentilateur)

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
