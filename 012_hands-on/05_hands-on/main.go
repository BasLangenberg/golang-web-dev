package main

import (
	"text/template"
	"os"
	"log"
)

// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

var tpl *template.Template

type dish struct {
	Name  string
	Price int
}

type menu struct {
	Breakfast []dish
	Lunch     []dish
	Dinner    []dish
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	dishes := menu{
		Breakfast: []dish{
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
		},
		Lunch: []dish{
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
		},
		Dinner: []dish{
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
			dish{"Vreten", 10},
		},
	}

	err := tpl.Execute(os.Stdout, dishes)
	if err != nil {
		log.Fatal(err)
	}
}
