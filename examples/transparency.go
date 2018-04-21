package main

import (
	".."
	"../cairo"
)

func main() {
	root, err := vtk.New()
	if err != nil {
		panic(err)
	}
	defer root.Destroy()

	win, err := root.NewWindow("vtk transparency example", 0, 0, 800, 600)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	cr := win.Cairo()
	x, y := 0.0, 0.0
	pat := cairo.PatternCreateRadial(0, 0, 0, 0, 0, 200)
	pat.AddColorStopRGBA(0, .4, 1, .4, 1)
	pat.AddColorStopRGBA(.3, 0, 1, 0, .5)
	pat.AddColorStopRGBA(1, 0, 1, 0, 0)

	win.SetEventHandler(vtk.Close, func(ev vtk.Event) {
		win.Close()
	})

	win.SetEventHandler(vtk.Draw, func(ev vtk.Event) {
		cr.Translate(x, y)
		cr.SetSource(pat)
		cr.Paint()
	})

	win.SetEventHandler(vtk.MouseMove, func(ev vtk.Event) {
		m := ev.(vtk.MouseMoveEvent)
		mx, my := m.Pos()
		x, y = float64(mx), float64(my)
		win.Redraw()
	})

	win.Mainloop()
}
