package certificate

import (
	"github.com/phpdave11/gofpdf"
	"image/color"
)

func makeBorder(pdf *gofpdf.Fpdf, style string, pColor color.RGBA, sColor color.RGBA) {
	w, h := pdf.GetPageSize()
	//to draw triangular border like https://github.com/gophercises/pdf
	if style == "triangle" {

		//to draw upper triangle of secondary color
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0, 0}, {0, h / 9.0}, {w - w/6.0, 0}}, "F")

		//to draw upper triangle of primary color
		drawPolygon(pdf, pColor, []gofpdf.PointType{{w, 0}, {w, h / 9.0}, {w / 18.0, 0}}, "F")

		//to draw lower triangle of secondary color
		drawPolygon(pdf, sColor, []gofpdf.PointType{{w, h}, {w / 6.0, h}, {w, h - h/9.0}}, "F")

		// to draw lower triangle of primary color
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0, h}, {0, h - h/9.0}, {w - w/18.0, h}}, "F")

		// to draw uniform border all around
	} else if style == "square" {

		// to draw upper border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0.0, 0.0}, {w, 0}, {w, 24}, {0, 24}}, "F")

		//to draw left border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0.0, 0.0}, {24, 0}, {24, h}, {0, h}}, "F")

		//to draw bottom border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0, h}, {0, h - 24}, {w, h - 24}, {w, h}}, "F")

		//to draw right border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{w, 0.0}, {w - 24, 0}, {w - 24, h}, {w, h}}, "F")

		// to draw borders like https://graphicriver.net/item/certificate-template/31516409
	} else if style == "squareAdvanced" {

		// to draw upper border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0.0, 0.0}, {w, 0}, {w, 18}, {0, 18}}, "F")

		//to draw left border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0.0, 0.0}, {18, 0}, {18, h - h/3.0}, {0, h - h/3.0}}, "F")

		// to draw left (increased width) border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{18, 356}, {46, h - (h/3.0 + 1)}, {46, h}, {0, h}, {0.0, h - (h/3.0 + 1)}}, "F")

		//to draw right border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{w, 0.0}, {w - 18, 0}, {w - 18, h - h/3.0}, {w, h - h/3.0}}, "F")

		//to draw right (increased width) border

		drawPolygon(pdf, sColor, []gofpdf.PointType{{w - 18, 356}, {w - 46, h - (h/3.0 + 1)}, {w - 46, h}, {w, h}, {w, h - (h/3.0 + 1)}}, "F")

		//bottom border left
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0, h - 46}, {198, h - 46}, {277, h}, {0, h}}, "F")

		//bottom right border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{w, h - 46}, {w - 198, h - 46}, {w - 277, h}, {w, h}}, "F")

		//primary color overlay
		//left border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0, 356}, {30, h - (h/3.0 + 1)}, {30, h}, {0, h}}, "F")
		//left bottom
		drawPolygon(pdf, pColor, []gofpdf.PointType{{0, h - 30}, {182, h - 30}, {261, h}, {0, h}}, "F")

		//right border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{w, 356}, {w - 30, h - (h/3.0 + 1)}, {w - 30, h}, {w, h}}, "F")
		//bottom right border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{w, h - 30}, {w - 182, h - 30}, {w - 261, h}, {w, h}}, "F")

		//border based on https://www.postermywall.com/index.php/art/template/f4c8c7303c1ba2f677c987d55cdfb971/editable-certificate-template
	} else if style == "circle" {

		// left ellipse
		pdf.SetFillColor(int(pColor.R), int(pColor.G), int(pColor.B))
		pdf.Ellipse(0, h/2.0, 237, h/2.0, 0, "F")

		//upper circular figure
		pdf.SetFillColor(int(sColor.R), int(sColor.G), int(sColor.B))
		pdf.CurveBezierCubic(78, 0, 716, 416, 184, 208, 792, 0, "F")

		//based of https://www.slideshare.net/AniketRaj34/certificate-templates-design
	} else {
		//upper border
		drawPolygon(pdf, pColor, []gofpdf.PointType{{198, 0}, {237, 55}, {w, 55}, {w, 0}}, "F")
		//upper overlay
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0, 0}, {0, 79}, {255, 79}, {199, 0}}, "F")

		//bottom border
		drawPolygon(pdf, sColor, []gofpdf.PointType{{0, h}, {0, h - 55}, {w - 198, h - 55}, {w - 237, h}}, "F")
		//bottom overlay
		drawPolygon(pdf, pColor, []gofpdf.PointType{{w, h}, {w, h - 79}, {w - 199, h - 79}, {w - 255, h}}, "F")

	}
}
