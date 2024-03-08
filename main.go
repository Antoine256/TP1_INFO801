package main

import (
	"TP1_INFO801/agents/gaz"
	"fmt"
)
import . "github.com/pspaces/gospace"
import . "TP1_INFO801/agents/capteur"

func main() {
	ts := NewSpace("ts")

	go Capteur_ch4(&ts)
	go Capteur_co(&ts)
	go Capteur_h2o(&ts)
	go gaz.Surveillance_gaz_haut(&ts, 50, 50)
	go gaz.H2o_bas(&ts, 50)

	// Put a message into the space.
	ts.Put("Hello world!")

	// Get a message from the space
	// via pattern matching.
	var message string
	t, _ := ts.Get(&message)

	fmt.Println((t.GetFieldAt(0)).(string))
}
