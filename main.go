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

	totalCount := 0
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
	// Канал с буфером на 5 элементов. Чтобы ограничить создание гоРутин
	chanel := make(chan int, 5)
	// Группа ождиания, чтобы главная гоРутина ожидала выполнения остальных
	wg := sync.WaitGroup{}
	// Мьютекс для синхронизации работы с критической секцией
	m := sync.Mutex{}

	// Хочу посмотреть, за сколько выполняется логика программы
	timeStart := time.Now()

	// Начинаем перебирать слайс с URL
	for _, url := range sliceUrl {
		wg.Add(1)
		// Занести данные в канал. Когда буфер будет заполнен - тут встанем на ожидание момента, когда в буфере освободится место
		chanel <- 1
		// Вызов фунции для получения данных по URL и подсчёта кол-ва вхождения подстроки
		go searchSubstring(&totalCount, url, chanel, &wg, &m)
		// Посмотрим - правда ли у нас тут 5 гоРутин работает
		//fmt.Println("Кол-во гоРутин: ", len(chanel))
	}

	wg.Wait()

	fmt.Println("Всего кол-во вхождений = ", totalCount)

	fmt.Println("Время выполнения работы: ", time.Now().Sub(timeStart))
}

// Функция для подсчёта кол-во подстрок go в строке
func searchSubstring(totalCount *int, url string, chanel chan int, wg *sync.WaitGroup, m *sync.Mutex) {

	// Вызываем функцию после окончания работы родительской функции
	defer func() {
		// Считываем данные из буфера
		<-chanel
		// Сигналим о том, что задача завершена
		wg.Done()
	}()

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
	fmt.Println("Кол-во вхождуения для:", url, " = ", countSubstring)

	// Блокируем работу с критической секцией. всё-таки с одним элементом 5 потоков работают. Хочу синхронизацию
	m.Lock()
	*totalCount += countSubstring
	// Пологировать можем, корректная ли работа с критической секцией
	//fmt.Println(*totalCount)
	m.Unlock()

}
