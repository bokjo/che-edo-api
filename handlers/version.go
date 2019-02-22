package handlers

import (
	"net/http"

	"gitlab.com/bokjo/test_edo/model"
)

// VersionHandler struct
type VersionHandler struct {
	VersionService model.VersionService
}

// GetVersion handle retrieves the current version
func (vh *VersionHandler) GetVersion(w http.ResponseWriter, r *http.Request) {

	Version, err := vh.VersionService.GetVersion()

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid version request!")
		return
	}

	jsonRespond(w, http.StatusOK, Version)

}
