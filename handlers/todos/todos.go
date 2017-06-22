package todos

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/justcoders/go-api-example/db"
	"github.com/justcoders/go-api-example/models"
)

func Create(c *gin.Context) {
	movie := models.Todo{}

	if err := c.Bind(&movie); err != nil {
		c.Error(err)
		return
	}

	if err := db.Todos.Insert(movie); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func GetList(c *gin.Context) {
	movies := []models.Todo{}

	if err := db.Todos.Find(nil).All(&movies); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, movies)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	movie := models.Todo{}

	if !bson.IsObjectIdHex(id) {
		c.Error(errors.New("Not valid id"))
		return
	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Error(errors.New("Not valid id"))
		return
	} else if err := db.Todos.Find(bson.M{"_id": obj}).One(&movie); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, movie)
}
