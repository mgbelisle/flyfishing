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

type CastLog struct {
	Location Location
	Fly      Fly
	Fish     Fish
}

// Methods can be added to a struct by any file in the package.
func (l Lake) CastLogsToSVG(castLogs []CastLog) io.Reader {
	buffer := bytes.NewBufferString("")
	templ := template.Must(template.New("").Parse(svgTemplate))
	templ.Execute(buffer, svgTemplateVals{l, castLogs})
	return buffer
}

// Private objects/properties cannot be used outside the package
// because they start with a lower case letter.
type svgTemplateVals struct {
	Lake     Lake
	CastLogs []CastLog
}

const svgTemplate = `<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <rect width="{{.Lake.Length}}" height="{{.Lake.Width}}" fill="#0066FF" />
  {{range .CastLogs}}{{if .Fish}}
  <circle cx="{{.Location.X}}" cy="{{.Location.Y}}" r="1" fill="#00FF00" />
  {{else}}
  <!-- No fish at ({{.Location.X}}, {{.Location.Y}}) -->
  {{end}}{{end}}
</svg>
`
