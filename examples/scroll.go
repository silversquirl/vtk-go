package main

import (
	"../"
	"math"
)

func main() {
	root, err := vtk.New()
	if err != nil {
		panic(err)
	}
	defer root.Destroy()

	win, err := root.NewWindow("vtk scroll example", 0, 0, 800, 600)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	cr := win.Cairo()
	angle := 0.0

	win.SetEventHandler(vtk.Close, func(ev vtk.Event) {
		win.Close()
	})

	win.SetEventHandler(vtk.Draw, func(ev vtk.Event) {
		width, height := win.Size()
		w, h := float64(width), float64(height)

		cr.Translate(0, 0)
		cr.Rotate(0)
		cr.Rectangle(0, 0, w, h)
		cr.SetSourceRGB(0, 0, 0)
		cr.Fill()

		cr.Translate(w/2, h/2)
		cr.Rotate(angle)
		cr.MoveTo(-w/2, -h/2)
		cr.LineTo(w/2, h/2)
		cr.SetSourceRGB(1, 1, 1)
		cr.SetLineWidth(2)
		cr.Stroke()
	})

	win.SetEventHandler(vtk.Scroll, func(ev vtk.Event) {
		s := ev.(vtk.ScrollEvent)
		angle -= s.Amount() / (8 * math.Pi)
		win.Redraw()
	})

	win.Mainloop()
}
