package handler

import (
	"github.com/go-chi/chi/v5"
	"tinder-cloning/pkg/middleware"
	"tinder-cloning/services/swipe/usecase"
)

type SwipeHandler struct {
	swipeService usecase.SwipesUseCase
	pathName     string
}

func NewSwipeHandler(swipeService usecase.SwipesUseCase) *SwipeHandler {
	return &SwipeHandler{swipeService: swipeService, pathName: "/swipe"}
}

func (http *SwipeHandler) RegisterRoute(r chi.Router) chi.Router {
	r.Route(http.pathName, func(r chi.Router) {
		r.Use(middleware.JwtAuthMiddleware)
		r.Route("/v1", func(r chi.Router) {
			r.Get("/list", http.GetAllProfileHandler)
			r.Post("/action", http.CreateReactionSwipesHandler)
		})
	})
	return r
}
