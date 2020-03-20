package button

 
type Button interface {
	
	// Vor.: 
	// Erg.: 
	//rr,rg,rb := Farbe Rahmen
	//r,g,b := Farbe Button
	//tr,tg,tb := Farbe Text
	// ausrichtungh := horizontale Ausrichtung: o-oben;m-mitte;u-unten
	// ausrichtungv := vertikale Ausrichtung: l-links;m-mitte;r-rechts	
	// New (x,y uint16, breite,hoehe uint16, rr,rg,rb uint8, r,g,b uint8, tr,tg,tb uint8, Text string, ausrichtungh,ausrichtungv rune) Datei 
	// entspricht dem sonst zu verwendenden New
	
	//Vor.: -
	//Erg.: Die horizontale Ausrichtung des Textes ist geändert.
	//		'o':=oben
	//		'm':=mitte
	//		'u':=unten
	SetzeAusruchtungHorizontal (ausrichtungh rune)
	
	//Vor.: -
	//Erg.: Die horizontale Ausrichtung des Textes ist geliefert.
	GibAusruchtungHorizontal () rune
	
	//Vor.: -
	//Erg.: Die vertikale Ausrichtung des Textes ist geändert.
	//		'l':=links
	//		'm':=mitte
	//		'r':=rechts
	SetzeAusruchtungVertikal (ausrichtungv rune)
	
	//Vor.: -
	//Erg.: Die vertikale Ausrichtung des Textes ist geliefert.
	GibAusruchtungVertikal () rune
	
	//Vor.: -
	//Erg.: Der Text des Buttons ist auf "Text" geändert.
	SetzeText (Text string)
	
	//Vor.: -
	//Erg.: Der aktuelle Text des Buttons ist geliefert.
	GibText () string
	
	//Vor.: -
	//Erg.: Die Textfarbe ist auf "tr (rot), tg (gruen), tb (blau)" geändert.
	SetzeTextfarbe (tr,tg,tb uint8)
	
	//Vor.: -
	//Erg.: Die aktuelle Textfarbe "tr (rot), tg (gruen), tb (blau)" ist geliefert.
	GibTextfarbe () (tr,tg,tb uint8)
	
	//Vor.: -
	//Erg.: Die Rahmenfarbe ist auf "rr (rot), rg (gruen), rb (blau)" geändert.
	SetzeRamenfarbe (rr,rg,rb uint8)
	
	//Vor.: -
	//Erg.: Die aktuelle Rahmenfarbe "tr (rot), tg (gruen), tb (blau)" ist geliefert.
	GibRahmenfarbe () (rr,rg,rb uint8)
	
	//Vor.: -
	//Erg.: Die Hintergrundfarbe ist auf "r (rot), g (gruen), b (blau)" geändert.
	SetzeHintergundfarbe (r,g,b uint8)
	
	//Vor.: -
	//Erg.: Die aktuelle Hintergrundfarbe "r (rot), g (gruen), b (blau)" ist geliefert.
	GibHintergrundfarbe () (r,g,b uint8)
	
	//Vor.: -
	//Erg.: Die Position des Buttons ist auf "x,y" geändert.
	SetzePosition (x,y uint16)
	
	//Vor.: -
	//Erg.: Die aktuelle Position des Buttons ist geliefert.
	GibPosition () (x,y uint16)

	//Vor.: -
	//Erg.: Die aktuelle X-Position des Buttons ist geliefert.
	GibXPosition () (x uint16)
	
	//Vor.: -
	//Erg.: Die aktuelle Y-Position des Buttons ist geliefert.
	GibYPosition () (y uint16)
		
	//Vor.: -
	//Erg.: Die Breite des Buttons ist auf "b" geändert.
	SetzeBreite (breite uint16)
	
	//Vor.: -
	//Erg.: Die Breite des Buttons ist geliefert.
	GibBreite () (breite uint16)
		
	//Vor.: -
	//Erg.: Die Hoehe des Buttons ist auf "h" geändert.
	SetzeHoehe (hoehe uint16)
	
	//Vor.: -
	//Erg.: Die Hoehe des Buttons ist geliefert.
	GibHoehe () (hoehe uint16)
	
	//Vor.: -
	//Erg.: Der Button wird im gfx-Fenster gezeichnet.
	Zeichnen ()
	
	//Vor.: -
	//Erg.: Der Button ist durch ein Rechteck der Untergrundfarbe ersetzt (und damit faktisch entfernt).
	Entfernen (Untergrundfarbe uint8)
	
}
