/*
 * @Author: Wujiahuo
 * @Date: 2023-03-14 16:49:59
 * @LastEditors: OBKoro1
 * @LastEditTime: 2023-03-20 15:25:26
 * @FilePath: \alien\collision.go
 * @Description:
 * Copyright (c) 2023 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

func CheckCollision(entityA, entityB Entity) bool {
	//计算外星人的上下左右四角的坐标
	alienTop, alienLeft := entityA.Y(), entityA.X()
	alienBottom, alienRight := entityA.Y()+float64(entityA.Height()), entityA.X()+float64(entityA.Width())
	// //子弹左上角
	x, y := entityB.X(), entityB.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	//子弹右上角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	//子弹左小角
	x, y = entityB.X(), entityB.Y()+float64(entityB.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	//子弹右小角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()+float64(entityB.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	return false
}
