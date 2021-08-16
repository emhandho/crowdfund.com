package handler

import (
	"net/http"
	"strconv"

	"crowdfund.com/campaign"
	"crowdfund.com/helper"
	"github.com/gin-gonic/gin"
)

// catch params from client in handler
// handler to service
// service determine which repository that will called
// repository: FindAll & FindByUserID
// repository access to db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _:= strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	// fmt.Println(campaigns[0])
	response := helper.APIResponse("Successfuly get campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}