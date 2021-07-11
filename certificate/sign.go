package certificate

import (
	"bytes"
	"github.com/phpdave11/gofpdf"
	"io"
	"log"
	"mime/multipart"
)

func addSign(pdf *gofpdf.Fpdf, style string, pos *position, sF multipart.File) {

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, sF); err != nil {
		log.Println(err)
	}

	pdf.Rect(514, 490, 198, 2, "F")

	changePosAbs(pdf, 574, 495, pos)
	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt := pdf.GetFontSize()
	pdf.WriteAligned(0, lineHt, "Instructor", gofpdf.AlignCenter)

	sig, err := gofpdf.SVGBasicParse(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	changePosAbs(pdf, 534, 445, pos)
	pdf.SVGBasicWrite(&sig, 0.5)

}
