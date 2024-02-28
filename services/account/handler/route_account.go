package handler

import (
	"github.com/go-chi/chi/v5"
	"tinder-cloning/services/account/usecase"
)

type AccountHandler struct {
	accountService usecase.AccountUseCase
	pathName       string
}

func NewAccountHandler(accountService usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{accountService: accountService, pathName: "/account"}
}

func (http *AccountHandler) RegisterRoute(r chi.Router) chi.Router {
	r.Route(http.pathName, func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", http.RegisterHandler)
			r.Post("/login", http.LoginHandler)
		})
	})
	return r
}
