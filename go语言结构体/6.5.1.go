package main

import "math"

// 二维矢量模拟玩家移动
type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X+other.X,
		Y: v.Y+other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X-other.X,
		Y: v.Y-other.Y,
	}
}

func (v Vec2) Scale(other float32) Vec2 {
	return Vec2{
		X: v.X*other,
		Y: v.Y*other,
	}
}

func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := other.X - v.X
	dy := other.Y - v.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag >0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y* oneOverMag}
	} else {
		return Vec2{0,0}
	}
}


