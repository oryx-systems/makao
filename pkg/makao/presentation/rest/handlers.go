package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
	"github.com/oryx-systems/makao/pkg/makao/usecases"
)

// AcceptedContentTypes is a list of all the accepted content types
var AcceptedContentTypes = []string{"application/json", "application/x-www-form-urlencoded"}

// PresentationHandlers represents all the REST API logic
type PresentationHandlers interface {
	HandleIncomingMessages() gin.HandlerFunc
}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	usecases       usecases.Usecases
	infrastructure infrastructure.Interactor
}

// NewPresentationHandlers initializes a new rest handlers usecase
func NewPresentationHandlers(infrastructure infrastructure.Interactor, usecases usecases.Usecases) PresentationHandlers {
	return &PresentationHandlersImpl{infrastructure: infrastructure, usecases: usecases}
}

// HandleIncomingMessages handles and processes data posted by AIT to its callback URL
func (p PresentationHandlersImpl) HandleIncomingMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := c.Request.Context()
		// c.Accepted = append(c.Accepted, AcceptedContentTypes...)

		// payload := &dto.IncomingSMSPayload{}
		// err := c.Request.ParseForm()
		// if err != nil {
		// 	utils.ReportErr(c.Writer, err, http.StatusBadRequest)
		// }

		// payload.ID = c.Request.Form.Get("id")
		// payload.LinkID = c.Request.Form.Get("linkId")
		// payload.Date = c.Request.Form.Get("date")
		// payload.From = c.Request.Form.Get("from")
		// payload.To = c.Request.Form.Get("to")
		// payload.Text = c.Request.Form.Get("text")
		// payload.NetworkCode = c.Request.Form.Get("networkCode")

		// err = p.usecases.HandleIncomingMessages(ctx, payload)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{"status": "Successfully processed short code sms"})
	}
}
