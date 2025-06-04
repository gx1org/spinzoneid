package app

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func handleDetail(c *gin.Context) {
	spin := Spin{}
	DB.Where("id", c.Param("id")).First(&spin)
	if spin.ID == "" {
		c.HTML(404, "404.html", gin.H{
			"error": "Spin not found or already deleted",
		})
		return
	}

	result := Result{}
	token := c.Query("token")
	if token != "" {
		DB.Where("spin_id", spin.ID).Where("token", token).First(&result)
		if result.ID == "" {
			result.Spin = &spin
			result.Token = token
			generateResult(&result)
		}
	}

	c.HTML(200, "detail.html", gin.H{
		"spin":    spin,
		"options": parseOptions(spin.Options),
		"result":  result,
	})
}

func generateResult(result *Result) {
	result.ID = generateID()
	result.SpinID = result.Spin.ID
	result.SelectedOption = pickRandomOption(parseOptions(result.Spin.Options))
	result.Spin = nil
	if err := DB.Create(&result); err != nil {
		result.ID = generateID()
		DB.Create(&result)
	}
}

func handleCreate(c *gin.Context) {
	req := SpinPayload{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	curatedOptions := strings.Join(parseOptions(req.Options), "\n")
	spin := Spin{
		ID:       generateID(),
		Name:     req.Name,
		Options:  curatedOptions,
		Comment:  req.Comment,
		Password: req.Password,
	}
	if err := DB.Create(&spin); err != nil {
		spin.ID = generateID()
		DB.Create(&spin)
	}

	c.JSON(200, gin.H{
		"spin": spin,
	})
}

func handleUpdate(c *gin.Context) {
	req := SpinPayload{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	spin := Spin{}
	DB.Where("id", c.Param("id")).First(&spin)
	if spin.ID == "" {
		c.JSON(404, gin.H{
			"error": "Spin not found or already deleted",
		})
		return
	}

	if spin.Password != req.Password {
		c.JSON(401, gin.H{
			"error": "Password mismatch",
		})
		return
	}

	if req.IsDelete {
		DB.Delete(&spin)
	} else {
		spin.Comment = req.Comment
		DB.Save(&spin)
	}

	c.JSON(200, gin.H{
		"message": "Success!",
	})
}
