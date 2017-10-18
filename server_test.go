// Copyright 2017, Project ArteMisc
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
