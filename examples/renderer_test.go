package examples

import (
	"path/filepath"
	"testing"

	"github.com/kernel-punk/go-pdf-engine/pdfgen"
)

func TestRendererSupportsMoreRowsThanSeedServerList(t *testing.T) {
	t.Parallel()

	_, err := pdfgen.PdfGenerate(pdfgen.PdfGenerateData[[]*ServerTestData]{
		OutDir:          filepath.Join(t.TempDir(), "reports"),
		Data:            MultipleRandomTests(len(ServerIpS) + 1),
		BeforeFirstPage: InitReportAssets,
		AfterFirstPage:  ReportHeaderRender,
		Renderer:        Renderer,
		PdfName:         "servers",
		TimeFormat:      "20060102",
		PdfConfig: pdfgen.PDFConfig{
			Header:      PageHeaderRender,
			Footer:      FooterRender,
			OnPageBreak: PageBreakRender,
		},
	})
	if err != nil {
		t.Fatalf("PdfGenerate returned error: %v", err)
	}
}
