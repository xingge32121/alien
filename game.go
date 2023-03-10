package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input   *Input
	cfg     *Config
	ship    *Ship
	bullets map[*Bullet]struct{}
}

func (g *Game) Update() error {
	for bullet := range g.bullets {
		bullet.y -= bullet.speedFactor
	}
	g.input.Update(g, g.ship, g.cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}
func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		input:   &Input{},
		ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
	}
}
