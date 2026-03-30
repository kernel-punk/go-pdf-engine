package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kernel-punk/go-pdf-engine/examples"
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
	"log"
	"os"
)

func main() {

	_ = godotenv.Load()

	fmt.Println("Generating PDF with logo and table...")

	outputDir := os.Getenv("PDF_OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "output/"
		fmt.Println("PDF out dir not set. Use ./output by default")
	} else {
		fmt.Printf("PDF out dir is %s (from .env)\n", outputDir)

	}

	pdfName := os.Getenv("PDF_OUTPUT_NAME")
	if pdfName == "" {
		pdfName = "pdf"
	}

	data := examples.MultipleRandomTests(10) // []*examples.ServerTestData

	err := pdfgen.PdfGenerate(pdfgen.PdfGenerateData[[]*examples.ServerTestData]{
		OutDir:     outputDir,
		Data:       data,
		Renderer:   examples.Renderer,
		PdfName:    pdfName,
		TimeFormat: "20060102_150405",
		PdfConfig: pdfgen.PDFConfig{
			Header: examples.ColonTitleRender,
			Footer: examples.FooterRender,
		},
	})
	if err != nil {
		log.Println("PDF generation failed:", err)
	} else {
		fmt.Println("PDF created successfully!")
	}
}
