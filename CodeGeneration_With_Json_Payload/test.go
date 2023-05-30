package main

import (
	_ "encoding/json"
	"fmt"
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

	err := executeTemplate("model.txt", "modelTemplate", params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate model template")
	}

	err = executeTemplate("method.txt", "methodTemplate", params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate methods template")
	}

	err = executeTemplate("main.txt", "mainTemplate", params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate main template")
	}

	return c.String(http.StatusOK, "Setup completed successfully")
}

func executeTemplate(templFile string, templateName string, data interface{}) error {
	loadedTemplat, err := ioutil.ReadFile(templFile)
	if err != nil {
		return err
	}
	tmpl, err := template.New(templateName).Parse(string(loadedTemplat))
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.go", templateName)

	file, err := os.Create(fileName)

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
