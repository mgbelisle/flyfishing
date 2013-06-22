package flyfishing

import (
	"bytes"
	"io"
	"text/template"
)

// Attributes of the same type can be done in one declaration.
type Location struct {
	X, Y float64
}

// Methods can be added to a struct by any file in the package.
func (lake Lake) LocationsToSVG(locations []Location) io.Reader {
	buffer := bytes.NewBufferString("")
	templ := template.Must(template.New("").Parse(svgTemplate))
	templ.Execute(buffer, svgTemplateVals{lake, locations})
	return buffer
}

// Private objects/properties cannot be used outside the package
// because they start with a lower case letter.
type svgTemplateVals struct {
	Lake      Lake
	Locations []Location
}

const svgTemplate = `<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <rect width="{{.Lake.Length}}" height="{{.Lake.Width}}" fill="#0066FF" />
  {{range .Locations}}
  <circle cx="{{.X}}" cy="{{.Y}}" r="2" fill="#FFFFFF" />
  {{end}}
</svg>
`
