package helpers

import "encoding/json"

func DataParser[T1 any, T2 any] (src T1, dst T2) {
	bytData, err := json.Marshal(src)
	if err != nil {
		return
	}
	json.Unmarshal(bytData,dst)
}