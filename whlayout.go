package wly

import "github.com/twpayne/go-geom"


type LayoutParameters struct{
	Horizontal bool
	HLayout []float64
	VLayout []float64
	HDocks []float64
	VDocks []float64
	RackSize float64
	AisleSize float64
	Padding float64
	BinsPerRack int
}

type Rack struct{
	Id string
	Outline geom.Polygon
	AsGeoJSON string
	StorageBins []StorageBin
	BinsAsGeoJSON string
}

type StorageBin struct{
	Id string
	Outline geom.Polygon
	AsGeoJSON string
	Center geom.Point
}

type Dock struct{
	Id string
	Outline geom.Polygon
	AsGeoJSON string
	Center geom.Point	
}

type Warehouse struct{
	Outline geom.Polygon
	AsGeoJSON string
	Racks []Rack
	RacksAsGeoJSON string
	Docks []Dock
	DocksAsGeoJSON string
} 


func GenerateLayout(polygon *geom.Polygon, layout LayoutParameters) Warehouse {

	var wh Warehouse
	var points []geom.Point = edgePoints(polygon)
	var padding = meterToDeg(layout.Padding)
		
	var p1 = addPadding(points[0], padding, 1)
	var p2 = addPadding(points[1], padding, 2)
	var p3 = addPadding(points[2], padding, 3)
	var p4 = addPadding(points[3], padding, 4) 

	wh.Racks = []Rack{}
	wh.Outline = *polygon
	wh.AsGeoJSON = toPolygonGeoJSON([]geom.Point{p1,p2,p3,p4,p1})

	wh.Racks, wh.RacksAsGeoJSON = generateRacks(wh, layout)
	wh.Docks, wh.DocksAsGeoJSON = generateDocks(wh, layout)

	return wh
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
