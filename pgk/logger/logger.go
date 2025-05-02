package logger

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

//### Задача 3: Логгер сообщений
//**ТЗ: Гибкая система логирования**
//
//**Описание**:
//Разработайте систему логирования, которая поддерживает разные способы вывода сообщений (в консоль и в файл). Логгер должен поддерживать уровни сообщений (Info, Warning, Error).
//
//**Требования**:
//1. Определите интерфейс `Logger` с методом `Log(level string, message string) error`, который записывает сообщение с указанным уровнем.
//2. Реализуйте две структуры:
//- `ConsoleLogger` — выводит сообщения в консоль в формате: `[<уровень>] <сообщение>`.
//- `FileLogger` — записывает сообщения в строку (для упрощения вместо файла используйте поле типа `strings.Builder` для хранения логов).
//3. `FileLogger` должен иметь метод `GetLogs() string`, возвращающий все записанные сообщения.
//4. Создайте функцию `TestLogger(l Logger)`, которая записывает три сообщения с уровнями "Info", "Warning", "Error".
//5. В функции `main` создайте экземпляры `ConsoleLogger` и `FileLogger`, вызовите `TestLogger` для каждого, а затем выведите содержимое логов для `FileLogger`.
//
//**Ограничения**:
//- Уровни сообщений — только "Info", "Warning", "Error".
//- Для `FileLogger` сообщения сохраняются в `strings.Builder` вместо реального файла.
//- Обработка ошибок не требуется (метод `Log` всегда возвращает `nil`).
//
//**Пример вывода**:
//```
//[Info] Starting process
//[Warning] Resource low
//[Error] Failed to complete
//File logs:
//[Info] Starting process
//[Warning] Resource low
//[Error] Failed to complete
//```

type Logger interface {
	Log(level string, message string) error
}

type ConsoleLogger struct {
}

func (logger *ConsoleLogger) Log(level string, message string) error {
	fmt.Printf("[%s] %s\n", level, message)
	return nil
}

type FileLogger struct {
	Path string
}

func (fl *FileLogger) Log(level string, message string) error {
	file, err := os.OpenFile(fl.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	lineString := fmt.Sprintf("[%s] %s\n", level, message)
	_, err = file.WriteString(lineString)
	if err != nil {
		return err
	}
	return nil
}

func (fl *FileLogger) GetLogs() string {
	file, err := os.Open(fl.Path)
	defer file.Close()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("[Error] Файл %s не найден", fl.Path)
		}
		fmt.Printf("[Error] Не удалось открыть файл: %s", err)
	}
	result := ""
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				result += str
				break
			}
			fmt.Printf("[Error] %s\n", err)
			break
		}
		result += str
	}
	return result
}

func TestLogger(l Logger) {
	err := l.Log("Info", "Starting process")
	if err != nil {
		fmt.Printf("[Error.] %s\n", err)
	}
	err = l.Log("Warning", "Resource low")
	if err != nil {
		fmt.Printf("[Error.] %s\n", err)
	}
	err = l.Log("Error", "Failed to complete")
	if err != nil {
		fmt.Printf("[Error.] %s\n", err)
	}
}
