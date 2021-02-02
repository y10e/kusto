package main

import (
	"fmt"
	"context"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/data/table"
)

func main(){

	// Create an authorizer with an Device Auth. (you have to create aad app.)
	config := auth.NewDeviceFlowConfig("8ee3ad48-8b30-42d3-b767-f0fd7bbea294","common")
	config.Resource = "https://help.kusto.windows.net"
	fmt.Println(config.Resource)
	auth, err := config.Authorizer()
	if err != nil {
		fmt.Println(err)
		panic("add error handling")
	}

	authorizer := kusto.Authorization{
		Authorizer: auth,
	}

	/*
	authorizer := kusto.Authorization{
		Config: auth.NewClientCredentialsConfig("badd67f9-74de-413a-b8c4-93b7b8d71f06", "oT5YTqb4ScVyy4ncsp_xvt651M2uiP84Z8", "72f988bf-86f1-41af-91ab-2d7cd011db47"),
	}
	*/

	client, err := kusto.New("https://help.kusto.windows.net/;Fed=true", authorizer)
	if err != nil {
		panic("add error handling")
	}

	// table package is: github.com/Azure/azure-kusto-go/kusto/data/table

	ctx := context.Background()

	// Query our database table "systemNodes" for the CollectionTimes and the NodeIds.
	iter, err := client.Query(ctx, "Samples", kusto.NewStmt("PopulationData | project State, Population"))
	if err != nil {
		fmt.Println(err)
		panic("add error handling")
	}
	defer iter.Stop()

	// .Do() will call the function for every row in the table.
	err = iter.Do(
		func(row *table.Row) error {
			fmt.Println(row) // As a convenience, printing a *table.Row will output csv
			return nil
		},
	)
	if err != nil {
		panic("add error handling")
	}

}