package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"go.opentelemetry.io/otel"

	"github.com/hammer-code/lms-be/app"
	"github.com/hammer-code/lms-be/config"
	_ "github.com/hammer-code/lms-be/docs"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/ngelog"
	"github.com/hammer-code/lms-be/utils"

	// _ "swagger-mux/docs"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/swaggo/swag"
)

var serveHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "launches an HTTP server",
	Long:  "the serveHttp command initiates an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// load add package serve http here
		ctx := context.Background()

		// Init OpenTelemetry

		cfg := config.GetConfig()

		app := app.InitApp(cfg)

		// route
		router := registerHandler(app)

		// build cors
		muxCorsWithRouter := cors.New(cors.Options{
			AllowedOrigins:   cfg.CORS_ALLOWED_ORIGINS,
			AllowedHeaders:   cfg.CORS_ALLOWED_HEADERS,
			AllowedMethods:   cfg.CORS_ALLOWED_METHODS,
			AllowCredentials: true,
		}).Handler(router)

		srv := &http.Server{
			Addr:    cfg.APP_PORT,
			Handler: muxCorsWithRouter,
		}

		go func() {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			ngelog.Info(ctx, "service shutdown")
			if err := srv.Shutdown(ctx); err == context.DeadlineExceeded {
				ngelog.Error(ctx, "svr.Shutdown: context deadline exceeded", err)
			}
		}()

		ngelog.Info(ctx, fmt.Sprintf("server started, running on port %s", cfg.APP_PORT))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ngelog.Fatal(ctx, "starting server failed", err)
		}
	},
}

func LoadJSON(path string) string {
	jsonBytes, err := os.ReadFile(path)

	// jsonBytes, err := os.ReadFile("documentation/users.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return ""
	}
	return string(jsonBytes)
}

func LoadSwagger() {
	userTemplate := LoadJSON("documentation/users.json")
	var UsersSwaggerInfo = &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  userTemplate,
	}
	swag.Register(UsersSwaggerInfo.InstanceName(), UsersSwaggerInfo)
}

func init() {
	// LoadSwagger()
	rootCmd.AddCommand(serveHttpCmd)

}

func health(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("Test Trace")

	ctx, span := tracer.Start(r.Context(), "health controller")
	defer span.End()

	ngelog.Info(ctx, "service health good")
	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "good",
		Data:    nil,
	}, w)
}

func registerHandler(app app.App) *mux.Router {

	router := mux.NewRouter()
	router.Use(app.Middleware.LogMiddleware)
	router.HandleFunc("/health", health)

	return router
}
