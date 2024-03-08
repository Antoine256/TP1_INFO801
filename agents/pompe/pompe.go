package pompe

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func Pompe(ts *Space, etat *string) {
	fmt.Print("Pompe\n")
	_, errActivation := ts.Get("activation_pompe")
	_, errDesactivation := ts.Get("desactivation_pompe")
	if errActivation == nil {
		fmt.Println("Pompe activée")
		*etat = "activée"
		Pompe(ts, etat)
	}
	if errDesactivation == nil {
		fmt.Println("Pompe désactivée")
		*etat = "desactivée"
		Pompe(ts, etat)
	}
	Pompe(ts, etat)
}
