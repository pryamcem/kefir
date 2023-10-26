package main

import (
	"fmt"

	"github.com/pryamcem/kefir/aur"
)

func main() {
	//p, err := aur.SearchByMaintainer("naspeh")
	p, err := aur.Info("perevod-git")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
