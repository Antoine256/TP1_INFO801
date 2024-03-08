package ventilateur

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func Ventilateur(ts *Space, etat *string) {
	_, errActivation := ts.Get("activation_ventilateur")
	_, errDesactivation := ts.Get("desactivation_ventilateur")
	if errActivation == nil {
		fmt.Println("Ventilateur activée")
		*etat = "activée"
		Ventilateur(ts, etat)
	}
	if errDesactivation == nil {
		fmt.Println("Ventilateur désactivée")
		*etat = "désactivée"
		Ventilateur(ts, etat)
	}
	Ventilateur(ts, etat)
}
