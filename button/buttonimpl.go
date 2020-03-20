package button

import ( "gfx" )

type data struct {
	x,y uint16
	hoehe,breite uint16
	rr,rg,rb uint8
	r,g,b uint8
	tr,tg,tb uint8
	Text string
	ausrichtungh rune
	ausrichtungv rune
}
 	

func New (x,y uint16, breite,hoehe uint16, rr,rg,rb uint8, r,g,b uint8, tr,tg,tb uint8, Text string, ausrichtungh,ausrichtungv rune) *data {
	var neu = new(data)
	(*neu).x=x
	(*neu).y=y
	(*neu).hoehe=hoehe
	(*neu).breite=breite
	(*neu).r=r
	(*neu).g=g
	(*neu).b=b
	(*neu).rr=rr
	(*neu).rg=rg
	(*neu).rb=rb
	(*neu).tr=tr
	(*neu).tg=tg
	(*neu).tb=tb
	(*neu).ausrichtungh=ausrichtungh
	(*neu).ausrichtungv=ausrichtungv	
	(*neu).Text=Text	
	return neu
}
	
	//Vor.: -
	//Erg.: Der Text des Buttons ist auf "Text" geändert.
func (b *data)	SetzeText (Text string){
	(*b).Text=Text
}
	
	//Vor.: -
	//Erg.: Der aktuelle Text des Buttons ist geliefert.
func (b *data)	GibText () string {
	return (*b).Text
}

//Vor.: -
	//Erg.: Die horizontale Ausrichtung des Textes ist geändert.
	//		'o':=oben
	//		'm':=mitte
	//		'u':=unten
func (b *data) SetzeAusruchtungHorizontal (ausrichtungh rune){
	(*b).ausrichtungh=ausrichtungh
}
	
	//Vor.: -
	//Erg.: Die horizontale Ausrichtung des Textes ist geliefert.
func (b *data) GibAusruchtungHorizontal () rune{
	return (*b).ausrichtungh
	}
	
	//Vor.: -
	//Erg.: Die vertikale Ausrichtung des Textes ist geändert.
	//		'l':=links
	//		'm':=mitte
	//		'r':=rechts
func (b *data) SetzeAusruchtungVertikal (ausrichtungv rune){
	(*b).ausrichtungv=ausrichtungv
}
	
	//Vor.: -
	//Erg.: Die vertikale Ausrichtung des Textes ist geliefert.
func (b *data) GibAusruchtungVertikal () rune{
	return (*b).ausrichtungv
}
	
	//Vor.: -
	//Erg.: Die Textfarbe ist auf "tr (rot), tg (gruen), tb (blau)" geändert.
func (b *data) SetzeTextfarbe (tr,tg,tb uint8){
	(*b).tr=tr
	(*b).tg=tg
	(*b).tb=tb
}
	
	//Vor.: -
	//Erg.: Die aktuelle Textfarbe "tr (rot), tg (gruen), tb (blau)" ist geliefert.
func (b *data) GibTextfarbe () (tr,tg,tb uint8){
	return (*b).tr,(*b).tg,(*b).tb
}	
	//Vor.: -
	//Erg.: Die Rahmenfarbe ist auf "rr (rot), rg (gruen), rb (blau)" geändert.
func (b *data) SetzeRamenfarbe (rr,rg,rb uint8){
	(*b).rr=rr
	(*b).rg=rg
	(*b).rb=rb
}
	
	//Vor.: -
	//Erg.: Die aktuelle Rahmenfarbe "tr (rot), tg (gruen), tb (blau)" ist geliefert.
func (b *data) GibRahmenfarbe () (rr,rg,rb uint8){
	return (*b).rr,(*b).rg,(*b).rb
}
	
	//Vor.: -
	//Erg.: Die Hintergrundfarbe ist auf "r (rot), g (gruen), b (blau)" geändert.
func (bu *data) SetzeHintergundfarbe (r,g,b uint8){
	(*bu).r=r
	(*bu).g=g
	(*bu).b=b
}
	
	//Vor.: -
	//Erg.: Die aktuelle Hintergrundfarbe "r (rot), g (gruen), b (blau)" ist geliefert.
func (bu *data) GibHintergrundfarbe () (r,g,b uint8){
	return (*bu).r,(*bu).g,(*bu).b
}
	
	//Vor.: -
	//Erg.: Die Position des Buttons ist auf "x,y" geändert.
func (b *data) SetzePosition (x,y uint16){
	(*b).x=x
	(*b).y=y
}


	//Vor.: -
	//Erg.: Die aktuelle Position des Buttons ist geliefert.
func (b *data) GibPosition () (x,y uint16){
	return (*b).x,(*b).y
}
	
	//Vor.: -
	//Erg.: Die aktuelle X Position des Buttons ist geliefert.
func (b *data) GibXPosition () (x uint16){
	return (*b).x
}

	//Vor.: -
	//Erg.: Die aktuelle Y Position des Buttons ist geliefert.
func (b *data) GibYPosition () (y uint16){
	return (*b).y
}
	
	//Vor.: -
	//Erg.: Die Breite des Buttons ist auf "b" geändert.
func (b *data) SetzeBreite (breite uint16){
	(*b).breite=breite
}
	
	//Vor.: -
	//Erg.: Die Breite des Buttons ist geliefert.
func (b *data) GibBreite () (breite uint16){
	return (*b).breite
}
		
	//Vor.: -
	//Erg.: Die Hoehe des Buttons ist auf "h" geändert.
func (b *data)SetzeHoehe (hoehe uint16){
	(*b).hoehe=hoehe
}
	
	//Vor.: -
	//Erg.: Die Hoehe des Buttons ist geliefert.
func (b *data) GibHoehe () (hoehe uint16){
	return (*b).hoehe
}
			
	//Vor.: -
	//Erg.: Der Button wird im gfx-Fenster gezeichnet.
func (b *data)	Zeichnen (){
	if !gfx.FensterOffen(){ 
		gfx.Fenster(800,600)
		}
	gfx.Stiftfarbe((*b).rr,(*b).rg,(*b).rb)
	gfx.Vollrechteck((*b).x,(*b).y,(*b).breite,(*b).hoehe)
	gfx.Stiftfarbe((*b).r,(*b).g,(*b).b)
	gfx.Vollrechteck((*b).x+uint16(5),(*b).y+uint16(5),(*b).breite-uint16(10),(*b).hoehe-uint16(10))
	gfx.Stiftfarbe((*b).tr,(*b).tg,(*b).tb)
	switch (*b).ausrichtungh{
	case 'o':
		switch (*b).ausrichtungv{
		case 'l':
		gfx.Schreibe((*b).x+uint16(10),(*b).y+uint16(10),(*b).Text)
		case 'm':
		gfx.Schreibe(((*b).x+((*b).breite/2)-uint16(len((*b).Text))*uint16(4)),(*b).y+uint16(10),(*b).Text)
		case 'r':
		gfx.Schreibe(((*b).x+((*b).breite)-uint16(len((*b).Text))*uint16(8))-uint16(10),(*b).y+uint16(10),(*b).Text)
		}
	case 'm':
		switch (*b).ausrichtungv{
		case 'l':
		gfx.Schreibe((*b).x+uint16(10),(*b).y+((*b).hoehe/2)-uint16(4),(*b).Text)
		case 'm':
		gfx.Schreibe(((*b).x+((*b).breite/2)-uint16(len((*b).Text))*uint16(4)),(*b).y+((*b).hoehe/2)-uint16(4),(*b).Text)
		case 'r':
		gfx.Schreibe(((*b).x+((*b).breite)-uint16(len((*b).Text))*uint16(8))-uint16(10),(*b).y+((*b).hoehe/2)-uint16(4),(*b).Text)
		}
	case 'u':
		switch (*b).ausrichtungv{
		case 'l':
		gfx.Schreibe((*b).x+uint16(10),(*b).y+(*b).hoehe-uint16(15),(*b).Text)
		case 'm':
		gfx.Schreibe(((*b).x+((*b).breite/2)-uint16(len((*b).Text))*uint16(4)),(*b).y+(*b).hoehe-uint16(15),(*b).Text)
		case 'r':
		gfx.Schreibe(((*b).x+((*b).breite)-uint16(len((*b).Text))*uint16(8))-uint16(10),(*b).y+(*b).hoehe-uint16(15),(*b).Text)
		}		
	}
}
	
	//Vor.: -
	//Erg.: Der Button ist durch ein Rechteck der Untergrundfarbe ersetzt (und damit faktisch entfernt).
func (b *data)	Entfernen (ufr,ufg,ufb uint8){
	if !gfx.FensterOffen(){ 
		gfx.Fenster(800,600)
		}	
	gfx.Stiftfarbe(ufr,ufg,ufb)
gfx.Vollrechteck((*b).x,(*b).y,(*b).breite,(*b).hoehe)	
	}
