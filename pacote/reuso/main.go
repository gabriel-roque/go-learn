package main

// go get -u github.com/gabriel-roque/star-wars-pkg-go

import (
	"fmt"

	starwars "github.com/gabriel-roque/star-wars-pkg-go/star-wars"
)

func main() {
	fmt.Println(starwars.GetDefaultPersons())
}
