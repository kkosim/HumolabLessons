package interfacePlanets

import "fmt"

type Titan struct {
	Name string
}

func (t Titan) calculateWeight() {
	fmt.Printf("Weight on %v = %v\n", t.Name, mass*0.13)
}

const mass = 60

type ProximaCentauriB struct {
	Name string
}

func (p ProximaCentauriB) calculateWeight() {
	fmt.Printf("Weight on %v = %v\n", p.Name, mass*0.9)
}

type GravityFinder interface {
	calculateWeight()
}

func Weight(newWeight GravityFinder) {
	newWeight.calculateWeight()
}
