package main

import (
	"database/sql"
	"fmt"
	"log"

	_handler "github.com/KennyKur/CRUD_Todo/handler"
	"github.com/KennyKur/CRUD_Todo/repository"
	"github.com/KennyKur/CRUD_Todo/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type TodoHandler struct {
	TodoUsecase _handler.TodoUsecaseInterface
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	r := gin.Default()
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open(`postgres`, connection)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	repoTodo := repository.NewTodoRepository(dbConn)
	usecaseTodo := usecase.NewTodoUsecase(repoTodo)
	api := r.Group("/v1")
	_handler.NewTodoHandler(api, usecaseTodo)
	r.Run()
}
