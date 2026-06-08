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