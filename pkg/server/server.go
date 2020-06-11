package server

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/tsoonjin/raspy/internal/handlers"
    "github.com/tsoonjin/raspy/pkg/utils"
)

var host, port, gqlPath, gqlPgPath string

func init() {
    host = utils.MustGet("GQL_SERVER_HOST")
    port = utils.MustGet("GQL_SERVER_PORT")
    gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
    gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
}

// Run spins up the server
func Run() {
    endpoint := "http://" + host + ":" + port
    r := gin.Default()
    // Simple keep-alive/ping handler
    r.GET("/ping", handlers.Ping())

    r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
    log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)

    r.POST(gqlPath, handlers.GraphqlHandler())
    log.Println("GraphQL @ " + endpoint + gqlPath)

    // Inform the user where the server is listening
    log.Println("Running @ http://" + host + ":" + port)
    // Print out and exit(1) to the OS if the server cannot run
    log.Fatalln(r.Run(host + ":" + port))
}
