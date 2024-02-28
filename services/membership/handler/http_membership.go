package handler

import (
	netHttp "net/http"
	"tinder-cloning/pkg/util"
)

func (http *MembershipHandler) GetFeaturesHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	// get accountData from context
	accountData, ok := r.Context().Value("accountData").(util.AccountDataClaims)
	if !ok {
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload Token")
		return
	}

	features, err := http.membershipService.GetFeatureMembership(r.Context(), accountData.AccountID)
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, features)
}
