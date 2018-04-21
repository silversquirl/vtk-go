package vtk

/*
#cgo pkg-config: vtk
#include <stdlib.h>
#include <vtk.h>
void goEventHandler(vtk_event ev, void *u);
static vtk_event_type vtk_event_get_type(vtk_event ev) { return ev.type; }
static struct vtk_key_event vtk_event_get_key(vtk_event ev) { return ev.key; }
static struct vtk_mouse_move_event vtk_event_get_mouse_move(vtk_event ev) { return ev.mouse_move; }
static struct vtk_mouse_button_event vtk_event_get_mouse_button(vtk_event ev) { return ev.mouse_button; }
static struct vtk_scroll_event vtk_event_get_scroll(vtk_event ev) { return ev.scroll; }
*/
import "C"

import (
	"./cairo"
	"errors"
	"unsafe"

	// TODO: remove this dependency
	"github.com/mattn/go-pointer"
)

var (
	Success            error = nil
	AllocationFailed         = errors.New("Allocation failed")
	NoSuitableVisual         = errors.New("Could not find TrueColor visual for X display")
	XOpenDisplayFailed       = errors.New("Opening X display failed")
	UnknownError             = errors.New("Unknown error")
)

func goError(err C.vtk_err) error {
	switch err {
	case C.VTK_SUCCESS:
		return Success
	case C.VTK_ALLOCATION_FAILED:
		return AllocationFailed
	case C.VTK_NO_SUITABLE_VISUAL:
		return NoSuitableVisual
	case C.VTK_XOPENDISPLAY_FAILED:
		return XOpenDisplayFailed
	default:
		return UnknownError
	}
}

type Root struct{ r C.vtk }

func New() (Root, error) {
	var root Root
	err := C.vtk_new(&root.r)
	return root, goError(err)
}

func (root Root) Destroy() {
	C.vtk_destroy(root.r)
}

type Window struct {
	w        C.vtk_window
	pointers []unsafe.Pointer
}

func (root Root) NewWindow(title string, x, y, w, h int) (Window, error) {
	var win Window
	ctitle := C.CString(title)
	err := C.vtk_window_new(&win.w, root.r, ctitle, C.int(x), C.int(y), C.int(w), C.int(h))
	C.free(unsafe.Pointer(ctitle))
	return win, goError(err)
}

func (win Window) Destroy() {
	C.vtk_window_destroy(win.w)
	for _, ptr := range win.pointers {
		pointer.Unref(ptr)
	}
}

func (win Window) Close() {
	C.vtk_window_close(win.w)
}

func (win Window) Redraw() {
	C.vtk_window_redraw(win.w)
}

func (win Window) Mainloop() {
	C.vtk_window_mainloop(win.w)
}

func (win Window) SetTitle(title string) {
	ctitle := C.CString(title)
	C.vtk_window_set_title(win.w, ctitle)
	C.free(unsafe.Pointer(ctitle))
}

func (win Window) Size() (width, height int) {
	cwidth := C.int(width)
	cheight := C.int(height)
	C.vtk_window_get_size(win.w, &cwidth, &cheight)
	return int(cwidth), int(cheight)
}

func (win Window) Cairo() cairo.Cairo {
	return cairo.Cairo{cairo.CCairoT(unsafe.Pointer(C.vtk_window_get_cairo(win.w)))}
}

type EventType C.vtk_event_type

const (
	Close        EventType = C.VTK_EV_CLOSE
	Draw         EventType = C.VTK_EV_DRAW
	KeyPress     EventType = C.VTK_EV_KEY_PRESS
	KeyRelease   EventType = C.VTK_EV_KEY_RELEASE
	MouseMove    EventType = C.VTK_EV_MOUSE_MOVE
	MousePress   EventType = C.VTK_EV_MOUSE_PRESS
	MouseRelease EventType = C.VTK_EV_MOUSE_RELEASE
	Resize       EventType = C.VTK_EV_RESIZE
	Scroll       EventType = C.VTK_EV_SCROLL
)

type Event interface {
	Type() EventType
}

func (ev C.vtk_event) Type() EventType {
	return EventType(C.vtk_event_get_type(ev))
}

//export goEventHandler
func goEventHandler(cev C.vtk_event, ptr unsafe.Pointer) {
	cb := *pointer.Restore(ptr).(*func(Event))
	var ev Event
	switch cev.Type() {
	case KeyPress, KeyRelease:
		ev = KeyEvent(C.vtk_event_get_key(cev))
	case MouseMove:
		ev = MouseMoveEvent(C.vtk_event_get_mouse_move(cev))
	case MousePress, MouseRelease:
		ev = MouseButtonEvent(C.vtk_event_get_mouse_button(cev))
	case Scroll:
		ev = ScrollEvent(C.vtk_event_get_scroll(cev))
	}
	cb(Event(ev))
}

func (win Window) SetEventHandler(t EventType, cb func(Event)) {
	ptr := pointer.Save(&cb)
	win.pointers = append(win.pointers, ptr)
	C.vtk_window_set_event_handler(win.w, C.vtk_event_type(t), C.vtk_event_handler(C.goEventHandler), ptr)
}

type Key C.vtk_key

const (
	Backspace Key = C.VTK_K_BACKSPACE
	Tab       Key = C.VTK_K_TAB
	Return    Key = C.VTK_K_RETURN
	Escape    Key = C.VTK_K_ESCAPE
	Space     Key = C.VTK_K_SPACE
	Delete    Key = C.VTK_K_DELETE
	Insert    Key = C.VTK_K_INSERT

	PageUp   Key = C.VTK_K_PAGE_UP
	PageDown Key = C.VTK_K_PAGE_DOWN
	Home     Key = C.VTK_K_HOME
	End      Key = C.VTK_K_END
	Up       Key = C.VTK_K_UP
	Down     Key = C.VTK_K_DOWN
	Left     Key = C.VTK_K_LEFT
	Right    Key = C.VTK_K_RIGHT
)

type Modifier C.vtk_modifiers

type EventWithModifiers interface {
	Mods() Modifier
}

func HasMod(e EventWithModifiers, m Modifier) bool {
	return Modifier(e.Mods())&m != 0
}

const (
	Shift    Modifier = C.VTK_M_SHIFT
	CapsLock Modifier = C.VTK_M_CAPS_LOCK
	Control  Modifier = C.VTK_M_CONTROL
	Alt      Modifier = C.VTK_M_ALT
	Super    Modifier = C.VTK_M_SUPER

	LeftButton   Modifier = C.VTK_M_LEFT_BTN
	MiddleButton Modifier = C.VTK_M_MIDDLE_BTN
	RightButton  Modifier = C.VTK_M_RIGHT_BTN
)

type KeyEvent C.struct_vtk_key_event

func (k KeyEvent) Type() EventType {
	return EventType(k._type)
}

func (k KeyEvent) Key() Key {
	return Key(k.key)
}

func (k KeyEvent) Mods() Modifier {
	return Modifier(k.mods)
}

type MouseMoveEvent C.struct_vtk_mouse_move_event

func (m MouseMoveEvent) Mods() Modifier {
	return Modifier(m.mods)
}

func (m MouseMoveEvent) Pos() (x, y int) {
	return int(m.x), int(m.y)
}

func (m MouseMoveEvent) Type() EventType {
	return EventType(m._type)
}

type MouseButtonEvent C.struct_vtk_mouse_button_event

func (b MouseButtonEvent) Btn() Modifier {
	return Modifier(b.btn)
}

func (b MouseButtonEvent) Mods() Modifier {
	return Modifier(b.mods)
}

func (m MouseButtonEvent) Pos() (x, y int) {
	return int(m.x), int(m.y)
}

func (b MouseButtonEvent) Type() EventType {
	return EventType(b._type)
}

type ScrollEvent C.struct_vtk_scroll_event

func (s ScrollEvent) Amount() float64 {
	return float64(s.amount)
}

func (s ScrollEvent) Type() EventType {
	return EventType(s._type)
}
