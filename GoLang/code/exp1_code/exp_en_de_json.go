package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	if err := json.Unmarshal(data, &result); err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Printf("%T\n", result["status"])	// float64
//	var status = result["status"].(int)	// 类型断言错误
//	fmt.Println("Status value: ", status)
//}

//// 将 decode 的值转为 int 使用
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	if err := json.Unmarshal(data, &result); err != nil {
//		log.Fatalln(err)
//	}
//
//	var status = uint64(result["status"].(float64))
//	fmt.Println("Status value: ", status)
//}

//// 指定字段类型
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	var decoder = json.NewDecoder(bytes.NewReader(data))
//	decoder.UseNumber()
//
//	if err := decoder.Decode(&result); err != nil {
//		log.Fatalln(err)
//	}
//
//	var status, _ = result["status"].(json.Number).Int64()
//	fmt.Println("Status value: ", status)
//}

//// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
//// 将数据转为 decode 为 string
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//	var decoder = json.NewDecoder(bytes.NewReader(data))
//	decoder.UseNumber()
//	if err := decoder.Decode(&result); err != nil {
//		log.Fatalln(err)
//	}
//	var status uint64
//	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status);
//	checkError(err)
//	fmt.Println("Status value: ", status)
//}

// struct 中指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	//checkError(err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %+v", result)
}