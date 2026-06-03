package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	quote := flag.String("q", "none", "Тип кавычек: none, single, double")
	separator := flag.String("sep", ", ", "Разделитель")
	sepType := flag.String("sep-type", "custom", "Тип разделителя: comma, space, newLine, sql, custom")
	flag.Parse()

	var ids []string

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				ids = append(ids, line)
			}
		}
	} else if len(flag.Args()) > 0 {
		ids = flag.Args()
	} else {
		fmt.Println("Введите ID (каждый на новой строке). Для завершения введите пустую строку:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				break
			}
			ids = append(ids, line)
		}
	}

	if len(ids) == 0 {
		fmt.Println("ID не переданы")
		fmt.Println("Пример: ids.exe -quote single 12312312 232131 334332")
		fmt.Println("Примеры разделителей:")
		fmt.Println("  -sep-type comma   -> 1, 2, 3")
		fmt.Println("  -sep-type space   -> 1 2 3")
		fmt.Println("  -sep-type newline -> 1\\n2\\n3")
		fmt.Println("  -sep-type sql     -> 1,\\n2,\\n3")
		flag.PrintDefaults()
		return
	}

	var result []string
	countIds := 0
	switch *quote {
	case "single":
		for _, id := range ids {
			result = append(result, fmt.Sprintf("'%s'", id))
			countIds++
		}
	case "double":
		for _, id := range ids {
			result = append(result, fmt.Sprintf("\"%s\"", id))
			countIds++
		}
	default:
		result = ids
		countIds = len(ids)
	}

	var finalSeparator string
	switch *sepType {
	case "comma":
		finalSeparator = ", "
	case "space":
		finalSeparator = " "
	case "newLine":
		finalSeparator = "\n"
	case "sql":
		finalSeparator = ",\n"
	case "semicolon":
		finalSeparator = "; "
	case "tab":
		finalSeparator = "\t"
	default:
		finalSeparator = *separator
	}

	output := strings.Join(result, finalSeparator)
	fmt.Println(output)
	fmt.Println("Обработано IDs: ", countIds)
}
