package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ServerDbFieldAddition struct {
	DatabaseName string `json:"database_name"`
	FieldName    string `json:"field_name"`
	FieldType    string `json:"field_type"`
	Comment      string `json:"comment"`
	Option       string `json:"option"`
}

func runServer() {
	r := gin.Default()
	const srcDir = "/app/sourcedir"

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server running!",
		})
	})

	//GET request - just a quick ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Server Info - send back a bunch of info to display on the main page
	r.GET("/server-info", func(c *gin.Context) {
		const srcDir = "/app/sourcedir"
		dirExists := false
		definitionFilesCount := 0
		files, err := ioutil.ReadDir(srcDir)
		if err == nil {
			dirExists = true
			for _, file := range files {
				if strings.HasSuffix(file.Name(), "_definitions.h") {
					definitionFilesCount++
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"directory_exists":  dirExists,
			"definitions_files": definitionFilesCount,
		})
	})

	//Check file existence - will be used before sending any modifications
	r.GET("/check-file/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		const srcDir = "/app/sourcedir"
		exists := false

		files, err := ioutil.ReadDir(srcDir)
		if err == nil {
			for _, file := range files {
				if file.Name() == filename+"_definitions.h" {
					exists = true
					break
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"file_exists": exists,
		})
	})

	r.POST("/add-db-field", func(c *gin.Context) {
		var input ServerDbFieldAddition

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filePath := fmt.Sprintf("%s/%s_definitions.h", srcDir, input.DatabaseName)

		if input.Option == "REC" {
			addRecField(filePath,
				input.FieldName,
				input.FieldType,
				input.Comment)
		} else if input.Option == "MEM" {
			addMemField(filePath,
				input.FieldName,
				input.FieldType,
				input.Comment)
		} else if input.Option == "NONDB" {
			addNonDbField(filePath,
				input.FieldName,
				input.FieldType,
				input.Comment)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Processing complete.",
		})
	})

	r.GET("/dblist", func(c *gin.Context) {
		definitions, status := findDefinitionFiles()
		c.String(status, definitions)
	})

	/* LAST ONE: ALL ROUTES */
	r.GET("/routes", func(c *gin.Context) {
		routes := r.Routes()
		var routesInfo []map[string]string
		for _, route := range routes {
			routeInfo := map[string]string{
				"method": route.Method,
				"path":   route.Path,
			}
			routesInfo = append(routesInfo, routeInfo)
		}

		c.JSON(http.StatusOK, gin.H{
			"routes": routesInfo,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r.Run(":" + port) // listen and serve on the specified port
}
