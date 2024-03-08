package ventilateur

import (
	. "github.com/pspaces/gospace"
)

func Ventilateur(ts *Space, etat *string) {
	_, errActivation := ts.GetP("activation_ventilateur")
	_, errDesactivation := ts.GetP("desactivation_ventilateur")
	if errActivation == nil {
		*etat = "activée"
		Ventilateur(ts, etat)
	}
	if errDesactivation == nil {
		*etat = "désactivée"
		Ventilateur(ts, etat)
	}
	Ventilateur(ts, etat)
}
