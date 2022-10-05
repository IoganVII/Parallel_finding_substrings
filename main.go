package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	// Урлы ведут на Стенд Elma365, там АПИ модуля. Апи возвращает строку.
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

	// Канал, куда будут слоены урлы. Канал необходим воркеру для чтения
	chanelWithUrl := make(chan string)
	// Канал, куда будут складываться кол-во вхождений с каждой обработки url. Он будет буферизированный, чтоьы можно было складывать значения, а потом считать их.
	// Сколькими разными способами можно эту задачи решить...иех
	chanelResult := make(chan int, len(sliceUrl))

	// Нужная, чтобы дожаться, пока Вворкеры закончат обрабатывать chanelWithUrl...Это очень грустно
	wg := sync.WaitGroup{}

	// Хочу посмотреть, за сколько выполняется логика программы
	timeStart := time.Now()

	// Запускаем воркеры - прогреть движки!
	for i := 0; i < 5 && i <= len(sliceUrl); i++ {
		go workerForSearchSubstringGo(i, chanelWithUrl, chanelResult, &wg)
	}

	// Начинаем перебирать слайс с URL
	for _, url := range sliceUrl {
		wg.Add(1)
		chanelWithUrl <- url
	}
	wg.Wait()
	close(chanelWithUrl)

	totalCount := 0
	for i := len(chanelResult); i > 0; i-- {
		totalCount += <-chanelResult
	}

	fmt.Println("Всего кол-во вхождений = ", totalCount)

	fmt.Println("Время выполнения работы: ", time.Now().Sub(timeStart))
}

// Функция для подсчёта кол-во подстрок go в строке
func workerForSearchSubstringGo(id int, chanelWithUrl <-chan string, chanelResult chan<- int, wg *sync.WaitGroup) {
	// Начинаем брать значение из канала, пока они там есть и канал открыт
	for url := range chanelWithUrl {

		// Получаем Данные по URL
		resp, er := http.Get(url)
		if er != nil {
			panic("Ошибка при вызове АПИ по url: " + url)
		}
		// Получаем байтМассив тела ответа
		responseBody, er := io.ReadAll(resp.Body)
		if er != nil {
			panic("Ошибка при вызове АПИ по url: " + url)
		}
		// Получаем кол-во вхождений
		countSubstring := strings.Count(string(responseBody), "go")
		fmt.Println(id, ": Кол-во вхождуения для:", url, " = ", countSubstring)
		// Записываем результат в канал для...Результатов
		chanelResult <- countSubstring
		wg.Done()
	}
}
