package gridcell

import (
	"math"
	"math/rand"
)

type Grid_layer struct {
	Cellnum int64
	Grlay   []Grid_cell
}

func (grly *Grid_layer) Init(num int64) {
	grly.Cellnum = num
	grly.Grlay = make([]Grid_cell, num)
}

type Grid_cell struct {
	cycle     float64
	offset    [2]float64
	spacesize float64
	k         float64
	placedev  float64     //place deviation \sigma_place
	hmean     float64     //hight mean
	hdev      float64     //hight deviation
	Mh        [][]float64 //matrix for peal value
	Firerate  float64
	CurX      float64
	CurY      float64
}

func (gr *Grid_cell) Init(cyc, x1off, y1off, spsiz, hm, hd, pd float64) bool {
	gr.cycle = cyc
	gr.spacesize = spsiz
	gr.hmean = hm
	gr.hdev = hd
	gr.k = gr.spacesize / gr.cycle
	gr.placedev = pd * gr.k
	gr.offset[0], gr.offset[1] = gr.TDeltaInv(x1off, y1off) //x1off, y1off [0 1]. Use rand.Float64() to generate number from 0-1.
	gr.Mh = make([][]float64, int(gr.cycle*3))
	for i := range gr.Mh {
		gr.Mh[i] = make([]float64, int(gr.cycle*3))
	}
	//	for i :=0;i<(gr.cycle+2.0);i++{
	//		for j:=0;j<(gr.cycle+2.0);j++{
	//			gr.M_h[i][j]=0.0
	//		}
	//	}
	return true
}

func (gr *Grid_cell) TDelta(x, y float64) (x1, y1 float64) {
	//M = (1/k)[[1 (-1)/(sqrt(3))];[0 2/(sqrt(3))]]
	x1 = (1 / gr.k) * (x - y*(1/(math.Sqrt(3))))
	y1 = (1 / gr.k) * y * (2 / math.Sqrt(3))
	return
}

func (gr *Grid_cell) TDeltaInv(x1, y1 float64) (x, y float64) {
	//M_inv = k[[1 0.5];[0 sqrt(3)/2]]
	x = gr.k * (x1 + 0.5*y1)
	y = gr.k * y1 * (math.Sqrt(3) / 2)
	return
}

func (gr *Grid_cell) Fireact(x, y float64) float64 {
	gr.CurX = x
	gr.CurY = y
	x = x - gr.offset[0]
	y = y - gr.offset[1]
	x1, y1 := gr.TDelta(x, y)
	xstartidx := math.Floor(x1)
	ystartidx := math.Floor(y1)
	x0, y0 := gr.TDeltaInv(xstartidx, ystartidx)
	min := math.Pow((x-x0), 2) + math.Pow((y-y0), 2)
	minplacex := 0
	minplacey := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			x0, y0 := gr.TDeltaInv(xstartidx+float64(i), ystartidx+float64(j))
			curVal := math.Pow((x-x0), 2) + math.Pow((y-y0), 2)
			if min > curVal {
				min = curVal
				minplacex = i
				minplacey = j
			}
		}
	}
	curpeaksitex := int(xstartidx+gr.cycle+1.0) + minplacex
	curpeaksitey := int(ystartidx+gr.cycle+1.0) + minplacey
	//	fmt.Println("curpeaksitex", curpeaksitex)
	//	fmt.Println("curpeaksitey", curpeaksitey)

	if gr.Mh[curpeaksitex][curpeaksitey] == 0.0 {
		gr.Mh[curpeaksitex][curpeaksitey] = gr.hmean + gr.hdev*(rand.Float64()-0.5)*2
		if gr.Mh[curpeaksitex][curpeaksitey] > 1 {
			gr.Mh[curpeaksitex][curpeaksitey] = 1
		} else if gr.Mh[curpeaksitex][curpeaksitey] < 0 {
			gr.Mh[curpeaksitex][curpeaksitey] = 0
		}
	}
	gr.Firerate = gr.Mh[curpeaksitex][curpeaksitey] * math.Exp(-(min / (2 * math.Pow(gr.placedev, 2))))
	return gr.Firerate
}