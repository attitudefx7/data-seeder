package pkg

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"strconv"
	"time"
)


func RandomTimestamp() time.Time {
	// 1 month 30 * 24 * 60 * 60
	randomTime := rand.Int63n(time.Now().Unix() - 10*24*60*60) + 10*24*60*60

	randomNow := time.Unix(randomTime, 0)

	return randomNow
}

func GetIndex(base int) int {
	// index started with 0
	//rand.Seed(time.Now().UnixNano())
	//
	//return rand.Intn(base)
	return gofakeit.Number(0, base - 1)
}

func GenName(str string) string {
	return str + " " + gofakeit.Name()
}


func GenDate() string {
	timeObj := RandomTimestamp()
	return timeObj.Format("2006-01-02 15:04:05")
}

func GenDate2() string {
	//timeObj := RandomTimestamp()
	//return timeObj.Format("2006-01-02")

	return "2021-" + strconv.Itoa(gofakeit.Month()) + "-" + strconv.Itoa(gofakeit.Day())
}

func Rand() int {
	// 10000 - 99999
	return int(gofakeit.Uint8())
}

func RandFloat32() int {
	return int(gofakeit.Uint8() + 1)
}

func RandCampaignId() string {
	// 14
	return gofakeit.CreditCardNumber(nil)
}

func RandGroupId() string {
	// 14
	return gofakeit.CreditCardNumber(nil)
}

func RandProductId() string {
	// 14
	return gofakeit.CreditCardNumber(nil)
}

func RandTargetId() string {
	// 14
	return gofakeit.CreditCardNumber(nil)
}
