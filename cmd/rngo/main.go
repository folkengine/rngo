package main

import (
	"flag"
	"fmt"
	"github.com/folkengine/rngo/pkg/rngo"
)

func main() {

	elvenFlag := flag.Bool("elven", false, "Create Elven names")
	fantasyFlag := flag.Bool("fantasy", false, "Create Fantasy names")
	//goblinFlag := flag.Bool("goblin", false, "Create Goblin names")
	//romanFlag := flag.Bool("roman", true, "Create Roman names")
	flag.Parse()

	var myname rngo.Rngo

	if *elvenFlag {
		myname = rngo.New(rngo.ElvenMap)
	} else if *fantasyFlag {
		myname = rngo.New(rngo.FantasyMap)
	} else {
		myname = rngo.New(rngo.GoblinMap)
	}

	fmt.Println(myname.FirstLast())
}
