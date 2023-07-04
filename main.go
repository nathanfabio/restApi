package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Client struct {
	ID        int `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Age       int    `json:"age,omitempty"`
	Order     *Order `json:"order,omitempty"`
}

type Order struct {
	Cake   string `json:"cake,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

var clients []Client
var clientID int

func getAllClients(c *gin.Context) {
	c.JSON(http.StatusOK, clients)
}


func getClientByID(c *gin.Context) {
	idParam := c.Param("id")
	clientID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Client ID"})
		return
	}

	for _, client := range clients {
		if client.ID == clientID {
			c.JSON(http.StatusOK, client)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
}

func createClient(c *gin.Context) {
	var newClient Client

	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	clientID++
	newClient.ID = clientID
	clients = append(clients, newClient)

	c.JSON(http.StatusCreated, newClient)
}


func deleteClient(c *gin.Context) {
	idParam := c.Param("id")
	clientID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Client ID"})
		return
	}

	for i, client := range clients {
		if client.ID == clientID {
			clients = append(clients[:i], clients[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
} 


func main() {
	r := gin.Default()

	r.GET("/clients", getAllClients)

	r.GET("/clients/:id", getClientByID)

	r.POST("/clients", createClient)

	r.DELETE("/clients/:id", deleteClient)

	clientID = 0
	clients = []Client{}

	r.Run(":8080")
}