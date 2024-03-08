package ventilateur

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func Ventilateur(ts *Space, etat string) {
	_, errActivation := ts.Get("activation_ventilateur")
	_, errDesactivation := ts.Get("desactivation_ventilateur")
	if errActivation == nil {
		fmt.Println("Ventilateur activée")
		Ventilateur(ts, "activée")
	}
	if errDesactivation == nil {
		fmt.Println("Ventilateur désactivée")
		Ventilateur(ts, "désactivée")
	}
	Ventilateur(ts, etat)
}
