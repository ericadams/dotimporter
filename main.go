package main

import (
	"fmt"

	. "github.com/ericadams/dotimporter/inner"
)

func main() {
	fmt.Println("vim-go")
	i := inner{member: thingy}
	I := newPublicMember()
	fmt.Println(i, I)
}
