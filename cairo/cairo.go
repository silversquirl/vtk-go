package cairo // import "go.vktec.org.uk/vtk/cairo"

/*
#cgo pkg-config: cairo
#include <stdlib.h>
#include <cairo.h>
*/
import "C"

import "image/color"

type CCairoT *C.cairo_t
type Cairo struct{ Cr *C.cairo_t }
type Pattern struct{ pat *C.cairo_pattern_t }

func (cr Cairo) Save() {
	C.cairo_save(cr.Cr)
}

func (cr Cairo) Restore() {
	C.cairo_restore(cr.Cr)
}

func (cr Cairo) PushGroup() {
	C.cairo_push_group(cr.Cr)
}

func (cr Cairo) PopGroup() Pattern {
	return Pattern{ C.cairo_pop_group(cr.Cr) }
}

func (cr Cairo) PopGroupToSource() {
	C.cairo_pop_group_to_source(cr.Cr)
}

// TODO: set_operator

func (cr Cairo) SetSource(pat Pattern) {
	C.cairo_set_source(cr.Cr, pat.pat)
}

func (cr Cairo) SetSourceRGB(r, g, b float64) {
	C.cairo_set_source_rgb(cr.Cr, C.double(r), C.double(g), C.double(b))
}

func (cr Cairo) SetSourceRGBA(r, g, b, a float64) {
	C.cairo_set_source_rgba(cr.Cr, C.double(r), C.double(g), C.double(b), C.double(a))
}

func (cr Cairo) SetSourceColor(c color.Color) {
	ir, ig, ib, ia := c.RGBA()
	r, g, b, a := float64(ir), float64(ig), float64(ib), float64(ia)

	// Scale the values to 0.0-1.0
	r /= 0xffff
	g /= 0xffff
	b /= 0xffff
	a /= 0xffff

	// Un-premultiply the values
	r /= a
	g /= a
	b /= a

	// Set the source
	cr.SetSourceRGBA(r, g, b, a)
}

// TODO: set_source_surface
// TODO: set_tolerance
// TODO: set_antialias
// TODO: set_fill_rule

func (cr Cairo) SetLineWidth(width float64) {
	C.cairo_set_line_width(cr.Cr, C.double(width))
}

// TODO: set_line_cap
// TODO: set_line_join
// TODO: set_miter_limit

func (cr Cairo) Translate(tx, ty float64) {
	C.cairo_translate(cr.Cr, C.double(tx), C.double(ty))
}

func (cr Cairo) Scale(sx, sy float64) {
	C.cairo_scale(cr.Cr, C.double(sx), C.double(sy))
}

func (cr Cairo) Rotate(angle float64) {
	C.cairo_rotate(cr.Cr, C.double(angle))
}

type Matrix struct{ XX, YX, XY, YY, X0, Y0 float64 }

func (mat Matrix) toC() (cmat *C.cairo_matrix_t) {
	C.cairo_matrix_init(
		cmat,
		C.double(mat.XX),
		C.double(mat.YX),
		C.double(mat.XY),
		C.double(mat.YY),
		C.double(mat.X0),
		C.double(mat.Y0),
	)
	return
}

func (cr Cairo) Transform(mat Matrix) {
	C.cairo_transform(cr.Cr, mat.toC())
}

func (cr Cairo) SetMatrix(mat Matrix) {
	C.cairo_set_matrix(cr.Cr, mat.toC())
}

func (cr Cairo) IdentityMatrix() {
	C.cairo_identity_matrix(cr.Cr)
}

func (cr Cairo) UserToDevice(x, y float64) (float64, float64) {
	cx, cy := C.double(x), C.double(y)
	C.cairo_user_to_device(cr.Cr, &cx, &cy)
	return float64(cx), float64(cy)
}

func (cr Cairo) UserToDeviceDistance(dx, dy float64) (float64, float64) {
	cx, cy := C.double(dx), C.double(dy)
	C.cairo_user_to_device_distance(cr.Cr, &cx, &cy)
	return float64(cx), float64(cy)
}

func (cr Cairo) DeviceToUser(x, y float64) (float64, float64) {
	cx, cy := C.double(x), C.double(y)
	C.cairo_device_to_user(cr.Cr, &cx, &cy)
	return float64(cx), float64(cy)
}

func (cr Cairo) DeviceToUserDistance(dx, dy float64) (float64, float64) {
	cx, cy := C.double(dx), C.double(dy)
	C.cairo_device_to_user_distance(cr.Cr, &cx, &cy)
	return float64(cx), float64(cy)
}

func (cr Cairo) NewPath() {
	C.cairo_new_path(cr.Cr)
}

func (cr Cairo) MoveTo(x, y float64) {
	C.cairo_move_to(cr.Cr, C.double(x), C.double(y))
}

func (cr Cairo) NewSubPath() {
	C.cairo_new_sub_path(cr.Cr)
}

func (cr Cairo) LineTo(x, y float64) {
	C.cairo_line_to(cr.Cr, C.double(x), C.double(y))
}

func (cr Cairo) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(cr.Cr, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(x3), C.double(y3))
}

func (cr Cairo) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(cr.Cr, C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (cr Cairo) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(cr.Cr, C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (cr Cairo) RelMoveTo(dx, dy float64) {
	C.cairo_rel_move_to(cr.Cr, C.double(dx), C.double(dy))
}

func (cr Cairo) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(cr.Cr, C.double(dx), C.double(dy))
}

func (cr Cairo) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(cr.Cr, C.double(dx1), C.double(dy1), C.double(dx2), C.double(dy2), C.double(dx3), C.double(dy3))
}

func (cr Cairo) Rectangle(x, y, w, h float64) {
	C.cairo_rectangle(cr.Cr, C.double(x), C.double(y), C.double(w), C.double(h))
}

func (cr Cairo) ClosePath() {
	C.cairo_close_path(cr.Cr)
}

func (cr Cairo) PathExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_path_extents(cr.Cr, &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

func (cr Cairo) Paint() {
	C.cairo_paint(cr.Cr)
}

func (cr Cairo) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(cr.Cr, C.double(alpha))
}

func (cr Cairo) Mask(pat Pattern) {
	C.cairo_mask(cr.Cr, pat.pat)
}

// TODO: mask_surface

func (cr Cairo) Stroke() {
	C.cairo_stroke(cr.Cr)
}

func (cr Cairo) StrokePreserve() {
	C.cairo_stroke_preserve(cr.Cr)
}

func (cr Cairo) Fill() {
	C.cairo_fill(cr.Cr)
}

func (cr Cairo) FillPreserve() {
	C.cairo_fill_preserve(cr.Cr)
}

func (cr Cairo) CopyPage() {
	C.cairo_copy_page(cr.Cr)
}

func (cr Cairo) ShowPage() {
	C.cairo_show_page(cr.Cr)
}

func (cr Cairo) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(cr.Cr, C.double(x), C.double(y)) != 0
}

func (cr Cairo) InFill(x, y float64) bool {
	return C.cairo_in_fill(cr.Cr, C.double(x), C.double(y)) != 0
}

func (cr Cairo) InClip(x, y float64) bool {
	return C.cairo_in_clip(cr.Cr, C.double(x), C.double(y)) != 0
}

func (cr Cairo) StrokeExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_stroke_extents(cr.Cr, &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

func (cr Cairo) FillExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_fill_extents(cr.Cr, &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

func (cr Cairo) ResetClip() {
	C.cairo_reset_clip(cr.Cr)
}

func (cr Cairo) Clip() {
	C.cairo_clip(cr.Cr)
}

func (cr Cairo) ClipPreserve() {
	C.cairo_clip_preserve(cr.Cr)
}

func (cr Cairo) ClipExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_clip_extents(cr.Cr, &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

// vktec suggests using Pango over Cairo's font API so he hasn't wrapped it

// TODO: get_operator

func (cr Cairo) Source() Pattern {
	return Pattern{ C.cairo_get_source(cr.Cr) }
}

func (cr Cairo) Tolerance() float64 {
	return float64(C.cairo_get_tolerance(cr.Cr))
}

// TODO: get_antialias

func (cr Cairo) HasCurrentPoint() bool {
	return C.cairo_has_current_point(cr.Cr) != 0
}

func (cr Cairo) CurrentPoint() (x, y float64) {
	var cx, cy C.double
	C.cairo_get_current_point(cr.Cr, &cx, &cy)
	return float64(cx), float64(cy)
}

// TODO: get_fill_rule

func (cr Cairo) LineWidth() float64 {
	return float64(C.cairo_get_line_width(cr.Cr))
}

// TODO: get_line_cap
// TODO: get_line_join

func (cr Cairo) MiterLimit() float64 {
	return float64(C.cairo_get_miter_limit(cr.Cr))
}

func (cr Cairo) DashCount() float64 {
	return float64(C.cairo_get_dash_count(cr.Cr))
}

func (cr Cairo) Dash() (dashes, offset float64) {
	var cdashes, coffset C.double
	C.cairo_get_dash(cr.Cr, &cdashes, &coffset)
	return float64(cdashes), float64(coffset)
}

// TODO: get_target
// TODO: get_group_target

// TODO: path_t API
// TODO: error handling
// TODO: surface_t API
// TODO: raster source API

func PatternCreateRGB(r, g, b float64) Pattern {
	return Pattern{ C.cairo_pattern_create_rgb(C.double(r), C.double(g), C.double(b)) }
}

func PatternCreateRGBA(r, g, b, a float64) Pattern {
	return Pattern{ C.cairo_pattern_create_rgba(C.double(r), C.double(g), C.double(b), C.double(a)) }
}

// TODO: pattern_create_for_surface

func PatternCreateLinear(x0, y0, x1, y1 float64) Pattern {
	return Pattern{ C.cairo_pattern_create_linear(C.double(x0), C.double(y0), C.double(x1), C.double(y1)) }
}

func PatternCreateRadial(x0, y0, r0, x1, y1, r1 float64) Pattern {
	return Pattern{ C.cairo_pattern_create_radial(C.double(x0), C.double(y0), C.double(r0), C.double(x1), C.double(y1), C.double(r1)) }
}

func PatternCreateMesh() Pattern {
	return Pattern{ C.cairo_pattern_create_mesh() }
}

func (pat Pattern) Destroy() {
	C.cairo_pattern_destroy(pat.pat)
}

// TODO: pattern_status

type PatternType C.cairo_pattern_type_t
const (
	Solid        PatternType = C.CAIRO_PATTERN_TYPE_SOLID
	Surface      PatternType = C.CAIRO_PATTERN_TYPE_SURFACE
	Linear       PatternType = C.CAIRO_PATTERN_TYPE_LINEAR
	Radial       PatternType = C.CAIRO_PATTERN_TYPE_RADIAL
	Mesh         PatternType = C.CAIRO_PATTERN_TYPE_MESH
	RasterSource PatternType = C.CAIRO_PATTERN_TYPE_RASTER_SOURCE
)

func (pat Pattern) Type() PatternType {
	return PatternType(C.cairo_pattern_get_type(pat.pat))
}

func (pat Pattern) AddColorStopRGB(offset, r, g, b float64) {
	C.cairo_pattern_add_color_stop_rgb(pat.pat, C.double(offset), C.double(r), C.double(g), C.double(b))
}

func (pat Pattern) AddColorStopRGBA(offset, r, g, b, a float64) {
	C.cairo_pattern_add_color_stop_rgba(pat.pat, C.double(offset), C.double(r), C.double(g), C.double(b), C.double(a))
}

func (pat Pattern) MeshBeginPatch() {
	C.cairo_mesh_pattern_begin_patch(pat.pat)
}

func (pat Pattern) MeshEndPatch() {
	C.cairo_mesh_pattern_end_patch(pat.pat)
}

func (pat Pattern) MeshCurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_mesh_pattern_curve_to(pat.pat, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(x3), C.double(y3))
}

func (pat Pattern) MeshLineTo(x1, y1 float64) {
	C.cairo_mesh_pattern_line_to(pat.pat, C.double(x1), C.double(y1))
}

func (pat Pattern) MeshMoveTo(x1, y1 float64) {
	C.cairo_mesh_pattern_move_to(pat.pat, C.double(x1), C.double(y1))
}

func (pat Pattern) MeshSetControlPoint(pointNum uint, x, y float64) {
	C.cairo_mesh_pattern_set_control_point(pat.pat, C.uint(pointNum), C.double(x), C.double(y))
}

func (pat Pattern) MeshSetCornerColorRGB(cornerNum uint, r, g, b float64) {
	C.cairo_mesh_pattern_set_corner_color_rgb(pat.pat, C.uint(cornerNum), C.double(r), C.double(g), C.double(b))
}

func (pat Pattern) MeshSetCornerColorRGBA(cornerNum uint, r, g, b, a float64) {
	C.cairo_mesh_pattern_set_corner_color_rgba(pat.pat, C.uint(cornerNum), C.double(r), C.double(g), C.double(b), C.double(a))
}

// TODO: set_matrix
// TODO: get_matrix

type Extend C.cairo_extend_t
const (
	ExtendNone Extend = C.CAIRO_EXTEND_NONE
	ExtendRepeat Extend = C.CAIRO_EXTEND_REPEAT
	ExtendReflect Extend = C.CAIRO_EXTEND_REFLECT
	ExtendPad Extend = C.CAIRO_EXTEND_PAD
)

func (pat Pattern) SetExtend(extend Extend) {
	C.cairo_pattern_set_extend(pat.pat, C.cairo_extend_t(extend))
}

func (pat Pattern) Extend() Extend {
	return Extend(C.cairo_pattern_get_extend(pat.pat))
}

// TODO: the rest of the pattern API (line 2900 and down in cairo.h)

// TODO: matrix_t API
// TODO: region_t API
