package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)
func showText(){
	w:= a.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(600,600))
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Ninja Text Editor"),
		),
	)
	i := 1
	content.Add(widget.NewButton("Add More Files", func() {
		content.Add(widget.NewLabel("New File "+strconv.Itoa(i)))
		i++
	}))
	input := widget.NewMultiLineEntry()
	saveBtn :=  widget.NewButton("Save Text File", func() {
	saveFileDialog:=dialog.NewFileSave(func(uwc fyne.URIWriteCloser,_ error){
		textData := []byte(input.Text)
		uwc.Write(textData)
	}, w)
	saveFileDialog.SetFileName("New File "+strconv.Itoa(i-1))
	saveFileDialog.Show()
	})
	openBtn :=  widget.NewButton("Open Text File", func() {
		openFileDialog:=dialog.NewFileOpen(func(urc fyne.URIReadCloser,_ error){
			readData,_ := ioutil.ReadAll(urc)
			viewData := widget.NewMultiLineEntry()
			viewData.SetText(string(readData))
			nw := fyne.CurrentApp().NewWindow("New File")
			saveBtn :=  widget.NewButton("Save Text File", func() {
				saveFileDialog:=dialog.NewFileSave(func(uwc fyne.URIWriteCloser,_ error){
					textData := []byte(input.Text)
					uwc.Write(textData)
			 }, nw)
			 saveFileDialog.Show()
			})
			nw.SetContent(container.NewScroll(
				container.NewVBox(
					viewData,
				    saveBtn,
				),
				))
			nw.Resize(fyne.NewSize(400,400))
			nw.Show()
		}, w)
		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
		})
	c := container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		 )
	w.SetContent(container.NewBorder(desktBtn,nil,nil,nil,c))
	w.Show()
}