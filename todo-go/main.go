package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Goの言語仕様を学習する", Completed: false},
	{ID: "2", Item: "Goのチュートリアルを実践する", Completed: false},
	{ID: "3", Item: "Goで動くアプリを開発する", Completed: false},
	{ID: "4", Item: "QiitaにGoに関するアウトプットを行う", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)

	router.Run("localhost:8080")
}
