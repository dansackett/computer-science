package factory

import "testing"

func TestFactory(t *testing.T) {
	var appliances []Appliance

	factory := NewApplianceFactory()
	blender := factory.MakeBlender()
	microwave := factory.MakeMicrowave()

	appliances = append(appliances, blender)
	appliances = append(appliances, microwave)

	namesShouldEqual := []string{"Blender", "Microwave"}

	for i, appliance := range appliances {
		if appliance.GetName() != namesShouldEqual[i] {
			t.Errorf("Appliance names do not match")
		}
	}
}
