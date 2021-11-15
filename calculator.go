package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func showCalculator() {
	w := a.NewWindow("Calculator")
	output:= ""
	input:= widget.NewLabel(output)
	var hist []string
	histString := ""
	isHist := false
	history := widget.NewLabel(histString)
	historyBtn:= widget.NewButton("History",func(){
		if isHist{
			histString = ""
		}else{
			for i:=len(hist)-1;i>=0;i--{
				histString += hist[i]
				histString += "\n"
			}
		}
		isHist = !isHist
		history.SetText(histString)
	})
	backBtn:= widget.NewButton("Back",func(){
		sz:=len(output)
		if sz>0{
			output = output[:sz-1]
		}
		input.SetText(output)
	})
	clearBtn:= widget.NewButton("Clear",func(){
		output = "";
		input.SetText(output);
	})
	closeBtn:= widget.NewButton(")",func(){
		output += ")";
		input.SetText(output);
	})
	OpenBtn:= widget.NewButton("(",func(){
		output += "(";
		input.SetText(output);
	})
	divideBtn:= widget.NewButton("/",func() {
		output += "/";
		input.SetText(output);
	})
	subtractBtn:= widget.NewButton("-",func(){
        output += "-";
		input.SetText(output);
	})
	plusBtn:= widget.NewButton("+",func(){
		output += "+";
		input.SetText(output);
	})
	multiplyBtn:= widget.NewButton("*",func(){
		output += "*";
		input.SetText(output);
	})
	dotBtn:= widget.NewButton(".",func(){
		output += ".";
		input.SetText(output);
	})
	equalBtn:= widget.NewButton("=",func(){
		functions := map[string]govaluate.ExpressionFunction {
			"strlen": func(args ...interface{}) (interface{}, error) {
				length := len(args[0].(string))
				return (float64)(length), nil
			},
		}
		expression, err := govaluate.NewEvaluableExpressionWithFunctions(output, functions)
		if(err==nil){
			result, err := expression.Evaluate(nil)
			if(err==nil){
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strApp := output+" = "+ans
				hist = append(hist,strApp)
				output = ans
			}else{
				output = "error"
			}
		}else{
		   output = "error"
		}
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0",func(){
		output += "0";
		input.SetText(output);
	})
	oneBtn := widget.NewButton("1",func(){
		output += "1";
		input.SetText(output);
	})
	twoBtn := widget.NewButton("2",func(){
		output += "2";
		input.SetText(output);
	})
	threeBtn := widget.NewButton("3",func(){
        output += "3";
		input.SetText(output);
	})
	fourBtn := widget.NewButton("4",func(){
        output += "4";
		input.SetText(output);
	})
	fiveBtn := widget.NewButton("5",func(){
		output += "5";
		input.SetText(output);
	})
	sixBtn := widget.NewButton("6",func(){
		output += "6";
		input.SetText(output);
	})
	sevenBtn := widget.NewButton("7",func(){
        output += "7";
		input.SetText(output);
	})
	eightBtn := widget.NewButton("8",func(){
        output += "8";
		input.SetText(output);
	})
	nineBtn := widget.NewButton("9",func(){
        output += "9";
		input.SetText(output);
	})
	c := container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					historyBtn,
					backBtn,
				),
				container.NewGridWithColumns(4,
					clearBtn,
					OpenBtn,
					closeBtn,
					divideBtn,
				),
				container.NewGridWithColumns(4,
					sevenBtn,
					eightBtn,
					nineBtn,
					multiplyBtn,
				),
				container.NewGridWithColumns(4,
					fourBtn,
					fiveBtn,
					sixBtn,
					subtractBtn,
				),
				container.NewGridWithColumns(4,
					oneBtn,
					twoBtn,
					threeBtn,
					plusBtn,
				),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
					),
					equalBtn,
				),
			),	 
		)
	w.SetContent(container.NewBorder(desktBtn,nil,nil,nil,c),)
	w.Show()
}