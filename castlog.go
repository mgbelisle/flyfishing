package flyfishing

import (
	"html/template"
	"os"
)

type CastLog struct {
	Location Location
	Fish Fish
}

func (l Lake) ShowCastLogs(castLogs []CastLog) {
	t := template.Must(template.ParseFiles("castlogs.html"))
	t.Execute(os.Stdout, CastLog{})
}
