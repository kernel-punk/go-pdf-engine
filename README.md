# pdfGen

A generic PDF report generation library for Go.

## Overview

`pdfgen` is a lightweight, type-safe PDF generation engine.
The library is decoupled from report content — you bring your own data types,
rendering logic, and assets. The engine handles page lifecycle, header/footer
callbacks, and file output.

A ready-to-run example (Server Diagnostics Report) is included in `examples/`.

## Features

- Generic engine — works with any data type via `PdfGenerateData[T]`
- Split lifecycle hooks before and after the first page
- Custom renderers — pass your own `func(pdf, data)` render function
- Custom page header/footer via `PDFConfig`
- Custom page-break callback via `PDFConfig`
- Configurable page format, units, margins, and bottom margin
- Embedded assets via `embed.FS`
- Configurable output directory, filename, and timestamp format
- Output directory is created automatically when needed

## Project Structure
```
pdfgen/       — core engine (PDF lifecycle, page breaks, config)
examples/     — example renderer: Server Diagnostics Report
main.go       — entry point demonstrating library usage
```

## Quick Start
```bash
git clone https://github.com/kernel-punk/go-pdf-engine
cd go-pdf-engine
go build -o ./build/pdfGen .
./build/pdfGen
```

Or without building:
```bash
go run .
```

Configure output directory via `.env`:
```
PDF_OUTPUT_DIR=/tmp
PDF_OUTPUT_NAME=server_report
```

## Usage
```text
package main

import (
	"fmt"
	"log"

	"github.com/kernel-punk/go-pdf-engine/examples"
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
)

func main() {
	outputDir := "./output"
	data := examples.MultipleRandomTests(10)

	fileName, err := pdfgen.PdfGenerate(pdfgen.PdfGenerateData[[]*examples.ServerTestData]{
		OutDir:          outputDir,
		Data:            data,
		BeforeFirstPage: examples.InitReportAssets,
		AfterFirstPage:  examples.ReportHeaderRender,
		Renderer:        examples.Renderer,
		PdfName:         "server_report",
		TimeFormat:      "20060102_150405",
		PdfConfig: pdfgen.PDFConfig{
			Header:      examples.PageHeaderRender,
			Footer:      examples.FooterRender,
			OnPageBreak: examples.PageBreakRender,
			Margins: &pdfgen.PDFMargins{
				Left:  10,
				Top:   20,
				Right: 20,
			},
			BottomMargin: 20,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("created:", fileName)
}
```


## Credits

- Server icon by OpenClipart-Vectors from Pixabay
- PDF icon by OpenClipart-Vectors from Pixabay

## License

MIT
