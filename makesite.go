package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	//     "text/template"
)

// Page that hold all the information we need to generate a new page
//  HTML page form a text file on the filesystem
type Page struct {
	textFilePath string
	textFileName string
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
	fileNameWithoutExtension := strings.Split(filePath, ".")[0]
	// instantiate a new Page object
	// Populate each field and return the date
	return Page{
		textFilePath: filePath,
		textFileName: fileNameWithoutExtension,
		HTMLPageName: fileNameWithoutExtension + ".html",
		Content:      string(fileContents),
	}

}

func renderTemplateFromPage(templateFilePath string, page Page) {

	// Create a new template in memory names "template.tmpl"
	// when the template is executed, it will parse template.tmpl
	// looking for the { } where we can inject the content.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// create a new blank HTML file
	newFile, err := os.Create(page.HTMLPageName)
	if err != nil {
		panic(err)
	}
	// Executing the template injects the Page instance's data
	// allowing us to render the content of one text file
	// Upon Execution , the template will be saved inside the new file we create earlier
	t.Execute(newFile, page)
	fmt.Println(" Generated File: ", page.HTMLPageName)

}

func main() {
	// The flas represents the name of anu '.txt' file in the same directory as your program
	// Run './makesite --file=latest-post.txt to test
	textFilepath := "latest-post.txt"

	fileFlag := flag.String("file", textFilepath, "path a text file")
	flag.Parse()

	if *fileFlag != "" {
		textFilepath = *fileFlag
	}

	// // Make sure the 'file' flag isn't blank
	// if textFilePath == "" {
	// 	panic("Sorry, You are missing the --file flag!Please provide one.")
	// }

	newPage := createPageFromTextFile(textFilepath)

	// Use the struct to generate a new HTML page based on teh provided template
	renderTemplateFromPage("template.tmpl", newPage)

}
