// Autor: Friederike Richter und Moritz Raake
// Datum: 11.03.20
// Zweck: Softwarepraktikum - Testdatei zu den Spielfeldmethoden

package main 

import ("../spielfeld";"fmt";"../gelaende";"gfx";"time";"../einheiten")

func main () {
	// Spielfeld wird aufgebaut und mit Wasser initialisiert
	fmt.Println ("Ein neues Spielfeld mit Wasserfeldern wird erstellt")
	var dummyspielfeld spielfeld.Spielfeld 
	dummyspielfeld = spielfeld.New (20,12,50,"Dummyfeld")			// XXXXX muss verändern werden, wenn Variabilität verloren geht
	fmt.Println ("Das Spielfeld wurde deklariert")
	for y:=uint16(0);y<12;y++ {
		for x:=uint16(0);x<20;x++ {
			var neuesgelände gelaende.Gelände
			//fmt.Println ("Ein Geländefeld wurde deklariert")
			neuesgelände = gelaende.New (x,y,"Wald") //XXXX man beachte, dass Wasser ist "animiert", ob das im fertigen Spiel auch funktionier kann ich nicht sagen
			//fmt.Println ("Das Geländefeld wurde als Wasser initialisiert")
			dummyspielfeld.SetzeGelände(neuesgelände)
			//fmt.Println ("Ein Geländefeld wurde in das Spielfeld an der Stelle x:",x," y:",y,"eingefügt")
		}
	
	}
	var neueeinheit einheiten.Einheit
	neueeinheit = einheiten.New(2,2,"Tank",1)
	dummyspielfeld.SetzeEinheit (neueeinheit)
  fmt.Println(dummyspielfeld.GibReichweite(neueeinheit)) 
	
	
	fmt.Println ("Das Spielfeld wurde komplett mit Wasser initialisiert")
	fmt.Println ("Das Spielfeld wird gezeichnet")
	dummyspielfeld.Zeichnen()
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	// Test der Geländemethoden
	fmt.Println ("\nDie GeländeMethoden des Spielfelds werden getestet:")
	var x0y0 gelaende.Gelände = gelaende.New(0,0,"Wald")
	// GibGelände wird getestet
	x0y0 = dummyspielfeld.GibGelände(0,0)
	// Der Typ des ursprünglichen Waldfeldes muss nun Wasser sein.
	fmt.Println ("Das Gelände an der Koordinate x:0 y:0 ist vom Typ (erwartet Wasser):",x0y0.GibTyp())
	// Setze Gelände wird getestet
	x0y0.SetzeTyp("Wueste")
	dummyspielfeld.SetzeGelände(x0y0)
	fmt.Println ("Das Wüstenfeld wird an der Koordinate x:0 y:0 eingefügt.")
	var x0y0test gelaende.Gelände = gelaende.New(0,0,"Wald")
	x0y0test = dummyspielfeld.GibGelände(0,0)
	fmt.Println ("Das Gelände an der Koordinate x:0 y:0 ist nun vom Typ (erwartet Wüste):",x0y0test.GibTyp())
	/*// Test der Einheitenmethoden
	fmt.Println ("\nDie EinheitenMethoden des Spielfelds werden getestet:")
	var x0y0einheit einheiten.Einheit = einheiten.New(0,0,"Artillerie",'A')
	// SetzeEinheit wird getestet
	dummyspielfeld.SetzeEinheit(0,0,x0y0einheit)
	fmt.Println (x0y0einheit)
	fmt.Println (dummyspielfeld.GibEinheit(0,0))
	var x0y1einheit einheiten.Einheit = einheiten.New(0,1,"Infanterie",'A')
	dummyspielfeld.SetzeEinheit(0,1,x0y1einheit)
	fmt.Println ("Das Spielfeld wird gezeichnet")
	dummyspielfeld.Zeichnen()
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)
	x0y1einheit.SetzeTyp("Artillerie")
	// GibEinheit wird getestet
	var x0y1einheittest einheiten.Einheit = einheiten.New(0,1,"Infanterie",'A')
	x0y1einheittest= dummyspielfeld.GibEinheit(0,1)
	// Der Typ der ursprünglichen Infanterie muss nun auch Artillerie sein.
	fmt.Println ("Die Einheit an der Koordinate x:0 y:1 ist vom Typ (erwartet Artillerie):",x0y1einheittest.GibTyp())
	// EntferneEinheit wird getestet
	x0y0=dummyspielfeld.GibGelände(0,0)
	fmt.Println ("Das Geländefeld an der Koordinate x:0 y:0 hat folgenden Belegungsstatus(erwartet true):",x0y0.GibBelegung())
	dummyspielfeld.EntferneEinheit(0,0)
	x0y0=dummyspielfeld.GibGelände(0,0)
	fmt.Println ("Das Geländefeld an der Koordinate x:0 y:0 hat nun folgenden Belegungsstatus(erwartet false):",x0y0.GibBelegung())
	fmt.Println ("Das Spielfeld wird gezeichnet")
	dummyspielfeld.Zeichnen()
	gfx.TastaturLesen1()
	time.Sleep(2*1e9)*/
	// Test der Kodieren Methode:
	fmt.Println ("Das Spielfeld wird kodiert")
	var bytespielfeld []byte = make([]byte,0)
	bytespielfeld= dummyspielfeld.Kodieren()
	//fmt.Println (bytespielfeld)
	// Test der Dekodieren Methode:
	var spielfeldkopie spielfeld.Spielfeld 
	spielfeldkopie = spielfeld.New (20,12,50,"Spielfeldkopie")
	gfx.FensterAus()
	time.Sleep(2*1e9)
	fmt.Println ("Das Spielfeld wird dekodiert")
	spielfeldkopie.Dekodieren(bytespielfeld)
	fmt.Println ("Das Spielfeld wird gezeichnet")
	spielfeldkopie.Zeichnen()						
	gfx.TastaturLesen1()
}
