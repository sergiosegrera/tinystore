package main

import (
	"database/sql"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

// TODO: User authentication
// TODO: Create data struct

type Data struct {
	code    int
	message map[string]string
}

var database *sql.DB

func main() {
	database, _ = sql.Open("sqlite3", "./store.db")
	defer database.Close()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS images (id INTEGER PRIMARY KEY, uuid TEXT, location TEXT, time DATETIME)")
	statement.Exec()

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/store", "./store")
	router.POST("/upload", upload)
	router.GET("/image/:uuid", image)

	router.Run(":8080")
}

func image(c *gin.Context) {
	var location string
	uuid := c.Param("uuid")
	err := database.QueryRow("SELECT location FROM images WHERE uuid=?", uuid).Scan(&location)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"location": location,
	})
}

// TODO: Add automatic compression
func upload(c *gin.Context) {
	file, err := c.FormFile("image")
	splitFileName := strings.Split(file.Filename, ".")
	fileType := splitFileName[len(splitFileName)-1]
	log.Println(fileType)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	id := shortuuid.New()

	statement, err := database.Prepare("INSERT INTO images (uuid, location, time) VALUES (?, ?, datetime('now'))")
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	statement.Exec(id, "/store/"+id+"."+fileType)
	err = c.SaveUploadedFile(file, "./store/"+id+"."+fileType)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"uuid":     id,
		"location": "/store/" + id + "." + fileType,
		"filetype": fileType,
		"size":     file.Size,
	})
}
