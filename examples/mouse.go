package main

import (
	"../"
	"fmt"
	"os"
)

func myMain() int {
	root, err := vtk.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer root.Destroy()

	win, err := root.NewWindow("vtk mouse example", 0, 0, 800, 600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer win.Destroy()

	cr := win.Cairo()
	var sx, sy, ex, ey float64

	win.SetEventHandler(vtk.Close, func(ev vtk.Event) {
		win.Close()
	})

	win.SetEventHandler(vtk.Draw, func(ev vtk.Event) {
		width, height := win.Size()
		w, h := float64(width), float64(height)
		cr.Rectangle(0, 0, w, h)
		cr.SetSourceRGB(0, 0, 0)
		cr.Fill()

		cr.MoveTo(sx, sy)
		cr.LineTo(ex, ey)
		cr.SetSourceRGB(1, 1, 1)
		cr.SetLineWidth(2)
		cr.Stroke()
	})

	win.SetEventHandler(vtk.MouseMove, func(ev vtk.Event) {
		m := ev.(vtk.MouseMoveEvent)
		if vtk.HasMod(m, vtk.LeftButton) {
			ex, ey = m.Pos()
			win.Redraw()
		}
	})

	win.SetEventHandler(vtk.MousePress, func(ev vtk.Event) {
		b := ev.(vtk.MouseButtonEvent)
		if b.Btn() == vtk.LeftButton {
			sx, sy = b.Pos()
			ex, ey = sx, sy
			win.Redraw()
		}
	})

	win.Mainloop()

	return 0
}

func main() {
	os.Exit(myMain())
}
