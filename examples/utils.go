package examples

import (
	"fmt"
	"git.proxeet.com/kernel/pdfgen/pdfgen"
	"math/rand"
	"strconv"
	"strings"
)

func ColorFunc(value string) RGB {
	fields := strings.Fields(value)
	if len(fields) > 0 {
		if num, err := strconv.Atoi(fields[0]); err == nil {
			switch {
			case num <= 50:
				return PureGreen
			case num <= 75:
				return Amber
			case num <= 100:
				return Crimson
			}
		}
	}
	switch value {
	case "Full", "UP", "OK", "YES", "WORK", "Not Required":
		return PureGreen

	case "DOWN", "ERROR", "NO", "Required":
		return Crimson

	case "HEADER":
		return DarkTurquoise

	default:
		return ColorDefault
	}
}

func AddLabeledValue(pdf *pdfgen.PDF, label string, labelWidth *float64, value string) {

	startX := pdf.GetX()
	pdf.SetFont(MainFont, "B", 14)

	if labelWidth == nil {
		value := float64(len(label))*2.1 + 10
		labelWidth = &value
	}

	pdf.CellFormat(*labelWidth, 5, label, "", 0, "L", false, 0, "")

	color := ColorFunc(value)
	pdf.SetTextColor(color.R, color.G, color.B)

	pdf.SetFont(MainFont, "", 14)
	pdf.CellFormat(0, 5, value, "", 1, "L", false, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(4)
	pdf.SetX(startX)
}

func randomUptime() string {
	hours := rand.Intn(720)
	minutes := rand.Intn(60)
	seconds := rand.Intn(60)

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
