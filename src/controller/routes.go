package controller

type routesPaths string

const (
	layoutPath routesPaths = "layout"
	testPath   routesPaths = "test"
	tabPath    routesPaths = "tabs"
	samplePath routesPaths = "sample"
)

func (c *Controller) generatePath(paths ...routesPaths) string {
	end := "/"
	for _, p := range paths {
		end += "/"
		end += string(p)
	}
	return end
}
