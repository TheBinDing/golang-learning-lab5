package main

import "fmt"

type People struct {
	name string
}

type Animal interface {
	Type() string
}

type Dog struct{}

func (d Dog) Type() string {
	return "dog"
}

type Cat struct{}

func (c Cat) Type() string {
	return "cat"
}

type Bat struct{}

func (b Bat) Type() string {
	return "bat"
}

func main() {
	/* 1 */
	dog := Dog{}
	cat := Cat{}
	bat := Bat{}

	var countDog int
	var countCat int
	var countBat int

	animals := []Animal{dog, cat, bat, cat}
	for _, animal := range animals {
		if animal.Type() == "dog" {
			countDog += 1
		} else if animal.Type() == "cat" {
			countCat += 1
		} else if animal.Type() == "bat" {
			countBat += 1
		}
	}
	fmt.Printf("Dog: %d\nCat: %d\nBat: %d\n\n", countDog, countCat, countBat)

	/* 2 */
	vowels := make(map[string]interface{})
	vowels["a"] = "a"
	vowels["e"] = "e"
	vowels["i"] = "i"
	vowels["o"] = "o"
	vowels["u"] = "u"
	vowels["A"] = "A"
	vowels["E"] = "E"
	vowels["I"] = "I"
	vowels["O"] = "O"
	vowels["U"] = "U"

	var ThaiPeople People = People{
		name: "Saritrat Jirakulphondchai",
	}

	fmt.Println(ThaiPeople.name)

	var count int

	for j := 0; j < len(ThaiPeople.name); j++ {
		if _, ok := vowels[ThaiPeople.name[j:j+1]]; ok {
			if ok == true {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
