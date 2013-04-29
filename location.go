package flyfishing

import (
	"html/template"
	"os"
)

type Location struct {
	X, Y float64
}

func (l Lake) LocationsToSVG(locations []Location) {
	t := template.Must(template.New("main").Parse(svgTemplate))
	t.Execute(os.Stdout, svgTemplateVals{l, locations})
}

type svgTemplateVals struct {
	Lake Lake
	Locations []Location
}

const svgTemplate =
`<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <rect width="{{.Lake.Length}}" height="{{.Lake.Width}}" fill="blue" fill-opacity="0.2" />
  {{range .Locations}}
  <circle cx="{{.X}}" cy="{{.Y}}" r="2" fill="green" />
  {{end}}
</svg>
`
