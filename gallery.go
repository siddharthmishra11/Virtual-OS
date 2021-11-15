package main

import (
	// "image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"

	// "fmt"
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2/canvas"
	"strconv"
)

func showGallery(w fyne.Window) {
	w.Resize(fyne.NewSize(2400,800))
	rootsrc:= "/Users/siddharthmishra/Desktop/Root File/NARUTO"
	files, err := ioutil.ReadDir(rootsrc)
    if err != nil {
        log.Fatal(err)
	}
	tabs := container.NewAppTabs()
	idx:=1
    for _, f := range files {
		if !f.IsDir(){
			extension:= strings.Split(f.Name(),".")[1]
			if extension=="png"||extension=="jpeg"||extension=="jpg"{
				fname:=rootsrc+"/"+f.Name()
				img :=  canvas.NewImageFromFile(fname)
				tabs.Append(container.NewTabItem("Image "+strconv.Itoa(idx),img))
				idx++
			}

		}
	}
	// tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tabs),)
	w.Show()
}