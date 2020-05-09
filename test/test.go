package main

import (
	"../mmc"
	"image/color"
)

func main(){
	var mmc1 mmc.MMC
	var mmc2 mmc.MMC
	var mmc3 mmc.MMC
	var mmc4 mmc.MMC
	var mmc5 mmc.MMC
	var mmc6 mmc.MMC
	var mmc7 mmc.MMC
	var mmc8 mmc.MMC
	var mmc9 mmc.MMC
	var mmc10 mmc.MMC
	var mmc11 mmc.MMC
	var mmc12 mmc.MMC
	var mmc13 mmc.MMC
	var mmc14 mmc.MMC

	fg:=color.RGBA{255,255,255,255}
	bg:=color.RGBA{0,0,0,255}

	//Black-Box Boundary-Value Analysis and Equivalence Partitioning
	//size
	e1:=mmc1.Init(-10,0,0,bg,fg)   //size invalid
	e2:=mmc2.Init(0,0,0,bg,fg)     //size invalid
	e3:=mmc3.Init(1,0,0,bg,fg)     //size valid
	e4:=mmc4.Init(10,0,0,bg,fg)    //size valid
	//padding
	e5:=mmc5.Init(10,-10,0,bg,fg)  //padding invalid
	e6:=mmc6.Init(10,-1,0,bg,fg)   //padding invalid
	e7:=mmc7.Init(10,0,0,bg,fg)    //padding valid
	e8:=mmc8.Init(10,3,0,bg,fg)    //padding valid
	e9:=mmc9.Init(10,5,0,bg,fg)    //padding valid
	e10:=mmc10.Init(10,6,0,bg,fg)  //padding invalid
	e11:=mmc11.Init(10,10,0,bg,fg) //padding invalid

	//White-Box
	e12:=mmc12.Init(-10,0,0,bg,fg) //size invalid, padding is not important
	e13:=mmc13.Init(10,-5,0,bg,fg) //size valid, padding invalid
	e14:=mmc14.Init(10,2,0,bg,fg)  //size valid, padding valid

	if(e1 == 1){ mmc1.Create(200,2,0); mmc1.Save("test1_output.png"); }
	if(e2 == 1){ mmc2.Create(200,2,0); mmc2.Save("test2_output.png"); }
	if(e3 == 1){ mmc3.Create(200,2,0); mmc3.Save("test3_output.png"); }
	if(e4 == 1){ mmc4.Create(200,2,0); mmc4.Save("test4_output.png"); }
	if(e5 == 1){ mmc5.Create(200,2,0); mmc5.Save("test5_output.png"); }
	if(e6 == 1){ mmc6.Create(200,2,0); mmc6.Save("test6_output.png"); }
	if(e7 == 1){ mmc7.Create(200,2,0); mmc7.Save("test7_output.png"); }
	if(e8 == 1){ mmc8.Create(200,2,0); mmc8.Save("test8_output.png"); }
	if(e9 == 1){ mmc9.Create(200,2,0); mmc9.Save("test9_output.png"); }
	if(e10 == 1){ mmc10.Create(200,2,0); mmc10.Save("test10_output.png"); }
	if(e11 == 1){ mmc11.Create(200,2,0); mmc11.Save("test11_output.png"); }
	if(e12 == 1){ mmc12.Create(200,2,0); mmc12.Save("test12_output.png"); }
	if(e13 == 1){ mmc13.Create(200,2,0); mmc13.Save("test13_output.png"); }
	if(e14 == 1){ mmc14.Create(200,2,0); mmc14.Save("test14_output.png"); }
}
