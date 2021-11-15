package main

import (
	"fmt"
	// "image/color"
	"log"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err!=nil{
		log.Fatal(err)
	}
}
var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("MANU OS")
var wbtn fyne.Widget
var tbtn fyne.Widget
var timbtn fyne.Widget
var mbtn fyne.Widget
var gbtn fyne.Widget
var gabtn fyne.Widget
var calcbtn fyne.Widget
var img fyne.Canvas
var desktBtn fyne.Widget
var waBtn fyne.Widget
var panelContent *fyne.Container
func main(){
	w.Resize(fyne.NewSize(800,720))
	a.Settings().SetTheme(theme.LightTheme())
	img := canvas.NewImageFromFile("nge.jpeg")
	img.FillMode = canvas.ImageFillOriginal
	img.Resize(fyne.NewSize(800,680))
	r, _ := fyne.LoadResourceFromURLString("https://www.cnet.com/a/img/mdR2XJTUO7Fn7wXpBQhT68G0qXA=/940x0/2014/03/11/a677e134-8837-4353-ac51-1bbe9c497994/chrome-logo-large.jpg")
	gbtn = widget.NewButtonWithIcon("",r,func(){
		openbrowser("https://www.google.co.in/")
		})
	r2,_ := fyne.LoadResourceFromURLString("https://i.pinimg.com/originals/06/c4/f7/06c4f70ec5931e2342e703e8a3f0a253.png")
	wbtn = widget.NewButtonWithIcon("", r2,func(){
		showWeather()
		})
	r3, _ := fyne.LoadResourceFromURLString("https://icon-library.com/images/time-tracking-icon/time-tracking-icon-9.jpg")
	timbtn = widget.NewButtonWithIcon("",r3,func(){
		showTime(w)
		})
	r4, _ := fyne.LoadResourceFromURLString("https://www.sublimehq.com/images/sublime_text.png")
	tbtn = widget.NewButtonWithIcon("",r4,func(){
		showText()
		})
	r5, _ := fyne.LoadResourceFromURLString("https://upload.wikimedia.org/wikipedia/commons/thumb/1/19/Spotify_logo_without_text.svg/1024px-Spotify_logo_without_text.svg.png")
	mbtn = widget.NewButtonWithIcon("",r5,func(){
			showMusic()
		})
	r6, _ := fyne.LoadResourceFromURLString("https://cdn.osxdaily.com/wp-content/uploads/2015/04/photos-app-icon-mac.jpg?w=640")
		gabtn = widget.NewButtonWithIcon("",r6,func(){
			showGallery(w)
		})
	r7, _ := fyne.LoadResourceFromURLString("https://lh3.googleusercontent.com/proxy/j2KskndVKxg8HHByvpPBxjtZPTlQafddQtCFVgdfDZCq7kl59OLZlu0_zFEvFHRGIvoeJgzhomejkPPoKnAyoi31Ur7OPXc7L8Whff2Ggvf8lgLZp_waQTd10NY3Eab4")
		calcbtn = widget.NewButtonWithIcon("",r7,func(){
			showCalculator()
		})
	r8, _ := fyne.LoadResourceFromURLString("https://media.idownloadblog.com/wp-content/uploads/2018/07/Apple-logo-black-and-white.png")
	desktBtn = widget.NewButtonWithIcon("",r8,func(){
		w.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
		w.Resize(fyne.NewSize(800,720))
	})
	r9, _ := fyne.LoadResourceFromURLString("https://previews.123rf.com/images/vasiffeyzullazadeh/vasiffeyzullazadeh1904/vasiffeyzullazadeh190400745/122592474-change-icon-vector.jpg")
	waBtn = widget.NewButtonWithIcon("",r9,func(){
		openBtn :=  widget.NewButton("Change Wallpaper", func() {
			openFileDialog:=dialog.NewFileOpen(func(urc fyne.URIReadCloser,_ error){
				img = canvas.NewImageFromURI(urc.URI())
				w.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
			},w)
			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png",".jpeg",".jpg"}))
			openFileDialog.Show()
		})
		img2 := canvas.NewImageFromFile("naruto.jpeg")
		img2.FillMode = canvas.ImageFillOriginal
		img.FillMode = canvas.ImageFillOriginal
		c := container.NewVBox(
			img2,
			openBtn,
		)
		w.SetContent(container.NewBorder(panelContent,nil,nil,nil,container.NewScroll(c)),)
	})
	panelContent  = container.NewVBox(container.NewGridWithColumns(8,desktBtn,gbtn,wbtn,timbtn,tbtn,mbtn,gabtn,calcbtn,waBtn),)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	w.CenterOnScreen()
	w.ShowAndRun()
}
 