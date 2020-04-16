package spielablauf

type Spielablauf interface {


// Vor.: ---
// Erg.: Der aktuelle Spieler ist geliefert. Spieler sind 0 oder 1. 
GibSpieler() uint8

// Vor.: 
// Eff.:
SetzeSpieler(s uint8)

// Vor.: 
// Erg.: 
GibModus() uint8

// Vor.: 
// Eff.:
SetzeModus (m uint8) 

/*// Vor.:
// Eff.: 
ZeichneBewegungsReichweite () 
--> Als Methode oder Hilfsfunktion??
// Vor.:
// Eff.: 
ZeichneAngriffsReichweite () 
--> Als Methode oder Hilfsfunktion??*/
// Vor.: 
// Eff.: 
Angreifen (x,y uint16) 

Bewegen (x,y uint16)

Zeichnen()
}









