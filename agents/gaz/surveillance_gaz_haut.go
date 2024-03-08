package gaz

import . "github.com/pspaces/gospace"

func Surveillance_gaz_haut(ts *Space, seuil_ch4 float64, seuil_co float64) {
	var x float64
	var y float64
	ts.Get("detection_h2o_bas")
	ts.Get("niveau_ch4", &x)
	ts.Get("niveau_co", &y)
	if x < seuil_ch4 && y < seuil_co {
		Surveillance_gaz_haut(ts, seuil_ch4, seuil_co)
	} else {
		ts.Put("activation_ventilateur")
		ts.Get("detection_gaz_haut")
		Surveillance_gaz_haut(ts, seuil_ch4, seuil_co)
	}
}
