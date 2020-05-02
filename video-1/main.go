package main

import "fmt"

//Interfaces
// implicit interfaces
// struct -> interface
func main() {
	fmt.Println("I want a coffee...")

	var (
		i ItalianCoffeeMachine
		c ColombianCoffeeMachine
	)

	italianCoffee := GetCoffee(&i, 10)

	italianCoffee.PrintCoffee()

	colombianCoffee := GetCoffee(&c, 25)

	colombianCoffee.PrintCoffee()

	// ...

	machine := CoffeeMachine{&c}

	machineCCoffee := machine.MakeCoffee(39)

	machineCCoffee.PrintCoffee()

	// ...
	supreme := Supreme{}
	supremeCoffee, status := SupremeCaller(&supreme)

	supremeCoffee.PrintCoffee()
	fmt.Println("status is: " + status)

}

type Coffee struct {
	Intensity int
	Region    string
}

func (c *Coffee) PrintCoffee() {
	fmt.Println(fmt.Sprintf("This coffee is from %s and intensity is %d", c.Region, c.Intensity))
}

// CoffeeMaker
type CoffeeMaker interface {
	MakeCoffee(intensity int) Coffee
}

type ItalianCoffeeMachine struct {
}

type ColombianCoffeeMachine struct {
}

func (i *ItalianCoffeeMachine) Print() {

}

func (i *ItalianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Italy"}
}

func (c *ColombianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Colombia"}
}

// .....

func GetCoffee(coffeeMaker CoffeeMaker, i int) Coffee {
	return coffeeMaker.MakeCoffee(i)
}

type CoffeeMachine struct {
	CoffeeMaker
}

type Supreme struct{}

type SupremeMachine interface {
	CoffeeMaker
	CheckMachine() string
}

func (s *Supreme) MakeCoffee(i int) Coffee {
	return Coffee{Intensity: i, Region: "unknown but tasty"}
}

func (s *Supreme) CheckMachine() string {
	return "looks good"
}

func SupremeCaller(sm SupremeMachine) (coffee Coffee, status string) {
	coffee = sm.MakeCoffee(99)
	status = sm.CheckMachine()
	return
}
