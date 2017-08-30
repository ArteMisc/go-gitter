package gitter

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"text/template"
)

// tmpl is the
const (
	templateHtml = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.Name}} git {{.Git.RepoUrl}}">
		<meta name="go-source" content="{{.Name}} {{.Git.RepoUrl}}/ {{.Git.RepoUrl}}/tree/{{.Git.Branch}}{/dir} {{.Git.RepoUrl}}/blob/{{.Git.Branch}}{/dir}/{file}#L{line}">
		<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.Name}}">
	</head>
	<body>
		Nothing to see here, <a href="https://godoc.org/{{.Name}}">move along</a>
	</body>
</html>`
)

var (
	tmplt = template.Must(template.New("git").Parse(templateHtml))
)

// Config holds all the relevant configuration values for the
type Config struct {
	// Host holds the host on which the gitter Server should listen.
	Host string `json:"host"`

	// Post holds the port on which the gitter Server should listen.
	Port uint16 `json:"port"`

	// Tls configuration, including certificate path and key path for
	// http.ListenAndServeTLS.
	Tls *TlsConfig `json:"tls"`

	// Packages holds a list of packages for which the gitter Server should
	// serve redirects.
	Packages []*Package `json:"packages"`
}

// TlsConfig specifies the server's TLS configuration. If the Config field is
// left empty, the package's defaults are used.
type TlsConfig struct {
	CertPath string `json:"cert_path"`
	KeyPath  string `json:"key_path"`
	*tls.Config
}

// Package contains a path (package name) match with a git repository.
type Package struct {
	// Name holds the package name
	Name string `json:"name"`

	// Git describes a git repository.
	Git Repo `json:"git"`
}

// Repo describes a git repository
type Repo struct {
	// Host holds the host domain of the git repository, e.g. "github.com".
	// GitHub is currently the only supported host.
	Host string `json:"host"`

	// Username is the username of the repository's owner on the git platform.
	Username string `json:"username"`

	// Package holds the name of the package/repo, e.g. "go-gitter"
	Package string `json:"package"`

	// Branch holds the name of the branch to refer to, e.g. "master".
	Branch string `json:"branch"`
}

func (r *Repo) RepoUrl() string {
	return fmt.Sprintf("https://%s/%s/%s", r.Host, r.Username, r.Package)
}

func (p *Package) HtmlBody() (result []byte) {
	body := bytes.NewBuffer(make([]byte, 0, 1024))

	err := tmplt.Execute(body, p)
	if err != nil {
		panic(err)
	}

	result = body.Bytes()
	return
}

// HttpHandler returns a http.Handler that responds with the HTML needed to
// redirect a 'go get' command to the right git repository.
func (p *Package) HttpHandler() http.Handler {
	data := p.HtmlBody()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(data)
	})
}

// Handle adds a Handler to the mux. The Package's full name (domain and path)
// are used as route. The result of p.HttpHandler() is used as the route's
// handler.
func (p *Package) Handle(mux *http.ServeMux) {
	mux.Handle(p.Name, p.HttpHandler())
}
