/*
 * @Author: Wujiahuo
 * @Date: 2023-03-14 15:57:59
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 16:03:54
 * @FilePath: \alien\alien.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Alien struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./images/alien.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	alien := &Alien{
		image: img,
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      0,
			y:      0,
		},
		speedFactor: cfg.AlienSpeedFactor,
	}
	return alien
}
func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) outOfScreen(cfg *Config) bool {

	return alien.y > float64(cfg.ScreenHeight)

}
