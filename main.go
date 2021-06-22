package main

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
)

func main () {
	fmt.Println ("Hello world!")
	d, errX := docx.ReadDocxFile ("./f1.docx")
	if errX != nil {
		fmt.Println ("Error occured:", errX)
		return
	}
	
	dY := d.Editable ()
	fmt.Println ("Content:", dY.GetContent ())
	dY.SetContent ("Hello")
	dY.WriteToFile ("./f2.docx")
}
