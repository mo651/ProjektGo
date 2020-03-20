package main
//XXX Speicher fehlt noch 
import ("../gelaende";"gfx";"strconv";"time")

func maussteuerung (s []gelaende.Gelände){
for{	
		taste, status, mausX, mausY:= gfx.MausLesen1 ()
		if taste == 1 && status == 1 { //linke Maustaste gerade gedrückt
			if mausX<1000 && mausY<600 {
				var x,y uint16
				x= mausX/50
				y= mausY/50
				for _,w := range s {
					xs,ys:=w.GibKoordinaten()
					w.SetzeGewaehlt(false)
					if x==xs && y==ys {
					w.SetzeGewaehlt(true)
					}
				}	
			}

		}
	}
}

		
		
		
// ein Komentar als Test!	




func main () {
//Aufbau des Fensters:	
	gfx.Fenster(1200,700)
 
	
	//gfx.Vollrechteck(1050,600,100,50)
	//gfx.Stiftfarbe(255,255,255)
	//gfx.Schreibe (1085, 620, "save")
	
	
	
	var spielfeld []gelaende.Gelände
	spielfeld = make ([]gelaende.Gelände,0)
		for i:=uint16(0);i<20;i++ {
			for j:=uint16(0);j<12;j++ {
				var neuesgelände gelaende.Gelände
				neuesgelände = gelaende.New (i,j,"Wasser") //man beachte, dass Wasser ist "animiert", ob das im feritgen Spiel auch funktionier kann ich nicht sagen
				spielfeld=append(spielfeld,neuesgelände)
				}	
			}
	 for i,_:=range spielfeld{ //Ein Wald entsteht
		if (i==0)||(i==1)||(i==2)||(i==3)||(i==12)||(i==13)||(i==14)||(i==24)||(i==25)||(i==36)||(i==37)||(i==48){
		spielfeld[i].SetzeTyp("Wald")
		}
	}
		spielfeld[18].SetzeTyp("Wueste")
		spielfeld[19].SetzeTyp("Wueste")
		spielfeld[44].SetzeTyp("Berg")
		spielfeld[45].SetzeTyp("Berg")
	 
	go maussteuerung (spielfeld)
	for {
	gfx.UpdateAus ()	
	gfx.Stiftfarbe(255,255,255)
	gfx.Vollrechteck(1000,0,200,200)	
	gfx.Stiftfarbe(0,0,0)
	gfx.Schreibe (1030, 70, "Bonus Angriff")
	gfx.Schreibe (1030, 100, "Bonus Bewegung")
	gfx.Schreibe (1030, 130, "Bonus Verteidigung")
		for _,w:=range spielfeld {
			w.Zeichnen()
			if w.GibGewaehlt(){
						//gfx.Sperren()	

						//hier muss ein Schloss hin, wenn die Farbe nicht random sein soll
						gfx.Stiftfarbe(0,0,0)
						gfx.Schreibe (1050, 50, w.GibTyp())
						gfx.Schreibe (1050, 80, strconv.Itoa(int(w.GibBonusAngriff())))
						gfx.Schreibe (1050, 110, strconv.Itoa(int(w.GibBonusBewegung())))
						gfx.Schreibe (1050, 140, strconv.Itoa(int(w.GibBonusVerteidigung())))
						//gfx.Entsperren()
				}
		}
		
		
		
		gfx.Stiftfarbe(0,0,0)
		for i:=uint16(0);i<20;i++ {
			gfx.Linie(i*50,0,i*50,599)
		}
		for i:=uint16(0);i<12;i++ {
			gfx.Linie(0,i*50,999,i*50)
		}
		gfx.UpdateAn ()
		time.Sleep(time.Duration(90 * 1e5)) // Immer ca. 100 FPS !!
	}	
}
