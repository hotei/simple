// toml_test.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	// below is go std lib 1.X only
	"fmt"
	"log"
	"testing"

	// external
	"github.com/BurntSushi/toml"
)

var (
	md  toml.MetaData
	err error
)

func Test_0001(t *testing.T) {
	fmt.Printf("Test_0001\n")
	type Config_02 struct {
		A []int
		B []int
	}

	var tomlBlob = `
A = [ 100, 200 ]
B = [ 300, 400 ]
`
	var myConfig Config_02
	md, err := toml.Decode(tomlBlob, &myConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", md)
	fmt.Printf("myConfig.A = %v\n", myConfig.A)
	for _, val := range myConfig.A {
		fmt.Printf("A(%d)\n", val)
	}
	fmt.Printf("Test_0001 PASS\n\n")
}

func Test_0002(t *testing.T) {
	fmt.Printf("Test_0002\n")
	var tomlBlob = `
title = "a mouse"
sx = 0.001

[mm]		# mouse data
    clip = [ 20, 30 ]
    scale = 10.000001
`
	type mm struct {
		Clip  []int
		Scale float64
	}

	type tomlConfig struct {
		Mouse mm `toml:"mm"` // adding this made it work - but what the hell is it doing???
		Title string
		Sx    float64
	}

	var config tomlConfig
	md, err := toml.Decode(tomlBlob, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", md)
	fmt.Printf("config.Mouse.Clip = %v\n", config.Mouse.Clip)
	fmt.Printf("config.Mouse.Scale = %v\n", config.Mouse.Scale)
	fmt.Printf("config.Title = %v\n", config.Title)
	fmt.Printf("config.Sx = %v\n", config.Sx)
	fmt.Printf("Test_0002 PASS\n\n")
}

/*
mm1.rconf:
	sourcefile = "mm1.heatmap.something"
	ringlevel = 700
	ringwidth = 10
	colorfile = "more.cmap"
	cliphilo = [0, 1e+9]  # clip first then scale
	scale = 1e-6	 # -1 for auto; 0 to skip; 1.0 is default
	hiclipcolor = "Red"	# do this if clipped on high end
	loclipcolor = "Yellow" # do this if clipped on low end

*/

// load the rings.conf file
func Test_0003(t *testing.T) {
	fmt.Printf("Test_0003\n")
	var tomlBlob = `
# rings.conf
	numberFont = "TrueMono.ttf"
	fontDir = "Desktop/MYGO/src/fonts"
	textFont = "barbie.ttf"
	imageXY = [ 3000, 3000 ]  # note - [3000,3000] failed
	pngName = "file01.png"
	renderList = [
		"mm1",
		"mm2",
		"mm3",
		]
	colorfile = "colormap.cmap"
`
	type ringConfig struct {
		NumberFont string
		FontDir    string
		TextFont   string
		ImageXY    []int
		PngName    string
		RenderList []string
		ColorFile  string
	}

	var rc ringConfig

	md, err := toml.Decode(tomlBlob, &rc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("md %v\n", md)
	fmt.Printf("------------------\n")
	fmt.Printf("FontDir(%s)\n", rc.FontDir)
	fmt.Printf("NumberFont(%s) \n", rc.NumberFont)
	fmt.Printf("TextFont(%s) \n", rc.TextFont)
	fmt.Printf("ImageXY %v\n", rc.ImageXY)
	fmt.Printf("PngName %s\n", rc.PngName)
	fmt.Printf("RenderList %v\n", rc.RenderList)
	fmt.Printf("ColorFile %v\n", rc.ColorFile)
}
