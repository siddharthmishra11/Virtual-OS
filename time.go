package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
)
func showTime(w fyne.Window){
		h,m,_ := time.Now().Clock()
		Time := "AM"
		if(h>=12){
		  Time = "PM"
		   h -= 12
		}
		_,month,date := time.Now().Date()
	ti := strconv.Itoa(h)+":"+strconv.Itoa(m)+Time
	dat := strconv.Itoa(date)+" "+fmt.Sprint(month)
	text1 := canvas.NewText(dat, color.Black)
	text1.Alignment = fyne.TextAlignCenter
	text2 := canvas.NewText(ti, color.Black)
	text2.Alignment = fyne.TextAlignCenter
	image := canvas.NewImageFromFile("download.png")
	image.FillMode = canvas.ImageFillOriginal
	c := container.NewVBox(
		image,
		text1,
		text2,
	)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,c))
	w.Resize(fyne.NewSize(300,300))
	w.Show()
}