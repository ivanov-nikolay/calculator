# Сервис подсчёта арифметических выражений

Арифметический калькулятор (Arithmetic calculator) позвляет выполнять арифметические вычисления над целыми числами и числами с десятичной запятой<br>
(разделитель дробной и целой части - точка), приоритет вычисления подвыражений выделяется круглыми скобками<br>

Арифметическое выражение передается в теле запроса POST
## Пример:
```json
{
    "expression": "выражение, которое ввёл пользователь"
}
```

### Получение результата для выражения "(2+7)*4"
```json
{
    "expression": "(2+7)*4"
}
```

### Получение результата для выражения с делением "10*2"
```json
{
    "expression": "10*2"
}
```

## Результат вычисления возвращается в виде json
```json
{
  "result": "36.00"
}
```
и 
```json
{
  "result": "20.00"
}
```
cоответственно.

Eсли при вычислении значения возникает ошибка, то ответ будет тоже в виде json:<br>
```json
{
    "error": "expression is not valid"
}
```
В случае возникновения внутренней ошибки приложения возвращается ошибка<br>

```json
{
    "error": "internal server error"
}
```

## Коды ответа

- успешное вычисление значения: `200 (OK)`<br>
- некорректной ввод арифметичесого выражения: `422 (Status Unprocessable Entity)`<br>
- внутренняя ошибка приложения: `500 (Internal Server Error)`<br>

## Логирование

Все действия приложения фиксируются в журнале: **файл** ***logs.log***

- успешное вычисление значения: 
```
2024/12/01 21:53:17 (2+7)*4 = 36.00
```
- ошибка (например, деление на 0): 
```
2024/12/01 22:12:43 the key 'answer' is missing from the answer: map[error:divided by zero]
2024/12/01 22:12:43 10/0 = processing error
```
- арифметическое выражение отсутствует:
```
2024/12/01 22:13:05 the key 'result' is missing from the response: map[error:expression is required]
2024/12/01 22:13:05 the query does not contain an expression
```

### Запуск приложения:
```
go run cmd/main.go
```
#### Пример curl:
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "(5+2)/4+(2+2)*8+(4-3)*10"
}'
```

Ответ:
```json
{"result":"43.75"}
```

Запуск тестов:
```
go test ./...
```
#### Приложение работает на порту 8080, номер порта записан в файле .env:
```
SERVER_PORT=:8080
```
Файл .env должен быть обязательно, если порт не задан, то приложение по умолчанию запускается на порту :8080

##### Контакт для связи:
telegram @ivanovnickv