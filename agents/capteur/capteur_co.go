package capteur

import (
	"fmt"
	. "github.com/pspaces/gospace"
	"math/rand"
	"time"
)

func Capteur_co(ts *Space) {
	valeur_CO := rand.Float64() * 100
	ts.Put("niveau_co", valeur_CO)
	fmt.Print("Capteur_co: ", valeur_CO, "\n")
	time.Sleep(5 * time.Second)
	Capteur_co(ts)
}
