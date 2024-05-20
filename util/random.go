package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets="abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64{
	r := max-min+1
	return min + rand.Int63n(r)
}

func RandomString(n int) string {
	var sb strings.Builder
	k :=len(alphabets)

	for i:=0; i<n; i++ {
		c:=alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}