package decorator

import "testing"

func TestDecorator(t *testing.T) {
	var toppings []string
	icecream := NewVanillaIceCream()
	toppings = icecream.GetToppings()

	if len(toppings) != 0 {
		t.Error("Plain vanilla ice cream should have no toppings")
	}

	icecream = AddSprinkles(icecream)
	toppings = icecream.GetToppings()

	if len(toppings) != 1 {
		t.Error("Vanilla ice cream should have one topping")
	}

	if toppings[0] != "Sprinkles" {
		t.Error("Sprinkles should be in the toppings list")
	}

	icecream = AddChocolateSauce(icecream)
	toppings = icecream.GetToppings()

	if len(toppings) != 2 {
		t.Error("Vanilla ice cream should have two toppings")
	}

	if toppings[1] != "Chocolate Sauce" {
		t.Error("Chocolate Sauce should be in the toppings list")
	}
}
