package gaz

import . "github.com/pspaces/gospace"

func Gaz_bas(ts *Space, seuil_ch4 float64, seuil_co float64) {
	var y float64
	var z float64
	ts.Get("detection_h2o_bas")
	ts.Get("niveau_ch4", &y)
	ts.Get("niveau_co", &z)
	if y < seuil_ch4 && z < seuil_co {
		ts.Put("activation_pompe")
		ts.Put("detection_h2o_bas")
		ts.Get("detection_gaz_bas")
		ts.Put("detection_gaz_haut")
		Gaz_bas(ts, seuil_ch4, seuil_co)
	} else {
		Gaz_bas(ts, seuil_ch4, seuil_co)
	}

}
