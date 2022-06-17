package explorer

import (
	"fmt"
	"github.com/mynameisdaun/squirtlecoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(writer http.ResponseWriter, request *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(writer, "home", data)
}

func add(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(writer, "add", nil)
	case "POST":
		request.ParseForm()
		data := request.FormValue("data")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}

}

func Start(aPort int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", aPort), handler))
}
