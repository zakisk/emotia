package handlers

import (
	"net/http"

	"github.com/zakisk/emotia/models"
	"github.com/zakisk/emotia/pkg/utils"
)

func (h *Handler) GetEmotions(w http.ResponseWriter, r *http.Request, options models.GetCommentsOptions) {
	response, err := h.Service.GetTopLevelComments(&options)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.ToJSONIndent(response, "", "  ", w)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
