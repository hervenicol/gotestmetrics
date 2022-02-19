package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
)

// Get a random joke from icanhazdadjoke.com
// returns:
// - the joke (text)
// - the http status code (int)
// - an error
func getRandomJoke() (string, int, error) {

    req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

    // for json response
    //req.Header.Set("Accept", "application/json")

    // for just the joke, as text
    req.Header.Set("Accept", "text/plain")

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
        fmt.Printf("Error getRandomJoke: %s\n", err)
        return fmt.Sprint("Error: %s", err), 500, err
    }

    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)

    return bodyString, resp.StatusCode, err
}
