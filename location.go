package flyfishing

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
)

// Attributes of the same type can be done in one declaration.
type Location struct {
	X, Y float64
}

// Methods can be added to a struct in any file in the package.
func (l Lake) LocationsToSVG(locations []Location) *os.File {
	fname := fmt.Sprintf("%s/locations_%d.svg", os.TempDir(), rand.Int())
	file, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("").Parse(svgTemplate))
	t.Execute(file, svgTemplateVals{l, locations})
	return file
}

// Private objects/properties cannot be used outside the package
// because they start with a lower case letter.
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

type CastLog struct {
	Location Location
	Fish Fish
}
