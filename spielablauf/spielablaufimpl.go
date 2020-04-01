package spielablauf

import "../spielfeld"

type data struct {
  aktspieler uint8    // 0=Spieler 1 ; 1=Spieler2
  anzeinh1,anzeinh2   // Anzahl der Einheiten von Spieler 1 und 2 --> wenn eine auf 0 gibt es einen Sieger
  modus uint8         // 0=neutral ; 1=bewegen ; 2=angreifen // Start eines neues Spiels eigener Modus??--> eher über 
                      // Auswahlmenü am Anfang? speichern und laden? Optionsmenü? erweiterbar durch bauen,
                      // Itemeinsatz etc...
      
  aktiv bool          // Idee...ob es bereits einen Sieger gibt...ggf. ins Menü zurück wechseln --> ev. überflüssig, 
                      // da dies auch über anzeinh geregelt werden kann???
  landschaft string   // oder eher Level?? mit Switch-Case welche Datei dann eingespielt wird?? 
  spielfeld spielfeld.Spielfeld
}

func New () *data {
  var neu *data=new(data)
  // Erst zu erledigen wenn Struktur grob steht
  return neu
}

func (sa *data) GibSpieler () uint8 {
  return (*sa).aktspieler
}

func (sa *data) SetzeSpieler (s uint8)  {
  (*sa).aktspieler=s
}

func (sa *data) GibModus () uint8 {
  return (*sa).modus
}

func (sa *data) SetzeModus (m uint8) {
  (*sa).modus=m
}

func (sa *data) GibReichweite(einheit einheiten.Einheit)(erg,erg2 [][]bool) {
  y:=(*sa).spielfeld.GibHoehe()
  x:=(*sa).spielfeld.GibBreite()
	  erg=make ([][]bool,y)
	for i:=uint16(0);i<y;i++{
		erg[i]= make([]bool,x) //standartmäßig Matrix mit false
	}
	erg2=make ([][]bool,y)
	for i:=uint16(0);i<y;i++{
		erg2[i]= make([]bool,x) //standartmäßig Matrix mit false
	}
	// Achtung es ist noch nicht darauf geachtet, dass Eiheiten auf den zu betretenden Felder stehen
	xein,yein:=einheit.GibKoordinaten()
	println("Koordinaten:",xein,"/",yein)
	moeglicheschritte:=einheit.GibBewegungsReichweite()
	println("Bewegungsreichweite: ",moeglicheschritte)
	//gelände:=(*s).gelaendematrix[y][x]
	//bonusBewegung:=gelände.GibBonusBewegung()
	var reichweite func(xein,yein,schritte uint16)
	reichweite = func (xein,yein,schritte uint16){
		println("Anzahl Schritte:", schritte)
		if xein>x||yein>y&&schritte<0 {println ("Kann nicht weiter laufen");return}		//Abbruch, da am Rand, bzw. keine Schritte mehr übrig
		gelände:=(*sa).spielfeld.GibGelände(x,y)
		malusBewegung:=gelände.GibMalusBewegung() 
		println("Geländemalus: ",malusBewegung)			
		if erg[y][x] == false{
			println ("Feldinhalt wird geändert")
			erg[y][x]=true
			println (erg[y][x])
			if (int16(schritte)-malusBewegung)<0{println ("Kann nicht weiter laufen");return}	//Abbruch, da Abzug zu groß...
			// Für jeden weiteren Aufruf muss überprüft werden, ob sich auf diesem Feld eine
			// Einheit befindet
			if (*sa).spielfeld.GibEinheit(x+1,y)==nil {
				reichweite(x+1,y,schritte-uint16(malusBewegung))	// Darf ja jetzt nicht nur einen Schritt weniger weit gehen...
			} else {
        if (*sa).spielfeld.GibEinheit(x+1,y).GibSpieler()!=einheit.GibSpieler() {
				  erg2[y][x+1]=true
        }
			}
			if (*sa).spielfeld.GibEinheit(x-1,y)==nil {
				reichweite(x-1,y,schritte-uint16(malusBewegung))	// ...sondern so viele Schritte weniger, wie ihm durch den Malus 
			} else {
        if (*sa).spielfeld.GibEinheit(x-1,y).GibSpieler()!=einheit.GibSpieler() {
				   erg2[y][x-1]=true
        }
			}
			if (*sa).spielfeld.GibEinheit(x,y+1)==nil {
				reichweite(x,y+1,schritte-uint16(malusBewegung))	// ...abgezogen werden
			} else {
        if (*sa).spielfeld.GibEinheit(x,y+1).GibSpieler()!=einheit.GibSpieler() {
				  erg2[y+1][x]=true
        }
			}
			if (*sa).spielfeld.GibEinheit(x,y-1) ==nil {
				reichweite(x,y-1,schritte-uint16(malusBewegung))
			} else {
        if (*sa).spielfeld.GibEinheit(x,y-1).GibSpieler()!=einheit.GibSpieler() {
				  erg2[y-1][x]=true
        }
			}
		}
	} 
	reichweite(x,y,moeglicheschritte)
	return 
}