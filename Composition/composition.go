/*
	GoLang does not support inheritance but it something like it called COMPOSITION.
	In composition we include a struct in another struct.
	For example if we have
	A{
		---
	}
	B{
		---
		A
	}
	Then all the member variables & methods are A in accessible through B(Important: But not the other way round).
*/
package main

import (
	"fmt"
)

type Car struct {
	Name  string
	Model string
	Engine
	Wheel
}

func (c *Car) setName(s string) {
	c.Name = s
}

func (c *Car) getName() {
	fmt.Println("Car Name: %s", c.Name)
}

func (c *Car) setModel(m string) {
	c.Model = m
}

func (c *Car) getModel() {
	fmt.Println("Car Name: %s", c.Model)
}

type Engine struct {
	FuelType string
	Power    float32
	Torque   float32
}

func (e *Engine) setFuelType(ft string) {
	e.FuelType = ft
}

func (e *Engine) getFuelType() {
	fmt.Println("Engine FuelType: %s", e.FuelType)
}

func (e *Engine) setPower(p float32) {
	e.Power = p
}

func (e *Engine) getPower() {
	fmt.Println("Engine Power: %f", e.Power)
}

func (e *Engine) setTorque(t float32) {
	e.Torque = t
}

func (e *Engine) getTorque() {
	fmt.Println("Engine Torque: %f", e.Torque)
}

type Wheel struct {
	Name string
	Type string
	Size float32
}

func (w *Wheel) setType(t string) {
	w.Type = t
}

func (w *Wheel) getType() {
	fmt.Println("Wheel Type: %s", w.Type)
}

func (w *Wheel) setName(n string) {
	w.Name = n
}

func (w *Wheel) getName() {
	fmt.Println("Wheel Name: %f", w.Name)
}

func main() {

	var c = &Car{
		Name:  "Tata Indigo",
		Model: "TA30",
	}

	c.setName("Honda")    // Changes car.name not wheel.name
	c.setType("Tubeless") // Wheel Type
	c.setPower(100.57)    // Engine power
	c.setTorque(200)      // Engine torque
	fmt.Printf("%v", *c)  // {Honda TA30 { 100.57 200} { Tubeless 0}}

}
