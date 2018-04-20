package cairo

/*
#cgo pkg-config: cairo
#include <stdlib.h>
#include <cairo.h>
*/
import "C"

type CCairoT *C.cairo_t
type Cairo struct{ Cr *C.cairo_t }

func (cr Cairo) LineTo(x, y float64) {
	C.cairo_line_to(cr.Cr, C.double(x), C.double(y))
}

func (cr Cairo) MoveTo(x, y float64) {
	C.cairo_move_to(cr.Cr, C.double(x), C.double(y))
}

func (cr Cairo) Rectangle(x, y, w, h float64) {
	C.cairo_rectangle(cr.Cr, C.double(x), C.double(y), C.double(w), C.double(h))
}

func (cr Cairo) SetSourceRGB(r, g, b float64) {
	C.cairo_set_source_rgb(cr.Cr, C.double(r), C.double(g), C.double(b))
}

func (cr Cairo) SetLineWidth(width float64) {
	C.cairo_set_line_width(cr.Cr, C.double(width))
}

func (cr Cairo) Stroke() {
	C.cairo_stroke(cr.Cr)
}

func (cr Cairo) Fill() {
	C.cairo_fill(cr.Cr)
}
