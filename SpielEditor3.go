package main
//XXX Speicher fehlt noch 
import (/*"fmt";*/"./einheiten";"./gelaende";"gfx";"fmt";"./Hilfspakete/dateien";"./button";"./spielfeld")
var speichern bool
var laden bool
//var modus string
var ok bool= true
		
func maussteuerung (s spielfeld.Spielfeld, a []gelaende.Gelände, b []einheiten.Einheit) {
	var agewaehlt string
	var bgewaehlt string
	var spieler uint8
	var element uint8 // bei 0 soll ein Gelände eingefügt werden, bei 1 soll eine Einheit eingefügt werden
	a[0].SetzeGewaehlt(true)
	for { 
		// Es wird überprüft, welcher der Geländetypen ausgewählt ist
		for _,w:=range a{
			if w.GibGewaehlt(){
				agewaehlt = w.GibTyp()
			}
		}
		// Es wird überprüft, welcher der Einheitentypen ausgewählt ist
		for _,w:=range b{
			if w.GibGewaehlt(){
				bgewaehlt = w.GibTyp()
			}
		}
	// Die Mausabfrage beginnt:
	// Wird ein Geländefeld auf dem Spielfeld gewählt:	
		taste, status, mausX, mausY:= gfx.MausLesen1 ()
		if taste == 1 && status == 1 { //linke Maustaste gerade gedrückt
			if mausX<1000 && mausY<600 {				// XXX wenn Variabel hier und im folgenden Spielfeldmethoden
				var x,y uint16							// XXX GibX GibY GibGroesse einbauen
				x= mausX/50
				y= mausY/50
				if element == 0 {
					var einzufgel gelaende.Gelände = gelaende.New (x,y,agewaehlt)
					s.SetzeGelände (x,y,einzufgel)	
				} else if element == 1 {
					var einzufein einheiten.Einheit = einheiten.New (x,y,bgewaehlt,spieler)
					s.SetzeEinheit (x,y,einzufein)	
				}		
			}
	// Oder ein neuer Geländetyp gewählt, welcher gesetzt werden soll			
			if mausX>1050 && mausX<1150 && mausY<100 {
				var x,y uint16
				x= mausX/50
				y= mausY/50
				for _,w := range a {
					xs,ys:=w.GibKoordinaten()
					w.SetzeGewaehlt(false)
					if x==xs && y==ys {
						w.SetzeGewaehlt(true)
					}
				}	
			}
	// Oder ein neuer Einheitentyp gewählt, welcher gesetzt werden soll			
			if mausX>1050 && mausX<1150 && mausY<100 {
				var x,y uint16
				x= mausX/50
				y= mausY/50
				for _,w := range b {
					xs,ys:=w.GibKoordinaten()
					w.SetzeGewaehlt(false)
					if x==xs && y==ys {
						w.SetzeGewaehlt(true)
						spieler=w.GibSpieler()
					}
				}	
			}
	// Oder es wird Speichern ausgewählt
			if mausX>1050 && mausX<1150 && mausY>600 && mausY<650{
				speichern=true
				ok=true
			}
	// Oder es wird Laden ausgewählt
			if mausX>1050 && mausX<1150 && mausY>400 && mausY<450{
				ok=true
				laden=true
			}						
		}
	}
}

func main () {
// Aufbau des Fensters:	
	gfx.Fenster(1200,700)								// XXX Sofern variabel gestaltet mit Spielfeld-Methoden abfragen
// Zur Auswahl stehende Geländetypen werden aufgebaut	
	var auswahlgelaende []gelaende.Gelände 
	var auswahleinheit []einheiten.Einheit
	auswahlwueste := gelaende.New (21,0,"Wueste")
	auswahlwasser := gelaende.New (22,0,"Wasser")
	auswahlwald := gelaende.New (21,1,"Wald")
	auswahlberg := gelaende.New (22,1,"Berg")	
	auswahlgelaende=append(auswahlgelaende,auswahlwasser)
	auswahlgelaende=append(auswahlgelaende,auswahlwald)	
	auswahlgelaende=append(auswahlgelaende,auswahlwueste)
	auswahlgelaende=append(auswahlgelaende,auswahlberg)
	auswahltankA := einheiten.New (21,2,"Tank",0)
	auswahltankB := einheiten.New (22,2,"Tank",1)
	auswahlinfanterieA := einheiten.New (21,3,"Infanterie",0)
	auswahlinfanterieB := einheiten.New (22,3,"Infanterie",1)	
	auswahleinheit=append(auswahleinheit,auswahltankA)
	auswahleinheit=append(auswahleinheit,auswahltankB)	
	auswahleinheit=append(auswahleinheit,auswahlinfanterieA)
	auswahleinheit=append(auswahleinheit,auswahlinfanterieB)
	
// Button zum Speichern wird erzeugt
	save:=button.New(1050,600,100,50,0,0,0,255,255,255,0,0,0,"save",'m','m')
	save.Zeichnen()	
	
// Button zum Laden wird erzeugt
	load:=button.New(1050,400,100,50,0,0,0,255,255,255,0,0,0,"load",'m','m')
	load.Zeichnen()		
	
// Spielfeld wird aufgebaut
	var dummyspielfeld spielfeld.Spielfeld 
	dummyspielfeld = spielfeld.New (20,12,50,"Dummyfeld")			// XXXXX muss verändern werden, wenn Variabilität verloren geht
	
	for y:=uint16(0);y<12;y++ {
		for x:=uint16(0);x<20;x++ {
			var neuesgelände gelaende.Gelände
			neuesgelände = gelaende.New (x,y,"Wasser") //XXXX man beachte, dass Wasser ist "animiert", ob das im fertigen Spiel auch funktionier kann ich nicht sagen
			dummyspielfeld.SetzeGelände(x,y,neuesgelände)	
		}	
	}
// Mausabfrage findet nebenläufig statt	 
	go maussteuerung (dummyspielfeld,auswahlgelaende,auswahleinheit)

//  Endlosschleife mit Programmsteuerung startet
	for {
//	Editiermodus (nicht speichern) XXX muss noch geändert werden, da auch geladen werden soll		
		if !speichern{
			
			if laden{
			var level dateien.Datei 						
			gfx.UpdateAus()
	//Rechteck für das Speichern-Menü		
			lade_menue:=button.New(400,200,200,100,0,0,0,255,255,255,0,0,0,"Dateiname",'o','l')
			lade_menue.Zeichnen()			
	//Rechteck für den Enter Hinweis			
			hinweisl:=button.New(400,275,200,25,0,0,0,255,255,255,0,0,0,"OK: Enter; Abbruch ESC",'m','m')
			hinweisl.Zeichnen()				
			gfx.UpdateAn()
			var erg string
			for ok {
				taste,gedrueckt,_:=gfx.TastaturLesen1 () 
				if gedrueckt==0{
					if taste==13{
						ok=false
					}else if taste == 27{
						ok=false
						speichern=false	
					}else{		
						erg=erg+string(rune((int(taste))))		
						fmt.Println(erg)
						gfx.UpdateAus()
						gfx.Stiftfarbe(255,255,255)
						gfx.Vollrechteck(420,230,100,20)
						gfx.Stiftfarbe(0,0,0)
						for i,w:= range erg {
							gfx.Schreibe((410+uint16(8*i)),220,string(w))
						}
						gfx.UpdateAn()
					}
				}		
			}
// Datei laden
			level = dateien.Oeffnen ("./Level/level_"+erg , 'l' )
			var bs []byte = make([]byte,0)
				for !level.Ende() {
					bs=append(bs,level.Lesen())
				}
				dummyspielfeld.Dekodieren (bs)			
				level.Ende()
			}		
			laden=false	
			gfx.UpdateAus()
//Spielwelt wird gezeichnet
			dummyspielfeld.Zeichnen()
			for _,w:=range auswahlgelaende {
				w.Zeichnen()
			}
			for _,w:=range auswahleinheit{
				w.Zeichnen()
			}
			gfx.UpdateAn()
		}else{
			gfx.UpdateAus()
//Rechteck für das Speichern-Menü		
			save_menue:=button.New(400,200,200,100,0,0,0,255,255,255,0,0,0,"Dateiname",'o','l')
			save_menue.Zeichnen()			
//Rechteck für den Enter Hinweis			
			hinweis:=button.New(400,275,200,25,0,0,0,255,255,255,0,0,0,"OK: Enter; Abbruch ESC",'m','m')
			hinweis.Zeichnen()				
			gfx.UpdateAn()
			var erg string
			for ok {
				taste,gedrueckt,_:=gfx.TastaturLesen1 () 
				if gedrueckt==0{
					if taste==13{
						ok=false
					}else if taste == 27{
						ok=false
						speichern=false	
					}else{		
						erg=erg+string(rune((int(taste))))		
						fmt.Println(erg)
						gfx.UpdateAus()
						gfx.Stiftfarbe(255,255,255)
						gfx.Vollrechteck(420,230,100,20)
						gfx.Stiftfarbe(0,0,0)
						for i,w:= range erg {
							gfx.Schreibe((410+uint16(8*i)),220,string(w))
						}
						gfx.UpdateAn()
					}
				}		
			}
//Datei speichern
			if len(erg)>0{
				var level dateien.Datei
				level = dateien.Oeffnen ("./Level/level_"+erg , 's' )
				var bs []byte = make([]byte,0)
				bs=dummyspielfeld.Kodieren()
				for _,w:=range bs {
					level.Schreiben(w)
				}		
				level.Ende()
			}	
			speichern=false
		}
	}	
}
