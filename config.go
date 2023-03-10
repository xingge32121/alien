package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	BulletColor       color.RGBA `json:"bulletColor"`
	MaxBulletNum      int        `json:"maxBulletNum"`
}

func loadConfig() *Config {
	r, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}
	var cfg Config
	err = json.NewDecoder(r).Decode(&cfg)
	if err != nil {
		log.Fatalf("cannot NewDecoder  %v", err)
	}
	return &cfg
}
