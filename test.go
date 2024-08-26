package test

import (
	dog "github.com/dineshs90/demo3"
)

func Bark() string {
	return "Woof !!!"
}

func Barks() string {
	return "Dog Barks!!!"

}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
