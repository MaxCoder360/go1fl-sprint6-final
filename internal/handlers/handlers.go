package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler отдаёт HTML-страницу с формой.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler принимает файл из формы, конвертирует содержимое
// и сохраняет результат в локальный файл.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму, максимум 10 MB
	r.ParseMultipartForm(10 << 20)

	// Получаем файл из формы.
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error receiving file", http.StatusInternalServerError)
		return
	}
	// Закрываем файл при выходе из функции
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	// Определяем, что это и конвертируем
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "error conversion", http.StatusInternalServerError)
		return
	}

	// Формируем имя файла: время в UTC + расширение исходного файла
	fileName := time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)

	// Создаём файл для записи.
	outFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		http.Error(w, "error creating file", http.StatusInternalServerError)
		return
	}
	// Закрываем файл при выходе из функции
	defer outFile.Close()

	// Записываем результат конвертации в файл
	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, "error writing file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат пользователю
	fmt.Fprint(w, result)
}
