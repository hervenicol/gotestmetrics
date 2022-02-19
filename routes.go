package main

import ( 
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

func getRoot(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Welcome, "})
}

func getHello(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Hello world!"})
}

type JokeParameters struct {
    Quantity int `json:"qty" binding:"required"`
}

func getJoke(c *gin.Context) {
    var params JokeParameters
    var jokes = make(map[int]string)

    var maxQty = 10
    var minQty = 1
    var jokesQty = 1

    // If parameters format is good, set custom quantity
    if err := c.ShouldBindJSON(&params); err == nil {
        jokesQty = params.Quantity
    }

    fmt.Printf("parameters.qty=%d\n", jokesQty)

    // Error if quantity is out of bounds
    if jokesQty > maxQty || jokesQty < minQty {
        fmt.Printf("parameters.qty out of bounds\n")
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("qty must be between %d and %d", minQty, maxQty)})
        return
    }

    // retrieve n jokes
    for i:=0; i<jokesQty; i++ {
        jokeText, jokeStatus, err := getRandomJoke()

        // fail and stop looping if error
        if err != nil {
            fmt.Printf("error: %s\n", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("jokes provider failed with error %s" , err)})
            return
        }

        // fail and stop looping if error status
        if jokeStatus != 200 {
            fmt.Printf("error code: %d - text: %s\n", jokeStatus, jokeText)
            c.JSON(jokeStatus, gin.H{"error": fmt.Sprintf("jokes provider failed with error \"%s\"" , jokeText)})
            return
        }


        jokes[i] = jokeText
        fmt.Printf("joke %d = %s\n", i, jokes[i])
    }

    c.JSON(http.StatusOK, jokes)
}

