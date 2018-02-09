# WarehouseLayout Generator
This lib allows you to generate a warehouse layout based on a given polygon and with layout parameters

## Getting Started
`import "github.com/md-golibs/whlayout"`

In order to call the function GenerateLayout you need to pass the a []byte with the GeoJSON containing a polygon and the layout definition (type whlayout.LayoutParameters).

Remark: The polygon should only consist of four points and the order of the points should be: lower left corner, lower right corner, upper right corner, upper left corner. This is currently still a limitation (sorry).

```Go
    //Create the LayoutParamters
	layout := wly.LayoutParameters{ }
	layout.Horizontal = false
	layout.BinsPerRack = 10
	layout.HLayout = []float64{1.075,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.05,0.05,1.075}
	layout.VLayout = []float64{1.1,0.35,1.1,0.35,1.1}
	layout.HDocks  = []float64{0.05,1.9,0.05}
	layout.VDocks  = []float64{1.1,0.2,1.1,0.2,1.1,0.2,1.1}
	
	var wh wly.Warehouse = wly.GenerateLayout(polygonGeoJSON, layout)
```