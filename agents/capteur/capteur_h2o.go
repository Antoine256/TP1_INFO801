package capteur

import (
	"TP1_INFO801/global"
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

var valeur_H2O = rand.Float64() * 100

func get_valeur_h2o(etatPompe string) float64 {
	switch etatPompe {
	case "activée":
		valeur_H2O -= rand.Float64() * 5
		if valeur_H2O < 0.0 {
			valeur_H2O = 0.0
		}
		return valeur_H2O
	default:
		valeur_H2O += rand.Float64() * 3
		return valeur_H2O
	}
}

func Capteur_h2o(ts *Space, etatPompe *string) {
	valeur_h2o := get_valeur_h2o(*etatPompe)
	global.Add(ts, "niveau_h2o", valeur_h2o)
	fmt.Print("Capteur_h2o: ", valeur_h2o, "\n")
	time.Sleep(global.WaitTime)
	Capteur_h2o(ts, etatPompe)
}
