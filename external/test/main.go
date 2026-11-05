package main

import (
	"fmt"

	"github.com/Triplebond1/secure-notes/internal/auth"
)

func main() {
	fmt.Println(auth.GenerateToken())
}
