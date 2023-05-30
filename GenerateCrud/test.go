package main

import (
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
	"text/template"
)

type DatabaseInfo struct {
	URL          string `json:"url"`
	DatabaseName string `json:"databasename"`
	Collection   string `json:"collection"`
}

var databaseInfos []DatabaseInfo

func main() {
	e := echo.New()
	mongoUrl := "mongodb+srv://chamith:123@cluster0.ujlq82i.mongodb.net/?retryWrites=true&w=majority"
	databseName := "sample_db2"
	collection := "Users2"

	type CrudParams struct {
		MongoUrl     string
		DatabaseName string
		Collection   string
	}

	params := CrudParams{
		MongoUrl:     mongoUrl,
		DatabaseName: databseName,
		Collection:   collection,
	}

	//tmpl, err := template.New("modelTemplate").Parse(methodsTemplate)
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//
	//err = tmpl.Execute(os.Stdout, params)
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	err := executeTemplate("model.txt", "model.go", "modelTemplate", params)

	if err != nil {
		e.Logger.Fatal(err)
	}

	err = executeTemplate("method.txt", "method.go", "methodTemplate", params)

	if err != nil {
		e.Logger.Fatal(err)
	}

	err = executeTemplate("main.txt", "main.go", "mainTemplate", params)

	if err != nil {
		e.Logger.Fatal(err)
	}

}
func executeTemplate(tempFile string, outputFile string, templateName string, data interface{}) error {
	templateLoaded, err := ioutil.ReadFile(tempFile)
	if err != nil {
		return err
	}
	tmpl, err := template.New(templateName).Parse(string(templateLoaded))
	if err != nil {
		return err
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		return err
	}

	return nil
}
