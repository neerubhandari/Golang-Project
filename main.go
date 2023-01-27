package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct{
   ID string `json:"id"`
   Name string `json:"name"`
   Price int `json:"price"`
   Imahe string `json:"image"`
}

var product [] Product
func init(){
   product =make ([]Product ,0)
   file,_:=ioutil.ReadFile("back.json")
   _=json.Unmarshal([]byte(file),&product)
}

func getProducts(c *gin.Context) {
   c.JSON(http.StatusOK, product)
}
func main() {
   router := gin.Default()
   router.GET("/products", getProducts)
   router.Run()}