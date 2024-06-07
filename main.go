package main

import (
    "log"
    "os"
    "path/filepath"
    "text/template"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    templates, err := template.ParseFiles(
        filepath.Join("templates", "base.txt"),
        filepath.Join("templates", "home.txt"),
        filepath.Join("templates", "about.txt"),
    )
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }

    
    homeData := PageData{
        Title:   "Home Page",
        Message: "This is the home page rendered using Go templates.",
    }

    aboutData := PageData{
        Title:   "About Us",
        Message: "This is the about page rendered using Go templates.",
    }

    
    renderAndSaveTemplate(templates, "base.txt", homeData, "rendered_home.txt")
    renderAndSaveTemplate(templates, "base.txt", aboutData, "rendered_about.txt")
}

func renderAndSaveTemplate(tmpl *template.Template, templateName string, data PageData, outputFileName string) {
    file, err := os.Create(outputFileName)
    if err != nil {
        log.Printf("Error creating file %s: %v", outputFileName, err)
        return
    }
    defer file.Close()

    err = tmpl.ExecuteTemplate(file, templateName, data)
    if err != nil {
        log.Printf("Error rendering template %s: %v", templateName, err)
    }

    log.Printf("Template %s rendered and saved to %s", templateName, outputFileName)
}
