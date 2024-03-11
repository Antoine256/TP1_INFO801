package global

import (
	. "github.com/pspaces/gospace/space"
)

func Out[T any](ts *Space, name string, value T) {
	var x float64
	ts.Put(name, value) // On initialise le tuple si jamais il n'existe pas
	ts.Get(name, &x)    // On récupère le tuple
	ts.Put(name, value) // On met à jour la nouvelle valeur
}
