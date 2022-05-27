package stats

import (
	"encoding/json"
	"net/http"
)

type getUsageResponse struct {
	DailyUsage []int `json:"dailyUsage"`
}

func (c *statsController) getUsage(w http.ResponseWriter, r *http.Request) {
	usageByDay := c.db.GetPagesPrintedOverLastMonth()

	jsonResponse := getUsageResponse{DailyUsage: usageByDay}
	json.NewEncoder(w).Encode(&jsonResponse)
}
