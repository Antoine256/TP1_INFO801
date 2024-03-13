package pompe

import (
	"TP1_INFO801/global"
	. "github.com/pspaces/gospace"
)

func Pompe(ts *Space, etat *string) {
	_, errActivation := ts.GetP("activation_pompe")
	_, errDesactivation := ts.GetP("desactivation_pompe")
	if errActivation == nil {
		global.SendToConn("Pompe activée")
		*etat = "activée"
		Pompe(ts, etat)
	}
	if errDesactivation == nil {
		global.SendToConn("Pompe desactivée")
		*etat = "desactivée"
		Pompe(ts, etat)
	}
	Pompe(ts, etat)
}
