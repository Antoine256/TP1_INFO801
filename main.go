package main

import (
	"TP1_INFO801/agents/capteur"
	"TP1_INFO801/agents/pompe"
	"fmt"
)
import . "github.com/pspaces/gospace"

func main() {
	var etatPompe string
	ts := NewSpace("ts")

	go capteur.Capteur_ch4(&ts)
	go capteur.Capteur_co(&ts)
	go capteur.Capteur_h2o(&ts)
	go pompe.Pompe(&ts, etatPompe)

	// Put a message into the space.
	ts.Put("Hello world!")

	// Get a message from the space
	// via pattern matching.
	var message string
	t, _ := ts.Get(&message)

	fmt.Println((t.GetFieldAt(0)).(string))
}
