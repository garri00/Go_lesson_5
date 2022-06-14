// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

package main

import "fmt"

func main() {

	var (
		apple_price float32 = 5.99
		pear_price  float32 = 7
		money       float32 = 23
	)

	fmt.Println("Грощей в кишені ", money)
	fmt.Println("Ціна яблука ", apple_price)
	fmt.Println("Ціна груші ", pear_price)

	var money_to_buy_9apples_8pears float32 = apple_price*8 + pear_price*9

	var _2apple_price float32 = apple_price * 2
	var _2pear_price float32 = pear_price * 2

	var how_many_apples_can_we_buy = money / apple_price

	var how_many_pears_can_we_buy = money / pear_price

	var can_we_buy_2aaples_2pears bool = money >= _2apple_price+_2pear_price

	fmt.Println("1. Скільки грошей треба витратити щоб купити 9 яблук та 8 груш = ", money_to_buy_9apples_8pears)
	fmt.Println("2. Скільки груш ми можемо купити ", int(how_many_apples_can_we_buy)) // Округлюємо значення int ()
	fmt.Println("1. Скільки яблук ми можемо купити ", int(how_many_pears_can_we_buy)) // Округлюємо значення
	fmt.Println("1. Чи ми можемо купити 2 яблука та 2 груші ", can_we_buy_2aaples_2pears)
}
