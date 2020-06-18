package controllers

import (
	"github.com/thienkimlove/phunghanh_api/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from users")

	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": err})
	}

}

func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=? LIMIT 1", id)

	if err == nil {
		userId, _ := strconv.ParseInt(id, 0, 64)

		content := &models.User{
			Id:        userId,
			Password:  user.Password,
			Name: user.Name,
			Email: user.Email,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": err})
	}
}

func Login(c *gin.Context) {
	var user models.User
	errBind := c.Bind(&user)
	if errBind != nil {
		log.Println(errBind)
	}
	err := dbmap.SelectOne(&user, "select * from users where Username=? LIMIT 1", user.Email)

	if err == nil {
		userId := user.Id

		content := &models.User{
			Id:        userId,
			Password:  user.Password,
			Name: user.Name,
			Email: user.Email,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": err})
	}

}

func PostUser(c *gin.Context) {
	var user models.User
	errBind := c.Bind(&user)
	if errBind != nil {
		log.Println(errBind)
	}
	log.Println(user.Name)

	if user.Password != "" && user.Name != "" && user.Email != "" {

		if insert, errInsert := dbmap.Exec(`INSERT INTO users (password, name, email) VALUES (?, ?, ?)`, user.Password, user.Name, user.Email); insert != nil {
			userId, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:        userId,
					Password:  user.Password,
					Name: user.Name,
					Email: user.Email,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		} else {
			c.JSON(400, gin.H{"error": errInsert})
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=?", id)

	if err == nil {
		var json models.User

		errBind := c.Bind(&json)
		if errBind != nil {
			log.Println(errBind)
		}

		userId, _ := strconv.ParseInt(id, 0, 64)

		user := models.User{
			Id:        userId,
			Password:  user.Password,
			Name: json.Name,
			Email: user.Email,
		}

		if user.Name != "" {
			_, err = dbmap.Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
