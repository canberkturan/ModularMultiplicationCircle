package main

import (
	"./mmc"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"
)

// ColorValue has an attribute that shows if rgba values are default or not
type ColorValue struct {
	rgba  color.RGBA
	isSet bool
}

func (v *ColorValue) String() string {
	return string(v.rgba.R) + "," + string(v.rgba.G) + "," + string(v.rgba.B)
}

// Set rgb or rgba values with string value
func (v *ColorValue) Set(s string) error {
	rgb := []uint8{ 0, 0, 0, 255 }
	parsed := strings.Split(s, " ")
	if len(parsed) != 3 && len(parsed) != 4 {
		fmt.Println("Please enter 3 or 4 values for color paramethers")
		os.Exit(2)
	} else {
		for i := 0; i < len(parsed); i++ {
			val, err := strconv.Atoi(parsed[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			} else {
				rgb[i] = uint8(val)
			}
		}
		v.rgba = color.RGBA{rgb[0], rgb[1], rgb[2], rgb[3}
		v.isSet = true
	}
	return nil
}
func main() {
	var mmc mmc.MMC
	var fg ColorValue
	var bg ColorValue

	size := flag.Int("s", 1080, "Size of picture")
	pad := flag.Int("p", 0, "Padding between picture borders and circle (default 0)")
	flag.Var(&fg, "fg", "Foreground Color (default 255 255 255)(ex: -fg \"255 0 0\")")
	flag.Var(&bg, "bg", "Background Color (default 0 0 0)(ex: -bg \"0 0 0\")")
	rot := flag.Float64("r", 0, "Rotation of circle (default 0)")
	dot := flag.Float64("d", 200, "Dot count")
	mult := flag.Float64("m", 2, "Multiplier")
	bias := flag.Float64("b", 0, "Bias (default 0)")
	out := flag.String("o", "output.png", "Output file path")
	flag.Parse()

	if !fg.isSet {
		fg.rgba = color.RGBA{255, 255, 255, 255}
	}
	if !bg.isSet {
		bg.rgba = color.RGBA{0, 0, 0, 255}
	}

	returnCode := mmc.Init(*size, *pad, *rot, bg.rgba, fg.rgba)
	if returnCode == -1 {
		os.Exit(2)
	}
	mmc.Create(*dot, *mult, *bias)
	mmc.Save(*out)
}
