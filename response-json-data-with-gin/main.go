package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	Username     string `json:"username"`
	Guests       Guests `json:"guests"`
	RoomType     string `json:"roomType"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
}

type Guests struct {
	Person []Person `json:"person"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	r := gin.New()
	r.GET("/test", func(c *gin.Context) {

		p1 := Person{
			Firstname: "Tim",
			Lastname:  "Gid",
		}

		p2 := Person{
			Firstname: "Tim2",
			Lastname:  "Gid2",
		}

		guest := Guests{
			Person: []Person{
				p1, p2,
			},
		}

		f := CreateParams{
			Username:     "test",
			Guests:       guest,
			RoomType:     "test",
			CheckinDate:  "2019-08-08",
			CheckoutDate: "2019-09-09",
		}

		c.IndentedJSON(http.StatusOK, f)
	})
	r.Run(":8080")
}
