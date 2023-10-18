package utils

// 配置信息
type ConfigInfo struct {
	AccessKey string `json:"accessKey"`
	Cookie    string `json:"cookie"`
	RoomId    int    `json:"roomId"`
}

// bilibili api 普遍返回数据格式
type ApiResponseCommon struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
