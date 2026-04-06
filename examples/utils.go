package examples

import (
	"fmt"
	"github.com/kernel-punk/go-pdf-engine/pdfgen"
	"time"
)

func StatusColor(value string) RGB {
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

func LinkColor(linkUp bool) RGB {
	if linkUp {
		return PureGreen
	}

	return Crimson
}

func UsageColor(value *int) RGB {
	if value == nil {
		return ColorDefault
	}

	switch {
	case *value <= 50:
		return PureGreen
	case *value <= 75:
		return Amber
	case *value <= 100:
		return Crimson
	default:
		return ColorDefault
	}
}

func PingColor(value *int) RGB {
	if value == nil {
		return ColorDefault
	}

	switch {
	case *value <= 100:
		return PureGreen
	case *value <= 250:
		return Amber
	default:
		return Crimson
	}
}

func AddLabeledValue(pdf *pdfgen.PDF, label string, labelWidth *float64, value string, valueColor RGB) {

	startX := pdf.GetX()
	pdf.SetFont(MainFont, "B", 14)

	if labelWidth == nil {
		value := float64(len(label))*2.1 + 10
		labelWidth = &value
	}

	pdf.CellFormat(*labelWidth, 5, label, "", 0, "L", false, 0, "")

	pdf.SetTextColor(valueColor.R, valueColor.G, valueColor.B)

	pdf.SetFont(MainFont, "", 14)
	pdf.CellFormat(0, 5, value, "", 1, "L", false, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(4)
	pdf.SetX(startX)
}

func randomUptime(rng interface {
	Intn(n int) int
}) time.Duration {
	hours := rng.Intn(720)
	minutes := rng.Intn(60)
	seconds := rng.Intn(60)

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
}

func formatLinkStatus(linkUp bool) string {
	if linkUp {
		return "YES"
	}

	return "NO"
}

func formatOptionalInt(value *int, suffix string) string {
	if value == nil {
		return "N/A"
	}

	if suffix == "" {
		return fmt.Sprintf("%d", *value)
	}

	return fmt.Sprintf("%d %s", *value, suffix)
}

func formatOptionalString(value string) string {
	if value == "" {
		return "N/A"
	}

	return value
}

func formatOptionalDuration(value *time.Duration) string {
	if value == nil {
		return "N/A"
	}

	totalSeconds := int(value.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func defaultServerName(index int) string {
	if index < len(ServerIpS) {
		return ServerIpS[index]
	}

	offset := index - len(ServerIpS)

	return fmt.Sprintf("10.255.%d.%d", (offset/254)+1, (offset%254)+1)
}

func serverName(testResult *ServerTestData, index int) string {
	if testResult != nil && testResult.Server != "" {
		return testResult.Server
	}

	return defaultServerName(index)
}
