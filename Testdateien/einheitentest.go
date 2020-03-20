// Autor: Friederike Richter und Moritz Raake
// Datum: 11.03.20
// Zweck: Softwarepraktikum - Testten der Einheiten Methoden

package main 

import "../einheiten"
import "gfx"
import "fmt"
import "time"

func main () {
	fmt.Println ("Die Zeichen-Methode wird nach Einfügen der Einheiten getestet")
	var einheitauffeld []einheiten.Einheit
	einheitauffeld = make ([]einheiten.Einheit,0)
	var einheit1,einheit2,einheit3,einheit4 einheiten.Einheit 
	einheit1 = einheiten.New (0,5,"Tank",0)
	einheit2 = einheiten.New (15,7,"Tank",1)
	einheit3 = einheiten.New (3,5,"Infanterie",0)
	einheit4 = einheiten.New (10,7,"Infanterie",1)
	einheitauffeld = append (einheitauffeld,einheit1,einheit2,einheit3,einheit4)
	for _,w:=range einheitauffeld {
		w.Zeichnen()
	}
	gfx.Stiftfarbe(0,0,0)
	for i:=uint16(0);i<20;i++ {
		gfx.Linie(i*50,0,i*50,599)
	}
	for i:=uint16(0);i<12;i++ {
		gfx.Linie(0,i*50,999,i*50)
	}
	fmt.Println ("\nDie Angriffs-Methoden der Einheiten werden getestet")
	fmt.Println ("Der Angriffswert der Einheit 1 beträgt: ",einheit1.GibAngriffswert())
	fmt.Println ("Der Angriffswert der Einheit 1 wird auf 2 verändert.")
	einheit1.SetzeAngriffswert(2)
	fmt.Println ("Der Angriffswert der Einheit 1 beträgt nun: ",einheit1.GibAngriffswert())
	
	fmt.Println ("\nDie Verteidigungs-Methoden der Einheiten werden getestet")
	fmt.Println ("Der Verteidigungswert der Einheit 1 beträgt: ",einheit1.GibVerteidigungswert())
	fmt.Println ("Der Verteidigungswert der Einheit 1 wird auf 2 verändert.")
	einheit1.SetzeVerteidigungswert(2)
	fmt.Println ("Der Angriffswert der Einheit 1 beträgt nun: ",einheit1.GibVerteidigungswert())
	
	fmt.Println ("\nDie Angriffsreichweite-Methoden der Einheiten werden getestet")
	fmt.Println ("Die Angriffsreichweite der Einheit 1 beträgt: ",einheit1.GibAngriffsReichweite())
	fmt.Println ("Die Angriffsreichweite der Einheit 1 wird auf 2 verändert.")
	einheit1.SetzeAngriffsReichweite(2)
	fmt.Println ("Der Angriffswert der Einheit 1 beträgt nun: ",einheit1.GibAngriffsReichweite())
	
	fmt.Println ("\nDie Bewegungsreichweite-Methoden der Einheiten werden getestet")
	fmt.Println ("Die Bewegungsreichweite der Einheit 1 beträgt: ",einheit1.GibBewegungsReichweite())
	fmt.Println ("Die Bewegungsreichweite der Einheit 1 wird auf 2 verändert.")
	einheit1.SetzeBewegungsReichweite(2)
	fmt.Println ("Der Bewegungswert der Einheit 1 beträgt nun: ",einheit1.GibBewegungsReichweite())

	fmt.Println ("\nDie Typ-Methoden der Einheiten werden getestet")
	fmt.Println ("Der Typ der Einheit 1 ist: ",einheit1.GibTyp())
	fmt.Println ("Der Typ der Einheit 1 wird auf 'Infanterie' verändert.")
	einheit1.SetzeTyp("Infanterie")
	fmt.Println ("Der Typ der Einheit 1 ist nun: ",einheit1.GibTyp())	
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	fmt.Println("ich zeichne neu")
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	for _,w:=range einheitauffeld {
		w.Zeichnen()
	}
	gfx.Stiftfarbe(0,0,0)
	for i:=uint16(0);i<20;i++ {
		gfx.Linie(i*50,0,i*50,599)
	}
	for i:=uint16(0);i<12;i++ {
		gfx.Linie(0,i*50,999,i*50)
	}	
	gfx.TastaturLesen1()
	
	
	fmt.Println ("\nDie Orts-Methoden der Einheiten werden getestet")
	x,y:=einheit1.GibKoordinaten()
	fmt.Println ("Der Ort der Einheit 1 ist: ",x,y)
	fmt.Println ("Der Ort der Einheit 1 wird auf 7,7 verändert.")
	einheit1.SetzeKoordinaten(7,7)
	x,y=einheit1.GibKoordinaten()
	fmt.Println ("Der Ort der Einheit 1 ist nun: ",x,y)	
	
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	fmt.Println("ich zeichne neu")
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	for _,w:=range einheitauffeld {
		w.Zeichnen()
	}
	gfx.Stiftfarbe(0,0,0)
	for i:=uint16(0);i<20;i++ {
		gfx.Linie(i*50,0,i*50,599)
	}
	for i:=uint16(0);i<12;i++ {
		gfx.Linie(0,i*50,999,i*50)
	}	
	gfx.TastaturLesen1()
	

	
	fmt.Println ("\n Die Spielerwechsel-Methoden der Einheiten wird getestet")
	fmt.Println ("Der Spieler der Einheit 1 ist: ", einheit1.GibSpieler())
	fmt.Println ("Der Spieler der Einheit 1 wird auf 2 verändert.")
	einheit1.SetzeSpieler(2)
	fmt.Println ("Der Spieler der Einheit 2 ist nun: ",einheit1.GibSpieler())	
	
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	fmt.Println("ich zeichne neu")
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	for _,w:=range einheitauffeld {
		w.Zeichnen()
	}
	gfx.Stiftfarbe(0,0,0)
	for i:=uint16(0);i<20;i++ {
		gfx.Linie(i*50,0,i*50,599)
	}
	for i:=uint16(0);i<12;i++ {
		gfx.Linie(0,i*50,999,i*50)
	}	
	gfx.TastaturLesen1()
	
	
}
