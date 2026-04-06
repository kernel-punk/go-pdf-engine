package pdfgen

import "codeberg.org/go-pdf/fpdf"

type PdfGenerateData[T any] struct {
	OutDir          string
	Data            T
	BeforeFirstPage func(pdf *PDF) error
	AfterFirstPage  func(pdf *PDF) error
	Renderer        func(pdf *PDF, data T) error
	PdfConfig       PDFConfig
	PdfName         string
	TimeFormat      string
}

type PDF struct {
	*fpdf.Fpdf
	bottomMargin float64
	onPageBreak  func(p *PDF)
}

type PDFMargins struct {
	Left  float64
	Top   float64
	Right float64
}

type PDFConfig struct {
	Orientation  string
	Unit         string
	Size         string
	Margins      *PDFMargins
	BottomMargin float64
	Header       func(p *PDF)
	Footer       func(p *PDF)
	OnPageBreak  func(p *PDF)
}
