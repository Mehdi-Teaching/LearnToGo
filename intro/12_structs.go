package main

import "fmt"

// Defin person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {
	// Init person using struct
	// you can put values for fields you want

	person1 := Person{
		firstName: "Mehdi",
		lastName:  "Teymorian",
		city:      "Tehran",
		gender:    "Male",
		age:       22, // in multiline initialization, comma is necessary for last field
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName)

	// function defined in line 37
	println(person1.greet())
	// function defined in line 45
	person1.incrementAge()
	fmt.Println(person1)

	// function defined in line 50
	person1.changeLastName("Green")
	fmt.Println(person1)

	betterPerson, err := createPerson("Mehti", "Timurid", "Tehran", "Male", 22)

	// for a better perception of struct
	fmt.Printf("%+v %d\n", betterPerson, err)
}

// Greeting method for Person Struct
// value receiver
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName
}

// pointer receiver
func (person *Person) incrementAge() {
	person.age++
}

func (person *Person) changeLastName(newName string) {
	person.lastName = newName
}

// =====================================

func createPerson(firstName, lastName, city, gender string, age int) (*Person, error) {
	if firstName == "" || lastName == "" {
		return nil, fmt.Errorf("firstName or lastName cannot be empty")
	}

	person := &Person{
		firstName: firstName,
		lastName:  lastName,
		city:      city,
		gender:    gender,
		age:       age,
	}
	return person, nil
}

// =====================================
// shorter way to define struct
// define field with same type in one line
type Animal struct {
	name, gender, owner string
	age                 int
}
