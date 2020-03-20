package objekte

import "gfx"

type data struct {
	x,y uint16				// spätere Feldkoordinate im Spielfeld
	r,g,b uint8				// Farbe des Objektes ; später einfügen eines Bildes --> bild string
	sichtbar bool			// Das Objekt kann u.U. nicht sichtbar sein
	gewaehlt bool 			// Das Objekt kann aus- oder abgewählt werden
	geschwindigkeit uint16	// Das Objekt hat eine Geschwindigkeit (Anzahl Felder) mit der es sich bewegen kann ---->>>>>>> FÜR SPÄTERE ANIMATION???
}

func New (x,y uint16, r,g,b uint8) *data {
	var neu *data = new(data)
	(*neu).x=x
	(*neu).y=y
	(*neu).r=r
	(*neu).g=g
	(*neu).b=b
	(*neu).sichtbar=true
	(*neu).gewaehlt=false 		// eigentlich standardmäßig auf false gesetzt
	(*neu).geschwindigkeit=0	// eigentlich standardmäßig auf 0 gesetzt ---->>>>>>> FÜR SPÄTERE ANIMATION???
	return neu
}

// Vor.: ---
	// Erg.: die Feldkoordinaten x und y
func (o *data) GibKoordinaten () (uint16,uint16) {
	return (*o).x,(*o).y
}

	// Vor.: ---
	// Eff.: Die Feldkoordinaten des Objekts sind auf (x,y) gesetzt.
func (o *data) SetzeKoordinaten (x,y uint16) {
	(*o).x=x
	(*o).y=y
}
	
	// Vor.: ---
	// Erg.: Die Farbkomponenten rot, grün und blau des Objektes sind geliefert.
func (o *data) GibFarbe () (uint8, uint8, uint8) {
	return (*o).r,(*o).g,(*o).b
}
	
// NOTWENDIGKEIT GEGEBEN???-------------------------------------------------------------
	// Vor.: ---
	// Eff.: Die Farbe des Objekts sind nach dem RGB-Modell gesetzt.
func (o *data) SetzeFarbe (r,g,b uint8) {
	(*o).r=r
	(*o).g=g
	(*o).b=b
}
	
	// Vor.: ---
	// Erg.: Der Abstand des Objekts zu einem zweitem Objekt
	//       ist geliefert. Mit Abstand ist der Abstand in Feldern in x- und 
	//       y-Richtung gemeint. Bei Überschneidungen ist der Abstand 0 geliefert.
func (o *data) Distanz (o2 Objekt) (x,y int16) {
	x2,y2:=o2.GibKoordinaten()
	return (-1)*int16((*o).x)+int16(x2),(-1)*int16((*o).y)+int16(y2)	// hier mit int-Werten --> noch überprüfen ob Richtung relevant sein wird-----------
}
	
	// Vor.: ---
	// Eff.: Das Objekt ist im gfx-Grafikfenster gezeichnet.
func (o *data) Zeichnen () {
	if !gfx.FensterOffen() {
		gfx.Fenster(1000,600)
	}
	//Lassen wir erstmal noch offen --> an welcher Stelle muss gezeichnet werden????? Größe und Position abhängig von Spielfeld----------------------------
}
	
	// Vor.: ---
	// Erg.: true ist geliefert, gdw. das Objekt sichtbar ist. Andernfalls ist 
	//		 false geliefert.
func (o *data) IstSichtbar () bool {
	return (*o).sichtbar
}
	
	// Vor.: ---
	// Eff.: Die Sichtbarkeit des Objektes ist je nach übergebenem Wert geändert. 
	//		 Bei Übergabe des von true ist es sichtbar, bei Übergabe von false
	//		 ist es nicht sichtbar.
func (o *data) SetzeSichtbar (b bool) {
	(*o).sichtbar=b
}
	
	// Vor.: ---
	// Erg.: true ist geliefert, gdw. das Objekt ausgewählt ist. Andernfalls ist 
	//		 false geliefert.
func (o *data) GibGewaehlt () bool {
	return (*o).gewaehlt
}
	
	// Vor.: ---
	// Eff.: Das Objektes ist je nach übergebenem Wert gewählt oder nicht. 
	//		 Bei Übergabe des von true ist es gewählt, bei Übergabe von false
	//		 ist es nicht gewählt.
func (o *data) SetzeGewaehlt (b bool) {
	(*o).gewaehlt=b
}

	// Vor.: ---
	// Erg.: Die Geschwindigkeit des Objektes ist geliefert.
func (o *data) GibGeschwindigkeit () uint16 {
	return (*o).geschwindigkeit
}
	
	// Vor.: --
	// Eff.: Die Geschindigkeit des Objektes ist auf s gesetzt.
func (o *data) SetzeGeschwindigkeit (s uint16) {
	(*o).geschwindigkeit=s
}
