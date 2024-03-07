package pompe

import (
	. "github.com/pspaces/gospace"
)

func Pompe(ts *Space, etat string) {
	var pompe string
	select {
	case ts.Get("activation_pompe"):
		Pompe(ts, "activée")
	}
}
