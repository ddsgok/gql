/* Package gql provides a low level GraphQL client.

 // create a client (safe to share across requests)
 client := gql.NewClient("https://machinebox.io/gql")

 // make a request
 req := gql.NewRequest(`
     query ($key: String!) {
         items (id:$key) {
             field1
             field2
             field3
         }
     }
 `)

 // set any variables
 req.Var("key", "value")

 // run it and capture the response
 var respData ResponseStruct
 if err := client.Run(ctx, req, &respData); err != nil {
     log.Fatal(err)
 }

Specify client

To specify your own http.Client, use the WithHTTPClient option:
 httpclient := &http.Client{}
 client := gql.NewClient("https://machinebox.io/gql", gql.WithHTTPClient(httpclient))
*/package gql
