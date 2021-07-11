package certificate

import "github.com/phpdave11/gofpdf"

func makeSubHeading(pdf *gofpdf.Fpdf, style string, pos *position) {
	if style == "circle" {
		_, lineHt := pdf.GetFontSize()
		pdf.MoveTo(0, 100+lineHt*2)
		changePosRel(pdf, 100, lineHt*2.75, pos)
		pdf.SetFont("times", "", 28)
		_, lineHt = pdf.GetFontSize()
		pdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	} else {
		_, lineHt := pdf.GetFontSize()
		pdf.MoveTo(0, 100+lineHt*2)
		changePosRel(pdf, 0, lineHt*1.75, pos)
		pdf.SetFont("times", "", 28)
		_, lineHt = pdf.GetFontSize()
		pdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	}

}
