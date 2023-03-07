package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func main() {
	tmpl := parseTemplate()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		someData := [][]string{
			{"Name", "Position", "Office", "Age", "Start date", "Salary"},
			{"Tiger Nixon", "System Architect", "Edinburgh", "61", "2011-04-25", "$320,800"},
			{"Garrett Winters", "Accountant", "Tokyo", "63", "2011-07-25", "$170,750"},
			{"Ashton Cox", "Junior Technical Author", "San Francisco", "66", "2009-01-12", "$86,000"},
			{"Cedric Kelly", "Senior Javascript Developer", "Edinburgh", "22", "2012-03-29", "$433,060"},
			{"Airi Satou", "Accountant", "Tokyo", "33", "2008-11-28", "$162,700"},
			{"Brielle Williamson", "Integration Specialist", "New York", "61", "2012-12-02", "$372,000"},
			{"Herrod Chandler", "Sales Assistant", "San Francisco", "59", "2012-08-06", "$137,500"},
		}
		tmpl.Execute(w, someData)
	})
	http.ListenAndServe(":8080", nil)
}

func parseTemplate() *template.Template {
	return template.Must(
		template.
			New("layout.html").
			Funcs(template.FuncMap{
				"json": jsonencode,
			}).
			ParseFiles("templates/layout.html"),
	)
}

func jsonencode(v any) template.JS {
	b, _ := json.Marshal(v)
	return template.JS(b)
}
