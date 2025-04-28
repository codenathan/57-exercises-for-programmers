package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeFolder(path string) {
	err := os.Mkdir(path, 0755)

	if err != nil {
		fmt.Println("Error creating site folder:", err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Site name: ")
	siteName, _ := reader.ReadString('\n')
	siteName = strings.TrimSpace(siteName)

	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Do you want a folder for JavaScript? (y/n) ")
	jsFolder, _ := reader.ReadString('\n')
	jsFolder = strings.TrimSpace(strings.ToLower(jsFolder))

	fmt.Print("Do you want a folder for CSS? (y/n) ")
	cssFolder, _ := reader.ReadString('\n')
	cssFolder = strings.TrimSpace(strings.ToLower(cssFolder))

	makeFolder(siteName)
	fmt.Printf("Created ./%s\n", siteName)
	indexPath := fmt.Sprintf("./%s/index.html", siteName)
	indexFile, err := os.Create(indexPath)
	if err != nil {
		fmt.Println("Error creating index.html:", err)
		return
	}
	defer indexFile.Close()

	htmlContent := fmt.Sprintf(`<html>
<head>
    <title>%s</title>
    <meta name="author" content="%s">
</head>
<body>

</body>
</html>`, siteName, author)

	indexFile.WriteString(htmlContent)
	fmt.Printf("Created ./%s/index.html\n", siteName)

	if jsFolder == "y" {
		jsPath := fmt.Sprintf("./%s/js", siteName)
		makeFolder(jsPath)
		fmt.Printf("Created ./%s/js/\n", siteName)
	}

	if cssFolder == "y" {
		cssPath := fmt.Sprintf("./%s/css", siteName)
		makeFolder(cssPath)
		fmt.Printf("Created ./%s/css/\n", siteName)
	}
}
