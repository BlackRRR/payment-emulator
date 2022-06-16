# payment-emulator

### Реализован эмулятор платежного сервиса.

## Database

База данных выбрана **Postgresql**, поля в базе данных:

- ID транзакции (число) `transaction_id`
- UUID транзакции (было добавлено для реализации авторизации п.2 из *Rest API*) `transaction_hash`
- ID пользователя (число) `user_id`
- email пользователя `email`
- сумма `amount`
- валюта `currency`
- дата и время создания `date_of_creation`
- дата и время последнего изменения `date_of_last_change`
- статус `status`

## Rest API

Сервис принимает запросы через через REST API.
Ниже описаны действия API и указан метод http запроса:

1. Создание платежа. `HTTP METHOD POST`
2. Изменение статуса платежа платежной системой. (*Авторизация*) `HTTP METHOD PUT`
3. Проверка статуса платежа по ID. `HTTP METHOD GET`
4. Получение списка всех платежей пользователя по его ID. `HTTP METHOD GET`
5. Получение списка всех платежей пользователя по его e-mail. `HTTP METHOD GET`
6. Отмена платежа по его ID. `HTTP METHOD DELETE`

Каждый запрос возвращает `Result`, `Payload`, `Error`

Где `Result` это значения `200`, `400` тоесть OK и ERROR

`Payload` это какие то значения, например `transaction_hash` или `status`

Авторизация выполнена в виде передачи ключа. При создании платежа создается уникальный ключ транзакции (п.1), который
возвращается пользователю и для того чтобы изменить статус транзакции, пользователь должен передать ключ для
идентификации.
