package main 

import "../gelaende"
import "gfx"
import "fmt"

func main () {
	fmt.Println ("Die Zeichen-Methode wird nach Einfügen der Geländefelder getestet")
	var spielfeld []gelaende.Gelände
	spielfeld = make ([]gelaende.Gelände,0)
	for i:=uint16(0);i<20;i++ {
		for j:=uint16(0);j<12;j++ {
			var neuesgelände gelaende.Gelände
			x:="Wasser"

			
			neuesgelände = gelaende.New (i,j,x)
			spielfeld=append(spielfeld,neuesgelände)
			if j<2 && i==4 {
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
				//fmt.Println (neuesgelände.Kodieren())
			}
		}	
	}
	for i,_:=range spielfeld{ //Ein Wald entsteht
		if (i==0)||(i==1)||(i==2)||(i==3)||(i==12)||(i==13)||(i==14)||(i==24)||(i==25)||(i==36)||(i==37)||(i==48){
		spielfeld[i].SetzeTyp("Wueste")
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
}
