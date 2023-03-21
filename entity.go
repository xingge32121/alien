/*
 * @Author: Wujiahuo
 * @Date: 2023-03-20 14:59:22
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 15:00:27
 * @FilePath: \alien\entity.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}
