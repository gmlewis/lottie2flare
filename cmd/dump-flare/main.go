// dump-flare parses a flare JSON file and dumps it to a file.
// It is used for debugging.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/lottie2flare/flare"
)

var ()

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("Processing %v ...", arg)
		anim, err := flare.NewFile(arg)
		check("NewFile(%q): %v", arg, err)
		err = anim.WriteFile(arg + ".out.json")
		check("WriteFile(%q): %v", arg+".out.json", err)
	}

	log.Println("Done.")
}

func check(fmtStr string, args ...interface{}) {
	err := args[len(args)-1]
	if err != nil {
		log.Fatalf(fmtStr, args...)
	}
}
