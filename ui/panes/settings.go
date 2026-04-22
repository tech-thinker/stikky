package panes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tech-thinker/stikky/utils"
)

type SettingPane interface {
	Draw() *fyne.Container
}

type settingPane struct {
	pane *fyne.Container

	// widgets
	btnGenerateNewKeyPair *widget.Button
	prgProgress           *widget.ProgressBar
	etPublicKey           *widget.Entry
	etPrivateKey          *widget.Entry
}

func NewSettingPane() SettingPane {
	return &settingPane{
		pane: container.NewStack(),
	}
}

func (p *settingPane) generateKeyPairs() {
	// fyne.DoAndWait(func() {
	// 	tmp := p.btnGenerateNewKeyPair.Text
	// 	p.btnGenerateNewKeyPair.SetText("Generating...")
	// 	p.btnGenerateNewKeyPair.Disable()
	// 	pKey, pubKey, _ := utils.GenerateKeyPair(4096)
	// 	p.etPublicKey.SetText(pubKey)
	// 	p.etPrivateKey.SetText(pKey)
	// 	p.btnGenerateNewKeyPair.SetText(tmp)
	// 	p.btnGenerateNewKeyPair.Enable()
	// })
	tmp := p.btnGenerateNewKeyPair.Text
	p.btnGenerateNewKeyPair.SetText("Generating...")
	p.btnGenerateNewKeyPair.Disable()
	go func(btn *widget.Button, pub, prv *widget.Entry) {
		// fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		// p.prgProgress.SetValue(0)
		pKey, pubKey, _ := utils.GenerateKeyPair(4096)
		pub.SetText(pubKey)
		prv.SetText(pKey)
		btn.SetText(tmp)
		if btn.Disabled() {
			btn.Enable()
		}
		// p.prgProgress.SetValue(100)

		// }, false)

	}(p.btnGenerateNewKeyPair, p.etPublicKey, p.etPrivateKey)
}

func (p *settingPane) Draw() *fyne.Container {
	title := widget.NewRichTextFromMarkdown("# Settings")

	p.btnGenerateNewKeyPair = widget.NewButton("Generate New KeyPair", p.generateKeyPairs)
	p.prgProgress = widget.NewProgressBar()
	lblPublicKey := widget.NewLabel("Public Key")
	p.etPublicKey = widget.NewMultiLineEntry()
	p.etPublicKey.SetPlaceHolder("Type here.")
	p.etPublicKey.SetMinRowsVisible(10)

	lblPrivateKey := widget.NewLabel("Private Key")
	p.etPrivateKey = widget.NewMultiLineEntry()
	p.etPrivateKey.SetPlaceHolder("Type here...")
	p.etPrivateKey.SetMinRowsVisible(10)

	vb := container.NewVBox(
		title,
		container.NewHBox(widget.NewLabel("Generate New Keypair:"), p.btnGenerateNewKeyPair),
		// p.prgProgress,
		lblPublicKey,
		p.etPublicKey,
		lblPrivateKey,
		p.etPrivateKey,
	)
	pv := container.NewPadded(container.NewPadded(vb))
	p.pane.Add(pv)
	return p.pane
}
