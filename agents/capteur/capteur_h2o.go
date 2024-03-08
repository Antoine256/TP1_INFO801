package capteur

import (
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

func Capteur_h2o(ts *Space) {
	valeur_h2o := rand.Float64() * 100
	ts.Put("niveau_h2o", valeur_h2o)
	fmt.Print("Capteur_h2o: ", valeur_h2o, "\n")
	time.Sleep(5 * time.Second)
	Capteur_h2o(ts)
}
