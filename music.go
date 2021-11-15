package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)
func showMusic(){
    w := a.NewWindow("Music Player")
    var streamer beep.StreamSeekCloser
    var format beep.Format
    var pause bool = false
    w.Resize(fyne.NewSize(400,400))
    logo := canvas.NewImageFromFile("music.png")
    logo.FillMode = canvas.ImageFillOriginal
    toolbar := widget.NewToolbar(
        widget.NewToolbarSpacer(),
        widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
            speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
            speaker.Play(streamer)
        }),
        widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
            if !pause {
                pause = true
                speaker.Lock()
            } else if pause {
                pause = false
                speaker.Unlock()
            }
        }),
        widget.NewToolbarAction(theme.MediaStopIcon(), func() {
            speaker.Clear()
        }),
        widget.NewToolbarSpacer(),
    )
    label2 := widget.NewLabel("Play MP3")
    label2.Alignment = fyne.TextAlignCenter
    browse_files := widget.NewButton("Open Mp3", func() {
        fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
            streamer, format, _ = mp3.Decode(uc)
            label2.Text = uc.URI().Name()
            label2.Refresh()
        }, w)
        fd.Show()
        fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    })
    c :=  container.NewVBox(
            logo,
            browse_files, 
            label2, 
            toolbar,
        )
    w.SetContent(container.NewBorder(desktBtn,nil,nil,nil,c))
    w.Show()
}