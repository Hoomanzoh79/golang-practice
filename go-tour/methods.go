package main

import ("math")

type MyVertex struct{
	X float64 ;
	Y float64
}

func (v MyVertex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}