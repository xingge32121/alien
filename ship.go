/*
 * @Author: Wujiahuo
 * @Date: 2023-03-08 16:14:58
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-09 16:44:28
 * @FilePath: \alien\ship.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/bmp"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64
	y      float64
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, op)
}
func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./images/ship.bmp")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}
	return ship
}
