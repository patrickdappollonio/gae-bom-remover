package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 20) // ~32 MB
	file, meta, err := r.FormFile("bomfile")

	if err != nil {
		if err == http.ErrMissingFile {
			http.Error(w, "", http.StatusNoContent)
			return
		}

		http.Error(w, "Unable to process uploaded file. "+err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "Unable to read file for cleanup. "+err.Error(), http.StatusInternalServerError)
		return
	}

	if ct := http.DetectContentType(buf.Bytes()); !strings.Contains(ct, "text/") {
		http.Error(w, "Unable to process a non-JSON file. Type reported: "+ct, http.StatusInternalServerError)
		return
	}

	data := bytes.Trim(buf.Bytes(), "\xef\xbb\xbf")

	w.Header().Set("Content-Disposition", "attachment; filename="+meta.Filename)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(data)))

	fmt.Fprint(w, string(data))
}
