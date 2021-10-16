package code

import "encoding/json"

// 模板参数
type KnifeLine struct {
	FileName3    string `json:"FileName3"`
	PackageAlian string `json:"PackageAlian"`
}

type KnifeArgs struct {
	KnifeImport map[string]string `json:"KnifeImport"` // 别名 ->  包名
	Lines       []KnifeLine       `json:"KnifeLine"`
}

func (t KnifeArgs) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(data)
}
