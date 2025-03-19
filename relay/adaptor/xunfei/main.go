package xunfei

import (
	"errors"
	"strings"
)

// https://console.xfyun.cn/services/cbm
// https://www.xfyun.cn/doc/spark/HTTP%E8%B0%83%E7%94%A8%E6%96%87%E6%A1%A3.html

func getXunfeiDomain(modelName string) (string, error) {
	// 先检查是否直接匹配域名
	switch strings.ToLower(modelName) {
	case "lite", "generalv3", "pro-128k", "generalv3.5", "max-32k", "4.0ultra":
		return modelName, nil
	}

	// 如果不是直接域名，则尝试解析 spark-xxx 格式
	_, s, ok := strings.Cut(modelName, "-")
	if !ok {
		return "", errors.New("invalid model name")
	}
	switch strings.ToLower(s) {
	case "lite":
		return "lite", nil
	case "pro":
		return "generalv3", nil
	case "pro-128k":
		return "pro-128k", nil
	case "max":
		return "generalv3.5", nil
	case "max-32k":
		return "max-32k", nil
	case "4.0-ultra":
		return "4.0Ultra", nil
	}
	return "", errors.New("invalid model name")
}
