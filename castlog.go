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
	t := template.Must(template.New("castLogs").Parse(castLogsTemplate))
	t.Execute(os.Stdout, castLogs)
}

const castLogsTemplate =
`<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <rect width="500" height="300" fill="blue" fill-opacity="0.2" />
  {{range .}}{{if .Fish}}
  <circle cx="{{.Location.X}}" cy="{{.Location.Y}}" r="2" fill="green" />
  {{end}}{{end}}
</svg>
`
