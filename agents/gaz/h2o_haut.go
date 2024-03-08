package gaz

import . "github.com/pspaces/gospace"

func H2o_haut(ts *Space, seuil_h2o_haut float64) {
	var x float64
	ts.Query("activation_pompe")
	ts.Query("niveau_h2o", &x)
	if x >= seuil_h2o_haut {
		ts.Put("h2o_haut_detect")
		ts.Get("detection_h2o_haut")
		H2o_haut(ts, seuil_h2o_haut)
	} else {
		H2o_haut(ts, seuil_h2o_haut)
	}
}
