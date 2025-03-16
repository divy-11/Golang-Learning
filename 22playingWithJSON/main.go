package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //name changes in json data type
	Price    int
	Website  string   `json:"platform"`
	Password string   `json:"-"`              // - means i dont want this field to be shown
	Tags     []string `json:"tags,omitempty"` //omitempty means if nul doesnt show, ALSO NO SPACE BTW COMMA AND OMITEMPTY
}

func main() {
	// EncodingJSON()
	DecodeJSON()
}

func EncodingJSON() {
	myCourses := []course{
		{"Cohort 2.0", 3999, "udemy.com", "hCyaK18#", []string{"webd", "devops", "js", "ts"}},
		{"JS Completed", 999, "udemy.com", "hYshK17#", nil},
		{"Frontend Advanced", 2999, "udemy.com", "jKbud06#", []string{"reactjs", "threejs", "design", "animation"}},
	}

	//package above data as JSON data

	// data, err := json.Marshal(myCourses)
	data, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", data)
	fmt.Println(string(data))
}

func DecodeJSON() {
	jsonDataInput := []byte(`
		{
                "coursename": "Cohort 2.0",
                "Price": 3999,
                "platform": "udemy.com",
                "tags": ["webd","devops","js","ts"]
        }
	`)

	var courseCheck course
	checkValid := json.Valid(jsonDataInput)
	if checkValid {
		fmt.Println("Valid JSON.")
		json.Unmarshal(jsonDataInput, &courseCheck)
		// fmt.Println(courseCheck)
		fmt.Printf("%#v\n", courseCheck)
	} else {
		fmt.Println("Invalid JSON data.")
	}

	//Special Case: where we want to add data to key value pair
	var myconvertedData map[string]interface{} //for online data key will be string but value is not fixes hence interface
	json.Unmarshal(jsonDataInput, &myconvertedData)
	fmt.Printf("%#v\n", myconvertedData)

	for key, val := range myconvertedData {
		fmt.Printf("Key is %v and Value is %v || Type: %T \n", key, val, val)
	}
}
