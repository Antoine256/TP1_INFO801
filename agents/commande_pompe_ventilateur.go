package agents

import (
	. "github.com/pspaces/gospace"
)

func Commande_pompe_ventilateur(ts *Space, seuil_ch4 float64, seuil_co float64) {
	var x float64
	var y float64
	ts.Get("H2o_haut_detect")
	ts.Query("nivel_ch4", &x)
	ts.Query("nivel_co", &y)
	if x < seuil_ch4 && y < seuil_co {
		ts.Put("activation_pompe")
		ts.Put("detection_h2o_bas")
		ts.Put("detection_gaz_haut")
		Commande_pompe_ventilateur(ts, seuil_ch4, seuil_co)
	} else {
		ts.Put("activation_ventilateur")
		ts.Put("detection_gaz_bas")
		Commande_pompe_ventilateur(ts, seuil_ch4, seuil_co)
	}
}
