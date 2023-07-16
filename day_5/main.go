package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
		return p
	}
	p.age = initialAge
	return p
}

func (p person) yearPasses() person {
	p.age += 1
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 12 && p.age < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are too old")
	}
}
func main() {
	mtu := person{}
	fmt.Println(mtu.age)
	mtu = mtu.NewPerson(20)
	fmt.Println(mtu.age)
	for i := 1; i <= 4; i++ {
		mtu = mtu.yearPasses()
	}
	fmt.Println("Increased age for four years", mtu.age)
	mtu.amIOld()
}
