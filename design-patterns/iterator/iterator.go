package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Recipe struct {
	name        string
	ingredients []string
	index       int
}

func NewRecipe(name string, ingredients []string) *Recipe {
	return &Recipe{name: name, ingredients: ingredients, index: 0}
}

func (r *Recipe) HasNext() bool {
	if r.index < len(r.ingredients) {
		return true
	}
	return false
}

func (r *Recipe) Next() interface{} {
	if r.HasNext() {
		curIdx := r.index
		r.index++
		return r.ingredients[curIdx]
	}
	return nil
}
