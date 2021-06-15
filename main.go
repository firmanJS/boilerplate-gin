package main

import (
	"log"
	"strconv"

	helmet "github.com/danielkov/gin-helmet"
	config "github.com/firmanJS/boilerplate-gin/config"
	route "github.com/firmanJS/boilerplate-gin/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	  @description Setup Server
	*/
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	/**
	  @description Setup Server
	*/
	router := SetupRouter()
	/**
	  @description Run Server
	*/
	log.Fatal(router.Run(":" + strconv.Itoa(config.GO_PORT)))
}

func SetupRouter() *gin.Engine {
	/*
	    @description Setup Database Connection
	*/
	db := config.Connection()
	
    /*
	    @description Init Router
	*/
	router := gin.Default()
	
    /*
	    @description Setup Mode Application
	*/
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if config.GO_ENV != "production" && config.GO_ENV != "test" {
		gin.SetMode(gin.DebugMode)
	} else if config.GO_ENV == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	
    /*
	    @description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	
    /**
	    @description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	route.InitRoute(router)

	return router
}
