package main

import (
	"fmt"
	"os"
	"strings"
)

func acronym(s []string) (acr string) {
	for _, v := range s {
		acr = acr + strings.ToUpper(string(v[0]))
	}
	return acr
}

func main() {
	s := "Simple Test Of Pull Request"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(strings.Fields(s)[:]))
}
