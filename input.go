/*
 * @Author: Wujiahuo
 * @Date: 2023-03-08 14:38:14
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-10 11:43:31
 * @FilePath: \alien\input.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
}

func (i *Input) Update(g *Game, ship *Ship, cfg *Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ship.x -= cfg.ShipSpeedFactor
		if ship.x < -float64(ship.width)/2 {
			ship.x = -float64(ship.width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ship.x += cfg.ShipSpeedFactor
		if ship.x > float64(cfg.ScreenWidth)-float64(ship.width)/2 {
			ship.x = float64(cfg.ScreenWidth) - float64(ship.width)/2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if len(g.bullets) < cfg.MaxBulletNum {
			Bullet := NewBullet(g.cfg, g.ship)
			g.addBullet(Bullet)
		}
	}
}
