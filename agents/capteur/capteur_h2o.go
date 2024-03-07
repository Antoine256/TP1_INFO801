package capteur

import (
	. "github.com/pspaces/gospace"
	"math/rand"
)

func Capteur_h2o(ts *Space) {
	valeur_h2o := rand.Float64() * 100
	ts.Put("niveau_h2o", valeur_h2o)
	Capteur_h2o(ts)
}
