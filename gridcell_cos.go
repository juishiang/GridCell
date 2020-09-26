package gridcell

import (
	"math"
)

type GCcoslay struct {
	Cellnum int64
	Grlay   []Gridcellcos
}

func (grly *GCcoslay) Init(num int64) {
	grly.Cellnum = num
	grly.Grlay = make([]Gridcellcos, num)
}

type Gridcellcos struct {
	k          float64
	spacing    float64
	origori    float64
	ori        float64
	offx       float64
	offy       float64
	firingrate float64
	curx       float64
	cury       float64
}

func (grcos *Gridcellcos) Init(spacing, ori, offx, offy, origori float64) {
	grcos.spacing = spacing
	grcos.k = (4 * math.Pi) / (math.Sqrt(3.0) * spacing)
	grcos.origori = origori * (math.Pi / 180.0)
	grcos.ori = (origori + ori) * (math.Pi / 180.0)
	grcos.offx = offx * (math.Sqrt(3.0) / 2.0)
	grcos.offy = offy
}

func (grcos *Gridcellcos) Activation(curx, cury float64, plus bool) float64 {
	grcos.curx = curx
	grcos.cury = cury
	x := curx - grcos.offx
	y := cury - grcos.offy
	ori1 := grcos.ori
	ori2 := ori1 + (math.Pi/180.0)*60.0
	ori3 := ori2 + (math.Pi/180.0)*60.0
	Fk1 := math.Cos(grcos.k * ((x)*math.Sin(ori1) + (y)*math.Cos(ori1)))
	Fk2 := math.Cos(grcos.k * ((x)*math.Sin(ori2) + (y)*math.Cos(ori2)))
	Fk3 := math.Cos(grcos.k * ((x)*math.Sin(ori3) + (y)*math.Cos(ori3)))
	if plus {
		//grcos.firingrate = (Fk1 + Fk2 + Fk3 + 3.0) / 6.0 //(2.0 / 3.0) * ((1.0/3.0)*(Fk1+Fk2+Fk3) + 0.5)
		grcos.firingrate = (Fk1 + Fk2 + Fk3 - 1.0) / 2.0 //(2.0 / 3.0) * ((1.0/3.0)*(Fk1+Fk2+Fk3) + 0.5)
	} else {
		grcos.firingrate = (1.0 + Fk1) * (1.0 + Fk2) * (1.0 + Fk3) * (1.0 / 8.0)
	}
	if grcos.firingrate > 1 {
		return 1.0
	} else if grcos.firingrate <= 0 {
		return 0.0
	} else {
		return grcos.firingrate
	}
}
