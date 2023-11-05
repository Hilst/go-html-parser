package controller

type routesPaths string

const (
	layoutPath       routesPaths = "layout"
	nameVariablePath routesPaths = ":name"
	testPath         routesPaths = "test"
)

func (c *Controller) generatePath(paths ...routesPaths) string {
	end := "/"
	for _, p := range paths {
		end += "/"
		end += string(p)
	}
	return end
}

func (c *Controller) clearVariablePath(path routesPaths) string {
	return string(path)[1:]
}
