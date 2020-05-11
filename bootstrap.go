package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	handler "github.com/taufiqade/gowallet/handler"
	"github.com/taufiqade/gowallet/models"
	dbRepo "github.com/taufiqade/gowallet/repository/db"

	// redisRepo "github.com/taufiqade/gowallet/repository/redis"
	service "github.com/taufiqade/gowallet/service"
)

/////////////////INIT REPOSITORY/////////////////////
var dbUserRepository models.IUserRepository
var dbUserBalanceRepository models.IUserBalanceRepository
var dbUserBalanceHistoryRepository models.IUserBalanceHistoryRepository

// var redisAuthRepository models.IRedisAuthRepository

/////////////END OF INIT REPOSITORY//////////////////

//////////////////INIT SERVICE//////////////////////
var authService models.IAuthService
var userService models.IUserService
var transactionService models.ITransactionService

///////////////END OF INIT SERVICE//////////////////

var dbConn *gorm.DB
var dbOnce sync.Once
var redisConn *redis.Client

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

		// init redis conn
		dsn := fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOST"), 6379)
		fmt.Print(dsn)
		redisConn = redis.NewClient(&redis.Options{
			Addr: dsn, //redis port
		})
		_, errr := redisConn.Ping().Result()
		if errr != nil {
			log.Fatal(err)
		}
	})
}

func initRepository() {
	dbUserRepository = dbRepo.NewUserRepository(dbConn)
	dbUserBalanceRepository = dbRepo.NewUserBalanceRepository(dbConn)
	dbUserBalanceHistoryRepository = dbRepo.NewUserBalanceHistoryRepository(dbConn)
	// redisAuthRepository = redisRepo.NewRedisAuthRepository(redisConn)
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
