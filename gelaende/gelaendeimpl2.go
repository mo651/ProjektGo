package gelaende

import ("gfx";"../objekte"; "../Hilfspakete/zufallszahlen")

type objekt objekte.Objekt

type data struct {
	objekt
//	r,g,b uint8
//	zr,zg,zb uint8	// Zufallswerte für Gelände
//	rg,gg,bg uint8 // Farbe für gewählt
	bonusa,bonusv int16
	malusb int16
	belegt bool
	typ string
}

type untergrund struct{
	r,g,b uint8
	zr,zg,zb uint8
	rg,gg,bg uint8 			// Farbe für gewählt
	bonusa,bonusv int16
	malusb int16
}

var untergrundtab = map[string]untergrund{
	"Wasser": untergrund{
		0,0,200,
		30,30,30,
		0,0,0,
		-2,-2,
		2,
	},
	"Wald": untergrund{
		0,130,10,
		30,30,30,
		0,0,0,
		1,2,
		1,
	},
	"Berg": untergrund{
		130,130,130,
		30,30,30,
		0,0,0,
		1,2,
		1,
	},
	"Wueste": untergrund{
		180,180,10,
		30,30,30,
		0,0,0,
		2,2,
		2,
	},
}

/*
func New (x,y uint16, r,g,b uint8, bonusa uint16) *data {
	var neu *data = new(data)
	(*neu).objekt = objekte.New(x,y,r,g,b)
	(*neu).bonusa = bonusa
	(*neu).belegt = false 	// Standardmäßig mit false belegt
	(*neu).bonusb = 0		// Standardmäßig mit 0 belegt
	(*neu).bonusv = 0		// Standardmäßig mit 0 belegt
	return neu
}
*/

func New (x,y uint16, t string) *data {
	zufallszahlen.Randomisieren()
	var neu *data = new(data)
	(*neu).objekt = objekte.New (x,y,untergrundtab[t].r+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zr))),untergrundtab[t].g+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zg))),untergrundtab[t].b+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zb))))
//	(*neu).rg 	  = untergrundtab[t].rg
//	(*neu).gg 	  = untergrundtab[t].gg
//	(*neu).bg 	  = untergrundtab[t].bg
	(*neu).bonusa = untergrundtab[t].bonusa
	(*neu).bonusv = untergrundtab[t].bonusv
	(*neu).malusb = untergrundtab[t].malusb
	(*neu).belegt = false 	// Standardmäßig mit false belegt
	(*neu).typ=t
	return neu
}


	// Erg.: Der Bonus zum Angriffswert des Geländes ist geliefert.
func (g *data) GibBonusAngriff () int16 {
	return (*g).bonusa
}

// NOTWENDIGKEIT SETZEN???-------------------------------------------------------------	
	// Eff.: Der Bonus zum Angriffswert des Geländes beträgt nun n. 
func (g *data) SetzeBonusAngriff (n int16) {
	(*g).bonusa = n
}


	// Erg.: Der Bonus zum Verteidigungswert des Geländes ist geliefert.
func (g *data) GibBonusVerteidigung () int16 {
	return (*g).bonusv
}

// NOTWENDIGKEIT SETZEN???-------------------------------------------------------------	
	// Eff.: Der Bonus zum Verteidigungswert des Geländes beträgt nun n. 
func (g *data) SetzeBonusVerteidigung (n int16) {
	(*g).bonusv = n
}


	// Erg.: Der Bonus zum Bewegungsswert des Geländes ist geliefert.
func (g *data) GibMalusBewegung () int16 {
	return (*g).malusb
}

// NOTWENDIGKEIT SETZEN???-------------------------------------------------------------	
	// Eff.: Der Bonus zum Bewegungsswert des Geländes beträgt nun n. 
func (g *data) SetzeMalusBewegung (n int16) {
	(*g).malusb = n
}
	
	// Erg.: Es ist true geliefert, gdw. das Geländefeld belegt ist, andernfalls ist false geliefert.
func (g *data) GibBelegung () bool {
	return (*g).belegt
}
	
	// Eff.: Die Belegung des Geländefeldes ist je nach übergebenem Wert belegt oder nicht. 
	//		 Bei Übergabe des von true ist es belegt, bei Übergabe von false
	//		 ist es nicht belegt.
func (g *data) SetzeBelegung (b bool) {
	(*g).belegt=b
}

	// Erg.: 
func (g *data) GibTyp () string {
return (*g).typ
}

func (g *data) SetzeTyp (t string) {
	g.SetzeFarbe(untergrundtab[t].r+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zr))),untergrundtab[t].g+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zg))),untergrundtab[t].b+uint8(zufallszahlen.Zufallszahl(0,int64(untergrundtab[t].zb))))
	(*g).bonusa = untergrundtab[t].bonusa
	(*g).bonusv = untergrundtab[t].bonusv
	(*g).malusb = untergrundtab[t].malusb
	(*g).typ=t
}

func (g *data) Zeichnen () {
	if !gfx.FensterOffen() {
		gfx.Fenster(1000,600)
	}
	var r,gr,b uint8					// Variablen für Farbwerte
	var x,y uint16						// Variablen für Koordinaten
	r = untergrundtab[(*g).typ].r
	gr = untergrundtab[(*g).typ].g
	b = untergrundtab[(*g).typ].b		// Zuweisung der r,g,b Werte des Geländes
	x,y = (g).GibKoordinaten()  		// Abfragen der x,y-Koordinaten des Geländes
	if (*g).typ=="Wasser" {				// Wasser wird immer neu gezeichnet (Effekt der Wasseränderung)
		zufallszahlen.Randomisieren()
		gfx.Stiftfarbe (r+uint8(zufallszahlen.Zufallszahl(0,5)),gr+uint8(zufallszahlen.Zufallszahl(0,10)),b+uint8(zufallszahlen.Zufallszahl(0,20)))	
	}else{
		gfx.Stiftfarbe (r,gr,b)				// Festlegen der Geländefarbe
	}
												
// Zeichnen des Geländes
	if g.GibGewaehlt() {
		gfx.Stiftfarbe (untergrundtab[(*g).typ].rg,untergrundtab[(*g).typ].gg,untergrundtab[(*g).typ].bg)
		gfx.Vollrechteck (x*50,y*50,50,50)
		gfx.Stiftfarbe (r,gr,b)
		gfx.Vollrechteck (x*50+5,y*50+5,40,40)
	}else{
		gfx.Vollrechteck (x*50,y*50,50,50)	
	}
}

func (g *data) Kodieren () (b []byte) {
	b = make([]byte,0)
	var x,y uint16
	x,y = g.GibKoordinaten()
	b = append(b,byte(x/256),byte(x%256),byte(y/256),byte(y%256)) // Kodieren der Koordinaten
	if (*g).belegt {
		b = append(b,byte(1)) // Kodieren der Belegung wenn true
	} else {
		b = append(b,byte(0)) // Kodieren der Belegung wenn false
	}	
	var stringlaenge uint8
	// Kodieren der Anzahl an Bytes des Strings
	stringlaenge=uint8(len((*g).typ))					
	b = append(b,stringlaenge)
	// Kodieren des Geländetyps
	for _,w := range (*g).typ {								
		b = append(b,byte(w))
	}
	return
}

func (g *data) Dekodieren (b []byte) {
	g.SetzeKoordinaten(uint16 (b[0])*256+uint16(b[1]),uint16(b[2])*256+uint16(b[3])) // Dekodieren der Koordinaten
	if b[4]==1 {
		(*g).belegt = true
	} else {
		(*g).belegt = false
	}
	var stringlaenge int = int(b[5])
	// Dekodieren des Geländetyps
	var stringtyp string
	for i:=6;i<int(stringlaenge)+6;i++ {
		stringtyp=stringtyp+string(b[i])
	}
	g.SetzeTyp(stringtyp)
}
