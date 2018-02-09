package wly

import _"fmt"
import "github.com/twpayne/go-geom"
import "github.com/twpayne/go-geom/encoding/geojson"


func generateRacks(wh Warehouse, layout LayoutParameters) ([]Rack,string) {
	
	var polygon = wh.Outline
	var collection = geom.NewGeometryCollection()

	var racks = []Rack{}
	var axis1 = 1
	var axis2 = 0

	if(layout.Horizontal){
		axis1 = 0
		axis2 = 1
	} 

	var hPolys []geom.Polygon = SplitPolygonWithFactors(&polygon, axis1, layout.HLayout)
	
	for i := 0; i < len(hPolys); i++{
		
		var vPolys []geom.Polygon = SplitPolygonWithFactors(&hPolys[i], axis2, layout.VLayout)
		
		for j := 0; j < len(vPolys); j++{

			var rack Rack = Rack{}
			rack.Id = string("r" + string(i) + "_" + string(j))
			rack.Outline = vPolys[j]
			rack.AsGeoJSON = marshallPolygon(&vPolys[j])
			rack.StorageBins, rack.BinsAsGeoJSON = generateStorageBins(rack, layout)

			racks = append(racks,rack)
			collection.Push(&vPolys[j])
		}
	}

	result, err := geojson.Marshal(collection)
	check(err)

	return racks, string(result)
}