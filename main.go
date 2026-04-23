package main

import "encoding/json"
import "net/http"
import "os"
import "log"
import "fmt"

type Citation struct {
	Citing string `json:"citing"`
}

func CatchError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		var citations []Citation

		resp, err := http.Get("https://api.opencitations.net/index/v1/citations/" + arg)
		CatchError(err)
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&citations)
		CatchError(err)

		for _, citation := range citations {
			fmt.Println("https://doi.org/" + citation.Citing + "\n")
		}
	}
}
