package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondFile(writer http.ResponseWriter, request *http.Request, filename string) {
	Openfile, err := os.Open(filename)
	defer Openfile.Close()
	if err != nil {
		http.Error(writer, "Dosya bulunamadı.", 404)
		return
	}

	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	writer.Header().Set("filename", filename)
	writer.Header().Set("Content-Type", ".docx")
	writer.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile)
	return
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
