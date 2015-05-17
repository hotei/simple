// bugBounce.go (c) 2010 David Rook - released with Simplified BSD 2-clause Lincense

/*
   working ok as of 2010-04-20_10:12:00

   Intended as test for several go concepts
       'interface'
       'inheritance' via types which include types

        equiv c++

        class bug {
            float32 x
            int xdir
            float32 y
            int ydir
            float32 r
        public:
            moveto(float32 x, float32 y)
            moverel(float32 dx, float32 dy)
            gohome()
        }

        class public bug : xbug {
            int Color
        }

        In c++ : xbug inherits the ability to do gohome() from being derived of class bug

        In go : xbug we have to declare interface and duplicate the functions needed explicitly
        for xbug

        I guess the issue is that there's no 'organized collection of methods' like a C++ class
        - maybe godoc organizes it "well enough"
*/

package main

import (
	"fmt"
)

type Moveable interface {
	MoveTo(x, y float32)
	MoveRel(dx, dy float32)
}

const (
	MaxX float32 = 100.0
	MaxY float32 = 100.0
)

type bug struct {
	X      float32
	DirX   int
	Y      float32
	DirY   int
	Radius float32
	Speed  float32
}

type xbug struct {
	Bug   bug
	Color int
}

func (b *bug) Setup() {
	b.X = 20
	b.DirX = 1
	b.Y = 20
	b.DirY = -1
	b.Radius = 13
	b.Speed = 5
}

func (b *xbug) Show() {
	fmt.Printf("circle r(%g) x(%g) y(%g) color(%d)\n", b.Bug.Radius, b.Bug.X, b.Bug.Y, b.Color)
}

func (b *bug) Show() {
	fmt.Printf("circle r(%g) x(%g) y(%g) \n", b.Radius, b.X, b.Y)
}

func (b *bug) reverseX() {
	b.DirX = -b.DirX
}

func (b *bug) reverseY() {
	b.DirY = -b.DirY
}

func (b *bug) MoveTo(x, y float32) {
	b.X = x
	b.Y = y
}

func (b *bug) MoveRel(dx, dy float32) {
	b.X += dx
	b.Y += dy
}

func (b *xbug) MoveRel(dx, dy float32) {
	b.Bug.MoveRel(dx, dy)
}

func (b *xbug) MoveTo(x, y float32) {
	b.Bug.MoveTo(x, y)
}

func (b *bug) Step(speed float32) {
	b.MoveRel(float32(b.DirX)*speed, float32(b.DirY)*speed)
	b.confine()
}

func (b *bug) confine() {
	if b.X > MaxX {
		b.X = MaxX
		b.reverseX()
	}
	if b.Y > MaxY {
		b.Y = MaxY
		b.reverseY()
	}
	if b.X < 0 {
		b.X = 0
		b.reverseX()
	}
	if b.Y < 0 {
		b.Y = 0
		b.reverseY()
	}
}

// anything that satisfies the Moveable interface can be told to "GoHome"
func GoHome(t Moveable) {
	t.MoveTo(0.0, 0.0)
}

func main() {
	// create two different kinds of bug
	var mybug = new(bug)
	var myxbug = new(xbug)

	fmt.Printf("<start of bugBounce>\n")
	mybug.Setup()
	mybug.Show()

	fmt.Printf("<starting loop>\n")
	for i := 0; i < 3; i++ {
		mybug.Step(mybug.Speed)
		mybug.Show()
	}

	for i := 0; i < 3; i++ {
		fmt.Println()
	}

	myxbug.Bug.Setup()
	for i := 0; i < 3; i++ {
		myxbug.Bug.Step(myxbug.Bug.Speed)
		myxbug.Show()
	}
	fmt.Printf("\n\n")

	// demonstrate interface driven movement works for both types of bug
	GoHome(myxbug)
	myxbug.Show()

	GoHome(mybug)
	mybug.Show()

	fmt.Printf("<end of bugBounce>\n")
}
