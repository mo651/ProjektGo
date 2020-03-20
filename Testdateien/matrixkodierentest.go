// Programm zum Austesten der Kordierung und Dekodierung einer Matrix
package main

func main () {
	var test string = "Hallö"
	var testmatrix [][]uint8 = make ([][]uint8,5)
	var zaehler uint8
	for y:=0;y<5;y++ {
		var input []uint8 = make ([]uint8,6)
		for x:=0;x<6;x++ {
			input[x]=zaehler
			zaehler++
		}
		print ("\nNeue Reihe:")
		for _,w:=range input {
			print (w)
		}
		testmatrix[y]=input
		print("\n")
		print("aktuelle Matrix")
		for _,w:=range testmatrix {
			for _,w2:=range w{
				print (w2)
			}
		}
	}
	print("\n")
	println (test)
	println (len(test))
	println (testmatrix)
	// Kodieren einer Matrix
	//Deklaration des Feldes
	var b []byte = make([]byte,0)
	//Durchlaufen der Zeilen
	for _,y := range testmatrix {	
		//Durchlaufen der einzelnen Felder					
		for _,x := range y {
			//Größe des Feldes ermitteln...
			var xalsbyteslice []byte = make ([]byte,0)
			xalsbyteslice=append(xalsbyteslice,x)
			//Ermitteln der Slicegröße
			var feldgröße uint8 = uint8(len(xalsbyteslice))
			//...und anhängen
			b = append(b,byte(len(string(feldgröße))))	
			//Anhängen des kodierten Feldes
			b = append(b,byte(x))							
		}
	}
	for _,y:=range b {
		print (y," ")
	}
	println("\nLänge des Bytefeldes: ",len(b))
	println("\nStarte mit Erstellung einer neuen Matrix")
	// Dekodieren einer Matrix
	// Erstellen der Zielmatrix
	var testmatrix2 [][]uint8 = make([][]uint8,5)
	for y:=0;y<5;y++{
		var input []uint8 = make ([]uint8,6)
		//Erstellen der Zeilen...hier noch notwendig, da Matrix neu erstellt wird
		testmatrix2 [y] = input
	}
	println ("Neues Testfeld erfolgreich erstellt\n")
	// Hilfsvariable zum Durchlaufen des Bytefeldes
	var aktindex int 
	// Durchlaufen der Reihen
	for y:=0;y<5;y++{
		println("Starte mit Reihe: ",y)
		// Durchlaufen der einzelnen Felder
		for x:=0;x<6;x++ {
			//Ermitteln des Startindex des nächsten Feldinhalts
			println("Aktueller Index: ",aktindex)
			println("Größe des nächsten Feldes",b[aktindex])
			var feldgroesse int =int(b[aktindex])
			aktindex++
			println("Aktueller Index: ",aktindex)
			print("Einfügen des Feldinhalts: ",b[aktindex])
			//Dekodieren des Feldinhalts und einfügen in der Zielmatrix
			println(" an der Stelle x: ",x," y: ",y)
			testmatrix2[y][x]=uint8(b[aktindex])
			//Erhöhen des aktIndex, auf Bytefeld nach dem kodierten Feldinhalt
			aktindex= aktindex+feldgroesse		
		}
		println (y,". Reihe erfolgreich dekodiert")
	}
	print("Dekodierte Matrix:")
	for _,w:=range testmatrix2 {
		for _,w2:=range w{
			print (w2)
		}
	}	
}
