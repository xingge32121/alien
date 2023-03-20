/*
 * @Author: Wujiahuo
 * @Date: 2023-03-10 14:12:17
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 15:16:53
 * @FilePath: \alien\game.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeOver
)

type Game struct {
	input     *Input
	cfg       *Config
	ship      *Ship
	bullets   map[*Bullet]struct{}
	aliens    map[*Alien]struct{}
	mode      Mode
	failCount int
	overMsg   string
}

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
func (g *Game) init() {
	g.createAliens()
	g.CreateFonts()
	g.failCount = 0
	g.overMsg = ""
}
func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	g := &Game{
		input:   &Input{},
		ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		bullets: make(map[*Bullet]struct{}),
		cfg:     cfg,
		aliens:  make(map[*Alien]struct{}),
	}
	g.init()
	return g
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.input.IsKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		for bullet := range g.bullets {
			bullet.y -= bullet.speedFactor

		}
		for alien := range g.aliens {
			alien.y += alien.speedFactor
		}
		g.input.Update(g)
		g.CheckCollision()
		for bullet := range g.bullets {
			if bullet.outOfScreen() {
				delete(g.bullets, bullet)
			}
		}
		for alien := range g.aliens {
			if alien.outOfScreen(g.cfg) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}
			if CheckCollision(alien, g.ship) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}
		}
		if g.failCount >= 3 {
			g.overMsg = "Game Over!"
		} else if len(g.aliens) == 0 {
			g.overMsg = "You win!"
		}
		if len(g.overMsg) > 0 {
			g.mode = ModeOver
			g.aliens = make(map[*Alien]struct{})
			g.bullets = make(map[*Bullet]struct{})
		}
	case ModeOver:
		if g.input.IsKeyPressed() {
			g.init()
			g.mode = ModeTitle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	var titleTexts []string
	var titles []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		titles = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		g.ship.Draw(screen)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	case ModeOver:
		titles = []string{"", g.overMsg}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*int(g.cfg.TitleFontSize)) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*int(g.cfg.TitleFontSize), color.White)

	}
	for i, l := range titles {
		x := (g.cfg.ScreenWidth - len(l)*int(g.cfg.FontSize)) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*int(g.cfg.FontSize), color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func (g *Game) createAliens() {
	alien := NewAlien(g.cfg)

	availableSpaceX := g.cfg.ScreenWidth - 2*alien.width
	numAliens := availableSpaceX / (2 * alien.width)
	for j := 0; j < 2; j++ {
		for i := 0; i < numAliens; i++ {
			alien := NewAlien(g.cfg)
			alien.x = float64(alien.width + 2*alien.width*i)
			alien.y = float64(alien.height*j) * 1.5
			g.addAlien(alien)
		}
	}
}
func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(alien, bullet) {
				delete(g.bullets, bullet)
				delete(g.aliens, alien)
			}
		}
	}
}
