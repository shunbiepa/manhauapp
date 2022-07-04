package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"go.uber.org/zap"
	"io/ioutil"
	"pangolin/global"
)

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer func() {
		if err = reader.Close(); err != nil {
			global.GVA_LOG.Info("Gzip解压失败", zap.Error(err))
		}
	}()
	return ioutil.ReadAll(reader)
}

// GzipEncode gzip加密
func GzipEncode(str string) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return b.Bytes()
}
