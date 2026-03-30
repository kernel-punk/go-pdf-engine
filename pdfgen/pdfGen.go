package pdfgen

import (
	"fmt"
	"time"
)

func PdfGenerate[T any](input PdfGenerateData[T]) (string, error) {

	if input.PdfName == "" {
		input.PdfName = "report"
	}
	if input.TimeFormat == "" {
		input.TimeFormat = "20060102_150405"
	}

	pdf := NewPDF(input.PdfConfig)
	pdf.AddPage()

	err := input.Renderer(pdf, input.Data)
	if err != nil {
		return "", err
	}

	outputFileName := fmt.Sprintf("%s_%s.pdf", input.PdfName, time.Now().Format(input.TimeFormat)) //20060102_150405

	err = pdf.OutputFileAndClose(input.OutDir + outputFileName)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
