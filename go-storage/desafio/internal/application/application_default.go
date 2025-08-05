package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationDefault struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Addr is the server address.
	Addr string
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationDefault{
		cfgDb:   defaultCfg.Db,
		cfgAddr: defaultCfg.Addr,
	}
}

// ApplicationDefault is an implementation of the Application interface.
type ApplicationDefault struct {
	// cfgDb is the database configuration.
	cfgDb *mysql.Config
	// cfgAddr is the server address.
	cfgAddr string
	// db is the database connection.
	db *sql.DB
	// router is the chi router.
	router *chi.Mux
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - db: init
	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(a.db)
	rpProduct := repository.NewProductsMySQL(a.db)
	rpInvoice := repository.NewInvoicesMySQL(a.db)
	rpSale := repository.NewSalesMySQL(a.db)
	// - service
	svCustomer := service.NewCustomersDefault(rpCustomer)
	svProduct := service.NewProductsDefault(rpProduct)
	svInvoice := service.NewInvoicesDefault(rpInvoice)
	svSale := service.NewSalesDefault(rpSale)
	// - handler
	hdCustomer := handler.NewCustomersDefault(svCustomer)
	hdProduct := handler.NewProductsDefault(svProduct)
	hdInvoice := handler.NewInvoicesDefault(svInvoice)
	hdSale := handler.NewSalesDefault(svSale)

	// routes
	// - router
	a.router = chi.NewRouter()
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - endpoints
	a.router.Route("/customers", func(r chi.Router) {
		// - GET /customers
		r.Get("/", hdCustomer.GetAll())
		// - POST /customers
		r.Post("/", hdCustomer.Create())
	})
	a.router.Route("/products", func(r chi.Router) {
		// - GET /products
		r.Get("/", hdProduct.GetAll())
		// - POST /products
		r.Post("/", hdProduct.Create())
	})
	a.router.Route("/invoices", func(r chi.Router) {
		// - GET /invoices
		r.Get("/", hdInvoice.GetAll())
		// - POST /invoices
		r.Post("/", hdInvoice.Create())
	})
	a.router.Route("/sales", func(r chi.Router) {
		// - GET /sales
		r.Get("/", hdSale.GetAll())
		// - POST /sales
		r.Post("/", hdSale.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	defer a.db.Close()

	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}

// Insert Data
func (a *ApplicationDefault) InsertData() (err error) {
	customersData, err := os.ReadFile("docs/db/json/customers.json")
	if err != nil {
		return fmt.Errorf("failed to read customers data: %w", err)
	}

	var customers []handler.CustomerJSON
	if err := json.Unmarshal(customersData, &customers); err != nil {
		return fmt.Errorf("failed to unmarshal customers data: %w", err)
	}

	productsData, err := os.ReadFile("docs/db/json/products.json")
	if err != nil {
		return fmt.Errorf("failed to read products data: %w", err)
	}
	var products []handler.ProductJSON
	if err := json.Unmarshal(productsData, &products); err != nil {
		return fmt.Errorf("failed to unmarshal products data: %w", err)
	}

	invoicesData, err := os.ReadFile("docs/db/json/invoices.json")
	if err != nil {
		return fmt.Errorf("failed to read invoices data: %w", err)
	}

	var invoices []handler.InvoiceJSON
	if err := json.Unmarshal(invoicesData, &invoices); err != nil {
		return fmt.Errorf("failed to unmarshal invoices data: %w", err)
	}

	salesData, err := os.ReadFile("docs/db/json/sales.json")
	if err != nil {
		return fmt.Errorf("failed to read sales data: %w", err)
	}

	var sales []handler.SaleJSON
	if err := json.Unmarshal(salesData, &sales); err != nil {
		return fmt.Errorf("failed to unmarshal sales data: %w", err)
	}
	for _, c := range customers {
		_, err = a.db.Exec("INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (?, ?, ?, ?)", c.Id, c.FirstName, c.LastName, c.Condition)
		if err != nil {
			return fmt.Errorf("failed to insert customers data: %w", err)
		}
	}

	for _, i := range invoices {
		_, err = a.db.Exec("INSERT INTO invoices (id, customer_id, total) VALUES (?, ?, ?)", i.Id, i.CustomerId, i.Total)
		if err != nil {
			return fmt.Errorf("failed to insert invoices data: %w", err)
		}
	}

	for _, p := range products {
		_, err = a.db.Exec("INSERT INTO products (id, description, price) VALUES (?, ?, ?)", p.Id, p.Description, p.Price)
		if err != nil {
			return fmt.Errorf("failed to insert products data: %w", err)
		}
	}
	for _, s := range sales {
		_, err = a.db.Exec("INSERT INTO sales (id, invoice_id, product_id, quantity) VALUES (?, ?, ?, ?)", s.Id, s.InvoiceId, s.ProductId, s.Quantity)
		if err != nil {
			return fmt.Errorf("failed to insert sales data: %w", err)
		}
	}

	return nil
}
