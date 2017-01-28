package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/stefanscheidt/weatherservice-go/weather"
)

const tmpl string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Weather Forecast</title>
</head>
<body>
<h1>Weather Forecast for Cologne</h1>
<table>
    <tr><td>Station:</td><td>{{.Name}}</td></tr>
    <tr><td>Weather:</td><td>{{range .Weather}}<span>{{.Description}} </span>{{end}}</td></tr>
    <tr><td>Temperature:</td><td>{{.Temp}}</td></tr>
    <tr><td>Sunrise:</td><td>{{.Sunrise}}</td></tr>
    <tr><td>Sunset:</td><td>{{.Sunset}}</td></tr>
</table>
</body>
</html>`

func main() {
	http.HandleFunc("/weather", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Port 8080 is already used")
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	report, err := weather.GetForecast("cologne")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = t.Execute(w, ReportVM(report))
	if err != nil {
		log.Print(err)
	}
}

type ReportVM weather.Report

func (r ReportVM) Temp() string {
	return fmt.Sprintf("%.2f", r.Main.Temperature-273.15)
}

func (r ReportVM) Sunrise() string {
	return timestamp(r.Sys.Sunrise)
}

func (r ReportVM) Sunset() string {
	return timestamp(r.Sys.Sunset)
}

func timestamp(millis int) string {
	return fmt.Sprintf("%v", time.Unix(int64(millis), 0))
}
