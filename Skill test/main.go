package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func postCustomer(c *gin.Context) {
    var newCustomer customer

    if err := c.BindJSON(&newCustomer); err != nil {
        return
    }

    customers = append(customers, newCustomer)
    c.IndentedJSON(http.StatusCreated, newCustomer)
}

type customer struct {
    ID     string  `json:"id"`
    Username  string  `json:"username"`
    Pass string  `json:"pass"`
}

var customers = []customer{
    {ID: "1", Username: "Kenneth", Pass: "12345678"},
}

func main() {
    router := gin.Default()
    router.GET("/customers", getCustomer)
    router.POST("/customers", postCustomer)

    router.Run("localhost:8080")
}

func getCustomer(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, customers)
}