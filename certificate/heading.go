package certificate

import "github.com/phpdave11/gofpdf"

func makeHeading(pdf *gofpdf.Fpdf, style string, pos *position) {
	pdf.SetFont("times", "B", 40)
	pdf.SetTextColor(40, 40, 40)
	_, lineHt := pdf.GetFontSize()
	if style == "circle" {
		changePosAbs(pdf, 198, 100, pos)
		pdf.MultiCell(514.0, lineHt*1.125, "CERTIFICATE OF COMPLETION", "", gofpdf.AlignCenter, false)
	} else {
		changePosAbs(pdf, 0, 100, pos)
		pdf.WriteAligned(0, lineHt, "CERTIFICATE OF COMPLETION", gofpdf.AlignCenter)
	}

}
