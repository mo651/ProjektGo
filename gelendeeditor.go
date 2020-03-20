package main
//XXX Speicher fehlt noch 
import (/*"fmt";"./einheiten";*/"./gelaende";"gfx";"fmt";"./Hilfspakete/dateien";"./button")
var speichern bool
var laden bool
//var modus string
var ok bool= true
		
func maussteuerung (s []gelaende.Gelände, a []gelaende.Gelände) {
var agewaehlt string
a[0].SetzeGewaehlt(true)
// Es wird überprüft, welcher der Geländetypen ausgewählt ist
	for { 
		for _,w:=range a{
			if w.GibGewaehlt(){
				agewaehlt = w.GibTyp()
			}
		}
// Die Mausabfrage beginnt:
// Wir ein Element aus dem Spielfeld gewählt:	
		taste, status, mausX, mausY:= gfx.MausLesen1 ()
		if taste == 1 && status == 1 { //linke Maustaste gerade gedrückt
			if mausX<1000 && mausY<600 {
				var x,y uint16
				x= mausX/50
				y= mausY/50
				for _,w := range s {
					xs,ys:=w.GibKoordinaten()		
					if x==xs && y==ys {
						w.SetzeTyp(agewaehlt)
					}
				}	
			}
// Oder ein neuer Geländetyp			
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
			//Speichern
			if mausX>1050 && mausX<1150 && mausY>600 && mausY<650{
				speichern=true
				ok=true
				}
			//laden 1050,400,100,50
			if mausX>1050 && mausX<1150 && mausY>400 && mausY<450{
				ok=true
				laden=true
				}						
		}
	}
}

func main () {
// Aufbau des Fensters:	
	gfx.Fenster(1200,700)
// Zur Auswahl stehende Geländetype werden aufgebaut	
	var auswahlgelaende []gelaende.Gelände 
	auswahlwueste := gelaende.New (21,0,"Wueste")
	auswahlwasser := gelaende.New (22,0,"Wasser")
	auswahlwald := gelaende.New (21,1,"Wald")
	auswahlberg := gelaende.New (22,1,"Berg")	
	auswahlgelaende=append(auswahlgelaende,auswahlwasser)
	auswahlgelaende=append(auswahlgelaende,auswahlwald)	
	auswahlgelaende=append(auswahlgelaende,auswahlwueste)
	auswahlgelaende=append(auswahlgelaende,auswahlberg)
	
// Button zum Speichern wird erzeugt
	save:=button.New(1050,600,100,50,0,0,0,255,255,255,0,0,0,"save",'m','m')
	save.Zeichnen()	
	
// Button zum Laden wird erzeugt
	load:=button.New(1050,400,100,50,0,0,0,255,255,255,0,0,0,"load",'m','m')
	load.Zeichnen()		
	
// Spielfeld wird aufgebaut; XXX dass soll noch in eine Matrix mit Einheiten und Items, so dass gespeichert und geladen werden kann	
	var spielfeld []gelaende.Gelände
	spielfeld = make ([]gelaende.Gelände,0)
	
	for i:=uint16(0);i<20;i++ {
		for j:=uint16(0);j<12;j++ {
			var neuesgelände gelaende.Gelände
			neuesgelände = gelaende.New (i,j,"Wasser") //man beachte, dass Wasser ist "animiert", ob das im feritgen Spiel auch funktionier kann ich nicht sagen
			spielfeld=append(spielfeld,neuesgelände)
			}	
		}
// Mausabfrage findet nebeläufig statt	 
	go maussteuerung (spielfeld,auswahlgelaende)

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

				level = dateien.Oeffnen ("./Level/level_"+erg , 'l' )
				i:=0
				for !level.Ende(){
					var bs []byte
					ende:=int(level.Lesen()) 
					for a:=0; a<ende;a++{
						bs=append(bs,level.Lesen())
						}
					i++
					fmt.Println(i-1,bs)	
					spielfeld[i-1].Dekodieren(bs)		
					}
				level.Ende()	
				laden=false
				
				
				
				}
			gfx.UpdateAus()
//Spielwelt wird gezeichnet		
				for _,w:=range spielfeld {
					w.Zeichnen()
				}
				for _,w:=range auswahlgelaende {
					w.Zeichnen()
				}
//Spielraster wird gezeichnet		
			gfx.Stiftfarbe(0,0,0)
			for i:=uint16(0);i<20;i++ {
				gfx.Linie(i*50,0,i*50,599)
			}
			for i:=uint16(0);i<12;i++ {
				gfx.Linie(0,i*50,999,i*50)
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
//Datei schreiben
			if len(erg)>0{
				var level dateien.Datei
				level = dateien.Oeffnen ("./Level/level_"+erg , 's' )			
for _,w := range spielfeld {
					var bs []byte 
					bs = w.Kodieren()
					level.Schreiben(byte(len(bs)))
					for _,w2 := range bs {	
						level.Schreiben(w2)
					}
				}
				level.Ende()
			}		
			speichern=false
		}
	}	
}
