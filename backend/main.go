package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/weaviate/weaviate-go-client/v4/weaviate"
    "github.com/weaviate/weaviate-go-client/v4/weaviate/config"
)

var weaviateClient *weaviate.Client

func main() {
    initWeaviate()

    router := gin.Default()
    router.POST("/query", handleQuery)

    log.Println("Server running at http://localhost:8080")
    router.Run(":8080")
}

func initWeaviate() {
    cfg := config.Config{Scheme: "http", Host: "localhost:8081"}
    client, err := weaviate.New(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to Weaviate: %v", err)
    }
    weaviateClient = client
}

func handleQuery(c *gin.Context) {
    type Query struct {
        UserInput string `json:"userInput"`
    }

    var query Query
    if err := c.BindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    response := "Placeholder response for: " + query.UserInput
    c.JSON(http.StatusOK, gin.H{"response": response})
}
