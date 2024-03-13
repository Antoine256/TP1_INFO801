package ventilateur

import (
	"TP1_INFO801/global"
	. "github.com/pspaces/gospace"
)

func Ventilateur(ts *Space, etat *string) {
	_, errActivation := ts.GetP("activation_ventilateur")
	_, errDesactivation := ts.GetP("desactivation_ventilateur")
	if errActivation == nil {
		global.SendToConn("Ventilateur activé")
		*etat = "activé"
		Ventilateur(ts, etat)
	}
	if errDesactivation == nil {
		global.SendToConn("Ventilateur désactivé")
		*etat = "désactivé"
		Ventilateur(ts, etat)
	}
	Ventilateur(ts, etat)
}
