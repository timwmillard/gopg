package gen

import (
	"log"
	"os"
	"text/template"
)

// sweaters := Inventory{"wool", 17}
// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
// if err != nil { panic(err) }
// err = tmpl.Execute(os.Stdout, sweaters)
// if err != nil { panic(err) }
func Run(function Function) error {

	tmpl, err := template.New("test.go.tmpl").ParseFiles("gen/test.go.tmpl")
	// tmpl, err := template.New("exec_tmpl").Parse("Name  = {{ .Name }}")
	if err != nil {
		log.Println("ParseGlob err: ", err)
		return err
	}

	err = tmpl.Execute(os.Stdout, function)
	if err != nil {
		log.Println("Execute err: ", err)
		return err
	}

	return nil
}
