package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path"
)

// ImageAnnotationRecipe describes how a given image should be annotated.
type ImageAnnotationRecipe struct {
	// The name of the base image, as it's called in the embedded FS.
	InputImage string
	// The name of the font, as it's called in the embedded FS.
	CaptionFont string
	// Color of the text.
	CaptionColor string
	// The caption text.
	CaptionText string
	// Size of the text.
	TextDimensions string
	// Geometry of the text (whatever that means)
	TextGeometry string
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// Annotate turns the recipe into an image and returns the path to it.
func (r *ImageAnnotationRecipe) Annotate() string {
	// Generate a temporary file with random name

	tmpFile := path.Join(os.TempDir(), fmt.Sprintf("espondabot-%s.png", randomString(60)))
	assetsPath := extractAssetsIfNeeded()

	// Actual image annotation
	cmd := exec.Command("convert", path.Join(assetsPath, r.InputImage),
		"-size", r.TextDimensions,
		"-background", "none",
		"-font", path.Join(assetsPath, r.CaptionFont),
		"-fill", fmt.Sprintf("\"%s\"", r.CaptionColor),
		fmt.Sprintf("caption:\"%s\"", r.CaptionText),
		"-geometry", r.TextGeometry,
		"-composite", tmpFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(cmd.String())
		fmt.Println(string(out))
		panic(err)
	}
	return tmpFile
}
