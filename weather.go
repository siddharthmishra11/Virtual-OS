package main

import (
	"fmt"
	"io/ioutil"
	"fyne.io/fyne/v2"
	"encoding/json"
	"net/http"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)
func showWeather(){
	w := a.NewWindow("Weather")
	w.Resize(fyne.NewSize(2400,800))
	input := widget.NewEntry()
	label := canvas.NewImageFromFile("weather.png")
	label.FillMode = canvas.ImageFillOriginal
	input.SetPlaceHolder("Enter The City Name Here")
	content := container.NewVBox(
		label,
		input,
		widget.NewButton("Submit", func() {
		city:= input.Text
		str:="https://api.openweathermap.org/data/2.5/weather?q="+city+"&appid=073968190fdcb156b6a450af652f438b";
		res,err := http.Get(str)
	    if err!=nil{
			fmt.Println("Heck")
		    fmt.Println(err)
	    }
	    defer res.Body.Close()
	    body,err := ioutil.ReadAll(res.Body)
	    if err!=nil{
		     fmt.Println(err)
	    }
	    weather,err := UnmarshalWelcome(body)
	    if err!=nil{
			fmt.Println(err)
		}
		label1 := canvas.NewImageFromFile("y.png")
		label1.FillMode = canvas.ImageFillOriginal
		label2 := canvas.NewText("Weather Details",color.Black)
		label2.TextStyle = fyne.TextStyle{Bold:true}
		label2.Alignment = fyne.TextAlignCenter
		label3 := canvas.NewText(fmt.Sprintf("Country %s",weather.Sys.Country),color.Black)
	    label3.Alignment = fyne.TextAlignCenter
		label4 := canvas.NewText(fmt.Sprintf("Wind Speed %f",weather.Wind.Speed),color.Black)
		label4.Alignment = fyne.TextAlignCenter
		label5 := canvas.NewText(fmt.Sprintf("Temperature %f",weather.Main.Temp-273.15),color.Black)
		label5.Alignment = fyne.TextAlignCenter
		label6 := canvas.NewText(fmt.Sprintf("Humidity %d",weather.Main.Humidity),color.Black)
		label6.Alignment = fyne.TextAlignCenter
		label7 := canvas.NewText(fmt.Sprintf("Pressure %d",weather.Main.Pressure),color.Black)
		label7.Alignment = fyne.TextAlignCenter
		 
		c := container.NewVBox(
					label1,
					label2,
					label3,
					label4,
					label5,
					label6,
					label7,
				)
		w.SetContent(container.NewBorder(desktBtn,nil,nil,nil,container.NewScroll(c)),)
		w.Show()
	}),)
	w.SetContent(container.NewBorder(desktBtn,nil,nil,nil,container.NewScroll(content)),)
	w.Show()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}

