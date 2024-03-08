package capteur

import (
	"fmt"
	. "github.com/pspaces/gospace"
	"time"
)

func Capteur_ch4(ts *Space) {
	valeur_CH4 := 55.0 //rand.Float64() * 100
	ts.Put("niveau_ch4", valeur_CH4)
	fmt.Print("Capteur_ch4: ", valeur_CH4, "\n")
	time.Sleep(5 * time.Second)
	Capteur_ch4(ts)
}
