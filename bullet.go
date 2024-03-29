/*
 * @Author: Wujiahuo
 * @Date: 2023-03-10 10:31:01
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 15:22:14
 * @FilePath: \alien\bullet.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)
	return &Bullet{
		image: img,
		GameObject: GameObject{
			width:  cfg.BulletWidth,
			height: cfg.BulletHeight,
			x:      ship.x + float64(ship.width-cfg.BulletWidth)/2,
			y:      float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight),
		},
		speedFactor: cfg.BulletSpeedFactor,
	}
}
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) outOfScreen() bool {

	return bullet.y < -float64(bullet.height)

}
