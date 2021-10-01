package templ

import "encoding/json"

// 模板参数
type DistFileTemplArgs struct {
	PackageName    string `json:"PackageName"`
	FileBase64     string `json:"FileBase64"`
	FileRelavePath string `json:"FileRelavePath"`
	FileType       string `json:"FileType"`
	FileName       string `json:"FileName"`
	FileName2      string `json:"FileName2"`
}

func (t DistFileTemplArgs) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(data)
}
