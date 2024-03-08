package gaz

import (
	"fmt"
	. "github.com/pspaces/gospace"
)

func Surveillance_gaz_haut(ts *Space, seuil_ch4 float64, seuil_co float64) {
	var x float64
	var y float64
	ts.Query("detection_gaz_haut")

	fmt.Println("Surveillance gaz haut actif")
	tx, _ := ts.Query("niveau_ch4", &x)
	x = (tx.GetFieldAt(1)).(float64)
	ty, _ := ts.Query("niveau_co", &y)
	y = (ty.GetFieldAt(1)).(float64)
	if x < seuil_ch4 && y < seuil_co {
		Surveillance_gaz_haut(ts, seuil_ch4, seuil_co)
	} else {
		ts.Put("activation_ventilateur")
		ts.Get("detection_gaz_haut")
		Surveillance_gaz_haut(ts, seuil_ch4, seuil_co)
	}
}
