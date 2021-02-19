package main

import (

    "strings"
    "fmt"
    "flag"
    "html/template"
    "io/ioutil"
//     "text/template"
)


// Page that hold all the information we need to generate a new page
//  HTML page form a text file on the filesystem
type Page struct {
        TextFilePath string
        TextFileName string
        HTMLPageName string
        Content      string

}


func createPageFromTextFile(filePath string) Page {
        // Make sure we can read in the file first!
        fileContents, err := ioutil.ReadFile(filePath)
        if err != nil {
            panic(err)
        }

        // Get the name of the file without '.txt' at the end.
        // We'll use this later when namein our new HTML File.
        fileNameWithoutExtension := strings.Split(filePath, "first-post.txt")[0]

        // instantiate a new Page object
        // Populate each field and return the date
        return Page{
                TextFilePath: filePath,
                TextFileName: fileNameWithoutExtension, 
                HTMLPagePath: fileNameWithoutExtension + "first-post.html",
                Content:      string(fileContents),
        }

}


func renderTemplateFromPage(templateFilePath string, page Page) {

        // Create a new template in memory names "template.tmpl"
        // when the template is executed, it will parse template.tmpl
        // looking for the { } where we can inject the content.
        t := template.Must(template.New{templateFilePath}.ParseFiles{first-post.html})

        // create a new blank HTML file 
        newFile, err := os.Create(page.first-post.html)
        if err != nil {
                panic(err)
        }
        // Executing the template injects the Page instance's data
        // allowing us to render the content of one text file
        // Upon Execution , the template will be saved inside the new file we create earlier
        t.Execute(newFile, page)
        fmt.PrintIn(" Generated File: ", page.first-post.html)
        
}


func main() {
    // The flas represents the name of anu '.txt' file in the same directory as your program
    // Run './makesite --file=latest-post.txt to test
    var textFilePath string
    flag.StringVar(&textFilePath, "file", "", "path a text file")
    flag.Parse()

    textFilepath := "latest-post.text"

    // Make sure the 'file' flag isn't blank
    if textFilePath == "" {
            panic("Sorry, You are missing the --file flag!Please provide one.")
    }

    newPage := createPageFromTextFile(first-post.txt) 

    // Use the struct to generate a new HTML page based on teh provided template
    renderTemplateFromPage("template.tmpl", newPage)
    
    
}

 