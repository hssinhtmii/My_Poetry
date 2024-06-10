package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Poetry = []string{
	"دل، مست و مدهوش او بود و دگر مپنداشت که مقصد دیوانگیست",
	"درد من را درمانی، گویی چرا؟ هرکه درد باشد درمان نیز هست",
}

func main() {
	fmt.Println(ShowPoetry())
}

func ShowPoetry() string {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(len(Poetry))

	return Poetry[number]
}