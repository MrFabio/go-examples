package lang

type Animal interface {
	AngrySound() string
	HappySound() string
}

type Cat string

func (c Cat) Attack() {
	pl("Cat attacks")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() string {
	return "Hisssssss"
}

func (c Cat) HappySound() string {
	return "Purrrrrrr"
}

func InterfacesExamples() {
	var cat Animal = Cat("Meowth")
	pl(cat.AngrySound())
	pl(cat.HappySound())
	var kitty Cat = cat.(Cat)
	pl(kitty.Name())
	kitty.Attack()
}
