package utils

// Language содержит переводы сообщений на разные языки.
var Language = map[string]map[string]string{
	"en": {
		"usage":              "Usage: go run formatter.go <command> [options]",
		"commands":           "Commands:",
		"format_file":        "<filename_or_directory> [config_file] - Format a Go file or directory",
		"edit_config":        "edit <config_file> - Edit the configuration file",
		"no_config":          "No configuration file specified.",
		"error_opening":      "Error opening configuration file:",
		"config_updated":     "Configuration updated successfully.",
		"current_config":     "Current configuration:",
		"ignored_dirs":       "Ignored Directories:",
		"ignored_files":      "Ignored Files:",
		"enter_ignore_dirs":  "Enter directories to ignore (comma-separated, or leave blank to keep current):",
		"enter_ignore_files": "Enter file patterns to ignore (comma-separated, or leave blank to keep current):",
		"error_saving":       "Error saving configuration:",
		"error_reading":      "Error reading configuration:",
		"file_formatting":    "Formatting file:",
		"error_formatting":   "Error formatting file:",
		"success":            "Success:",
	},
	"ru": {
		"usage":              "Использование: go run formatter.go <команда> [опции]",
		"commands":           "Команды:",
		"format_file":        "<filename_or_directory> [config_file] - Форматировать файл или директорию Go",
		"edit_config":        "edit <config_file> - Редактировать конфигурационный файл",
		"no_config":          "Не указан файл конфигурации.",
		"error_opening":      "Ошибка открытия файла конфигурации:",
		"config_updated":     "Конфигурация успешно обновлена.",
		"current_config":     "Текущая конфигурация:",
		"ignored_dirs":       "Игнорируемые директории:",
		"ignored_files":      "Игнорируемые файлы:",
		"enter_ignore_dirs":  "Введите директории для игнорирования (через запятую, или оставьте пустым для сохранения текущих):",
		"enter_ignore_files": "Введите шаблоны файлов для игнорирования (через запятую, или оставьте пустым для сохранения текущих):",
		"error_saving":       "Ошибка сохранения конфигурации:",
		"error_reading":      "Ошибка чтения конфигурации:",
		"file_formatting":    "Форматирование файла:",
		"error_formatting":   "Ошибка форматирования файла:",
		"success":            "Успех:",
	},
}

// GetMessage возвращает сообщение на указанном языке.
func GetMessage(lang, key string) string {
	if message, exists := Language[lang][key]; exists {
		return message
	}
	return key // Возвращаем ключ, если перевод отсутствует
}
