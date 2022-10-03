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

	totalCount := 0
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
	// Канал с буфером на 5 элементов. Чтобы ограничить создание гоРутин
	chanel := make(chan int, 5)
	// Группа ождиания, чтобы главная гоРутина ожидала выполнения остальных
	wg := sync.WaitGroup{}
	// Мьютекс для синхронизации работы с критической секцией
	m := sync.Mutex{}
	start := time.Now()

	// Начинаем перебирать слайс с URL
	for _, value := range sliceUrl {
		wg.Add(1)
		// Занести данные в канал. Когда буфер будет заполнен - тут встанем на ожидание момента, когда в буфере освободится место
		chanel <- 1
		// Вызов фунции для получения данных по URL и подсчёта кол-ва вхождения подстроки
		go searchSubstring(&totalCount, value, chanel, &wg, &m)
	}
	/*
		Тут оставим цикл в одном потоке, чтобы сравнить - насколько мы быстры наши потоки
			for _, value := range sliceUrl {
				// Получаем Данные по URL
				resp, _ := http.Get(value)
				// Получаем байтМассив тела ответа
				body, _ := io.ReadAll(resp.Body)
				// Получаем кол-во вхождений
				countSubstring := strings.Count(string(body), "go")
				totalCount += countSubstring
			}*/

	wg.Wait()

	fmt.Println("Всего кол-во вхождений = ", totalCount)
	fmt.Println(time.Now().Sub(start))
}

func searchSubstring(totalCount *int, data string, chanel chan int, wg *sync.WaitGroup, m *sync.Mutex) {

	// Вызываем функцию после окончания работы родительской функции
	defer func() {
		// Считываем данные из буфера
		<-chanel
		// Сигналим о том, что задача завершена
		wg.Done()
	}()

	// Получаем Данные по URL
	resp, er := http.Get(data)
	if er != nil {
		panic("Ошибка при вызове АПИ по url: " + data)
	}
	// Получаем байтМассив тела ответа
	body, er := io.ReadAll(resp.Body)
	if er != nil {
		panic("Ошибка при вызове АПИ по url: " + data)
	}
	// Получаем кол-во вхождений
	countSubstring := strings.Count(string(body), "go")
	fmt.Println("Кол-во Горутин:", runtime.NumGoroutine())
	fmt.Println("Кол-во вхождуения для:", data, " = ", countSubstring)

	// Блокируем работу с критической секцией
	m.Lock()
	*totalCount += countSubstring
	fmt.Println(*totalCount)
	m.Unlock()

}
