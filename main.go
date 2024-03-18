package main

import (
	"TP1_INFO801/agents"
	"TP1_INFO801/agents/capteur"
	"TP1_INFO801/agents/gaz"
	"TP1_INFO801/agents/pompe"
	"TP1_INFO801/agents/ventilateur"
	"TP1_INFO801/global"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"sync"
	"time"
)

import . "github.com/pspaces/gospace"

var wg sync.WaitGroup

func printEtat() {
	fmt.Println("Pompe : " + global.EtatPompe)
	fmt.Println("Ventilateur : " + global.EtatVentilateur)
	time.Sleep(global.WaitTime)
	printEtat()
}

func main() {

	// Socket !

	http.HandleFunc("/", global.Handler)
	fmt.Println("Server listening on :8081")
	//--------------------

	ts := NewSpace("ts")

	ts.Put("detection_h2o_haut")

	wg.Add(11) // Attendre que tous les processus en parallèle se lancent

	go capteur.Capteur_ch4(&ts, &global.EtatVentilateur)
	go capteur.Capteur_co(&ts, &global.EtatVentilateur)
	go capteur.Capteur_h2o(&ts, &global.EtatPompe)
	go gaz.Surveillance_gaz_haut(&ts, 50.0, 50.0)
	go gaz.Gaz_bas(&ts, 50.0, 50.0)
	go gaz.H2o_bas(&ts, 50.0)
	go gaz.H2o_haut(&ts, 50.0)
	go pompe.Pompe(&ts, &global.EtatPompe)
	go ventilateur.Ventilateur(&ts, &global.EtatVentilateur)
	go agents.Commande_pompe_ventilateur(&ts, 50.0, 50.0)

	go printEtat()

	go global.HandleWrites()
	go log.Fatal(http.ListenAndServe(":8081", nil))

	wg.Wait()
}
