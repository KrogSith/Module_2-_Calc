# Calc
# Golang calculator

Данная программа - веб-калькулятор: пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат, в противном случае - ошибку.


## Запуск:
Для запуска введите в терминал команду: go run ./cmd/main.go

Адрес сервера: http://localhost:8080

Запустив программу, воспользуйтесь сервисом Postman и отошлите запрос вида:
{
    "expression": "выражение, которое вы хотите посчитать"
}

## Ограничения:
Калькулятор сломается, если выражение начинается с отрицательного числа


## Пример использования:
Адрес: ```http://localhost:8080/api/v1/calculate```

Запрос:
```bash
{
    "expression": "2+2"
}
```
Ответ:
```bash
{
    "id": 1
}
```


Адрес: ```http://localhost:8080/api/v1/calculate```

Запрос:
```bash
{
    "expression": "124-(253878*351753)/1251"
}
```
Ответ:
```bash
{
    "id": 2
}
```


Адрес: ```http://localhost:8080/api/v1/expressions```

Запрос:
```bash
{
    "expressions": ""
}
```
Ответ:
```bash
{
    "expressions": [
        {
            "id": 1
            "status": OK
            "result": 4
        },
        {
            "id": 2
            "status": OK
            "result": -7.138464669064748e+07
        },
    ]
}
```


Адрес: ```http://localhost:8080/api/v1/expressions/:id```

Запрос:
```bash
{
    "id":  2
}
```
Ответ:
```bash
{
    "expression":
        {
            "id": 1
            "status": OK
            "result": -7.138464669064748e+07
        }
}
```


Адрес: ```http://localhost:8080/internal/task```

Запрос:
```bash
{
    "task": "2+2"
}
```
Ответ:
```bash
{
    "task":
        {
            "id": 1
            "arg1": 2
            "operation": "+"
            "operation_time": 10
            "arg2": 2
        }
}
```
