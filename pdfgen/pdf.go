package pdfgen

import (
	"codeberg.org/go-pdf/fpdf"
)

func NewPDF(cfg PDFConfig) *PDF {

	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.SetMargins(0, 20, 20)
	marginBottom := 20.0
	pdf.SetAutoPageBreak(true, marginBottom)
	pdf.AliasNbPages("")

	p := &PDF{Fpdf: pdf, bottomMargin: marginBottom}

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
		p.Ln(30)

		return true
	}
	return false
}
