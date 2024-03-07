package capteur

import (
	. "github.com/pspaces/gospace"
	"math/rand"
)

func Capteur_co(ts *Space) {
	valeur_CO := rand.Float64() * 100
	ts.Put("niveau_co", valeur_CO)
	Capteur_co(ts)
}
