package datatypes

// MapsExamples demonstrates various map operations in Go
// including creation, manipulation, and common operations
func MapsExamples() {
	// Create an empty map using var and make
	var heroes map[string]string
	heroes = make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Spider-Man"] = "Peter Parker"

	// Create and initialize a map in one line
	villains := make(map[string]string)
	villains["Vibe Coding"] = "LLMs"
	villains["Penguin"] = "Oswald Cobblepot"
	villains["Green Goblin"] = "Norman Osborn"

	// Create and initialize a map with values
	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound"}

	pl(heroes)
	pl(villains)
	pl(superPets)

	// Delete a key from a map
	delete(heroes, "Superman")

	// Iterate over map keys and values
	pl("Heroes:")
	for k, v := range heroes {
		pl(k, v)
	}

	pl("Villains:")
	for k, v := range villains {
		pl(k, v)
	}

	pl("Super Pets:")

	for k, v := range superPets {
		pl(k, v)
	}

	// Access map values and check for existence
	pl(heroes["Batman"])
	pl(heroes["Superman"] == "")

}
