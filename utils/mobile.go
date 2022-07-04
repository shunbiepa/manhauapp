package utils

import (
	"math/rand"
	"pangolin/global"
	"pangolin/model/table"
	"strconv"
	"time"
)

var telStarts = []int{133, 149, 153, 173, 177, 180, 181, 189, 199, 130, 131, 132, 145, 155, 156, 166, 171, 175, 176, 185, 186, 166, 134, 135, 136, 137, 138, 139, 147, 150, 151, 152, 157, 158, 159, 172, 178, 182, 183, 184, 187, 188, 198, 170}

// GetRandomTel 生成随机手机号码
func GetRandomTel() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(telStarts) - 1)

	first := strconv.Itoa(telStarts[index])

	second := strconv.Itoa(rand.Intn(788) + 10100)
	second = second[1:]

	thrid := strconv.Itoa(rand.Intn(9100) + 10001)
	thrid = thrid[1:]

	return first + second + thrid
}

// GetActiveMobile 获取可用的手机号码
func GetActiveMobile() string {
	mobile := GetRandomTel()
	var user table.User
	if global.GVA_DB.Where("mobile = ?", mobile).First(&user).RecordNotFound() {
		return mobile
	}
	return GetActiveMobile()
}
