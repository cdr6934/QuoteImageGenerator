package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"image/color"
	"os"
	"path/filepath"
)


func main() {

	if err := run("dd_1_resample.jpg", "output/file_out.png"); err != nil {
		fmt.Fprintf(os.Stderr, "#s\n",err)
	os.Exit(1)
	}
	//example_zero()
	//example_one()
}

func run(filein string, fileout string ) error {

	dc := gg.NewContext(1200,628)
	backgroundImage, err := gg.LoadImage(filein)
	if err != nil {
		return errors.Wrap(err, "load background image")
	}
	dc.DrawImage(backgroundImage, 0 , 0)

	// full color border effect
	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width())- (2.0 * margin)
	h := float64(dc.Height())- (2.0 * margin)
 	dc.SetColor(color.RGBA{0,0,0,204})
	dc.DrawRectangle(x,y ,w,h)
	dc.Fill()

	// Add text
	fontPath := filepath.Join("fonts", "OpenSans-Bold.ttf")
	s := "PACE."
	marginX := 50.0
	marginY := -10.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x,y)

	// Add domain name
	textColor := color.White
	fontPath = filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 60); err != nil {
		return errors.Wrap(err, "lload Open_Sans")
	}

	r,g,b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	marginY = 30
	s = "https://www.cbayes.com"
	_, textHeight = dc.MeasureString(s)
	x = 70
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

	// Att title
	title := "Where she goes we will not know"
	textShadowColor := color.Black
	fontPath = filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 90); err !=  nil {
		return errors.Wrap(err, "load Playfair_Display")
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x = textRightMargin
	y = textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth , 1.5, gg.AlignLeft)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(title, x,y,0,0,maxWidth, 1.5, gg.AlignLeft)
	if err := dc.SavePNG(fileout); err != nil {
		return errors.Wrap(err, "save png")
	}

	return nil
}

