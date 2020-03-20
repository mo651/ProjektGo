package main 

import "../gelaende"
import "gfx"
import "fmt"
import "../einheiten"
import "time"

func main () {
	fmt.Println ("Die Zeichen-Methode wird nach Einfügen der Geländefelder getestet")
	var spielfeld []gelaende.Gelände
	spielfeld = make ([]gelaende.Gelände,0)
	for i:=uint16(0);i<20;i++ {
		for j:=uint16(0);j<12;j++ {
			var neuesgelände gelaende.Gelände
			neuesgelände = gelaende.New (i,j,"Wasser")
			spielfeld=append(spielfeld,neuesgelände)
			if j<2 && i==0 {
				fmt.Println ("\nDie Methoden von Karte ",i,"/",j,"werden getestet")
				fmt.Println ("\nTest der Bonus-Methoden")
				fmt.Println ("Der Bonuswert ist: ",neuesgelände.GibBonusAngriff())
				fmt.Println ("Der Bonuswert wird auf 5 gesetzt")
				neuesgelände.SetzeBonusAngriff(5)
				fmt.Println ("Der Bonuswert ist nun: ",neuesgelände.GibBonusAngriff())
				fmt.Println ("\nTest der Belegungs-Methoden")
				fmt.Println ("Die Belegung ist: ",neuesgelände.GibBelegung())
				fmt.Println ("Die Belegung wird auf true gesetzt")
				neuesgelände.SetzeBelegung(true)
				fmt.Println ("Die Belegung ist nun: ",neuesgelände.GibBelegung())
				fmt.Println ("\nTest der Gewählt-Methoden")
				fmt.Println ("Ist das Gelände gewählt?: ",neuesgelände.GibGewaehlt())
				fmt.Println ("Die Geländewahl wird auf true gesetzt")
				neuesgelände.SetzeGewaehlt(true)
				fmt.Println ("Ist das Gelände nun gewählt?: ",neuesgelände.GibGewaehlt())
			}
		}	
	}
	for _,w:=range spielfeld {
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
	fmt.Println ("Die Zeichen-Methode wird nach Einfügen der Einheiten getestet")
	var einheitauffeld []einheiten.Einheit
	einheitauffeld = make ([]einheiten.Einheit,0)
	var einheit1,einheit2,einheit3,einheit4 einheiten.Einheit 
	einheit1 = einheiten.New (0,5,"Tank",0)
	einheit2 = einheiten.New (15,7,"Tank",1)
	einheit3 = einheiten.New (3,5,"Infanterie",0)
	einheit4 = einheiten.New (19,11,"Infanterie",1)
	
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
	
	fmt.Println ("\nDie Distanz-Methoden der Einheiten werden getestet")
	
	x,y:=einheit1.Distanz(einheit2)
	fmt.Println("Die Distanz von Einheit 1 und 2 beträgt", x,y)
	
	
	
	
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	fmt.Println("ich zeichne neu")
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

