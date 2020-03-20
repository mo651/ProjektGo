package main

import "../button"
import "gfx"
//Testdatei it noch nicht vollständig!
func main(){
	a:=button.New(100,100,500,300,0,0,0,255,255,255,0,0,0,"Hallo du Herzchen ",'m','m')
	a.Zeichnen()
	gfx.TastaturLesen1()
	a.Entfernen(255,255,255)
	gfx.TastaturLesen1()
	gfx.TastaturLesen1()
	}
