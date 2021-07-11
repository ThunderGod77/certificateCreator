package certificate

import (
	"github.com/phpdave11/gofpdf"
	"image/color"
)

type position struct {
	x float64
	y float64
}

func changePosRel(p *gofpdf.Fpdf, incX float64, incY float64, pos *position) {
	(*pos).x = pos.x + incX
	pos.y = pos.y + incY
	p.MoveTo(pos.x, pos.y)
}

func changePosAbs(p *gofpdf.Fpdf, x float64, y float64, pos *position) {
	pos.x = x
	pos.y = y
	p.MoveTo(pos.x, pos.y)
}

func drawPolygon(p *gofpdf.Fpdf, c color.RGBA, points []gofpdf.PointType, style string) {
	p.SetFillColor(int(c.R), int(c.G), int(c.B))
	p.Polygon(points, style)
}
