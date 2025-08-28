package handler

import (
	"log"
	"net/http"
	"save-tamal/cms/paginator"
	"time"

	usergrpc "save-tamal/proto/users"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    string
	Name      string
	Batch     int32
	Email     string
	Password  string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}

type UserTemplateData struct {
	User       User
	List       []User
	Paginator  paginator.Paginator
	SearchTerm string
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	h.loadCreateForm(w, User{})
}

func (h *Handler) storeUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usr User
	err = h.decoder.Decode(&usr, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	passByte, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = h.uc.CreateUser(r.Context(), &usergrpc.CreateUserRequest{
		User: &usergrpc.User{
			Name:      usr.Name,
			Batch:     usr.Batch,
			Email:     usr.Email,
			Password:  string(passByte),
			CreatedBy: "",
			UpdatedBy: "",
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, userListPath, http.StatusTemporaryRedirect)
}

func (h *Handler) loadCreateForm(w http.ResponseWriter, usr User) {
	form := UserTemplateData{
		User: usr,
	}

	err := h.templates.ExecuteTemplate(w, "user-create.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loadEditForm(w http.ResponseWriter, usr User) {
	form := UserTemplateData{
		User: usr,
	}

	err := h.templates.ExecuteTemplate(w, "user-edit.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
