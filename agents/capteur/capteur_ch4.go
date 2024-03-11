package capteur

import (
	"TP1_INFO801/global"
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

var valeur_CH4 = rand.Float64() * 100

func get_valeur_ch4(etatVentilateur string) float64 {
	switch etatVentilateur {
	case "activé":
		valeur_CH4 -= rand.Float64() * 5
		if valeur_CH4 < 0.0 {
			valeur_CH4 = 0.0
		}
		return valeur_CH4
	default:
		valeur_CH4 += rand.Float64() * 3
		return valeur_CH4
	}
}

func Capteur_ch4(ts *Space, etatVentilateur *string) {
	valeur_CH4 := get_valeur_ch4(*etatVentilateur)
	ts.Put("niveau_ch4", valeur_CH4)
	fmt.Print("Capteur_ch4: ", valeur_CH4, "\n")
	time.Sleep(global.WaitTime)
	Capteur_ch4(ts, etatVentilateur)
}
