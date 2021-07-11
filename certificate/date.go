package certificate

import "github.com/phpdave11/gofpdf"

func addDate(pdf *gofpdf.Fpdf, style string, pos *position) {
	if style == "circle" {
		// to make bold line under date
		pdf.SetFillColor(100, 100, 100)
		pdf.Rect(237, 490, 198, 2, "F")

		// to enter "date"
		changePosAbs(pdf, 316, 495, pos)
		pdf.SetFont("times", "", 18)
		pdf.SetTextColor(100, 100, 100)
		_, lineHt := pdf.GetFontSize()
		pdf.WriteAligned(30, lineHt, "Date", gofpdf.AlignCenter)

		// write the date of certificate presented
		changePosAbs(pdf, 282, 465, pos)
		pdf.SetFont("times", "", 22)
		pdf.SetTextColor(150, 150, 150)
		_, lineHt = pdf.GetFontSize()
		pdf.WriteAligned(45, lineHt, " 12/03/2019 ", gofpdf.AlignCenter)

	} else {
		// to make bold line under date
		pdf.SetFillColor(100, 100, 100)
		pdf.Rect(79, 490, 198, 2, "F")

		// to enter "date"
		changePosAbs(pdf, 159, 495, pos)
		pdf.SetFont("times", "", 18)
		pdf.SetTextColor(100, 100, 100)
		_, lineHt := pdf.GetFontSize()
		pdf.WriteAligned(30, lineHt, "Date", gofpdf.AlignCenter)

		// write out the date of certification
		changePosAbs(pdf, 123, 465, pos)
		pdf.SetFont("times", "", 22)
		pdf.SetTextColor(150, 150, 150)
		_, lineHt = pdf.GetFontSize()
		pdf.WriteAligned(45, lineHt, " 12/03/2019 ", gofpdf.AlignCenter)
	}
}
