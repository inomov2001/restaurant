package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	structs "github.com/gokurs/Projects/restaurant/repository/struct"
)

func (h Server) FoddGet1(c *gin.Context) {
	// req := structs.MenyuJson{}
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	//  fmt.Println(err)
	// }

	r1, err := h.repo.Food1()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r1)
}

func (h Server) FoodGet2(c *gin.Context) {
	r2, err := h.repo.Food2()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r2)
}

func (h Server) FoodGet3(c *gin.Context) {
	r3, err := h.repo.Food3()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r3)

}

func (h Server) SaladGet(c *gin.Context) {
	r4, err := h.repo.Salad()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, r4)

}

func (h Server) DrinksGet(c *gin.Context) {
	r5, err := h.repo.Drinks()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, r5)
}

func (h Server) OpenChekPost(c *gin.Context) {
	t_id := c.Query("table_id")
	err := h.repo.OpenChek(t_id)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}

func (h Server) ShopPost(c *gin.Context) {
	var req structs.ShopStruct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := h.repo.Shop(req.TableId, req.FoodId, req.SaladId, req.DrinckId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (h Server) ChekGet(c *gin.Context) {
	a := c.Query("table_id")
	r8, err := h.repo.Chek(a)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r8)
}
