package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func test1() {
	type Movie struct {
		Title  string
		Year   uint32   `json:"released"`
		Color  bool     `json:"color,omitempty"` //color：标签 omitempty：零值时不输出成员到json中
		Actors []string //如果是小写，则json字符串中不会出现，只有可导出成员可以装换为json字段
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: true, Actors: []string{"Berg"}},
		{Title: "Basablanca", Year: 1928, Color: false, Actors: []string{"Gseg"}},
	}

	fmt.Println(movies)

	// if data, err := json.Marshal(movies); err != nil {
	// 	log.Fatalf("Json marshaling failed: %s", err)
	// }
	// fmt.Printf("%s\n", data) //data 访问不到

	if data, err := json.Marshal(movies); err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	//MarshalIndent格式化输出
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("Json masshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(data)

	//可以选择性解码部分成员
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Json unmashaling failed: %s", err)
	}
	fmt.Println(titles)
}

func test2() {
	jsonStr1 := `{"id":10001,"Name":"tony","group_id":10}`

	type Student struct {
		Id        uint32
		Name      string
		Group_Id  uint32 `json:group_id`
		GroupName string
	}

	var stu Student
	if err := json.Unmarshal([]byte(jsonStr1), &stu); err != nil {
		log.Fatalf("Json Unmarshaling failed: %s", err)
	}
	fmt.Println(stu)
	fmt.Println(stu.Group_Id)

	jsonStr2 := `[{"id":10001,"Name":"tony","group_id":10}]`
	var stus []Student
	if err := json.Unmarshal([]byte(jsonStr2), &stus); err != nil {
		log.Fatalf("Json unmarshaling failed: %s", err)
	}
	fmt.Println(stus)
}

func main() {
	test1()
	test2()
}
