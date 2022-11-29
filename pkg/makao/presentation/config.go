package presentation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure"
	"github.com/oryx-systems/makao/pkg/makao/presentation/rest"
	"github.com/oryx-systems/makao/pkg/makao/usecases"
)

const serverTimeoutSeconds = 120

// SMSServiceAllowedOrigins is a list of CORS origins allowed to interact with
// this service
var SMSServiceAllowedOrigins = []string{
	"http://localhost:8080",
	"https://oryx-staging-6hapifddxq-nw.a.run.app",
}

// SMSServiceAllowedHeaders is a list of CORS allowed headers for the clinical
// service
var SMSServiceAllowedHeaders = []string{
	"Accept",
	"Accept-Charset",
	"Accept-Language",
	"Accept-Encoding",
	"Origin",
	"Host",
	"User-Agent",
	"Content-Length",
	"Content-Type",
	"Authorization",
	"X-Authorization",
}

// PrepareServer sets up the HTTP server
func PrepareServer(
	ctx context.Context,
	port int,
	allowedOrigins []string,
) *http.Server {
	// start up the router
	r, err := StartGinRouter(ctx)
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	// Set allowed origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:     SMSServiceAllowedOrigins,
		AllowHeaders:     SMSServiceAllowedHeaders,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	}))

	// Use custom http to serve request via GIN
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
		WriteTimeout: serverTimeoutSeconds * time.Second,
	}

	return srv
}

// StartGinRouter sets up the GIN router
func StartGinRouter(ctx context.Context) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	infrastructure := infrastructure.NewInfrastructureInteractor()
	usecases := usecases.NewUseCasesInteractor(infrastructure)
	h := rest.NewPresentationHandlers(infrastructure, usecases)

	api := r.Group("/api/v1")
	{
		api.POST("/incoming_messages", h.HandleIncomingMessages())
	}

	return r, nil
}
