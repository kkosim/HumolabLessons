package interfacePlanets

import (
	"fmt"
	"math"
)

const G = 6.67430 / 100000000000
const mass = 100
const EarthMass = 5977220000000000000000000
const EarthRadius = 6375433

type Satellite struct {
	Name   string
	Mass   float64
	Radius float64
	Planet string
}

type Planet struct {
	Name   string
	Mass   float64
	Radius float64
	Star   string
}

type Star struct {
	Name   string
	Mass   float64
	Radius float64
	Galaxy string
}

func (s *Satellite) calculateWeight() {

	g := G * s.Mass * EarthMass / (s.Radius * s.Radius * EarthRadius * EarthRadius)
	currMass := mass * g / 9.8
	fmt.Printf("Weight on %v = %v in the System of %v\n", s.Name, math.Round(currMass), s.Planet)
}

func (p *Planet) calculateWeight() {

	g := G * p.Mass * EarthMass / (p.Radius * p.Radius * EarthRadius * EarthRadius)
	currMass := mass * g / 9.8
	fmt.Printf("Weight on %v = %v in the System of %v\n", p.Name, math.Round(currMass), p.Star)
}

func (s *Star) calculateWeight() {

	g := G * s.Mass * EarthMass / (s.Radius * s.Radius * EarthRadius * EarthRadius)
	currMass := mass * g / 9.8
	fmt.Printf("Weight on %v = %v in the System of %v\n", s.Name, math.Round(currMass), s.Galaxy)
}

type GravityFinder interface {
	calculateWeight()
}

func Weight(newWeight GravityFinder) {
	newWeight.calculateWeight()
}
