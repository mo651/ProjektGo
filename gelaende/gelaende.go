package gelaende

import "../objekte"
// XXX NEW Funktion fehlt bzw. ist falsch
// KONSTRUKTOR:
// Vor.: --
// Erg.: Ein neues Geländefeld mit den Feldkoordinaten (x,y),
//       mit der Farbe gemäß den RGB-Werten r,g,b und einem Bonuswert für Angriff ist geliefert.
// New (x,y uint16, typ string) *data 
// Das zurückgegebene Objekt erfüllt das Interface Gelände.

// -------- Verbale Beschreibung ergänzen, was ein Objekt konkret ist???-------------------

type Gelände interface {
	
	objekte.Objekt
	
	// Vor.: --
	// Erg.: Der Bonus zum Angriffswert des Geländes ist geliefert.
	GibBonusAngriff () int16
	
	// Vor.: --
	// Eff.: Der Bonus zum Angriffwert des Geländes beträgt nun n. 
	SetzeBonusAngriff (n int16)
	
	// Vor.: --
	// Erg.: Der Bonus zum Verteidigungswert des Geländes ist geliefert.
	GibBonusVerteidigung () int16
	
	// Vor.: --
	// Eff.: Der Bonus zum Verteidigungswert des Geländes beträgt nun n. 
	SetzeBonusVerteidigung (n int16)
	
	// Vor.: --
	// Erg.: Der Bonus zum Bewegungswert des Geländes ist geliefert.
	GibBonusBewegung () int16
	
	// Vor.: --
	// Eff.: Der Bonus zum Bewegungswert des Geländes beträgt nun n. 
	SetzeBonusBewegung (n int16)
	
	// Vor.: --
	// Erg.: Es ist true geliefert, gdw. das Geländefeld belegt ist, andernfalls ist false geliefert.
	GibBelegung () bool 
	
	// Vor.: --
	// Eff.: Die Belegung des Geländefeldes ist je nach übergebenem Wert belegt oder nicht. 
	//		 Bei Übergabe des von true ist es belegt, bei Übergabe von false
	//		 ist es nicht belegt.
	SetzeBelegung (b bool) 
	
		// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
	GibTyp () string
	
	// Vor.: ---
	// Eff.: XXX-------------------------------------------------------------------------------------------------
	SetzeTyp (t string)
	// SetztTyp nicht notwendig, da Typ nicht veränderbar????????????????????????????????????????
	
	// Vor.: ---
	// Erg.: Eine Folge von Bytes ist geliefert, die die aufrufende Instanz serialisiert darstellt.
	Kodieren () (b []byte)
	
	// Vor.: Der übergebene Slice b entspricht einem kodierten Gelände.
	// Eff.: Die aufrufende Instanz hat nun genau diejenige Eigenschaft, sodass 'Kodieren ()' exakt
	//		 b ergeben würde. Alle vorherigen Eigenschaften des Geländes hat er nicht mehr.
	Dekodieren (b []byte) 
}
