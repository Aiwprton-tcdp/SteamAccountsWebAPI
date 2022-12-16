package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// func (b example) PrintFields() string {
// 	val := reflect.ValueOf(b)
// 	for i := 0; i < val.Type().NumField(); i++ {
// 	   return val.Type().Field(i).Tag.Get("json")
// 	   break
// 	}
// }
