package wly

import "github.com/twpayne/go-geom"
import "github.com/twpayne/go-geom/encoding/geojson"

func toLineGeoJSON(points []geom.Point) string{
	var c []float64 = flattenPoints(points)
	var lineString = geom.NewLineStringFlat(geom.XY, c)
	result, err := geojson.Marshal(lineString)
	check(err)
	return string(result)
}

func toPolygonGeoJSON(points []geom.Point) string{
	var c []geom.Coord = flattenPointsToCoords(points)
	var polygon = geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{c})
	result, err := geojson.Marshal(polygon)
	check(err)
	return string(result)
}

func marshallPolygon(polygon *geom.Polygon) string{
	result, err := geojson.Marshal(polygon)
	check(err)
	return string(result)
}

func mustMarshallToGeoJSON(g geom.T) string{
	result, err := geojson.Marshal(g)
	check(err)
	return string(result)
}
