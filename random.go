package main

import (
	"math/rand"
	"time"
)

func RandomItem(gifs []GIF) GIF {
	rand.Seed(time.Now().Unix())
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return gifs[r.Intn(len(gifs))]
}
