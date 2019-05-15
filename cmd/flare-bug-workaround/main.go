// flare-bug-workaround reorders the JSON keys to work around the known
// bug in the flare JSON reader: https://github.com/2d-inc/Flare-Flutter/issues/90
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gmlewis/lottie2flare/flare"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("Processing %v ...", arg)
		anim, err := flare.NewFile(arg)
		check("NewFile(%q): %v", arg, err)
		err = writeFile(anim, arg+".fix.json")
		check("WriteFile(%q): %v", arg+".fix.json", err)
	}

	log.Println("Done.")
}

func writeFile(root *flare.Root, filename string) error {
	var artboards []string
	for _, artboard := range root.Artboards {
		artboards = append(artboards, workaround(artboard))
	}
	buf := fmt.Sprintf(`{"version":%v,"artboards":[%v]}`, *root.Version, strings.Join(artboards, ","))
	return ioutil.WriteFile(filename, []byte(buf), 0644)
}

func workaround(ab *flare.Artboard) string {
	anims := ab.Animations
	ab.Animations = nil
	buf, err := json.Marshal(ab)
	check("artboard Marshal: %v", err)
	animBuf, err := json.Marshal(anims)
	check("animations Marshal: %v", err)
	return fmt.Sprintf(`%s,"animations":%s}`, buf[:len(buf)-1], animBuf)
}

func check(fmtStr string, args ...interface{}) {
	err := args[len(args)-1]
	if err != nil {
		log.Fatalf(fmtStr, args...)
	}
}
