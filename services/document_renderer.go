package services

import (
	"rkl.io/latex-renderer/contract"
	"io/ioutil"
	"os"
	"os/exec"
)

type latexRenderer struct {
}

// NewDocumentRenderer return new document renderer
func NewDocumentRenderer() contract.RendererInterface {
	return &latexRenderer{}
}

// Render render a dto
func (r *latexRenderer) Render(input interface{}) ([]byte, error) {
	srcData := input.([]byte)

	tempDir, err := ioutil.TempDir(os.TempDir(), "renderer-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	srcFilename := "document.tex"
	srcPath := tempDir + string(os.PathSeparator) + srcFilename

	srcFile, err := os.Create(srcPath)
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	_, err = srcFile.Write(srcData)
	if err != nil {
		return nil, err
	}

	workingDir := tempDir + string(os.PathSeparator)

	if err = os.Chdir(workingDir); err != nil {
		return nil, err
	}

	// generate pdf document
	cmd := exec.Command("xelatex", "document")
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	cmd = exec.Command("xelatex", "document")
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	generatedPdfData, err := ioutil.ReadFile(workingDir + string(os.PathSeparator) + "/document.pdf")
	if err != nil {
		return nil, err
	}

	return generatedPdfData, nil
}
