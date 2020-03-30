package spielablauf

import "../spielfeld"

type data struct {
  aktspieler uint8    // 0=Spieler 1 ; 1=Spieler2
  modus uint8         // 0=neutral ; 1=bewegen ; 2=angreifen // speicher und laden? Optionsmenü? erweiterbar durch bauen,Itemeinsatz etc...

}

