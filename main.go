/*
 * @Author: Wujiahuo
 * @Date: 2023-03-08 11:08:01
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-09 15:16:37
 * @FilePath: \alien\main.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// if err := ebiten.RunGame(&Game{}); err != nil {
	// 	log.Fatal(err)
	// }
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
