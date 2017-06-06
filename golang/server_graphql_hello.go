package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	pgql "github.com/playlyfe/go-graphql"
)

// see https://echo.labstack.com/guide
func main() {
	e := echo.New()
	e.GET("/graphql-go", graphql0)
	e.GET("/go-graphql", graphql1)
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

//-----------------------------------------------------------------------------
// use https://github.com/playlyfe/go-graphql
var schema1 = `
	type QueryRoot {
	    hello: String
	}
	`
var resolvers1 = map[string]interface{}{
	"QueryRoot/hello": func(params *pgql.ResolveParams) (interface{}, error) {
		return "world", nil
	},
}
var executor1, _ = pgql.NewExecutor(schema1, "QueryRoot", "", resolvers1)

func graphql1(context echo.Context) error {
	// executor1.ResolveType = func(value interface{}) string {
	// 	if object, ok := value.(map[string]interface{}); ok {
	// 		return object["__typename"].(string)
	// 	}
	// 	return ""
	// }
	query := context.QueryParam("query")
	// body, err := ioutil.ReadAll(context.Request().Body())
	// if err != nil {
	// 	panic(err)
	// }
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	panic(err)
	// }
	variables := map[string]interface{}{}
	result, err := executor1.Execute(context, query, variables, "")
	if err != nil {
		panic(err)
	}
	rJSON, _ := json.Marshal(result)
	return context.String(http.StatusOK, fmt.Sprintf("%s", rJSON))
}
