package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()
	mainChain := alice.New(app.recoverPanic, app.addJsonHeader, app.logRequests, app.sessionManager.LoadAndSave, app.cors)
	withAuthChain := mainChain.Append(app.requireAuthentication)
	router.Handler(http.MethodGet, "/", mainChain.Then(http.HandlerFunc(app.Home)))

	// Snippets
	router.Handler(http.MethodGet, "/snippet", withAuthChain.Then(http.HandlerFunc(app.GetSnippets)))
	router.Handler(http.MethodGet, "/snippet/:id", http.HandlerFunc(app.GetSnippet))
	router.Handler(http.MethodDelete, "/snippet/:id", withAuthChain.Then(http.HandlerFunc(app.DeleteSnippet)))
	router.Handler(http.MethodPost, "/snippet", http.HandlerFunc(app.CreateSnippet))

	// Auth
	router.Handler(http.MethodPost, "/auth/login", http.HandlerFunc(app.Login))
	router.Handler(http.MethodPost, "/auth/register", http.HandlerFunc(app.Register))
	router.Handler(http.MethodPost, "/auth/logout", http.HandlerFunc(app.Logout))

	return mainChain.Then(router)
}
