package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "productsdb"
)

var natsConn stan.Conn
var db *sql.DB

func init() {
	var err error
	time.Sleep(5 * time.Second)
	natsConn, err = stan.Connect("test-cluster", "products-service", stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = natsConn.Subscribe("user.created", func(msg *stan.Msg) {
		log.Printf("Received user.created event: %s", string(msg.Data))
		msg.Ack()
		// Handle the event as needed
	}, stan.StartAt(0), stan.SetManualAckMode())

	if err != nil {
		log.Fatal(err)
	}

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

func main() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.POST("/products", addProduct)

	err := router.Run(":4000")
	if err != nil {
		log.Fatal(err)
	}
}

func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

func addProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, product)
}
