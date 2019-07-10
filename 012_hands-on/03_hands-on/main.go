package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := struct {
		Southern []hotel
		MidWest  []hotel
		Northern []hotel
	}{
		Southern: []hotel{
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
		},
		MidWest: []hotel{
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
		},
		Northern: []hotel{
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
			hotel{"RedNeck In", "RedNeck Street 12", "RedNeckCity", "123456"},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
