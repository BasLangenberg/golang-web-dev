package main

import (
	"log"
	"os"
	"text/template"
)

// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

var tpl *template.Template

type restaurant struct {
	Name string
	Menu menu
}

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
	restaurants := []restaurant{
		restaurant{
			Name: "Harries Vispaleis",
			Menu: menu{
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
			},
		},
		restaurant{
			Name: "Barries Burgerhut",
			Menu: menu{
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
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatal(err)
	}
}
