package main

import "encoding/json"
import "net/http"
import "os"
import "log"
import "strings"

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

		f, err := os.OpenFile(strings.ReplaceAll(arg, "/", "_") + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		CatchError(err)
		defer f.Close()

		for _, citation := range citations {
			_, err = f.WriteString("https://doi.org/" + citation.Citing + "\n")
			CatchError(err)
		}
	}
}
