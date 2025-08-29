package handler

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	usergrpc "save-tamal/proto/users"
)

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	session   *sessions.CookieStore
	assets    fs.FS
	assetFS   *hashfs.FS
	uc        usergrpc.UserServiceClient
}

func GetHandler(decoder *schema.Decoder, session *sessions.CookieStore, assets fs.FS, uc usergrpc.UserServiceClient) *mux.Router {
	hand := &Handler{
		decoder: decoder,
		session: session,
		assets:  assets,
		assetFS: hashfs.NewFS(assets),
		uc:      uc,
	}
	hand.GetTemplate()

	r := mux.NewRouter()
	r.HandleFunc(userCreatePath, hand.createUser)
	r.HandleFunc(userStorePath, hand.storeUser)
	r.HandleFunc(userEditPath, hand.editUser)
	r.HandleFunc(userUpdatePath, hand.updateUser)
	r.HandleFunc(userListPath, hand.listUser)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.FS(hand.assetFS))))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := hand.templates.ExecuteTemplate(w, "404.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/users/user-list.html",
		"cms/assets/templates/users/user-create.html",
		"cms/assets/templates/users/user-edit.html",
		"cms/assets/templates/404.html",
	))
}
