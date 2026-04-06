package pdfgen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var ErrRendererRequired = errors.New("pdf renderer is required")

func PdfGenerate[T any](input PdfGenerateData[T]) (string, error) {
	if input.Renderer == nil {
		return "", ErrRendererRequired
	}

	if input.PdfName == "" {
		input.PdfName = "report"
	}
	if input.TimeFormat == "" {
		input.TimeFormat = "20060102_150405"
	}

	pdf := NewPDF(input.PdfConfig)

	if input.BeforeFirstPage != nil {
		err := input.BeforeFirstPage(pdf)
		if err != nil {
			return "", err
		}
	}

	pdf.AddPage()

	if input.AfterFirstPage != nil {
		err := input.AfterFirstPage(pdf)
		if err != nil {
			return "", err
		}
	}

	err := input.Renderer(pdf, input.Data)
	if err != nil {
		return "", err
	}

	outputFileName := fmt.Sprintf("%s_%s.pdf", input.PdfName, time.Now().Format(input.TimeFormat)) //20060102_150405

	if input.OutDir != "" {
		err = os.MkdirAll(input.OutDir, 0o755)
		if err != nil {
			return "", err
		}
	}

	outputPath := filepath.Join(input.OutDir, outputFileName)

	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
