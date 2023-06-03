package main

import (
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type CrudParams struct {
	MongoUrl     string `json:"mongo_url"`
	DatabaseName string `json:"database_name"`
	Collection   string `json:"collection"`
}

func main() {
	e := echo.New()

	e.POST("/setup", setupHandler)

	e.Start(":8000")
}

func setupHandler(c echo.Context) error {
	params := new(CrudParams)
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	err := executeTemplate("crud.txt", "C:\\Users\\chami\\OneDrive\\Desktop\\Chamith\\Repos\\GoTemplateProjects\\GeneratedCode\\main.go", "CrudTemplate", params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate model template")
	}

	

	return c.String(http.StatusOK, "Setup completed successfully")
}

func executeTemplate(templFile string, generFile string, templateName string, data interface{}) error {
	loadedTemplat, err := ioutil.ReadFile(templFile)
	if err != nil {
		return err
	}
	tmpl, err := template.New(templateName).Parse(string(loadedTemplat))
	if err != nil {
		return err
	}

	outPutfile, err := os.Create(generFile)

	err = tmpl.Execute(outPutfile, data)
	if err != nil {
		return err
	}

	return nil
}
