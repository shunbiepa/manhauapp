package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"pangolin/global"
	"pangolin/model/params"
	"time"
)

// GetRegeo 根据经纬度获取区域信息
func GetRegeo(longitude, latitude string) (address params.AddressComponent, err error) {

	location := longitude + "," + latitude

	httpGet := fmt.Sprintf(global.GVA_CONFIG.GD.URL, location, global.GVA_CONFIG.GD.Key)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", httpGet, nil)

	if err != nil {
		global.GVA_LOG.Info("GetRegeo err", zap.Any("err", err))
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Info("GetRegeo err2", zap.Any("err", err))
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Info("GetRegeo err3", zap.Any("err", err))
		return
	}

	mas := new(params.Regeo)
	json.Unmarshal(data, mas)

	if mas.Status != "1" {
		global.GVA_LOG.Info("GetRegeo err4", zap.Any("info", mas.Info))
		return
	}
	address = mas.RegeoCode.AddressComponent
	return
}
