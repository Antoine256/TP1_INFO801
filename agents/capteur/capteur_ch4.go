package capteur

import (
	. "github.com/pspaces/gospace"
	"math/rand"
	_ "math/rand"
)

func Capteur_ch4(ts *Space) {
	valeur_CH4 := rand.Float64() * 100
	ts.Put("niveau_ch4", valeur_CH4)
	Capteur_ch4(ts)
}
