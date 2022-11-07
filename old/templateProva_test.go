package old

import (
	"bytes"
	"demo/core"
	_ "embed"
	"fmt"
	approvals "github.com/approvals/go-approval-tests"
	"log"
	"os"
	"testing"
	"text/template"
)

func Test_main(t *testing.T) {

	i := core.Item{
		Name: "ciao",
		Sku:  11,
	}
	//assegna come template quello che recupera nel path
	//templ, err := template.ParseFS(provaTemplate, "/prova.txt")
	templ, err := template.ParseFiles("../templates/prova.txt")
	fmt.Println(os.Getwd())
	if err != nil {
		log.Fatal(err)
	}

	//str := fmt.Sprintf("%s %d aa", i.Name, i.Sku)
	buff := bytes.Buffer{}
	//fmt.Println(str)
	if err := templ.Execute(&buff, i); err != nil {
		panic(err)
	}
	approvals.VerifyString(t, buff.String())
}
