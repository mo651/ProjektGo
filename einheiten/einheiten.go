// Autor: Friederike Richter und Moritz Raake
// Datum: 11.03.20
// Zweck: Softwarepraktikum - Definition der Einheiten-Schnittstelle

package einheiten

import "../objekte"

// KONSTRUKTOR:
// Vor.: --
// Erg.: Eine neue Einheit mit den Feldkoordinaten (x,y),
//       mit der Farbe gemäß den RGB-Werten r,g,b, dem Typ t, der Angriffsstärke a,
//		 der Verteidigungsstärke v, der Reichweit rei, dem Radius rad und der 
//		 Geschwindigkeit geschw ist geliefert.
// New (x,y uint16, r,g,b uint8, t string, a,v,reia,reib,rad,geschw uint16) *data <--Muss überarbeitet werdex XXX--------
// Das zurückgegebene Objekt erfüllt das Interface Gelände.

// -------- Verbale Beschreibung ergänzen, was ein Objekt konkret ist???-------------------

type Einheit interface {
	
	objekte.Objekt
	
	// Vor.: ---
	// Erg.: Der Angriffswert der Einheit ist geliefert.
	GibAngriffswert () uint16
	
	// Vor.: ---
	// Eff.: Der Angriffswert der Einheit beträgt nun a. 
	SetzeAngriffswert (a uint16)
	
	// Vor.: ---
	// Erg.: Der Verteidigungswert der Einheit ist geliefert.
	GibVerteidigungswert () uint16 
	
	// Vor.: ---
	// Eff.: Der Verteidigungswert der Einheit beträgt nun v. 
	SetzeVerteidigungswert (v uint16) 
	
	// Vor.: ---
	// Erg.: Die Angriffsreichweite (Anzahl der Felder) der Einheit ist geliefert.
	GibAngriffsReichweite () uint16
	
	// Vor.: ---
	// Eff.: Die Angriffsreichweite (Anzahl der Felder) der Einheit ist auf rei gesetzt.
	SetzeAngriffsReichweite (reia uint16)
	
	// Vor.: ---
	// Erg.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist geliefert.
	GibBewegungsReichweite () uint16
	
	// Vor.: ---
	// Eff.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist auf rei gesetzt.
	SetzeBewegungsReichweite (reib uint16)
	
	// Vor.: ---
	// Eff.: Die verbleibende Lebensenergie der Einheit ist geliefert.
	GibEnergie () uint16
	
	// Vor.: 
	// Eff.: Die Lebensernergie der Einheit ist auf en gesetzt.
	SetzeEnergie (en uint16)
/*	
//XXXXXXXX		
	// Vor.: ---
	// Erg.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist geliefert.
	Angreifen (e2 einheit) 

	// Vor.: ---
	// Erg.: Die Bewegungsreichweite (Anzahl der Felder) der Einheit ist geliefert.
	//Verteidigen (e2 einheit) 
	//Wie soll das gehen? über setze und gib 	
//XXXXXXXXX
*/

	// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
	GibTyp () string

		// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
	GibSpieler () byte
	
		// Vor.: ---
	// Erg.:  XXX-------------------------------------------------------------------------------------------------
	SetzeSpieler (byte) 

	// Vor.: ---
	// Eff.: XXX-------------------------------------------------------------------------------------------------
	SetzeTyp (t string)
	// SetztTyp nicht notwendig, da Typ nicht veränderbar????????????????????????????????????????
	
	// Vor.: ---
	// Eff.: XXX-------------------------------------------------------------------------------------------------
	Kodieren () []byte
	
	// Vor.: ---
	// Eff.: XXX-------------------------------------------------------------------------------------------------
	Dekodieren (b []byte)
	
	// Vor.: ---
	// Eff.: XXX MUSS NOCH ERGÄNZT WERDEN --> MORITZ WILL DAS SO!!-------------------------------------------------------
	//Zeichnen () // Benötigt man Methode Zeichnen??? Oder übernimmt dies das Spielfeld???----> Ja benötigt man----------------------------------------
}

