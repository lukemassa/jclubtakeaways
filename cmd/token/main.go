package main

import (
	"github.com/lukemassa/jclubtakeaways/internal/token"
	"fmt"
)

func main() {
	tokener := token.New()
	fmt.Println(tokener.Get())
}
