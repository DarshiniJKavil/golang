package main

import "fmt"

func conference() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func grade() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v,%v,%v", grade1, grade2, grade3)
}

func grades() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v", grades)
}

func students() {
	var students [3]string
	fmt.Printf("students : %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("students : %v\n", students)
}

func student() {
	var students [3]string
	fmt.Printf("student : %v\n", students)
	students[0] = "Lisa"
	students[1] = "Rose"
	students[2] = "Jeny"
	fmt.Printf("student : %v\n", students)
}

func student1() {
	var students [3]string
	fmt.Printf("student1 : %v\n", students)
	students[0] = "Lisa"
	students[1] = "Rose"
	students[2] = "Jeny"
	fmt.Printf("student1 #1: %v\n", students[1])
	fmt.Printf("Number of student1 : %v\n", len(students))
}

func matrix() {
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
}

func main1() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}

func main2() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func main3() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func main4() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func main5() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func main6() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:len(a)-1]
	fmt.Println(b)
}

func population() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"california":   8643579,
		"Texas":        8523580,
		"florida":      4796343,
		"newyork":      9245772,
		"pennsylvania": 9876543,
		"illinois":     1234579,
		"ohio":         3325787,
	}
	statePopulations["georgia"] = 57890643
	fmt.Println(statePopulations["georgia"])
}

type doctor struct {
	number     int
	actorName  string
	companions []string
}

func doctors() {
	adoctor := doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(adoctor)
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func animal() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 30,
		CanFly:   false,
	}

	fmt.Println(b.Name)
}

func main() {
	conference()
	grade()
	grades()
	students()
	student()
	student1()
	matrix()
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	population()
	doctors()
	animal()
}
