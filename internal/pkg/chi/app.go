package chi

import (
	chi2 "github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"rest-api/internal/app/book/delivery/web"
	"rest-api/internal/app/book/repository/repositoryimpl"
	"rest-api/internal/app/book/usecase/usecaseimpl"
	"rest-api/internal/pkg/mongo"
	"time"
)

func NewChiRoutes() http.Handler {
	chi := chi2.NewRouter()
	chi.Use(corsMiddleware())
	registerRoutes(chi)
	return http.TimeoutHandler(chi, 120*time.Minute, `{"Message": "Service Unavailable"}`)

}

func registerRoutes(r *chi2.Mux) *chi2.Mux {
	bookRoutes(r)
	return r
}

func bookRoutes(r *chi2.Mux) {
	config := mongo.MongoConfig
	bookRepository := repositoryimpl.NewBookRepository(repositoryimpl.BookRepositoryOpts{
		DB: config,
	})
	bookUsecaseimpl := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		bookRepository,
	})
	bookRoute := web.NewBookRegistry(web.BookRegistryOpts{
		bookUsecaseimpl,
	})
	bookRoute.RegisterRoutesTo(r)
}
func corsMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowedOrigins can be "*" for all origins, or specific origins
		AllowedOrigins:   []string{"https://example.com", "https://anotherdomain.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value for Access-Control-Max-Age
	})
}
