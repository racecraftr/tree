package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
)

type colorPrint func(any)string

var printColors = []colorPrint{
	color.InRed,
	color.InYellow,
	color.InGreen,
	color.InBlue,
	color.InCyan,
	color.InPurple,
}

const depthPart = "│ "
const depthEnd = "├─"

var currentPath string

var maxDepth int = 0xffff

func main() {

	flag.StringVar(&currentPath, "path", "", "path to look in.")
	flag.IntVar(&maxDepth, "depth", 0xffff, "specifies the depth of the three. Must be greater than 0.")
	flag.Parse()

	if maxDepth < 0 {
		maxDepth = 0xffff
	}

	if _, err := os.Stat(currentPath); err != nil {
		currentPath, _ = os.Getwd()
	}

	fmt.Println("current path:\n" + currentPath)
	fmt.Println(maxDepth)
	err := filepath.Walk(currentPath, func(path string, info fs.FileInfo, err error) error {

		slashIndex := strings.LastIndexByte(path, os.PathSeparator)
		folderStr, filestr := path[:slashIndex], path[slashIndex:]

		// depth := len(strings.Split(folderStr, string(os.PathSeparator))) + 1
		// depth -= len(strings.Split(currentPath, string(os.PathSeparator)))
		depth := 1
		for _, c := range currentPath {
			if c == os.PathSeparator {
				depth --
			}
		}
		for _, c := range folderStr {
			if c == os.PathSeparator {
				depth ++
				if depth > maxDepth {
					return nil
				}
			}
		}

		depthStr := ""
		for i := 0; i < depth; i++ {
			if i == depth - 1 {
				print(printColors[i % len(printColors)](depthEnd))
			} else {
				print(printColors[i % len(printColors)](depthPart))
			}

		}


		str := fmt.Sprintf("%v%v", depthStr, filestr)

		if !info.IsDir() {
			str += fmt.Sprintf("\t\t%v", info.Size())
		}

		str += "\t\t" + fmt.Sprint(info.Mode())

		println(printColors[depth % len(printColors)](str))
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}
