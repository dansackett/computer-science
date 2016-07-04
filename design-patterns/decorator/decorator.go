package decorator

type IceCream interface {
	GetToppings() []string
	AddTopping(string)
}

type VanillaIceCream struct {
	Toppings []string
}

func NewVanillaIceCream() IceCream {
	var toppings []string
	return &VanillaIceCream{Toppings: toppings}
}

func (i *VanillaIceCream) GetToppings() []string {
	return i.Toppings
}

func (i *VanillaIceCream) AddTopping(topping string) {
	i.Toppings = append(i.Toppings, topping)
}

type Sprinkles struct {
	icecream IceCream
}

func AddSprinkles(icecream IceCream) IceCream {
	icecream.AddTopping("Sprinkles")
	return &Sprinkles{icecream: icecream}
}

func (s Sprinkles) GetToppings() []string {
	return s.icecream.GetToppings()
}

func (s Sprinkles) AddTopping(topping string) {
	s.icecream.AddTopping(topping)
}

type ChocolateSauce struct {
	icecream IceCream
}

func AddChocolateSauce(icecream IceCream) IceCream {
	icecream.AddTopping("Chocolate Sauce")
	return &ChocolateSauce{icecream: icecream}
}

func (s ChocolateSauce) GetToppings() []string {
	return s.icecream.GetToppings()
}

func (s ChocolateSauce) AddTopping(topping string) {
	s.icecream.AddTopping(topping)
}
