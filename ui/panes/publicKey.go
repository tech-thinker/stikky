package panes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tech-thinker/stikky/internal/db"
)

type publicKeyPane struct {
	pane         *fyne.Container
	etTitle      *widget.Entry
	etPublicKeys *widget.Entry
	btnAdd       *widget.Button
	keyBox       *fyne.Container
	keys         []db.PublicKey
}

func NewPublicKeyPane() *fyne.Container {
	var p publicKeyPane
	// Initialize widgets
	p.pane = container.NewStack()
	p.etTitle = widget.NewEntry()
	p.etPublicKeys = widget.NewMultiLineEntry()
	p.btnAdd = widget.NewButton("Add", p.btnAdd_OnClick)
	p.keyBox = container.NewVBox()
	p.keys = []db.PublicKey{}

	// Layout
	vb := container.NewVBox()
	lbl := widget.NewLabel("Public Keys")
	fixed := container.NewGridWrap(fyne.NewSize(200, 40), p.etTitle)
	vb.Add(container.NewHBox(lbl, fixed, p.btnAdd))

	// p.etTitle.Resize(fyne.NewSize(200, p.etTitle.Size().Height))
	vb.Add(p.etPublicKeys)
	vb.Add(p.keyBox)
	p.RefreshList()
	p.pane.Add(vb)

	return p.pane
}

func (p *publicKeyPane) btnAdd_OnClick() {
	key := db.PublicKey{
		Name:      p.etTitle.Text,
		PublicKey: p.etPublicKeys.Text,
	}
	// p.keys = append(p.keys, key)
	p.keys = append([]db.PublicKey{key}, p.keys...)
	p.etTitle.SetText("")
	p.etPublicKeys.SetText("")
	p.RefreshList()
}

func (p *publicKeyPane) btnDelete_Click(index int) {
	p.keys = append(p.keys[:index], p.keys[index+1:]...)
	p.RefreshList()
}

func (p *publicKeyPane) RefreshList() {
	p.keyBox.RemoveAll()
	for i, key := range p.keys {
		p.keyBox.Add(widget.NewCard(key.Name, "", container.NewVBox(widget.NewLabel(key.PublicKey), widget.NewButton("Delete", func() { p.btnDelete_Click(i) }))))
	}
}
