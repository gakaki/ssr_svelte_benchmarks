package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(j interface{}, fileName string) {
	f, err := json.MarshalIndent(&j, "", " ")

	if err != nil {
		fmt.Println(err)
	} else {
		os.WriteFile(fmt.Sprintf("%s", fileName), f, 0777)
	}
}
func ReadJSON[T interface{}](jsonBytes []byte) (t T) {
	json.Unmarshal(jsonBytes, &t)
	return t
}
