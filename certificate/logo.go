package certificate

import (
	"github.com/google/uuid"
	"github.com/phpdave11/gofpdf"
	"mime/multipart"
)

func addLogo(pdf *gofpdf.Fpdf, style string, f multipart.File, imgT string) {

	logoId := uuid.New().String()
	var opt gofpdf.ImageOptions
	opt.ImageType = imgT
	opt.AllowNegativePosition = true
	pdf.RegisterImageOptionsReader(logoId, opt, f)

	if style == "circle" {
		pdf.ImageOptions(logoId, 55, 237, 100, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
	} else {
		pdf.ImageOptions(logoId, 340, 400, 100, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
	}

}
