package wly

import "github.com/twpayne/go-geom"
import "github.com/twpayne/go-geom/encoding/geojson"

func generateDocks(wh Warehouse, layout LayoutParameters) ([]Dock, string) {

	var docks []Dock;

	var polygon = wh.Outline
	var collection = geom.NewGeometryCollection()

	var axis1 = 1
	var axis2 = 0

	if(layout.Horizontal){
		axis1 = 0
		axis2 = 1
	} 

	var hPolys []geom.Polygon = SplitPolygonWithFactors(&polygon, axis1, layout.HDocks)
	
	for i := 0; i < len(hPolys); i++{
		
		var vPolys []geom.Polygon = SplitPolygonWithFactors(&hPolys[i], axis2, layout.VDocks)
		
		for j := 0; j < len(vPolys); j++{

			var dock Dock = Dock{}
			dock.Id = string("r" + string(i) + "_" + string(j))
			dock.Outline = vPolys[j]
			dock.AsGeoJSON = marshallPolygon(&vPolys[j])
			dock.Center = *getCenterOfPolygon(&vPolys[j])

			collection.Push(&dock.Outline)
			collection.Push(&dock.Center)
		}
	}

	result, err := geojson.Marshal(collection)
	check(err)

	return docks, string(result)
}