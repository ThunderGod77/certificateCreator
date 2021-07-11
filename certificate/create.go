package certificate

import (
	"github.com/phpdave11/gofpdf"
	"image/color"
	"mime/multipart"
)

func CreateCertificate(f multipart.File, imgT string, sF multipart.File, name, description, style string) *gofpdf.Fpdf {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	pos := position{
		x: 0,
		y: 0,
	}
	primaryColor := color.RGBA{
		R: 102,
		G: 178,
		B: 178,
		A: 255,
	}
	secondaryColor := color.RGBA{
		R: 178,
		G: 216,
		B: 216,
		A: 255,
	}

	pdf.AddPage()
	makeBorder(pdf, style, primaryColor, secondaryColor)
	makeHeading(pdf, style, &pos)
	makeSubHeading(pdf, style, &pos)
	addName(pdf, style, &pos, name)
	addDescription(pdf, style, &pos, description)
	addDate(pdf, style, &pos)
	addSign(pdf, style, &pos, sF)
	addLogo(pdf, style, f, imgT)

	return pdf
}
