package pdfgen

import "codeberg.org/go-pdf/fpdf"

type PDFRenderer interface {
	Render(pdf *PDF) error
}

type PdfGenerateData[T any] struct {
	OutDir     string
	Data       T
	Renderer   func(pdf *PDF, data T) error
	PdfConfig  PDFConfig
	PdfName    string
	TimeFormat string
}

type PDF struct {
	*fpdf.Fpdf
	bottomMargin float64
}

type PDFConfig struct {
	Header func(p *PDF)
	Footer func(p *PDF)
}
