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

func (grcos *Gridcellcos) init(spacing, ori, offx, offy, origori float64) {
	grcos.spacing = spacing
	grcos.k = spacing * (2 / math.Sqrt(3.0))
	grcos.origori = origori * (math.Pi / 180)
	grcos.ori = (origori + ori) * (math.Pi / 180)
	grcos.offx = offx
	grcos.offy = offy
}

func (grcos *Gridcellcos) Activation(curx, cury float64, plus bool) float64 {
	grcos.curx = curx
	grcos.cury = cury
	x := curx - grcos.offx
	y := cury - grcos.offy
	ori1 := grcos.ori
	ori2 := ori1 + (math.Pi/180)*60
	ori3 := ori2 + (math.Pi/180)*60
	Fk1 := (1 + math.Cos(grcos.k*((x)*math.Sin(ori1)+(y)*math.Cos(ori1)))) * 0.5
	Fk2 := (1 + math.Cos(grcos.k*((x)*math.Sin(ori2)+(y)*math.Cos(ori2)))) * 0.5
	Fk3 := (1 + math.Cos(grcos.k*((x)*math.Sin(ori3)+(y)*math.Cos(ori3)))) * 0.5
	if plus {
		grcos.firingrate = Fk1 + Fk2 + Fk3
	} else {
		grcos.firingrate = Fk1 * Fk2 * Fk3
	}
	return grcos.firingrate
}