/*
 * @Author: Wujiahuo
 * @Date: 2023-03-08 11:08:01
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 16:31:03
 * @FilePath: \alien\main.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

//go:generate go install github.com/hajimehoshi/file2byteslice
//go:generate file2byteslice -input ../images/ship.png -output resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input ../images/alien.png -output resources/alien.go -package resources -var AlienPng
//go:generate file2byteslice -input config.json -output resources/config.go -package resources -var ConfigJso
