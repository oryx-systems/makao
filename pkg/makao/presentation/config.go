package presentation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/application/extension"
	pgDB "github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
	"github.com/oryx-systems/makao/pkg/makao/presentation/graph"
	"github.com/oryx-systems/makao/pkg/makao/presentation/graph/generated"
	"github.com/oryx-systems/makao/pkg/makao/presentation/rest"
	"github.com/oryx-systems/makao/pkg/makao/usecases"
	"github.com/oryx-systems/makao/pkg/makao/usecases/house"
	"github.com/oryx-systems/makao/pkg/makao/usecases/otp"
	"github.com/oryx-systems/makao/pkg/makao/usecases/residence"
	"github.com/oryx-systems/makao/pkg/makao/usecases/user"
)

const serverTimeoutSeconds = 120

// MakaoServiceAllowedOrigins is a list of CORS origins allowed to interact with this service
var MakaoServiceAllowedOrigins = []string{
	"http://localhost:8080",
}

// MakaoServiceAllowedHeaders is a list of CORS allowed headers for the clinical
// service
var MakaoServiceAllowedHeaders = []string{
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
	router, err := StartGinRouter(ctx)
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	// Set allowed origins
	router.Use(cors.New(cors.Config{
		AllowOrigins:     MakaoServiceAllowedOrigins,
		AllowHeaders:     MakaoServiceAllowedHeaders,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	}))

	// Use custom http to serve request via GIN
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
		WriteTimeout: serverTimeoutSeconds * time.Second,
	}

	return srv
}

// HealthStatusCheck endpoint to check if the server is working.
func HealthStatusCheck(w gin.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(true)
	if err != nil {
		log.Fatal(err)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL IDE", "/v1/auth/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GQLHandler sets up a GraphQL resolver
func GQLHandler(ctx context.Context,
	usecase usecases.Makao,
) gin.HandlerFunc {
	resolver, err := graph.NewResolver(ctx, usecase)
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}

// StartGinRouter sets up the GIN router
func StartGinRouter(ctx context.Context) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.Recovery())

	pg, err := gorm.NewPGInstance()
	if err != nil {
		return nil, fmt.Errorf("can't instantiate repository in resolver: %v", err)
	}

	db := pgDB.NewDBService(pg, pg, pg)
	ext := extension.NewExtension()

	userUsecase := user.NewUseCasesUser(db, db, db, ext)
	otpUsecase := otp.NewUseCaseOTP(db, db)
	residenceUsecase := residence.NewResidence(db)
	houseUsecase := house.NewUsecaseHouse(db, db)

	usecases := usecases.NewMakaoUsecase(userUsecase, otpUsecase, residenceUsecase, houseUsecase)
	h := rest.NewPresentationHandlers(*usecases)

	api := r.Group("/v1/api")
	{
		api.GET("/login_by_phone", h.HandleLoginByPhone())
		api.POST("/sign_up", h.HandleRegistration())
		api.GET("/ide", PlaygroundHandler())
		api.POST("/pin", h.SetUserPIN())
		api.GET("/user", h.GetUserProfileByPhoneNumber())
	}

	// Authenticated routes
	auth := r.Group("/v1/auth")
	auth.Use(rest.AuthMiddleware())
	{
		auth.POST("/graphql", GQLHandler(ctx, *usecases))
	}

	return r, nil
}
