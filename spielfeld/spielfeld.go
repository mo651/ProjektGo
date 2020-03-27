package spielfeld

import ("../gelaende";"../einheiten" /*; "../items"*/)

// XXX NEW Funktion fehlt bzw. ist falsch
// KONSTRUKTOR:
// Vor.: --
// Erg.: Ein neues Spielfeld mit x Kästchen Breite, y Kästchen Höhe, quatratischen Kästchen mit 
//		 einer Seitenlänge von groesse Pixeln und dem Dateinamen dateiname ist geliefer.liefert.
// New (x,y,groesse uint16, dateiname string) *data 
// Das zurückgegebene Objekt erfüllt das Interface Spielfeld.

// XXXXXXXXXXX Verbale Beschreibung ergänzen, was ein Spielfeld konkret istXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX



type Spielfeld interface {

	
	//XXXXX GETTER UND SETTER FÜR SPIELFELDGRÖßE NOTWENDIG?????XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	
	// Vor.: ---
	// Erg.: Ein Verweis auf das an dem übergebenen Koordinaten befindliche Gelände ist geliefert.	
	GibGelände (x,y uint16) gelaende.Gelände
	
	// Vor.: ---
	// Eff.: Das übergebene Gelände befindet sich nun an den übergebenen Koordinaten. Das vorherig dort
	//		 befindliche Gelände ist entfernt.
	SetzeGelände (x,y uint16, gelände gelaende.Gelände) 
	
	// Vor.: Das Geländefeld muss mit einer Einheit besetzt sein.
	// Erg.: Ein Verweis auf die auf dem übergebenen Geländefeld befindlichen Instanz ist geliefert.
	GibEinheit (x,y uint16) (einheit einheiten.Einheit)
	
	// Vor.: Das Geländefeld mit den übergebenen Koordinaten darf nicht besetzt sein.
	// Eff.: Auf dem Geländefeld mit den übergebenen Koordinaten befindet sich nun die übergebene
	//		 Einheit. 
	SetzeEinheit (x,y uint16, einheit einheiten.Einheit) 
	
	// Vor.: Das Geländefeld mit den übergebenen Koordinaten muss mit einer Einheit besetzt sein.
	// Eff.: Die auf dem Geländefeld mit den übergebenen Koordinaten befindliche Einheit ist entfernt.
	EntferneEinheit (x,y uint16)
	
	/*GibItem (x,y uint16) items.Item
	
	SetzeItem (x,y uint16, item items.Item) 
		
	EntferneItem (x,y uint16) */
	
	//Reichweite in zu erreichenden Feldern
	GibReichweite(einheit einheiten.Einheit)[][]bool
	
	Zeichnen ()
	
	// Vor.: ---
	// Erg.: Eine Folge von Bytes ist geliefert, die die aufrufende Instanz serialisiert darstellt.
	Kodieren () (b []byte)
	
	// Vor.: Der übergebene Slice b entspricht einem kodierten Spielfeld gleicher Größe.
	// Eff.: Die aufrufende Instanz hat nun genau diejenige Eigenschaft, sodass 'Kodieren ()' exakt
	//		 b ergeben würde. Alle vorherigen Eigenschaften des Spielfeldes hat sie nicht mehr.
	Dekodieren (b []byte)
}
	
