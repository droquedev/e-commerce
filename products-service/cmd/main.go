package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/droquedev/e-commerce/pkg/nats"
	"github.com/droquedev/e-commerce/products-service/internal/listeners"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "productsdb"
)

var db *sql.DB

func init() {
	var err error
	time.Sleep(5 * time.Second)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database")
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product = []Product{
	{
		ID:    1,
		Name:  "Cheese",
		Price: 500,
	},
	{
		ID:    2,
		Name:  "Milk",
		Price: 200,
	},
}

func main() {
	router := gin.Default()

	natsConn := nats.GetNatsConn().NatsConn

	listeners.InitializeListeners(natsConn)

	router.GET("/products", getProducts)
	router.POST("/products", addProduct)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func addProduct(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = len(products) + 1
	products = append(products, product)

	c.JSON(http.StatusOK, product)
}
