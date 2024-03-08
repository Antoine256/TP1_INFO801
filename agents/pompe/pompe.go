package pompe

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func Pompe(ts *Space, etat string) {
	_, errActivation := ts.Get("activation_pompe")
	_, errDesactivation := ts.Get("desactivation_pompe")
	if errActivation == nil {
		fmt.Println("Pompe activée")
		Pompe(ts, "activée")
	}
	if errDesactivation == nil {
		fmt.Println("Pompe désactivée")
		Pompe(ts, "désactivée")
	}
	Pompe(ts, etat)
}
