package main

import (
	"go.vktec.org.uk/vtk"
	"image/color/palette"
)

func main() {
	root, err := vtk.New()
	if err != nil {
		panic(err)
	}
	defer root.Destroy()

	win, err := root.NewWindow("vtk mouse example", 0, 0, 800, 600)
	if err != nil {
		panic(err)
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
		cr.SetSourceColor(palette.Plan9[0])
		cr.Fill()

		cr.MoveTo(sx, sy)
		cr.LineTo(ex, ey)
		cr.SetSourceColor(palette.Plan9[255])
		cr.SetLineWidth(2)
		cr.Stroke()
	})

	win.SetEventHandler(vtk.MouseMove, func(ev vtk.Event) {
		m := ev.(vtk.MouseMoveEvent)
		if vtk.HasMod(m, vtk.LeftButton) {
			mx, my := m.Pos()
			ex, ey = float64(mx), float64(my)
			win.Redraw()
		}
	})

	win.SetEventHandler(vtk.MousePress, func(ev vtk.Event) {
		b := ev.(vtk.MouseButtonEvent)
		if b.Btn() == vtk.LeftButton {
			bx, by := b.Pos()
			sx, sy = float64(bx), float64(by)
			ex, ey = sx, sy
			win.Redraw()
		}
	})

	win.Mainloop()
}
