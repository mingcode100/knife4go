package templ

import "encoding/json"

// 模板参数
type KnifeArgs struct {
	KnifeImport map[string]string `json:"KnifeImport"`
	KnifeLine   []string            `json:"KnifeLine"`
}

func (t KnifeArgs) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(data)
}
