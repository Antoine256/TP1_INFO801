package main

import (
	"TP1_INFO801/agents"
	"TP1_INFO801/agents/capteur"
	"TP1_INFO801/agents/gaz"
	"TP1_INFO801/agents/pompe"
	"TP1_INFO801/agents/ventilateur"
	"TP1_INFO801/global"
	"fmt"
	"sync"
	"time"
)
import . "github.com/pspaces/gospace"

var wg sync.WaitGroup

var etatPompe string
var etatVentilateur string

func printEtat() {
	fmt.Println("Pompe : " + etatPompe)
	fmt.Println("Ventilateur : " + etatVentilateur)
	time.Sleep(global.WaitTime)
	printEtat()
}

func main() {
	ts := NewSpace("ts")

	ts.Put("detection_gaz_haut")
	ts.Put("detection_h2o_haut")

	wg.Add(11) // Attendre que tous les processus en parallèle se lancent

	go capteur.Capteur_ch4(&ts, &etatVentilateur)
	go capteur.Capteur_co(&ts, &etatVentilateur)
	go capteur.Capteur_h2o(&ts, &etatPompe)
	go gaz.Surveillance_gaz_haut(&ts, 50.0, 50.0)
	go gaz.Gaz_bas(&ts, 50.0, 50.0)
	go gaz.H2o_bas(&ts, 50.0)
	go gaz.H2o_haut(&ts, 50.0)
	go pompe.Pompe(&ts, &etatPompe)
	go ventilateur.Ventilateur(&ts, &etatVentilateur)
	go agents.Commande_pompe_ventilateur(&ts, 50.0, 50.0)

	go printEtat()

	wg.Wait()
}
