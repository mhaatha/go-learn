package main

import (
	"fmt"
)

func (e email) cost() float64 {
	bodyLen := float64(len(e.body))

	if e.isSubscribed {
		return 0.01 * bodyLen
	}

	return 0.05 * bodyLen
}

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}
