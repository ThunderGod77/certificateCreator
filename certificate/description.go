package certificate

import "github.com/phpdave11/gofpdf"

func addDescription(pdf *gofpdf.Fpdf, style string, pos *position, description string) {
	w, _ := pdf.GetPageSize()
	if style == "circle" {
		_, lineHt := pdf.GetFontSize()

		changePosRel(pdf, -60, lineHt*1.5, pos)
		pdf.SetFont("times", "", 25)
		_, lineHt = pdf.GetFontSize()
		pdf.MultiCell(546, lineHt, description, "", gofpdf.AlignCenter, false)
	} else {
		_, lineHt := pdf.GetFontSize()

		changePosRel(pdf, 22, lineHt*1.5, pos)
		pdf.SetFont("times", "", 25)
		_, lineHt = pdf.GetFontSize()
		pdf.MultiCell(w-39, lineHt, description, "", gofpdf.AlignCenter, false)
	}
}
