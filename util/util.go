package util

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
	"os"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init (){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i:= 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner () string{
	return RandomString(6)
}

func RandomMoney () int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// varDump will print out any number of variables given to it
// e.g. varDump("test", 1234)
func VarDump(myVar ...interface{}) {
	fmt.Printf("%v\n", myVar)

}

// dd will print out variables given to it (like varDump()) but
// will also stop execution from continuing.
func dd(myVar ...interface{}) {
	VarDump(myVar...)
	os.Exit(1)
}