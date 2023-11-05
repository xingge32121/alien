/*
 * @Author: Wujiahuo
 * @Date: 2023-03-08 16:14:58
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 15:22:53
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
//飞船
type Ship struct {
	GameObject
	image *ebiten.Image
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./images/ship.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image: img,
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      float64(screenWidth-width) / 2,
			y:      float64(screenHeight - height),
		},
	}
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}
