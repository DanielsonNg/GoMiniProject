package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/danielsonng/ecomgo/internal/adapters/postgresql/sqlc"
	"github.com/danielsonng/ecomgo/internal/orders"
	"github.com/danielsonng/ecomgo/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // !important for rate limiting
	r.Use(middleware.RealIP)    // !important for rate limiting & analytics & tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Heal Thy"))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/product/{id}", productHandler.GetProductById)

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)
	r.Post("/orders", orderHandler.PlaceOrder)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server Started at address %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	//logger
	//db driver
	db *pgx.Conn
}

type config struct {
	addr string //port
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
