package main

import(
	lib "github.com/klykou/libtest"
	"fmt"
)

func main() {
	m := lib.Init{"YOU"}
	fmt.Println(m.GetString)
}
