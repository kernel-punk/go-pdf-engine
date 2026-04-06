package examples

import (
	"bytes"
	"codeberg.org/go-pdf/fpdf"
	"embed"
	"fmt"
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
	"time"
)

//go:embed assets/*
var AssetsFS embed.FS

func PageHeaderRender(pdf *pdfgen.PDF) {
	pdf.SetFillColor(DarkTurquoise.R, DarkTurquoise.G, DarkTurquoise.B)
	pdf.Rect(0, 0, 210, 14, "F")

	if pdf.GetImageInfo("logo") != nil {
		pdf.ImageOptions("logo", 1, 1, 0, 12, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
	}

	pdf.SetFont(MainFont, "B", 16)
	pdf.SetTextColor(White.R, White.G, White.B)
	pdf.SetXY(22, 3)
	pdf.CellFormat(100, 8, "Go pdf generator", "", 0, "L", false, 0, "")
	pdf.SetTextColor(ColorDefault.R, ColorDefault.G, ColorDefault.B)
}

func InitReportAssets(pdf *pdfgen.PDF) error {
	logoBytes, err := AssetsFS.ReadFile(LogoPath)
	if err != nil {
		return err
	}

	pdf.RegisterImageOptionsReader("logo", fpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(logoBytes))

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

func ReportHeaderRender(pdf *pdfgen.PDF) error {
	pdf.SetLeftMargin(10)
	pdf.Ln(20)
	pdf.SetFont(MainFont, "B", 24)
	pdf.SetTextColor(ColorDefault.R, ColorDefault.G, ColorDefault.B)
	pdf.CellFormat(0, 1, "Server Diagnostics Report", "", 1, "L", false, 0, "")

	pdf.SetFont(MainFont, "", 12)
	pdf.CellFormat(0, 1, fmt.Sprintf("Generated on: %s", time.Now().Format("2006-01-02 15:04:05")), "", 1, "R", false, 0, "")
	pdf.Ln(10)

	return nil
}

func PageBreakRender(pdf *pdfgen.PDF) {
	pdf.SetLeftMargin(10)
	pdf.SetX(10)
	pdf.SetY(reportBodyStartY)
}

func BodyRender(pdf *pdfgen.PDF, data []*ServerTestData) error {

	for i, testResult := range data {
		if testResult == nil {
			testResult = &ServerTestData{}
		}

		pdf.SetLeftMargin(10)

		blockHeight := pdf.PointConvert(BaseFontSize)*10 + 80
		if pdf.CheckPageBreak(blockHeight) {

			pdf.SetLeftMargin(10)
			pdf.SetX(10)
		}

		pdf.SetX(10)

		pdf.SetFont(MainFont, "B", 16)
		pdf.CellFormat(0, 5, fmt.Sprintf("Server %s", serverName(testResult, i)), "", 1, "L", false, 0, "")

		pageWidth, _ := pdf.GetPageSize()
		leftMargin, _, rightMargin, _ := pdf.GetMargins()
		usableWidth := pageWidth - leftMargin - rightMargin

		pdf.SetLineWidth(0.3)
		color := DarkTurquoise
		pdf.SetDrawColor(color.R, color.G, color.B)

		x := leftMargin
		y := pdf.GetY() + 5
		pdf.Line(x, y, x+usableWidth, y)
		pdf.Ln(10)

		if testResult.LinkUp {
			pdf.ImageOptions("serverUp", pdf.GetX(), pdf.GetY()+3, 80, 0, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
		} else {
			pdf.ImageOptions("serverDown", pdf.GetX(), pdf.GetY()+3, 80, 0, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")
		}

		pdf.Ln(10)
		pdf.SetX(leftX100)

		cellWidth := float64(len("WEB Server Status:"))*2.1 + 10 // max text

		AddLabeledValue(pdf, "Link:", &cellWidth, formatLinkStatus(testResult.LinkUp), LinkColor(testResult.LinkUp))

		AddLabeledValue(pdf, "Ping:", &cellWidth, formatOptionalInt(testResult.PingMS, "ms"), PingColor(testResult.PingMS))
		AddLabeledValue(pdf, "SSD used:", &cellWidth, formatOptionalInt(testResult.SSDUsedPercent, "%"), UsageColor(testResult.SSDUsedPercent))
		AddLabeledValue(pdf, "RAM used:", &cellWidth, formatOptionalInt(testResult.RAMUsedPercent, "%"), UsageColor(testResult.RAMUsedPercent))

		AddLabeledValue(pdf, "WEB Server Status:", &cellWidth, formatOptionalString(testResult.WebServerState), StatusColor(testResult.WebServerState))
		AddLabeledValue(pdf, "Need update:", &cellWidth, formatOptionalString(testResult.NeedUpdate), StatusColor(testResult.NeedUpdate))
		AddLabeledValue(pdf, "OS:", &cellWidth, formatOptionalString(testResult.OperatingSystem), ColorDefault)
		AddLabeledValue(pdf, "Uptime:", &cellWidth, formatOptionalDuration(testResult.Uptime), ColorDefault)

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
