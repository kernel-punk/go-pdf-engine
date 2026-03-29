package examples

import "git.proxeet.com/kernel/pdfgen/pdfgen"

func Renderer(pdf *pdfgen.PDF, data []*ServerTestData) error {

	if err := HeaderRender(pdf); err != nil {
		return err
	}
	return BodyRender(pdf, data)

}
