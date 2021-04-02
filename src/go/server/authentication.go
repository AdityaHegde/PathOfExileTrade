package server

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// SetupAuthentication is exported
func SetupAuthentication(handler *mux.Router) {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/core/google/callback"),
	)

	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	store.Options.Domain = "localhost"
	store.Options.HttpOnly = true
	store.Options.MaxAge = 60 * 60 * 24
	gothic.Store = store

	authRouter := handler.PathPrefix("/core").Subrouter()
	authRouter.HandleFunc("/{provider}", providerHandler)
	authRouter.HandleFunc("/{provider}/callback", callbackHandler)
	authRouter.HandleFunc("/{provider}/logout", logoutHandler)

	m := make(map[string]string)
	m["google"] = "Google"
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}
	handler.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, providerIndex)
	})
}

func callbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	// http.Redirect(res, req, strings.Replace(req.URL.String(), "/callback", "", 1), 302)
	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusTemporaryRedirect)
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(res, user)
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func providerHandler(res http.ResponseWriter, req *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(res, req)
	// try to get the user without re-authenticating
	if err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, gothUser)
	} else {
		fmt.Println(err)
		gothic.BeginAuthHandler(res, req)
	}
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

const indexTemplate = `{{range $key,$value:=.Providers}}
<p><a href="/core/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

const userTemplate = `
<p><a href="/core/{{.Provider}}/logout">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
