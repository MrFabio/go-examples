package datatypes

import "fmt"

type Teaspoon float64
type Tablespoon float64
type Milliliter float64

func teaspoonToMilliliter(t Teaspoon) Milliliter {
	return Milliliter(t * 4.92)
}

func tablespoonToMilliliter(t Tablespoon) Milliliter {
	return Milliliter(t * 14.79)
}

func (tsp Teaspoon) ToMilliliter() Milliliter {
	return Milliliter(tsp * 4.92)
}

func (tbsp Tablespoon) ToMilliliter() Milliliter {
	return Milliliter(tbsp * 14.79)
}

func CompositionExamples() {
	ml1 := Milliliter(Teaspoon(3) * 4.92)
	ml2 := Milliliter(Tablespoon(3) * 14.79)
	pl("3 teaspoons =", ml1, "ml")
	pl("3 tablespoon =", ml2, "ml")
	pl("3 teaspoons + 2 tablespoons =", Teaspoon(3)+Teaspoon(2))
	fmt.Printf("3 teaspoons = %.2f mL\n", teaspoonToMilliliter(3))
	fmt.Printf("3 tablespoons = %.2f mL\n", tablespoonToMilliliter(3))

	fmt.Printf("3 teaspoons = %.2f mL\n", Teaspoon(3).ToMilliliter())
	fmt.Printf("3 tablespoons = %.2f mL\n", Tablespoon(3).ToMilliliter())
}
