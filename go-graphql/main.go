package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println("GraphQL Tutorial")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	// defines the object config
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	// creates our schema
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new graphql schema, err %v", err)
	}

	query := `
		{
			hello
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s /v", rJSON)
}
