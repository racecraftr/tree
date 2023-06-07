package main

import (
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

var currentPath = "/Users/avi/code/github/racecraftr/gradzilla"

func main() {
	fmt.Println("current path:\n" + currentPath)
	err := filepath.Walk(currentPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		slashIndex := strings.LastIndexByte(path, os.PathSeparator)
		folderStr, filestr := path[:slashIndex], path[slashIndex:]

		depth := len(strings.Split(folderStr, string(os.PathSeparator))) + 1
		depth -= len(strings.Split(currentPath, string(os.PathSeparator)))

		depthStr := ""
		for i := 0; i < depth; i++ {
			print(printColors[i % len(printColors)](depthPart))
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
