package main

import "fmt"

type People struct {
	name string
}

func main() {
	/* 1 */
	animal := make(map[string]interface{})
	animal["dog"] = "dog"
	animal["cat"] = "cat"
	animal["bat"] = "bat"

	var nameAnimal []string
	nameAnimal = append(nameAnimal, "cat", "dog", "bat", "cat")
	var countDog int
	var countCat int
	var countBat int
	for i := 0; i < len(nameAnimal); i++ {
		if n, ok := animal[nameAnimal[i]]; ok {
			if n == "dog" {
				countDog += 1
			} else if n == "cat" {
				countCat += 1
			} else if n == "bat" {
				countBat += 1
			}
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
		name: "Apichat",
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
