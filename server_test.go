package gitter

import "encoding/json"

//
func ExampleServer_HttpServer() {
	config := new(Config)

	err := json.Unmarshal([]byte(`{
		"host": "0.0.0.0",
		"port": 8080,
		"packages": [
			{
				"name": "go.artemisc.eu/gitter",
				"git": {
					"host": "github.com",
					"username": "ArteMisc",
					"package": "go-gitter",
					"branch": "master"
				}
			}
		]
	}`), config)

	if err != nil {
		panic(err)
	}

	s := New(config).HttpServer()
	defer s.Close()

	go s.ListenAndServe()
}
