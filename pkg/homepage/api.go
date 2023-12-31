package homepage

import (
	"net/http"
	"time"

	"github.com/NotCoffee418/GoWebsite-Boilerplate/internal/types"
	"github.com/gin-gonic/gin"
)

// HomePageData is the response for the time call
type HomePageData struct {
	Time string `json:"time"`
}

// HomeApiHandler Implements types.HandlerRegistrar interface
type HomeApiHandler struct{}

// Initialize is called before the handler is registered
func (h *HomeApiHandler) Initialize(_ *types.HandlerInitContext) {
	// Nothing to initialize
}

// Handler Implements PageRouteRegistrar interface
func (h *HomeApiHandler) Handler(engine *gin.Engine) {
	engine.GET("/api/home/get-server-time", h.get)
}

func (h *HomeApiHandler) get(c *gin.Context) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	resp := types.ApiResponseFactory.Ok(
		&HomePageData{Time: timeStr})

	// Render page
	c.JSON(http.StatusOK, resp)
}
