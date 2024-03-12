package gaz

import (
	. "github.com/pspaces/gospace"
)

func Gaz_bas(ts *Space, seuil_ch4 float64, seuil_co float64) {
	var y float64
	var z float64
	ts.Query("detection_gaz_bas")
	ty, _ := ts.Query("niveau_ch4", &y)
	y = (ty.GetFieldAt(1)).(float64)
	tz, _ := ts.Query("niveau_co", &z)
	z = (tz.GetFieldAt(1)).(float64)
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
