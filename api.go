package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func startApi() {
	db, err := gorm.Open(sqlite.Open("sqlite/racemanager.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	//CORS
	r.Use(cors.Default())

	// Migrate Schema
	// db.AutoMigrate(&Driver{}, &Event{}, &EventResult{}, &Division{}, &Mod{})

	// Create Sample Data
	// db.Create(&Division{Name: "Sportsman"})
	// db.Create(&Division{Name: "Late Model"})
	//
	// db.Create(&Mod{Name: "aero88_cts"})
	// db.Create(&Mod{Name: "lmpv2"})
	//
	// db.Create(&Event{Date: 1698883200, Name: "Snowbird Series: All American 120", Track: "Nashville WKC", DivisionID: 1, ModID: 1})
	// db.Create(&Event{Date: 1702519200, Name: "Snowbird Series: North Wilksboro", Track: "North Wilkesboro PST '94", DivisionID: 1, ModID: 1})
	// db.Create(&Event{Date: 1704934800, Name: "Snowbird Series: Myrtle Beach", Track: "Myrtle Beach", DivisionID: 1, ModID: 1})
	// db.Create(&Event{Date: 1709172000, Name: "Snowbird Series: New Smyrna", Track: "New Smyrna 2023 PST (Night)", DivisionID: 1, ModID: 1})
	// db.Create(&Event{Date: 1711324800, Name: "Snowbird Series: Bowman Gray", Track: "Madhouse", DivisionID: 1, ModID: 1})
	// db.Create(&Event{Date: 1712797200, Name: "Snowbird Series: Kingsport", Track: "Kingsport 2021 PST", DivisionID: 1, ModID: 1})
	//
	// db.Create(&EventResult{EventID: 1, DriverID: 1, Heat1: 1, Heat2: 2, Feature: 3})
	// db.Create(&EventResult{EventID: 1, DriverID: 2, Heat1: 2, Heat2: 3, Feature: 1})
	// db.Create(&EventResult{EventID: 1, DriverID: 3, Heat1: 3, Heat2: 1, Feature: 2})
	//
	// db.Create(&Driver{LastName: "McQueen", FirstName: "Lightning", CarNumber: "95", DivisionID: 1})
	// db.Create(&Driver{LastName: "Ramirez", FirstName: "Cruz", CarNumber: "51", DivisionID: 1})
	// db.Create(&Driver{LastName: "Weathers", FirstName: "Strip", CarNumber: "43", DivisionID: 1})
	// db.Create(&Driver{LastName: "Hicks", FirstName: "Chick", CarNumber: "86", DivisionID: 1})

	r.GET("/divisions", func(c *gin.Context) {
		var divisions []Division
		db.Preload(clause.Associations).Find(&divisions)
		c.IndentedJSON(http.StatusOK, divisions)
	})

	r.POST("/division", func(c *gin.Context) {
		var division Division

		if err := c.BindJSON(&division); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Create(&division); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&division)
		c.IndentedJSON(http.StatusCreated, division)
	})

	r.GET("/drivers", func(c *gin.Context) {
		var drivers []Driver
		db.Preload(clause.Associations).Find(&drivers)
		c.IndentedJSON(http.StatusOK, drivers)
	})

	r.GET("/driver/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var result Driver

		db.Preload(clause.Associations).First(&result, id)
		c.IndentedJSON(http.StatusOK, result)
	})

	r.POST("/driver", func(c *gin.Context) {
		var driver Driver

		if err := c.BindJSON(&driver); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Create(&driver); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&driver)
		c.IndentedJSON(http.StatusCreated, driver)
	})

	r.PUT("/driver/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var driver Driver

		if err := db.Where("id = ?", id).First(&driver); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}

		if err := c.BindJSON(&driver); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Save(&driver); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&driver)
		c.IndentedJSON(http.StatusCreated, driver)
	})

	r.GET("/events", func(c *gin.Context) {
		var events []Event
		db.Preload(clause.Associations).Find(&events)
		c.IndentedJSON(http.StatusOK, events)
	})

	r.GET("/event/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var event Event

		db.Preload(clause.Associations).First(&event, id)
		c.IndentedJSON(http.StatusOK, event)
	})

	r.POST("/event", func(c *gin.Context) {
		var event Event

		if err := c.BindJSON(&event); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Create(&event); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&event)
		c.IndentedJSON(http.StatusCreated, event)
	})

	r.PUT("/event/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var result Event

		if err := db.Where("id = ?", id).First(&result); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}

		if err := c.BindJSON(&result); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Save(&result); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&result)
		c.IndentedJSON(http.StatusCreated, result)
	})

	r.GET("/points", func(c *gin.Context) {
		var results []EventResult
		var points []DriverEventPoints
		var entries int64
		db.Preload(clause.Associations).Preload("Event.Division").Preload("Event.Mod").Find(&results)
		for i := 0; i < len(results); i++ {
			db.Model(&EventResult{}).Where("event_id = ?", results[i].EventID).Count(&entries)
			h1p := entries - (results[i].Heat1 - 1)
			h2p := entries - (results[i].Heat2 - 1)
			fp := (entries - (results[i].Feature - 1)) * 2
			p := h1p + h2p + fp
			r := DriverEventPoints{ID: uint(i), DriverID: results[i].DriverID, Driver: results[i].Driver, EventID: results[i].EventID, Event: results[i].Event, Points: p}
			points = append(points, r)
		}
		c.IndentedJSON(http.StatusOK, points)
	})

	r.GET("/results", func(c *gin.Context) {
		var results []EventResult
		db.Preload(clause.Associations).Preload("Event.Division").Preload("Event.Mod").Find(&results)
		c.IndentedJSON(http.StatusOK, results)
	})

	r.GET("/result/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var result EventResult

		db.Preload(clause.Associations).Preload("Event.Division").Preload("Event.Mod").First(&result, id)
		c.IndentedJSON(http.StatusOK, result)
	})

	r.POST("/result", func(c *gin.Context) {
		var result EventResult

		if err := c.BindJSON(&result); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Create(&result); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&result)
		c.IndentedJSON(http.StatusCreated, result)
	})

	r.PUT("/result/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var result EventResult

		if err := db.Where("id = ?", id).First(&result); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}

		if err := c.BindJSON(&result); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Save(&result); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&result)
		c.IndentedJSON(http.StatusCreated, result)
	})

	r.GET("/mods", func(c *gin.Context) {
		var mods []Mod
		db.Preload(clause.Associations).Find(&mods)
		c.IndentedJSON(http.StatusOK, mods)
	})

	r.POST("/mod", func(c *gin.Context) {
		var mod Mod

		if err := c.BindJSON(&mod); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := db.Create(&mod); err.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error})
			return
		}
		db.Preload(clause.Associations).Find(&mod)
		c.IndentedJSON(http.StatusCreated, mod)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.Run(":9091")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
