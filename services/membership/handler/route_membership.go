package handler

import (
	"github.com/go-chi/chi/v5"
	"tinder-cloning/pkg/middleware"
	"tinder-cloning/services/membership/usecase"
)

type MembershipHandler struct {
	membershipService usecase.MembershipUseCase
	pathName          string
}

func NewMembershipHandler(membershipService usecase.MembershipUseCase) *MembershipHandler {
	return &MembershipHandler{membershipService: membershipService, pathName: "/membership"}
}

func (http *MembershipHandler) RegisterRoute(r chi.Router) chi.Router {
	r.Route(http.pathName, func(r chi.Router) {
		r.Use(middleware.JwtAuthMiddleware)
		r.Route("/v1", func(r chi.Router) {
			r.Get("/features", http.GetFeaturesHandler)
		})
	})
	return r
}
