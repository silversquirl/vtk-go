package main

import "../"

func main() {
	root, err := vtk.New()
	if err != nil {
		panic(err)
	}
	defer root.Destroy()

	win, err := root.NewWindow("Hello, world!", 0, 0, 800, 600)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	cr := win.Cairo()
	flip := false

	win.SetEventHandler(vtk.Close, func(ev vtk.Event) {
		win.Close()
	})

	win.SetEventHandler(vtk.Draw, func(ev vtk.Event) {
		width, height := win.Size()
		w, h := float64(width), float64(height)
		cr.Rectangle(0, 0, w, h)
		cr.SetSourceRGB(0, 0, 0)
		cr.Fill()

		if flip {
			cr.MoveTo(w, 0)
			cr.LineTo(0, h)
		} else {
			cr.MoveTo(0, 0)
			cr.LineTo(w, h)
		}
		cr.SetSourceRGB(1, 1, 1)
		cr.SetLineWidth(2)
		cr.Stroke()
	})

	win.SetEventHandler(vtk.KeyPress, func(ev vtk.Event) {
		k := ev.(vtk.KeyEvent)
		if k.Key() == vtk.Escape {
			win.Close()
		} else if k.Key() == 'f' && k.Mods() == vtk.Control {
			flip = !flip
			win.Redraw()
		}
	})

	win.Mainloop()
}
