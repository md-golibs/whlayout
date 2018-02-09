package whlayout

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProcessing(t *testing.T) {

	t.Log("running test")

	polygonGeoJSON, err := ioutil.ReadFile("testdata/warehouse1.json")
	check(err)

	layout := LayoutParameters{}
	layout.Horizontal = false
	layout.Padding = 10
	layout.BinsPerRack = 10
	layout.HLayout = []float64{1.075, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.05, 0.05, 1.075}
	layout.VLayout = []float64{1.1, 0.35, 1.1, 0.35, 1.1}
	layout.HDocks = []float64{0.05, 1.9, 0.05}
	layout.VDocks = []float64{1.1, 0.2, 1.1, 0.2, 1.1, 0.2, 1.1}

	var wh Warehouse = GenerateLayout(polygonGeoJSON, layout)
	fmt.Println(wh.Racks[0].BinsAsGeoJSON)
}
