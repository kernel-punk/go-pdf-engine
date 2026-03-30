package examples

import (
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
)

func Renderer(pdf *pdfgen.PDF, data []*ServerTestData) error {

	if err := HeaderRender(pdf); err != nil {
		return err
	}
	return BodyRender(pdf, data)

}
