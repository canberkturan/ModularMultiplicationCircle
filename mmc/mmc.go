package mmc

import (
	"os"
	"fmt"
	"math"
	"image"
	"image/png"
	"image/color"
)

type MMC struct {
	img *image.RGBA
	size int
	fg color.RGBA
	bg color.RGBA
	pad int
	rot float64
}

func (mmc *MMC) Init(size, pad int, rot float64, bg, fg color.RGBA) (int){
        if (size<=0){
                fmt.Printf("Size value must be greater than zero.(%d given)\n",size)
                return -1
        }
        if (pad < 0 || pad > size/2){
                fmt.Printf("Padding value must be between 0 and half of the size.(%d given,size=%d) (Boundary values are valid too)\n", pad, size)
                return -1
        }
        mmc.size=size
        mmc.pad=pad
        mmc.rot=rot*(math.Pi/180)
        mmc.bg=bg
        mmc.fg=fg
        mmc.img=image.NewRGBA(image.Rect(0,0,size,size))
        for i:=0; i<size; i++ {
                for j:=0; j<size; j++ {
                        mmc.img.Set(i,j,bg)
                }
        }
        mmc.DrawCircle()
        return 1
}

func (mmc *MMC) Create(dotcount, multiplier, bias float64){
	center:=float64(mmc.size)/2
	radius:=center-float64(mmc.pad)
	angle:=(2*math.Pi)/dotcount
	for i:=float64(0); i<dotcount; i+=1 {
		to:=math.Mod(i*multiplier+bias,dotcount)
		x0:=int(center+radius*math.Cos(angle*i+mmc.rot))
		y0:=int(center+radius*math.Sin(angle*i+mmc.rot))
		x1:=int(center+radius*math.Cos(angle*to+mmc.rot))
		y1:=int(center+radius*math.Sin(angle*to+mmc.rot))
		mmc.Drawline(x0,y0,x1,y1)
	}
}

func (mmc *MMC) DrawCircle(){
	da:=1/float64(mmc.size)
	center:=float64(mmc.size)/2
	radius:=center-float64(mmc.pad)
	for a:=float64(0); a<2*math.Pi; a+=da {
		x:=int(center+radius*math.Cos(a))
		y:=int(center+radius*math.Sin(a))
		mmc.img.Set(x,y,mmc.fg)
	}
}
func (mmc *MMC) Drawline(x0, y0, x1, y1 int){
	dx := math.Abs(float64(x1-x0))
	dy := -math.Abs(float64(y1-y0))
	sx:=map[bool]int{true: 1, false: -1}[x0<x1]
	sy:=map[bool]int{true: 1, false: -1}[y0<y1]
	err := dx + dy
	for {
		mmc.img.Set(x0,y0,mmc.fg)
		if (x0==x1 && y0==y1) { break }
		e2 := 2*err
		if (e2 >= dy){
			err += dy
			x0 += sx
		}
		if (e2 <= dx){
			err +=dx
			y0 += sy
		}
	}
}
func (mmc *MMC) Save(path string){
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	png.Encode(file, mmc.img)
	file.Close()
}

