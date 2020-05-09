package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"sync"
	handler "test/handler"
	"test/models"
	repository "test/repository/db"
	service "test/service"
)

/////////////////INIT REPOSITORY/////////////////////
var dbUserRepository models.IUserRepository
var dbUserBalanceRepository models.IUserBalanceRepository
var dbUserBalanceHistoryRepository models.IUserBalanceHistoryRepository
/////////////END OF INIT REPOSITORY//////////////////

//////////////////INIT SERVICE//////////////////////
var authService models.IAuthService
var userService models.IUserService
var transactionService models.ITransactionService
///////////////END OF INIT SERVICE//////////////////



var dbConn *gorm.DB
var dbOnce sync.Once

func initDB() {
	dbOnce.Do(func() {
		connStr := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
		db, err := gorm.Open("mysql", connStr)
		if err != nil {
			log.Fatal(err)
		}
		dbConn = db

		defer db.Close()
	})
}

func initRepository() {
	dbUserRepository = repository.NewUserRepository(dbConn)
	dbUserBalanceRepository = repository.NewUserBalanceRepository(dbConn)
	dbUserBalanceHistoryRepository = repository.NewUserBalanceHistoryRepository(dbConn)
}

func initService() {
	authService = service.NewAuthService(dbUserRepository)
	userService = service.NewUserService(dbUserRepository)
	transactionService = service.NewTransactionService(dbUserRepository, dbUserBalanceRepository, dbUserBalanceHistoryRepository)
}

func serveHTTP() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi there",
		})
	})

	//init handler
	handler.NewAuthHandler(r, authService)
	handler.NewUserHandler(r, userService)
	handler.NewTransactionHandler(r, transactionService)

	r.Run() // by default listen and serve on 0.0.0.0:8080
}