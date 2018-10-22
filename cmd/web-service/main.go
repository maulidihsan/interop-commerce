package main

import (
	// "flag"
	// "io/ioutil"
	"log"
	// "net/http"
	// "strings"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/maulidihsan/flashdeal-webservice/pkg/api"
)
type ProductInfo struct {
	Item string `json:"item"`
	Images string `json:"images"`
	Stocks string `json:"stocks"`
}
func main() {
	router := gin.Default()
	router.GET("/api/promo/:src", func(c *gin.Context) {
		switch source := c.Param("src"); source {
			default:
				conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()

				service := api.NewPromoClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
				defer cancel()
				response, err := service.GetPromo(ctx, &api.Empty{})
				if err != nil {
					log.Fatalf("Create failed: %v", err)
				}
				var arr []ProductInfo
				var prod ProductInfo
				for _, product := range response.Product{
					prod = ProductInfo{
						Item: product.GetItem(),
						Images: product.GetImages(),
						Stocks: product.GetStocks(),
					}
					arr = append(arr, prod)
				}
				c.JSON(200, arr)
		}
	})
	router.Run(":8080")
	// t := time.Now().In(time.UTC)
	// pfx := t.Format(time.RFC3339Nano)

	// var body string

	// // Call Create
	// resp, err := http.Post(*restAddress+"/api/promo", "application/json", strings.NewReader(fmt.Sprintf(`
	// 	{
	// 		"api":"v1",
	// 		"toDo": {
	// 			"title":"title (%s)",
	// 			"description":"description (%s)",
	// 			"reminder":"%s"
	// 		}
	// 	}
	// `, pfx, pfx, pfx)))
	// if err != nil {
	// 	log.Fatalf("failed to call Create method: %v", err)
	// }
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Create response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// // parse ID of created ToDo
	// var created struct {
	// 	API string `json:"api"`
	// 	ID  string `json:"id"`
	// }
	// err = json.Unmarshal(bodyBytes, &created)
	// if err != nil {
	// 	log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
	// 	fmt.Println("error:", err)
	// }

	// // Call Read
	// resp, err = http.Get(fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID))
	// if err != nil {
	// 	log.Fatalf("failed to call Read method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Read response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Read response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// // Call Update
	// req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID),
	// 	strings.NewReader(fmt.Sprintf(`
	// 	{
	// 		"api":"v1",
	// 		"toDo": {
	// 			"title":"title (%s) + updated",
	// 			"description":"description (%s) + updated",
	// 			"reminder":"%s"
	// 		}
	// 	}
	// `, pfx, pfx, pfx)))
	// req.Header.Set("Content-Type", "application/json")
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatalf("failed to call Update method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Update response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Update response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// Call ReadAll
	// resp, err := http.Get(*restAddress + "/api/promo/all")
	// if err != nil {
	// 	log.Fatalf("failed to call ReadAll method: %v", err)
	// }
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read ReadAll response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("ReadAll response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// // Call Delete
	// req, err = http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID), nil)
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatalf("failed to call Delete method: %v", err)
	// }
	// bodyBytes, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	body = fmt.Sprintf("failed read Delete response body: %v", err)
	// } else {
	// 	body = string(bodyBytes)
	// }
	// log.Printf("Delete response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}