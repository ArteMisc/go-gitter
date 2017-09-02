package gitter

import (
	"testing"
)

func TestPackage_HtmlBody(t *testing.T) {
	result := []byte(`<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="go.artemisc.eu/gitter git https://github.com/ArteMisc/go-gitter">
		<meta name="go-source" content="go.artemisc.eu/gitter https://github.com/ArteMisc/go-gitter/ https://github.com/ArteMisc/go-gitter/tree/master{/dir} https://github.com/ArteMisc/go-gitter/blob/master{/dir}/{file}#L{line}">
		<meta http-equiv="refresh" content="0; url=https://godoc.org/go.artemisc.eu/gitter">
	</head>
	<body>
		Nothing to see here, <a href="https://godoc.org/go.artemisc.eu/gitter">move along</a>
	</body>
</html>`)

	body := (&Package{
		Name: "go.artemisc.eu/gitter",
		Git: Repo{
			Host:     "github.com",
			Username: "ArteMisc",
			Package:  "go-gitter",
			Branch:   "master",
		},
	}).HtmlBody()

	for i := range result {
		if body[i] != result[i] {
			t.Fatalf("incorrect template result at index %d [%x != %x]", i, body[i], result[i])
		}
	}
}
