package pompe

import (
	. "github.com/pspaces/gospace"
)

func Pompe(ts *Space, etat *string) {
	_, errActivation := ts.GetP("activation_pompe")
	_, errDesactivation := ts.GetP("desactivation_pompe")
	if errActivation == nil {
		*etat = "activée"
		Pompe(ts, etat)
	}
	if errDesactivation == nil {
		*etat = "desactivée"
		Pompe(ts, etat)
	}
	Pompe(ts, etat)
}
