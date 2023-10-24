package main

import (
	"os"
	"os/exec"
	"strings"
	"log"
)

func main() {
	htmlData, err := os.ReadFile("template.html")
	if err != nil {
		log.Println("err11: ",err)
	}


	// Input JSON data
	jsonData := `{"title": "سلام", "content": "این یک مثال از ایجاد PDF با Go است."}`

	// Replace a placeholder with the actual JSON data
	htmlContent := strings.Replace(string(htmlData), "{{jsonData}}", jsonData, -1)

	// Run wkhtmltopdf command with HTML content as a string
	cmd := exec.Command("wkhtmltopdf", "--enable-local-file-access", "-", "output.pdf")
	cmd.Stdin = strings.NewReader(htmlContent)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	if err := cmd.Run(); err != nil {
		log.Println("err2222: ", err)
	}
}
