package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// see https://echo.labstack.com/guide
func main() {
	e := echo.New()
	e.GET("/graphql", graphql0)
	e.Logger.Fatal(e.Start(":3003"))
}

//-----------------------------------------------------------------------------
//use github.com/graphql-go/graphql
var fields0 = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}
var rootQuery0 = graphql.ObjectConfig{Name: "RootQuery", Fields: fields0}
var schemaConfig0 = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery0)}
var schema0, _ = graphql.NewSchema(schemaConfig0)

func graphql0(context echo.Context) error {
	// Query
	query := context.QueryParam("query")
	params := graphql.Params{Schema: schema0, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		fmt.Printf("query: %s \n", query)
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	return context.String(http.StatusOK, fmt.Sprintf("%s", rJSON))
	//fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
