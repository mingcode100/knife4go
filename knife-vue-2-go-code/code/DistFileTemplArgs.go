package code

import "encoding/json"

// 模板参数
type DistFileTemplArgs struct {
	PackageName     string `json:"PackageName"`
	Base64OrContent string `json:"Base64OrContent"`
	FileDir         string `json:"FileDir"`
	FileRelavePath  string `json:"FileRelavePath"`
	FileType        string `json:"FileType"`
	FileName        string `json:"FileName"`
	FileName2       string `json:"FileName2"`
	FileName3       string `json:"FileName3"`
}

func (t DistFileTemplArgs) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(data)
}
