package iterator

import "testing"

func TestIterator(t *testing.T) {
	ingredients := []string{"eggs", "flour", "butter"}
	recipe := NewRecipe("Cake", ingredients)

	for i := 0; recipe.HasNext(); i++ {
		nextIngredient := recipe.Next()
		if nextIngredient != ingredients[i] {
			t.Errorf("Expected '%s' found '%s'", ingredients[i], nextIngredient.(string))
		}
	}

	if recipe.Next() != nil {
		t.Errorf("Iterator should be terminated")
	}
}
