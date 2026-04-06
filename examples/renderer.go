package examples

import (
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
)

func Renderer(pdf *pdfgen.PDF, data []*ServerTestData) error {
	return BodyRender(pdf, data)
}
