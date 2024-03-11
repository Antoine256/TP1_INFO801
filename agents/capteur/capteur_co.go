package capteur

import (
	"TP1_INFO801/global"
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

var valeur_CO = rand.Float64() * 100

func get_valeur_co(etatVentilateur string) float64 {
	switch etatVentilateur {
	case "activé":
		valeur_CO -= rand.Float64() * 5
		if valeur_CO < 0.0 {
			valeur_CO = 0.0
		}
		return valeur_CO
	default:
		valeur_CO += rand.Float64() * 3
		return valeur_CO
	}
}

func Capteur_co(ts *Space, etatVentilateur *string) {
	valeur_CO := get_valeur_co(*etatVentilateur)
	global.Add(ts, "niveau_co", valeur_CO)
	fmt.Print("Capteur_co: ", valeur_CO, "\n")
	time.Sleep(global.WaitTime)
	Capteur_co(ts, etatVentilateur)
}
