package flyfishing

import (
	"html/template"
	"os"
)

type CastLog struct {
	Location Location
	Fly Fly
	Fish Fish
}

func (l Lake) ShowCastLogs(castLogs []CastLog) {
	t := template.Must(template.New("svg").Parse(svgTemplate))
	t.Execute(os.Stdout, CastLog{})
}

const svgTemplate =
`<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <rect width="500" height="300" fill="blue" fill-opacity="0.5" />
  <circle cx="250" cy="150" r="2" fill="red" />
  <circle cx="250" cy="155" r="2" fill="green" />
</svg>
`
