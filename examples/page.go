package examples

import (
	"bytes"
	"codeberg.org/go-pdf/fpdf"
	"embed"
	"fmt"
	"git.proxeet.com/kernel/pdfgen/pdfgen"
	"time"
)

//go:embed assets/*
var AssetsFS embed.FS

func ColonTitleRender(pdf *pdfgen.PDF) {
	pdf.SetFillColor(DarkTurquoise.R, DarkTurquoise.G, DarkTurquoise.B)
	pdf.Rect(0, 0, 210, 14, "F")

	imgBytes, err := AssetsFS.ReadFile(LogoPath)
	if err != nil {
		return
	}

	pdf.RegisterImageOptionsReader("logo", fpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(imgBytes))
	pdf.ImageOptions("logo", 1, 1, 0, 12, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
	// 0 ширина = авто по пропорции, 12 = высота полоски минус отступ

	pdf.SetFont(MainFont, "B", 16)
	pdf.SetTextColor(White.R, White.G, White.B)
	pdf.SetXY(22, 3)
	pdf.CellFormat(100, 8, "Go pdf generator", "", 0, "L", false, 0, "")

}

func HeaderRender(pdf *pdfgen.PDF) error {

	pdf.SetLeftMargin(10)
	pdf.Ln(20)
	pdf.SetFont(MainFont, "B", 24)
	pdf.CellFormat(0, 1, "Server Diagnostics Report", "", 1, "L", false, 0, "")

	pdf.SetFont(MainFont, "", 12)
	pdf.CellFormat(0, 1, fmt.Sprintf("Generated on: %s", time.Now().Format("2006-01-02 15:04:05")), "", 1, "R", false, 0, "")
	pdf.Ln(10)

	bytesPortUpImg, err := AssetsFS.ReadFile(ServerUpImg)
	if err != nil {
		return err
	}

	bytesPortDownImg, err := AssetsFS.ReadFile(ServerDownImg)
	if err != nil {
		return err
	}

	pdf.RegisterImageOptionsReader("serverUp", fpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(bytesPortUpImg))
	pdf.RegisterImageOptionsReader("serverDown", fpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(bytesPortDownImg))

	return nil
}

func BodyRender(pdf *pdfgen.PDF, data []*ServerTestData) error {

	for i, testResult := range data {
		pdf.SetLeftMargin(10)

		blockHeight := pdf.PointConvert(BaseFontSize)*10 + 80
		if pdf.CheckPageBreak(blockHeight) {

			pdf.SetLeftMargin(10)
			pdf.SetX(10)
		}

		pdf.SetX(10)

		pdf.SetFont(MainFont, "B", 16)
		pdf.CellFormat(0, 5, fmt.Sprintf("Server %s", ServerIpS[i]), "", 1, "L", false, 0, "")

		pageWidth, _ := pdf.GetPageSize()
		leftMargin, _, rightMargin, _ := pdf.GetMargins()
		usableWidth := pageWidth - leftMargin - rightMargin

		pdf.SetLineWidth(0.3)
		color := ColorFunc("HEADER")
		pdf.SetDrawColor(color.R, color.G, color.B)

		x := leftMargin
		y := pdf.GetY() + 5
		pdf.Line(x, y, x+usableWidth, y)
		pdf.Ln(10)

		if testResult.E == "YES" {
			pdf.ImageOptions("serverUp", pdf.GetX(), pdf.GetY()+3, 80, 0, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
		} else {
			pdf.ImageOptions("serverDown", pdf.GetX(), pdf.GetY()+3, 80, 0, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
		}

		pdf.Ln(10)
		pdf.SetX(leftX100)

		cellWidth := float64(len("WEB Server Status:"))*2.1 + 10 // max text

		AddLabeledValue(pdf, "Link:", &cellWidth, testResult.E)

		if testResult.E == "YES" {
			AddLabeledValue(pdf, "Ping:", &cellWidth, fmt.Sprintf("%s ms", testResult.F))
			AddLabeledValue(pdf, "SSD used:", &cellWidth, fmt.Sprintf("%s %%", testResult.G))
			AddLabeledValue(pdf, "RAM used:", &cellWidth, fmt.Sprintf("%s %%", testResult.H))
		} else {
			AddLabeledValue(pdf, "Ping:", &cellWidth, fmt.Sprintf("%s", testResult.F))
			AddLabeledValue(pdf, "SSD used:", &cellWidth, fmt.Sprintf("%s", testResult.G))
			AddLabeledValue(pdf, "RAM used:", &cellWidth, fmt.Sprintf("%s", testResult.H))
		}

		AddLabeledValue(pdf, "WEB Server Status:", &cellWidth, testResult.B)
		AddLabeledValue(pdf, "Need update:", &cellWidth, testResult.A)
		AddLabeledValue(pdf, "OS:", &cellWidth, testResult.C)
		AddLabeledValue(pdf, "Uptime:", &cellWidth, testResult.D)

		pdf.SetX(leftX10)
		pdf.Ln(15)

	}

	return nil
}

func FooterRender(pdf *pdfgen.PDF) {
	pdf.SetY(-15)
	pdf.SetFont(MainFont, "I", 8)
	pdf.CellFormat(0, 10, fmt.Sprintf("%d of {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
}
