package pdfgen

import (
	"codeberg.org/go-pdf/fpdf"
)

const (
	defaultOrientation  = "P"
	defaultUnit         = "mm"
	defaultSize         = "A4"
	defaultBottomMargin = 20.0
)

var defaultMargins = PDFMargins{
	Left:  0,
	Top:   20,
	Right: 20,
}

func NewPDF(cfg PDFConfig) *PDF {
	orientation := cfg.Orientation
	if orientation == "" {
		orientation = defaultOrientation
	}

	unit := cfg.Unit
	if unit == "" {
		unit = defaultUnit
	}

	size := cfg.Size
	if size == "" {
		size = defaultSize
	}

	margins := defaultMargins
	if cfg.Margins != nil {
		margins = *cfg.Margins
	}

	marginBottom := cfg.BottomMargin
	if marginBottom <= 0 {
		marginBottom = defaultBottomMargin
	}

	pdf := fpdf.New(orientation, unit, size, "")
	pdf.SetMargins(margins.Left, margins.Top, margins.Right)
	pdf.SetAutoPageBreak(true, marginBottom)
	pdf.AliasNbPages("")

	p := &PDF{
		Fpdf:         pdf,
		bottomMargin: marginBottom,
		onPageBreak:  cfg.OnPageBreak,
	}

	if cfg.Header != nil {
		p.SetHeaderFunc(func() { cfg.Header(p) })
	}

	if cfg.Footer != nil {
		p.SetFooterFunc(func() { cfg.Footer(p) })
	}

	return p
}

func (p *PDF) GetPageHeight() float64 {
	_, h, _ := p.PageSize(p.PageNo())
	return h
}

func (p *PDF) MarginBottom() float64 {
	return p.bottomMargin
}

func (p *PDF) CheckPageBreak(height float64) bool {
	if p.GetY()+height > p.GetPageHeight()-p.MarginBottom() {
		p.AddPage()
		if p.onPageBreak != nil {
			p.onPageBreak(p)
		}

		return true
	}
	return false
}
