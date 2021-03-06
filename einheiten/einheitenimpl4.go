// Autoren: Friederike Richter und Moritz Raake
// Datum: 11.03.20
// Zweck: Softwarepraktikum - Einheiten-impl

package einheiten

import ("gfx";"../objekte")

type objekt objekte.Objekt

// Muss eine Einheit eine Farbe haben?; dass sollte doch eher eine Eigenschaft des Spielers sein, oder? Allerdings ist es eine Eigenschaft des Objektes....
// XXX es müssten noch Boni gegen verschiedene Typen aufgeführt werden
type data struct {
	objekt
	angriff,verteidigung uint16
	angriffreichweite uint16
	bewegungsreichweite uint16
	typ string
	energie uint16
	spieler uint8
}

type einheit struct{
	r,g,b uint8  
	rg,gg,bg uint8    
	a,v,reia,reib uint16
	energie uint16
}


var einheitentab = map[string]einheit{
	"Tank": einheit{
		0,0,255,
		0,0,0,
		5,1,4,2,
		10,
	},
	"Infanterie": einheit{
		255,0,0,
		0,0,0,
		1,1,3,1,
		10,
	},
}
 
func New (x,y uint16, t string, spieler uint8) *data {
	var neu *data = new(data)
	(*neu).objekt = objekte.New (x,y,einheitentab[t].r,einheitentab[t].g,einheitentab[t].b)
	(*neu).angriff = einheitentab[t].a
	(*neu).verteidigung = einheitentab[t].v
	(*neu).angriffreichweite = einheitentab[t].reia
	(*neu).bewegungsreichweite = einheitentab[t].reib
	if spieler==0||spieler==1 { // Aktuell zwei Spieler, kann aber beliebig erweitert werden
		(*neu).spieler=spieler
	}else{
		panic ("Falscher Spielername")
	}
	(*neu).typ=t
	(*neu).energie=einheitentab[t].energie
	return neu
}

	// Erg.: Der Angriffswert der Einheit ist geliefert.
func (e *data) GibAngriffswert () uint16 {
	return (*e).angriff
}
	
	// Eff.: Der Angriffswert der Einheit beträgt nun a. 
func (e *data) SetzeAngriffswert (a uint16) {
	(*e).angriff=a
}

	// Erg.: Der Verteidigungswert der Einheit ist geliefert.
func (e *data) GibVerteidigungswert () uint16 {
	return (*e).verteidigung
}
	
	// Eff.: Der Verteidigungswert der Einheit beträgt nun v. 
func (e *data) SetzeVerteidigungswert (v uint16) {
	(*e).verteidigung=v
}
	
	// Erg.: Die Angriffsreichweite (Anzahl der Felder) der Einheit ist geliefert.
func (e *data) GibAngriffsReichweite () uint16 {
	return (*e).angriffreichweite
}
	
	// Eff.: Die Angriffsreichweite (Anzahl der Felder) der Einheit ist auf rei gesetzt.
func (e *data) SetzeAngriffsReichweite (reia uint16) {
	(*e).angriffreichweite=reia
}
	
	// Erg.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist geliefert.
func (e *data) GibBewegungsReichweite () uint16 {
	return (*e).bewegungsreichweite
}
	
	// Eff.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist auf rei gesetzt.
func (e *data) SetzeBewegungsReichweite (reib uint16) {
	(*e).bewegungsreichweite=reib
}

	// Vor.: ---
	// Eff.: Die verbleibende Lebensenergie der Einheit ist geliefert.
func (e *data) GibEnergie () (en uint16){
	return (*e).energie			
}

func (e *data) SetzeEnergie (en uint16) {
	(*e).energie = en
}

	
	// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
func (e *data)	GibSpieler () (spieler uint8){
	return (*e).spieler
	}
	
	// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
func (e *data)	SetzeSpieler (spieler uint8){
	 (*e).spieler=spieler
//	 (e).Zeichnen()	
} 
	
	


	// Erg.: 
func (e *data) GibTyp () string {
return (*e).typ
}

func (e *data) SetzeTyp (t string) {
	e.SetzeFarbe(einheitentab[t].r,einheitentab[t].g,einheitentab[t].b)
	(*e).angriff = einheitentab[t].a
	(*e).verteidigung = einheitentab[t].v
	(*e).angriffreichweite = einheitentab[t].reia
	(*e).bewegungsreichweite = einheitentab[t].reib	
	(*e).energie=einheitentab[t].energie
/*
	if (*e).spieler==0{
		(*e).bild = einheitentab[t].bildspieler1
		}else{
			(*e).bild = einheitentab[t].bildspieler2
			}
*/ 
	(*e).typ=t
}

func (e *data) Kodieren () (b []byte) {
	b = make([]byte,0)
	var x,y uint16
	x,y = e.GibKoordinaten()
	// Kodieren der Koordinaten
	b = append(b,byte(x/256),byte(x%256),byte(y/256),byte(y%256)) 	
	b = append(b,byte((*e).energie/256),byte((*e).energie%256))
	b = append(b,byte((*e).spieler))
	var stringlaenge uint8
	// Kodieren der Anzahl an Bytes des Typs
	stringlaenge=uint8(len((*e).typ))					
	b = append(b,stringlaenge)
	// Kodieren des Geländetyps
	for _,w := range (*e).typ {								
		b = append(b,byte(w))
	}
	return
}

func (e *data) Dekodieren (b []byte) {
	// Dekodieren der Koordinaten
	e.SetzeKoordinaten(uint16 (b[0])*256+uint16(b[1]),uint16(b[2])*256+uint16(b[3]))
	// Dekodieren des Energiewertes
	(*e).energie=uint16(b[4])*256+uint16(b[5])
	// Dekodieren des Spielers
	e.SetzeSpieler(uint8(b[6]))
	var stringlaenge int = int(b[7])
	// Dekodieren des Einheitentyps
	var stringtyp string
	for i:=8;i<int(stringlaenge)+8;i++ {
		stringtyp=stringtyp+string(b[i])
	}
	e.SetzeTyp(stringtyp)
}


	
func (e *data) Zeichnen () {
	if !gfx.FensterOffen() {
		gfx.Fenster(1000,600)
	}
	//var r,gr,b uint8								// Variablen für Farbwerte
	var i uint16									// für das vergrößern der Zeichnung
	var x,y uint16									// Variablen für Koordinaten
	var spr,spg,spb uint8							// Spieler Farbe
	//r,gr,b = (e).GibFarbe()						// Abfragen der r,g,b Werte des Geländes
	x,y = (e).GibKoordinaten()  					// Abfragen der x,y-Koordinaten des Geländes
	//gfx.Stiftfarbe (r,gr,b)						// Festlegen der Geländefarbe
	//gfx.Vollkreis (x*50+25,y*50+25,(*e).radius)	// Zeichnen des Geländes
	//gfx.LadeBild(x*50+10,y*50+10,(*e).bild)
	//einheitentab[t].r


	x=x*50+10		//50 ist die Kästchengröße und kann später einfach durch eine Variable ersetzt werden 
	y=y*50+10		//10 ist der Abstand zum Rand, der zu zeichnenden Einheiten, in Pixeln 
	i=1				//i ist der "Zoomfaktor" der Einheit 1:= 30x30 Pixel --> 10:= 300x300 Pixel	

//Ramen um die Einheit
if e.GibGewaehlt() {
	gfx.Stiftfarbe(einheitentab[(*e).typ].rg,einheitentab[(*e).typ].gg,einheitentab[(*e).typ].bg)
	gfx.Vollrechteck(x-2,y-2,2,34)
	gfx.Vollrechteck(x+30,y-2,2,34)
	gfx.Vollrechteck(x-2,y-2,34,2)
	gfx.Vollrechteck(x-2,y+30,34,2)
	}
	



// Das sollte später die Objektfarbe werden?	
	if (*e).spieler==0{
		spr=255
		}else{
		spb=255	
		}
//gfx.Zeichnungen der einzelnen Eiheiten Typen
	switch (*e).typ{
		case "Tank":
			gfx.Stiftfarbe(0,0,0)
			gfx.Vollrechteck(x+6*i,y+7*i,8*i,3*i)
			gfx.Vollrechteck(x+3*i,y+10*i,16*i,7*i)
			gfx.Vollrechteck(x+10*i,y+11*i,19*i,3*i)
			gfx.Vollellipse(x+11*i,y+19*i,10*i,3*i)
			gfx.Stiftfarbe(50,50,50)
			gfx.Vollkreis(x+11*i,y+19*i,2*i)
			gfx.Vollkreis(x+16*i,y+19*i,2*i)
			gfx.Vollkreis(x+6*i,y+19*i,2*i)
			gfx.Stiftfarbe(spr,spg,spb)
			gfx.Vollrechteck(x+4*i,y+11*i,14*i,5*i)
			gfx.Vollrechteck(x+7*i,y+8*i,6*i,4*i)
			gfx.Vollrechteck(x+16*i,y+12*i,12*i,1*i)
		
		case "Infanterie":
			gfx.Stiftfarbe(0,0,0)
			gfx.Vollkreis(x+6*i,y+6*i,3*i)
			gfx.Vollkreis(x+15*i,y+6*i,3*i)
			gfx.Vollkreis(x+24*i,y+6*i,3*i)
			gfx.Vollrechteck(x+2*i,y+9*i,8*i,11*i)
			gfx.Vollrechteck(x+11*i,y+9*i,8*i,11*i)
			gfx.Vollrechteck(x+20*i,y+9*i,8*i,11*i)
			gfx.Vollrechteck(x+2*i,y+20*i,3*i,8*i)
			gfx.Vollrechteck(x+11*i,y+20*i,3*i,8*i)
			gfx.Vollrechteck(x+20*i,y+20*i,3*i,8*i)
			gfx.Vollrechteck(x+7*i,y+20*i,3*i,8*i)
			gfx.Vollrechteck(x+16*i,y+20*i,3*i,8*i)
			gfx.Vollrechteck(x+25*i,y+20*i,3*i,8*i)
			gfx.Stiftfarbe(spr,spg,spb)
			gfx.Vollrechteck(x+3*i,y+10*i,6*i,9*i)
			gfx.Vollrechteck(x+12*i,y+10*i,6*i,9*i)
			gfx.Vollrechteck(x+21*i,y+10*i,6*i,9*i)
			gfx.Vollrechteck(x+3*i,y+18*i,1*i,9*i)	
			gfx.Vollrechteck(x+12*i,y+18*i,1*i,9*i)
			gfx.Vollrechteck(x+21*i,y+18*i,1*i,9*i)
			gfx.Vollrechteck(x+8*i,y+18*i,1*i,9*i)
			gfx.Vollrechteck(x+17*i,y+18*i,1*i,9*i)
			gfx.Vollrechteck(x+26*i,y+18*i,1*i,9*i)	
	}
}




