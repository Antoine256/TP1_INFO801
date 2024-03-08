package gaz

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func H2o_haut(ts *Space, seuil_h2o_haut float64) {
	var x float64
	ts.Query("detection_h2o_haut")
	t, _ := ts.Query("niveau_h2o", &x)
	x = (t.GetFieldAt(1)).(float64)
	if x >= seuil_h2o_haut {
		fmt.Print("\nH2O haut détect")
		ts.Put("h2o_haut_detect")
		ts.Get("detection_h2o_haut")
		H2o_haut(ts, seuil_h2o_haut)
	} else {
		H2o_haut(ts, seuil_h2o_haut)
	}
}
