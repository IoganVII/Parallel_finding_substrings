package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {

	fmt.Println("Кол-во лог. ядер на машине:", runtime.NumCPU())

	var m sync.Mutex
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(5)

	start := time.Now()

	count := 0

	sliceUrl := []string{
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=4956bf29-3649-4f62-8587-1fac3a8d90cb",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=4956bf29-3649-4f62-8587-1fac3a8d90cb",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=4956bf29-3649-4f62-8587-1fac3a8d90cb",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=4956bf29-3649-4f62-8587-1fac3a8d90cb",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
		"https://gtc5awmle2xwu.elma365.ru/api/extensions/3954a56f-8569-40c6-8f51-54ef45731c89/script/getDataString?id=84367f05-50a6-4d86-bae5-d1a9fe7dbb67",
	}

	wg.Add(len(sliceUrl))

	for _, value := range sliceUrl {
		go searchSubstring(&count, value, &m, &wg)
	}

	wg.Wait()
	fmt.Println("Всего кол-во вхождений = ", count)

	fmt.Println(time.Now().Sub(start))
}

func searchSubstring(count *int, data string, m *sync.Mutex, wg *sync.WaitGroup) {

	resp, _ := http.Get(data)
	body, _ := io.ReadAll(resp.Body)
	countSubstring := strings.Count(string(body), "go")
	fmt.Println("Кол-во вхождуения для:", data, " = ", countSubstring)
	m.Lock()
	*count += countSubstring
	fmt.Println(*count)
	wg.Done()
	m.Unlock()
}
