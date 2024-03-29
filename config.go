/*
 * @Author: Wujiahuo
 * @Date: 2023-03-10 14:12:17
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-16 14:19:59
 * @FilePath: \alien\config.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
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
	BulletInterval    int64      `json:"bulletInterval"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
	TitleFontSize     float64    `json:"titleFontSize"`
	FontSize          float64    `json:"fontSize"`
	SmallFontSize     float64    `json:"smallFontSize"`
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
