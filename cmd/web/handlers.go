package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.con/GiorgosMarga/copyshare/cmd/internal/models"
	"github.con/GiorgosMarga/copyshare/cmd/internal/validator"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to copyshare API"))
}

// Snippets Logic

func (app *Application) GetSnippet(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	userId := app.sessionManager.GetInt(r.Context(), "user")
	if userId == 0 {
		fmt.Println("Querying for a guest's snippet")
	} else {
		fmt.Println("Querying for a user's snippet")
	}
	snippet, err := app.snippet.GetSnippet(id, userId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecond) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return
	}
	json.NewEncoder(w).Encode(snippet)
}

func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	s := &models.Snippet{}
	json.NewDecoder(r.Body).Decode(s)
	app.validator.CheckField(validator.NotBlank(s.Content), "content", "content can't be blank")
	app.validator.CheckField(validator.NotBlank(s.Title), "title", "title can't be blank")
	app.validator.CheckField(validator.MaxChars(s.Title, 300), "title", "title can't be longer than 300 chars")
	if !app.validator.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(app.validator.FormatErrors()))
		return
	}
	id := app.sessionManager.GetInt(r.Context(), "user")
	// if id = 0, the snippet was created by a guest.
	id, err := app.snippet.Create(s.Title, s.Content, id, 10)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))

}

func (app *Application) GetSnippets(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt(r.Context(), "user")

	snippets, err := app.snippet.GetLatestSnippets(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	json.NewEncoder(w).Encode(snippets)
}

func (app *Application) DeleteSnippet(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	userId := app.sessionManager.GetInt(r.Context(), "user")

	if err != nil || id < 1 {
		app.serverError(w, err)
		return
	}
	err = app.snippet.DeleteSnippet(id, userId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecond) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return
	}
	w.Write([]byte("Success"))
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginForm := LoginForm{}
	json.NewDecoder(r.Body).Decode(&loginForm)
	app.validator.CheckField(validator.NotBlank(loginForm.Email), "email", "email can't be empty")
	app.validator.CheckField(validator.NotBlank(loginForm.Password), "password", "password can't be empty")
	if !app.validator.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(app.validator.FormatErrors()))
		return
	}
	user, err := app.user.Login(loginForm.Email, loginForm.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, models.ErrInvalidCredentials) || errors.Is(err, models.ErrNoRecond) || errors.Is(err, models.ErrRestrictedAcc) {
			app.clientError(w, http.StatusUnauthorized)
			return
		}
		app.serverError(w, fmt.Errorf("internal server error"))
		return
	}
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), "user", user.ID)

	json.NewEncoder(w).Encode(user)

}

type RegisterForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	registerForm := RegisterForm{}
	json.NewDecoder(r.Body).Decode(&registerForm)
	app.validator.CheckField(validator.NotBlank(registerForm.Email), "email", "email can't be empty")
	app.validator.CheckField(validator.NotBlank(registerForm.Password), "password", "password can't be empty")
	app.validator.CheckField(validator.NotBlank(registerForm.Username), "username", "username can't be empty")
	app.validator.CheckField(validator.MinChars(registerForm.Password, 5), "password", "password should be at least 5 chars long")
	if !app.validator.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(app.validator.FormatErrors()))
		return
	}
	id, err := app.user.Register(registerForm.Email, registerForm.Password, registerForm.Username)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			app.clientError(w, http.StatusConflict)
			return
		}
		app.serverError(w, fmt.Errorf("internal server error"))
		return
	}
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), "user", id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"msg": "User created successfully"}`))

}
func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Remove(r.Context(), "user")
}

func (app *Application) WhoAmI(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt(r.Context(), "user")
	username, err := app.user.WhoAmI(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecond) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		app.serverError(w, err)
		return
	}
	w.Write([]byte(`{"username": "` + username + "\"}"))
}
