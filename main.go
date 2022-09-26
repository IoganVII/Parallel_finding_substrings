package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	count := 0

	sliceString := make([]string, 0, 4)
	sliceString = append(sliceString, "https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=4956bf29-3649-4f62-8587-1fac3a8d90cb")
	sliceString = append(sliceString, "https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67")

	for _, value := range sliceString {
		searchSubstring(&count, MakeRequest(value))
	}

	fmt.Println("Количество вхождения:", count)

}

func searchSubstring(count *int, data string) {
	*count += strings.Count(data, "go")
}

func MakeRequest(url string) string {
	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
