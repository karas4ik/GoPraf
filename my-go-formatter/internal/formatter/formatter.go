package formatter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"my-go-formatter/internal/logger"
	"my-go-formatter/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

// Config представляет конфигурацию форматировщика.
type Config struct {
	IgnoreDirs  []string `json:"ignore_dirs"`  // Список игнорируемых директорий
	IgnoreFiles []string `json:"ignore_files"` // Список игнорируемых файлов
}

// LoadConfig загружает конфигурацию из файла.
func LoadConfig(filename string) (Config, error) {
	var config Config
	if filename == "" {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return config, fmt.Errorf("ошибка открытия файла конфигурации: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("ошибка чтения конфигурации: %w", err)
	}
	return config, nil
}

// Format форматирует файл или директорию.
func Format(target string, config Config, lang string) error {
	info, err := os.Stat(target)
	if err != nil {
		return fmt.Errorf("ошибка доступа к пути: %w", err)
	}

	if info.IsDir() {
		return formatDirectory(target, config, lang)
	} else if strings.HasSuffix(info.Name(), ".go") {
		return formatFile(target, lang)
	} else {
		return fmt.Errorf("указанный путь не является файлом Go или директорией")
	}
}

// formatFile форматирует указанный файл.
func formatFile(filename string, lang string) error {
	inputFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// Проверяем отступы
	formattedLines := checkIndentation(lines)

	outputFile, err := os.Create(filename + ".formatted")
	if err != nil {
		return fmt.Errorf("ошибка создания выходного файла: %w", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range formattedLines {
		fmt.Fprintln(writer, line)
	}

	logger.Log(fmt.Sprintf("%s %s", utils.GetMessage(lang, "file_formatting"), filename)) // Используем функцию Log
	return writer.Flush()
}

func formatDirectory(dir string, config Config, lang string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path: %w", err)
		}
		// Ignore specified directories
		for _, ignoreDir := range config.IgnoreDirs {
			if info.IsDir() && strings.Contains(path, ignoreDir) {
				return filepath.SkipDir
			}
		}
		// Ignore specified files
		for _, ignoreFile := range config.IgnoreFiles {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ignoreFile) {
				return nil
			}
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			logger.Log(fmt.Sprintf("%s %s", utils.GetMessage(lang, "file_formatting"), path))
			return formatFile(path, lang)
		}
		return nil
	})
}

// checkIndentation проверяет правильность отступов для функций, условий и циклов.
func checkIndentation(lines []string) []string {
	var formattedLines []string
	indentLevel := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Увеличиваем уровень отступа для открывающих фигурных скобок
		if strings.HasSuffix(trimmed, "{") {
			formattedLines = append(formattedLines, line) // добавляем строку с текущим отступом
			indentLevel++
			continue
		}
		// Уменьшаем уровень отступа для закрывающих фигурных скобок
		if trimmed == "}" {
			indentLevel--
			formattedLines = append(formattedLines, line) // добавляем строку с текущим отступом
			continue
		}

		// Применяем отступ в зависимости от уровня
		formattedLines = append(formattedLines, strings.Repeat("    ", indentLevel)+line)
	}

	return formattedLines
}

// EditConfig позволяет редактировать конфигурационный файл.
func EditConfig(filename string, lang string) {
	if filename == "" {
		fmt.Println(utils.GetMessage(lang, "no_config"))
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(utils.GetMessage(lang, "error_opening"), err)
		return
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println(utils.GetMessage(lang, "error_reading"), err)
		return
	}

	fmt.Println(utils.GetMessage(lang, "current_config"))
	fmt.Println(utils.GetMessage(lang, "ignored_dirs"), config.IgnoreDirs)
	fmt.Println(utils.GetMessage(lang, "ignored_files"), config.IgnoreFiles)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(utils.GetMessage(lang, "enter_ignore_dirs"))
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		config.IgnoreDirs = strings.Split(input, ",")
	}

	fmt.Println(utils.GetMessage(lang, "enter_ignore_files"))
	scanner.Scan()
	input = scanner.Text()
	if input != "" {
		config.IgnoreFiles = strings.Split(input, ",")
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(utils.GetMessage(lang, "error_saving"), err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		fmt.Println(utils.GetMessage(lang, "error_saving"), err)
	} else {
		fmt.Println(utils.GetMessage(lang, "config_updated"))
	}
}
