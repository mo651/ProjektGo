// Autoren: Friederike Richter und Moritz Raake
// Datum: 11.03.2020
// Zweck: Softwarepraktikum 2020

package spielfeld

import ("../gelaende";"../einheiten";"gfx" /*; "../items"*/)



type data struct {
  x,y uint16						// Anzahl x- & y-Kästchen --> perspektivisch soll es variabel sein
	groesse uint16					// Pixelzahl Feldgröße --> perspektivisch soll es variabel sein
	dateiname string
	gelaendematrix [][]gelaende.Gelände	// 
	einheitenmatrix [][]einheiten.Einheit	
  //itemmatrix [][]items.Item		// --> später einzufügen....Optional je nach Spielentwicklung
}

func New (x,y,groesse uint16, dateiname string) *data {
	var neu *data = new(data)
  (*neu).gelaendematrix= make ([][]gelaende.Gelände,y)
	(*neu).einheitenmatrix=make ([][]einheiten.Einheit,y)
	(*neu).x= 20 // x --> später wenn variabel;
	(*neu).y= 12 // y --> später wenn variabel
	(*neu).groesse = 50 // groesse --> später wenn variabel
	(*neu).dateiname = dateiname
	for i:=uint16(0);i<(*neu).y;i++ {
		var neuegelaendezeile []gelaende.Gelände = make([]gelaende.Gelände,(*neu).x)
		(*neu).gelaendematrix[i]=neuegelaendezeile
	}
	for i:=uint16(0);i<(*neu).y;i++ {
		var neueeinheitenzeile []einheiten.Einheit = make([]einheiten.Einheit,(*neu).x)
		(*neu).einheitenmatrix[i]=neueeinheitenzeile
	}
  
	//(*neu).itemmatrix=make ([][]items.Item,y)	
	//for i:=0;i<y;i++ {
	//	var neueitemzeile []items.Item = make([]items.Item,x)
	//	(*neu).itemmatrix[i]=neueitemzeile
	//}																	// --> später einzufügen....Optional je nach Spielentwicklung
	return neu
}

	// Vor.: ---
	// Erg.: Ein Verweis auf das an dem übergebenen Koordinaten befindliche Gelände ist geliefert.	

func (s *data) GibGelände (x,y uint16) gelaende.Gelände {
	return (*s).gelaendematrix[y][x]
}

	// Vor.: ---
	// Eff.: Das übergebene Gelände befindet sich nun an den übergebenen Koordinaten. Das vorherig dort
	//		 befindliche Gelände ist entfernt.

func (s *data) SetzeGelände (gelände gelaende.Gelände) {
  x,y:=gelände.GibKoordinaten()
	(*s).gelaendematrix[y][x]=gelände
}
	
	// Vor.: Das Geländefeld muss mit einer Einheit besetzt sein.
	// Erg.: Ein Verweis auf die auf dem übergebenen Geländefeld befindlichen Instanz ist geliefert.
func (s *data) GibEinheit (x,y uint16) einheiten.Einheit {
	if (*s).einheitenmatrix[y][x] == nil {
		panic ("Geländefeld nicht belegt")
	}
	return (*s).einheitenmatrix[y][x]
}
	
	// Vor.: Das Geländefeld mit den übergebenen Koordinaten darf nicht besetzt sein.
	// Eff.: Auf dem Geländefeld mit den übergebenen Koordinaten befindet sich nun die übergebene
	//		 Einheit. 
func (s *data) SetzeEinheit (einh einheiten.Einheit) {
  x,y:=einh.GibKoordinaten ()
	if (*s).einheitenmatrix[y][x] == nil {	
		(*s).einheitenmatrix[y][x] = einh
		// Belegung des Geländefeldes muss auf 'true' geändert werden
		(*s).gelaendematrix[y][x].SetzeBelegung(true)
	}
}
	
	// Vor.: Das Geländefeld mit den übergebenen Koordinaten muss mit einer Einheit besetzt sein.
	// Eff.: Die auf dem Geländefeld mit den übergebenen Koordinaten befindliche Einheit ist entfernt.
func (s *data) EntferneEinheit (x,y uint16) {
	if (*s).einheitenmatrix[y][x] != nil {
		(*s).einheitenmatrix[y][x] = nil
		// Belegung des Geländefeldes muss auf 'false' geändert werden
    (*s).gelaendematrix[y][x].SetzeBelegung(false)
	}
}
	
	/*GibItem (x,y uint16) items.Item
	
	SetzeItem (x,y uint16, item items.Item) 
		
	EntferneItem (x,y uint16) */
	
func (s *data) Zeichnen () {
	if !gfx.FensterOffen() {
		gfx.Fenster((*s).x*(*s).groesse+200,(*s).y*(*s).groesse+100) // --> variabel, falls entfernen breite:1200 und Höhe 700
	}
	// Zeichnen der Geländeschicht
	for _,gelreihe := range (*s).gelaendematrix {
		for _,gel := range gelreihe {
			gel.Zeichnen ()
		}
	}
	// Einzeichnen der Einheiten
	for _,einhreihe := range (*s).einheitenmatrix {
		for _,einh := range einhreihe {
			if einh != nil {
				einh.Zeichnen ()
			}
    }
	}
	/* Einzeichnen der Items
	 * for _,itmreihe := range itemmatrix {
		for _,itm := range itmreihe {
			if itm != nil {
				itm.Zeichnen ()
			}
		}
	} *///--> sofern Items hinzugefügt werden sollen 
	// Das Gitterraster wird darüber gezeichnet
	gfx.Stiftfarbe(0,0,0)
	// Vertikale Linien werden gezeichnet
	for i:=uint16(0);i<(*s).x;i++ {											// Gitter ebenfalls variabel gezeichnet
		gfx.Linie(i*(*s).groesse,0,i*(*s).groesse,(*s).y*(*s).groesse-1)	// ebd.
	}
	// Horizontale Linien werden gezeichnet
	for i:=uint16(0);i<(*s).y;i++ {											// ebd.
		gfx.Linie(0,i*(*s).groesse,(*s).x*(*s).groesse-1,i*(*s).groesse)	// ebd.
	}
}	
	
	// Vor.: ---
	// Erg.: Eine Folge von Bytes ist geliefert, die die aufrufende Instanz serialisiert darstellt.
func (s *data) Kodieren () (b []byte){
	b = make ([]byte,0)
	b = append(b,byte((*s).x/256),byte((*s).x%256))					// Kodieren der X-Koordinaten ---> XXXX nur notwendig wenn variabel
	b = append(b,byte((*s).y/256),byte((*s).y%256)) 				// Kodieren der Y-Koordinaten ---> XXXX nur notwendig wenn variabel
	b = append(b,byte((*s).groesse/256),byte((*s).groesse%256)) 	// Kodieren der Größe---> XXXX nur notwendig wenn variabe
	var stringlaenge uint8
	// Kodieren der Anzahl an Bytes des Strings
	stringlaenge=uint8(len((*s).dateiname))					
	b = append(b,stringlaenge)
	// Kodieren des Dateinamens
	for _,w := range (*s).dateiname {								
		b = append(b,byte(w))
	}
	// Kodieren der Geländematrix, Durchlaufen der Zeilen
	for _,w := range (*s).gelaendematrix {	
		//Durchlaufen der einzelnen Einträge/Gelände					
		for _,gel := range w {
			// Kodieren des Geländes als []byte ...
			var gelaendebytes []byte = make([]byte,0)
			gelaendebytes = gel.Kodieren()
			// Kodieren und anhängen der Anzahl an Bytes des Geländes
			var gelaendegroesse uint16 = uint16(len(gelaendebytes))
			b = append(b,byte(gelaendegroesse/256),byte(gelaendegroesse%256))			
			// ... und anhängen des Geländeslices
			//b = append(b,gel.Kodieren()...) 	
			b = append (b,gelaendebytes...)					
		}
	}
	// Kodieren der Einheitenmatrix, Durchlaufen der Zeilen
	for _,w := range (*s).einheitenmatrix {	
		//Durchlaufen der einzelnen Einträge/Einheiten					
		for _,ein := range w {
			if ein != nil {
				// Kodieren einer 1, als Zeichen, dass sich eine Einheit auf dem Feld befunden hat
				b = append(b,1)
				// Kodieren der Einheit als []byte ...
				var einheitbytes []byte = make([]byte,0)
				einheitbytes = ein.Kodieren()
				// Kodieren und anhängen der Anzahl an Bytes des Geländes
				var einheitgroesse uint16 = uint16(len(einheitbytes))
				b = append(b,byte(einheitgroesse/256),byte(einheitgroesse%256))			
				// ... und anhängen des Einheitslices
				b = append (b,einheitbytes...)	
			} else {
				// Kodieren einer 0, als Zeichen, dass sich keine Einheit auf dem Feld befunden hat.
				b = append(b,0)
			}			
		}
	}
	return
}
	
	// Vor.: Der übergebene Slice b entspricht einem kodierten Spielfeld gleicher Größe.
	// Eff.: Die aufrufende Instanz hat nun genau diejenige Eigenschaft, sodass 'Kodieren ()' exakt
	//		 b ergeben würde. Alle vorherigen Eigenschaften des Spielfeldes hat sie nicht mehr.
func (s *data) Dekodieren (b []byte) {
	//println ("Die Länge des zu dekodierenden Bytefeldes ist: ",len(b))
	//Hilfsvariable zum Durchlaufen von b
	var aktindex int
	(*s).x=uint16(b[0])*256+uint16(b[1])
	//println ("Die X-Koordinate wurde dekodiert...")		
	//println ("...und lautet: ",(*s).x)
	aktindex=aktindex+2
	(*s).y=uint16(b[2])*256+uint16(b[3])
	//println ("Die X-Koordinate wurde dekodiert...")		
	//println ("...und lautet: ",(*s).y)
	aktindex=aktindex+2
	(*s).groesse=uint16(b[4])*256+uint16(b[5])
	//println ("Die Feldgröße wurde dekodiert")			
	//println ("...und lautet: ",(*s).groesse)
	aktindex=aktindex+2
	// Dekodieren der Anzahl an benötigter Bytes für Dateinamenstrings
	var stringlaenge uint8 = b[aktindex]
	aktindex++
	// Dekodieren des Dateinamens
	for i:=aktindex;i<int(stringlaenge)+aktindex;i++ {
		if i==aktindex {
			(*s).dateiname =string(b[i])
		} else {
		(*s).dateiname=(*s).dateiname+string(b[i])
		}
	}
	//println ("Der Dateiname wurde dekodiert...")			
	//println ("...und lautet: ",(*s).dateiname)
	//println ("Der Dateiname hat diese Größe: ",stringlaenge," und muss diese haben: ",len("Dummyfeld"))
	aktindex=aktindex+int(stringlaenge)
	// Dekodieren der Geländematrix
	// Durchlaufen der Reihen
	for y:=uint16(0);y<(*s).y;y++ {
		//println("Starte mit Reihe: ",y)
		//Durchlaufen der einzelnen Felder
		for x:=uint16(0);x<(*s).x;x++ {
			//Ermitteln des Startindex des nächsten Geländefeldes
			//println("Aktueller Index: ",aktindex)
			var gelaendegroesse int
			gelaendegroesse= int(b[aktindex])*256+int(b[aktindex+1])
			//println("Größe des nächsten Feldes",gelaendegroesse)
			aktindex=aktindex+2
			//println("Aktueller Index: ",aktindex)
			//print("Einfügen des Feldinhalts ab Index: ",aktindex)
			//println(" an der Stelle x: ",x," y: ",y)
			// Dekodieren des nächsten Geländes
			var dekodgel gelaende.Gelände 
			dekodgel=gelaende.New(x,y,"Wasser") 
			//println("Neues Gelände zum Einfügen erstellt...")				
			//println("Dekodieren des Bytefeldes von Index: ",aktindex," bis Index: ",aktindex+gelaendegroesse)
			var hilfsbytefeld []byte = make ([]byte,0)
			for i:=aktindex;i<aktindex+gelaendegroesse+1;i++ {
				if x==19&&y==11&&i==aktindex+gelaendegroesse {		
				} else {
					hilfsbytefeld = append (hilfsbytefeld,b[i])
				}
			}
			//println ("Hilfsbytefeld erstellt")
			//for _,w:=range hilfsbytefeld {
				//print (w," ")
			//}
			//print ("Alle Bytes gedruckt\n")
			dekodgel.Dekodieren(hilfsbytefeld)
			//println("...dekodiert und ...")
			s.SetzeGelände(dekodgel)
			//println("eingefügt")
			aktindex=aktindex+gelaendegroesse
		}
		//println (y,". Reihe erfolgreich dekodiert")
	}
	// Dekodieren der Einheitenmatrix
	// Durchlaufen der Reihen
	for y:=uint16(0);y<(*s).y;y++ {
		//println("Starte mit Reihe: ",y)
		//Durchlaufen der einzelnen Felder
		for x:=uint16(0);x<(*s).x;x++ {
			// Ermitteln ob eine Einheit eingetragen werden muss
			var besetzt byte
			besetzt = b[aktindex]
			aktindex++ 
			if besetzt == 1 {
				// Eine Einheit muss eingetragen werden
				//Ermitteln des Startindex der nächsten Einheit
				//println("Aktueller Index: ",aktindex)
				var einheitgroesse int
				einheitgroesse= int(b[aktindex])*256+int(b[aktindex+1])
				//println("Größe der nächsten Einheit",einheitgroesse)
				aktindex=aktindex+2
				//println("Aktueller Index: ",aktindex)
				//print("Einfügen der Einheit ab Index: ",aktindex)
				//println(" an der Stelle x: ",x," y: ",y)
				// Dekodieren der nächsten Einheit
				var dekodein einheiten.Einheit 
				dekodein=einheiten.New(x,y,"Artillerie",0) 
				//println("Neue Einheit zum Einfügen erstellt...")
				//println("Dekodieren des Bytefeldes von Index: ",aktindex," bis Index: ",aktindex+einheitgroesse)
				var hilfsbytefeld []byte = make ([]byte,0)
				for i:=aktindex;i<aktindex+einheitgroesse+1;i++ {
					if x==19&&y==11&&i==aktindex+einheitgroesse {
					} else {
						hilfsbytefeld = append (hilfsbytefeld,b[i])
					}
				}
				//println ("Hilfsbytefeld erstellt")
				//for _,w:=range hilfsbytefeld {
					//print (w," ")
				//}
				//print ("Alle Bytes gedruckt\n")
				dekodein.Dekodieren(hilfsbytefeld)
				//println("...dekodiert und ...")
				s.SetzeEinheit(dekodein)
				//println("eingefügt")
				aktindex=aktindex+einheitgroesse
			} else {
				// Es musste keine Einheit eingetragen werden
				(*s).einheitenmatrix[y][x]=nil
			} 	
		} 
		//println (y,". Reihe erfolgreich dekodiert")
	}
}

func (s *data) GibReichweite(einheit einheiten.Einheit)(erg[][]bool) {
	erg=make ([][]bool,(*s).y)
	for i:=uint16(0);i<(*s).y;i++{
		erg[i]= make([]bool,(*s).x) //standartmäßig Matrix mit false
		}
	// Achtung es ist noch nicht darauf geachtet, dass Eiheiten auf den zu betretenden Felder stehen
	x,y:=einheit.GibKoordinaten()
  println("Koordinaten:",x,"/",y)
	moeglicheschritte:=einheit.GibBewegungsReichweite()
  println("Bewegungsreichweite: ",moeglicheschritte)
	//gelände:=(*s).gelaendematrix[y][x]
	//bonusBewegung:=gelände.GibBonusBewegung()
	var reichweite func(x,y,schritte uint16)
	reichweite = func (x,y,schritte uint16){
    println("Anzahl Schritte:", schritte)
		if x>(*s).x||y>(*s).y&&schritte<0 {println ("Kann nicht weiter laufen");return}		//Abbruch, da am Rand, bzw. keine Schritte mehr übrig
		gelände:=(*s).gelaendematrix[y][x]
		malusBewegung:=gelände.GibMalusBewegung() 
    println("Geländemalus: ",malusBewegung)			
		if erg[y][x] == false{
      println ("Feldinhalt wird geändert")
			erg[y][x]=true
      println (erg[y][x])
			if (int16(schritte)-malusBewegung)<0{println ("Kann nicht weiter laufen");return}	//Abbruch, da Abzug zu groß...
			// Für jeden weiteren Aufruf muss überprüft werden, ob sich auf diesem Feld eine
			// Einheit befindet
			if (*s).einheitenmatrix[y][x+1] ==nil {
				reichweite(x+1,y,schritte-uint16(malusBewegung))	// Darf ja jetzt nicht nur einen Schritt weniger weit gehen...
			}
			if (*s).einheitenmatrix[y][x-1] ==nil {
				reichweite(x-1,y,schritte-uint16(malusBewegung))	// ...sondern so viele Schritte weniger, wie ihm durch den Malus 
			}
			if (*s).einheitenmatrix[y+1][x] ==nil {
				reichweite(x,y+1,schritte-uint16(malusBewegung))	// ...abgezogen werden
			}
			if (*s).einheitenmatrix[y-1][x] ==nil {
				reichweite(x,y-1,schritte-uint16(malusBewegung))
			}
		}
	} 
	reichweite(x,y,moeglicheschritte)
	return 
}

