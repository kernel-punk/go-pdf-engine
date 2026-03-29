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
- Custom renderers — pass your own `func(pdf, data)` render function
- Custom page header/footer via `PDFConfig`
- One-time report init hook for registering images and drawing a title block
- Embedded assets via `embed.FS`
- Configurable output directory, filename, and timestamp format

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
mkdir -p output
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
```go
err := pdfgen.PdfGenerate(pdfgen.PdfGenerateData[[]*examples.ServerTestData]{
    OutDir:     outputDir,
    Data:       data,
    Init:       examples.HeaderRender,
    Renderer:   examples.Renderer,
    PdfName:    "server_report",
    TimeFormat: "20060102_150405",
    PdfConfig: pdfgen.PDFConfig{
        Header: examples.ColonTitleRender,
        Footer: examples.FooterRender,
    },
})
```


## Credits

- Server icon by OpenClipart-Vectors from Pixabay
- PDF icon by OpenClipart-Vectors from Pixabay

## License

MIT