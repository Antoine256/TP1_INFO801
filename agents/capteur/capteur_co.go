package capteur

import (
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

var valeur_CO = rand.Float64() * 100

func get_valeur_co(etatVentilateur string) float64 {
	switch etatVentilateur {
	case "activée":
		valeur_CO -= rand.Float64() * 5
		return valeur_CO
	default:
		valeur_CO += rand.Float64() * 3
		return valeur_CO
	}
}

func Capteur_co(ts *Space, etatVentilateur *string) {
	valeur_CO := get_valeur_co(*etatVentilateur)
	ts.Put("niveau_co", valeur_CO)
	fmt.Print("Capteur_co: ", valeur_CO, "\n")
	time.Sleep(5 * time.Second)
	Capteur_co(ts, etatVentilateur)
}
