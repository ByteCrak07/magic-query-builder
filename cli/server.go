package cli

import (
	"fmt"
	"io"
	"log"
	"magic-query/generator"
	"net/http"
	"strings"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// send instructions on GET request
	case "GET":
		fmt.Fprintf(w, "Welcome to the Magic Query Builder\n")
		fmt.Fprintf(w, "Send the input as a body (raw text) via POST method for generated query")

	// on POST read raw text from body, covert it to string slice, generate query and send it back
	case "POST":

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		sliceData := strings.Split(string(bodyBytes), "\n")
		output := generator.HandleQuery(sliceData)
		fmt.Fprintf(w, "%v", output)

	default:
		fmt.Fprintf(w, "Only GET and POST methods supported!")
	}

}

func handleServer(port string) {
	http.HandleFunc("/", handleRequests)
	fmt.Println("Server started at http://localhost:" + port)
	fmt.Println("Use GET req to \"/\" for more informatation")

	// start http server at a given port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
