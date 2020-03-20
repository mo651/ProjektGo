package objekte

// KONSTRUKTOR:
// Vor.: --
// Erg.: Ein neues Objekt mit den Feldkoordinaten (x,y) und
//       mit der Farbe gemäß den RGB-Werten r,g,b ist geliefert.
// New (x,y uint16, r,g,b uint8) *data 
// Das zurückgegebene Objekt erfüllt das Interface Objekt.

// -------- Verbale Beschreibung ergänzen, was ein Objekt konkret ist???-------------------

type Objekt interface {
	// Vor.: ---
	// Erg.: die Feldkoordinaten x und y
	GibKoordinaten () (uint16,uint16) 
	
	// Vor.: ---
	// Eff.: Die Feldkoordinaten des Objekts sind auf (x,y) gesetzt.
	SetzeKoordinaten (x,y uint16)
	
	// Vor.: ---
	// Erg.: Die Farbkomponenten rot, grün und blau des Objektes sind geliefert.
	GibFarbe () (uint8, uint8, uint8)
	
// NOTWENDIGKEIT GEGEBEN???-------------------------------------------------------------
	// Vor.: ---
	// Eff.: Die Farbe des Objekts sind nach dem RGB-Modell gesetzt.
	SetzeFarbe (r,g,b uint8)
	
	// Vor.: ---
	// Erg.: Der Abstand des Objekts zu einem zweitem Objekt
	//       ist geliefert. Mit Abstand ist der Abstand in Feldern in x- und 
	//       y-Richtung gemeint. Bei Überschneidungen ist der Abstand 0 geliefert.
	Distanz (o2 Objekt) (x,y int16)
	
	// Vor.: ---
	// Eff.: Das Objekt ist im gfx-Grafikfenster gezeichnet.
	Zeichnen ()
	
	// Vor.: ---
	// Erg.: true ist geliefert, gdw. das Objekt sichtbar ist. Andernfalls ist 
	//		 false geliefert.
	IstSichtbar () bool
	
	// Vor.: ---
	// Eff.: Die Sichtbarkeit des Objektes ist je nach übergebenem Wert geändert. 
	//		 Bei Übergabe des von true ist es sichtbar, bei Übergabe von false
	//		 ist es nicht sichtbar.
	SetzeSichtbar (b bool)
	
	// Vor.: ---
	// Erg.: true ist geliefert, gdw. das Objekt ausgewählt ist. Andernfalls ist 
	//		 false geliefert.
	GibGewaehlt () bool
	
	// Vor.: ---
	// Eff.: Das Objektes ist je nach übergebenem Wert gewählt oder nicht. 
	//		 Bei Übergabe des von true ist es gewählt, bei Übergabe von false
	//		 ist es nicht gewählt.
	SetzeGewaehlt (b bool)
	
	// Vor.: ---
	// Erg.: Die Geschwindigkeit des Objektes ist geliefert.
	GibGeschwindigkeit () uint16
	
	// Vor.: --
	// Eff.: Die Geschindigkeit des Objektes ist auf s gesetzt.
	SetzeGeschwindigkeit (s uint16) 
}
