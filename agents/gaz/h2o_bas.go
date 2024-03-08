package gaz

import . "github.com/pspaces/gospace"

func H2o_bas(ts *Space, seuil_h2o_bas float64) {
	var x float64
	ts.Query("detection_h2o_bas")
	ts.Get("niveau_h2o", &x)
	if x < seuil_h2o_bas {
		ts.Put("desactivation_pompe")
		ts.Put("desactivation_ventilateur")
		ts.Get("detection_h2o_bas")
		ts.Put("detection_h2o_haut")
		H2o_bas(ts, seuil_h2o_bas)
	} else {
		H2o_bas(ts, seuil_h2o_bas)
	}

}
