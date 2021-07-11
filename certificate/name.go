package certificate

import "github.com/phpdave11/gofpdf"

func addName(pdf *gofpdf.Fpdf, style string, pos *position, name string) {

	pdf.SetFont("times", "B", 48)
	_, lineHt := pdf.GetFontSize()
	if style == "circle" {
		changePosRel(pdf, -20, lineHt*1.35, pos)
		pdf.MultiCell(380, lineHt, name, "", gofpdf.AlignCenter, false)
	} else {
		changePosRel(pdf, 0, lineHt*1.45, pos)
		pdf.WriteAligned(0, lineHt, name, gofpdf.AlignCenter)

	}
}
