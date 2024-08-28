package test

import (
	"fmt"

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

func From11() {
	fmt.Println("I am in version v1.0.0")
}

func From12() {
	fmt.Println("I am in version v2.0.0")
}
