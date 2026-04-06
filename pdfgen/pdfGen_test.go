package pdfgen

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestPdfGenerateReturnsErrorWhenRendererMissing(t *testing.T) {
	t.Parallel()

	_, err := PdfGenerate(PdfGenerateData[string]{Data: "payload"})
	if !errors.Is(err, ErrRendererRequired) {
		t.Fatalf("expected ErrRendererRequired, got %v", err)
	}
}

func TestPdfGenerateCreatesOutputDirAndWritesFile(t *testing.T) {
	t.Parallel()

	outDir := filepath.Join(t.TempDir(), "reports")

	fileName, err := PdfGenerate(PdfGenerateData[string]{
		OutDir:     outDir,
		Data:       "payload",
		Renderer:   func(pdf *PDF, data string) error { return nil },
		PdfName:    "report",
		TimeFormat: "20060102",
	})
	if err != nil {
		t.Fatalf("PdfGenerate returned error: %v", err)
	}

	outputPath := filepath.Join(outDir, fileName)
	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("expected generated pdf at %s: %v", outputPath, err)
	}
}

func TestPdfGenerateRunsLifecycleHooksInOrder(t *testing.T) {
	t.Parallel()

	var sequence []string

	_, err := PdfGenerate(PdfGenerateData[string]{
		OutDir: filepath.Join(t.TempDir(), "reports"),
		Data:   "payload",
		BeforeFirstPage: func(pdf *PDF) error {
			sequence = append(sequence, "before")
			return nil
		},
		AfterFirstPage: func(pdf *PDF) error {
			sequence = append(sequence, "after")
			return nil
		},
		Renderer: func(pdf *PDF, data string) error {
			sequence = append(sequence, "render")
			return nil
		},
		PdfName:    "report",
		TimeFormat: "20060102",
	})
	if err != nil {
		t.Fatalf("PdfGenerate returned error: %v", err)
	}

	want := []string{"before", "after", "render"}
	if len(sequence) != len(want) {
		t.Fatalf("expected %v lifecycle calls, got %v", want, sequence)
	}
	for i := range want {
		if sequence[i] != want[i] {
			t.Fatalf("expected lifecycle order %v, got %v", want, sequence)
		}
	}
}

func TestPdfGenerateReturnsBeforeFirstPageError(t *testing.T) {
	t.Parallel()

	beforeErr := errors.New("before first page failed")

	_, err := PdfGenerate(PdfGenerateData[string]{
		Data: "payload",
		BeforeFirstPage: func(pdf *PDF) error {
			return beforeErr
		},
		Renderer: func(pdf *PDF, data string) error {
			t.Fatal("renderer should not run when before first page fails")
			return nil
		},
	})
	if !errors.Is(err, beforeErr) {
		t.Fatalf("expected before first page error, got %v", err)
	}
}

func TestPdfGenerateReturnsAfterFirstPageError(t *testing.T) {
	t.Parallel()

	afterErr := errors.New("after first page failed")

	_, err := PdfGenerate(PdfGenerateData[string]{
		Data: "payload",
		AfterFirstPage: func(pdf *PDF) error {
			return afterErr
		},
		Renderer: func(pdf *PDF, data string) error {
			t.Fatal("renderer should not run when after first page fails")
			return nil
		},
	})
	if !errors.Is(err, afterErr) {
		t.Fatalf("expected after first page error, got %v", err)
	}
}

func TestCheckPageBreakInvokesConfiguredCallback(t *testing.T) {
	t.Parallel()

	callbackCalled := false

	pdf := NewPDF(PDFConfig{
		OnPageBreak: func(p *PDF) {
			callbackCalled = true
			p.SetY(42)
		},
	})
	pdf.AddPage()
	pdf.SetY(pdf.GetPageHeight() - pdf.MarginBottom() - 1)

	if !pdf.CheckPageBreak(5) {
		t.Fatal("expected page break to be triggered")
	}
	if !callbackCalled {
		t.Fatal("expected OnPageBreak callback to run")
	}
	if got := pdf.GetY(); got != 42 {
		t.Fatalf("expected callback to reposition content, got y=%v", got)
	}
}

func TestNewPDFUsesConfiguredPageSettings(t *testing.T) {
	t.Parallel()

	pdf := NewPDF(PDFConfig{
		Orientation: "L",
		Unit:        "pt",
		Size:        "Letter",
		Margins: &PDFMargins{
			Left:  12,
			Top:   18,
			Right: 24,
		},
		BottomMargin: 36,
	})

	left, top, right, _ := pdf.GetMargins()
	if left != 12 || top != 18 || right != 24 {
		t.Fatalf("expected margins 12/18/24, got %v/%v/%v", left, top, right)
	}
	if got := pdf.MarginBottom(); got != 36 {
		t.Fatalf("expected bottom margin 36, got %v", got)
	}
}
