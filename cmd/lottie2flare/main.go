// lottie2flare parses a lottie JSON file, converts it to a flare animation
// and writes it to a .flr file.
package main

import (
	"flag"
	"log"

	"github.com/gmlewis/lottie2flare/convert"
	"github.com/gmlewis/lottie2flare/lottie"
)

var ()

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("Processing %v ...", arg)
		anim, err := lottie.NewFile(arg)
		check("NewFile(%q): %v", arg, err)
		flr, err := convert.Lottie2Flare(anim)
		check("Lottie2Flare: %v", err)
		err = flr.WriteFile(arg + ".flr")
		check("WriteFile(%q): %v", arg+".flr", err)
	}

	log.Println("Done.")
}

func check(fmtStr string, args ...interface{}) {
	err := args[len(args)-1]
	if err != nil {
		log.Fatalf(fmtStr, args...)
	}
}
