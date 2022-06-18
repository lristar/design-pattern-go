package factoryFunc

import "fmt"

type IButton interface {
	render()
	onClick()
}

type HtmlButton struct {
}

func (h *HtmlButton) render() {
	fmt.Println("html test button")
	h.onClick()
}

func (h *HtmlButton) onClick() {
	fmt.Println("html say hello")
}

type WindowButton struct {
}

func (w *WindowButton) render() {
	fmt.Println("window test button")
	w.onClick()
}
func (w *WindowButton) onClick() {
	fmt.Println("window say hello")
}


var dialog IDialog

type IDialog interface {
	renderWindow()
	createButton() IButton
}

// ----------------- 通用Dialog
type Dialog struct {
}

func (d *Dialog) renderWindow() {
	button := dialog.createButton()
	button.render()
}
func (d *Dialog) createButton() IButton {
	return nil
}

type HtmlDialog struct {
	Dialog
}

func (h *HtmlDialog) createButton() IButton {
	return &HtmlButton{}
}

type WIndowDialog struct {
	Dialog
}

func (h *WIndowDialog) createButton() IButton {
	return &WindowButton{}
}

func configure(t string) {
	if t == "window" {
		dialog = &WIndowDialog{}
	} else {
		dialog = &HtmlDialog{}
	}
}

func runBusinessLogic() {
	dialog.renderWindow()
}
