// post_server/post_server.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	utils "post_server/check"
	"post_server/post"
)

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	utils.Check(scanner.Err())
	return lines
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("views/new.html")
	utils.Check(err)
	err = html.Execute(writer, nil)
	utils.Check(err)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("docs/signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("views/view.html")
	utils.Check(err)
	getPost := post.New(len(signatures), signatures)
	err = html.Execute(writer, getPost)
	utils.Check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	post := request.FormValue("post")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("docs/signatures.txt", options, os.FileMode(0600))
	utils.Check(err)
	_, err = fmt.Fprintln(file, post)
	utils.Check(err)
	err = file.Close()
	utils.Check(err)
	http.Redirect(writer, request, "/post", http.StatusFound)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/post/new", newHandler)
	http.HandleFunc("/post/create", createHandler)
	_, err := os.Stdout.Write([]byte("Server Start...."))
	utils.Check(err)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
