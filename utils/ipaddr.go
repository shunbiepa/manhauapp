package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"pangolin/global"
	"pangolin/model/params"
	"strings"
)

// CheckUserRegion 校验用户登陆IP是否归属地ip
func CheckUserRegion(c *gin.Context, region string) bool {

	loginIP := ClientIP(c.Request)
	if loginIP == "" {
		return true
	}

	ipAddr, err := IPAddr(loginIP)
	if err != nil {
		global.GVA_LOG.Error("获取ip归属地失败!", zap.Any("err", err))
		return true
	}

	global.GVA_LOG.Info("ip addr is ", zap.Any("addr", ipAddr.Data))
	if ipAddr.Data.City != "" && strings.Contains(region, ipAddr.Data.City) {
		return true
	}
	return false
}

// IPAddr 获取ip地址归属地
func IPAddr(ip string) (mas params.IPAddr, err error) {
	data := url.Values{}
	data.Add("ip", ip)

	url := "https://ipaddquery.market.alicloudapi.com/ip/address-query"
	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	appcode := fmt.Sprintf("APPCODE %s", global.GVA_CONFIG.AliAppCode)
	request.Header.Set("Authorization", appcode)
	request.Header.Set("content-type", `application/x-www-form-urlencoded`)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	statusCode := resp.StatusCode
	if statusCode != 200 {
		err = errors.New("获取ip地址失败")
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	json.Unmarshal(respBytes, &mas)

	if !mas.Success {
		err = errors.New("ip地址无法识别")
		return
	}
	return
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
