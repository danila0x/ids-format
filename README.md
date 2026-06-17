## Personal project for educational and practical purposes

### Installing

Clone the repository
`git clone https://github.com/danila0x/ids-format.git`

Open the directory in terminal

Compile: `go build -o ids.exe`

You can adding it to the system PATH to run from anywhere

1. Run CMD as Administrator
2. Execute the following command: `setx /M PATH "%PATH%;C:\path\to\folder\with\ids.exe"`
3. Restart CMD

After that, you can launch from anywhere.

Examples:
`ids.exe 12 2 23 21 12 3`
Output: 
```
12, 2, 23, 21, 12, 3
Обработано IDs:  6
```
`ids.exe -q single -sql -table users 123 12 1  23 4 42 1123 33212`
Output: 
```
SELECT * FROM users WHERE id IN ('123', '12', '1', '23', '4', '42', '1123', '33212')
Обработано IDs:  8
```
`ids.exe -q single -sql -table users -column user_id 123 12 1  23 4 42 1123 33212  211 2 222`
Output: 
```
SELECT * FROM users WHERE user_id IN ('123', '12', '1', '23', '4', '42', '1123', '33212', '211', '2', '222')
Обработано IDs:  11
```

You can also insert an ID from the buffer.
Example:
```
ids.exe -q single -sql -table users
Введите ID (каждый на новой строке). Для завершения введите пустую строку:
444
43423
77777
343
888875
```
Output:
```
SELECT * FROM users WHERE id IN ('444', '43423', '77777', '343', '888875')
Обработано IDs:  5
```

To get information of all the flags, you can press enter twice after starting the program.
Example:
```
>>ids.exe
Введите ID (каждый на новой строке). Для завершения введите пустую строку:

ID не переданы
Пример: ids.exe -q single 12312312 232131 334332
Примеры разделителей:
  -sep-type comma   -> 1, 2, 3
  -sep-type space   -> 1 2 3
  -sep-type newline -> 1\n2\n3
  -sep-type sql     -> 1,\n2,\n3

SQL режим:
  -sql                - включить формирование SQL запроса
  -table имя_таблицы  - указать таблицу
  -column имя_колонки - указать колонку (по умолчанию 'id')
  Пример: ids.exe -sql -table users -column user_id 123 456
  -trim-ext .txt    - удалить расширение .txt у файлов
  -trim-ext .png,.txt - удалить несколько расширений
  -column string
        Имя колонки для SQL запроса (по умолчанию 'id')
  -q string
        Тип кавычек: none, single, double (default "none")
  -sep string
        Разделитель (default ", ")
  -sep-type string
        Тип разделителя: comma, space, newLine, sql, custom
  -sql
        Режим SQL запроса: true or false
  -table string
        Имя таблицы для SQL запроса
  -trim-all-ext
        Удалить все расширения файлов
  -trim-ext string
        Удалить расширение файлов (например: .png, .txt)

```