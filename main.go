package main

import (
	"TP1_INFO801/agents"
	"TP1_INFO801/agents/capteur"
	"TP1_INFO801/agents/gaz"
	"TP1_INFO801/agents/pompe"
	"TP1_INFO801/agents/ventilateur"
	"fmt"
	"sync"
	"time"
)
import . "github.com/pspaces/gospace"

var wg sync.WaitGroup

func printEtat(etatPompe *string, etatVentilateur *string) {
	fmt.Println("Pompe : " + *etatPompe)
	fmt.Println("Ventilateur : " + *etatVentilateur)
	time.Sleep(5 * time.Second)
	printEtat(etatPompe, etatVentilateur)
}

func main() {
	var etatPompe string
	var etatVentilateur string
	ts := NewSpace("ts")

	ts.Put("detection_gaz_haut")
	ts.Put("detection_h2o_haut")

	wg.Add(11) // Attendre que tous les processus en parallèle se lancent

	go capteur.Capteur_ch4(&ts)
	go capteur.Capteur_co(&ts)
	go capteur.Capteur_h2o(&ts)
	go gaz.Surveillance_gaz_haut(&ts, 50.0, 50.0)
	go gaz.Gaz_bas(&ts, 50.0, 50.0)
	go gaz.H2o_bas(&ts, 50.0)
	go gaz.H2o_haut(&ts, 50.0)
	go pompe.Pompe(&ts, &etatPompe)
	go ventilateur.Ventilateur(&ts, &etatVentilateur)
	go agents.Commande_pompe_ventilateur(&ts, 50.0, 50.0)

	go printEtat(&etatPompe, &etatVentilateur)

	wg.Wait()
}
