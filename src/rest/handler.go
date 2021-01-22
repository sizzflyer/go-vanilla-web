package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sizzflyer/go-vanilla-web/src/dblayer"
)

type HandlerInterface interface {
	Main(c *gin.Context) 
}

type Handler struct {
	db dblayer.DBLayer
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func NewHandler() (HandlerInterface, error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	mysqlDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	mysqlDB, err := dblayer.NewORM(mysqlDSN)
	if err != nil {
		return nil, err
	}

	return &Handler{
		db:     mysqlDB,
	}, nil
}

func (h *Handler) Main(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello go world")
}