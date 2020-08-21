package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	_ "github.com/casept/esponda/statik"

	"github.com/rakyll/statik/fs"
)

// Extracts the image template and font to a temporary directory and return the path to this directory.
func extractAssetsIfNeeded() string {
	// Create dir if it doesn't exist
	dirPath := path.Join(os.TempDir(), "espondabot-assets")
	_, err := os.Stat(dirPath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Creating directory and extracting assets")
		err = os.Mkdir(dirPath, 0700)
		if err != nil {
			panic(err)
		}

		// And extract assets into it
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}
		files := []string{"SchillerBlank.png", "Raleway-ExtraBold.ttf", "EspondaBlank.png"}
		for _, file := range files {
			r, err := statikFS.Open(strings.Join([]string{"/", file}, ""))
			if err != nil {
				panic(err)
			}
			defer r.Close()

			tempFile, err := os.Create(path.Join(dirPath, file))
			if err != nil {
				panic(err)
			}
			defer tempFile.Close()
			_, err = io.Copy(tempFile, r)
			if err != nil {
				panic(err)
			}
		}
	}

	return dirPath
}
