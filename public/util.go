package public

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
)

var chars = [64]string{"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z"}

//GetUuid : generate uuid, length 32
func GetUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

//GetShortUuid : generate short uuid, length 8
func GetShortUuid() string {
	var strBuilder strings.Builder
	id := GetUuid()
	for i := 0; i < 8; i++ {
		str := id[i*4 : i*4+4]
		x, _ := strconv.ParseInt(str, 16, 32)
		strBuilder.WriteString(chars[x%0x3E]) //mod 62
	}
	return strBuilder.String()
}

//PrintIfErr : print err
func PrintIfErr(err error, service string) bool {
	if err != nil {
		err = errors.New(service + "--" + err.Error())
		log.Println(err)
		return true
	}
	return false
}

//PanicIfErr : throw err, usually use in response/controller
func PanicIfErr(err error, service string) {
	if err != nil {
		err = errors.New(service + "--" + err.Error())
		log.Println(err)
		panic(err)
	}
}
