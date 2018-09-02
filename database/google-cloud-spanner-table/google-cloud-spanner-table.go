package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/spanner"
)

func insert(client *spanner.Client) error {
	_, err := client.Apply(context.Background(), []*spanner.Mutation{
		spanner.InsertOrUpdate("JeffTable",
			[]string{"name", "age"},
			[]interface{}{"lea", 26}),
		spanner.InsertOrUpdate("JeffTable",
			[]string{"name", "age"},
			[]interface{}{"stepan", 26}),
	})
	return err
}

func readAndPrint(client *spanner.Client) error {
	stmt := spanner.NewStatement("SELECT name, age FROM JeffTable")
	iter := client.Single().Query(context.Background(), stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		var name string
		var age int64

		err = row.Columns(&name, &age)
		if err != nil {
			return err
		}

		fmt.Printf("name=%s, age=%d\n", name, age)
	}
	return nil
}

func main() {

	spannerInstance := "projects/" + os.Getenv("PROJECTNAME") + "/instances/s-spanner/databases/jeff-test"

	client, err := spanner.NewClient(context.Background(), spannerInstance)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := insert(client); err != nil {
		panic(err)
	}

	if err := readAndPrint(client); err != nil {
		panic(err)
	}
}
