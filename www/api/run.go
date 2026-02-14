package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	Gor "github.com/iwhitebird/Gor"
)

type requestBody struct {
	Code string `json:"code"`
}

type responseBody struct {
	Output string `json:"output"`
	AST    string `json:"ast,omitempty"`
	Error  string `json:"error,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, responseBody{Error: "Method not allowed"})
		return
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 64*1024))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, responseBody{Error: "Failed to read request body"})
		return
	}
	defer r.Body.Close()

	var req requestBody
	if err := json.Unmarshal(body, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, responseBody{Error: "Invalid JSON"})
		return
	}

	if len(req.Code) == 0 {
		writeJSON(w, http.StatusBadRequest, responseBody{Error: "No code provided"})
		return
	}

	if len(req.Code) > 50000 {
		writeJSON(w, http.StatusBadRequest, responseBody{Error: "Code too large (max 50KB)"})
		return
	}

	var result Gor.Result
	var execErr string

	done := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				execErr = fmt.Sprintf("%v", r)
			}
			close(done)
		}()
		res := <-Gor.RunFromInput(req.Code)
		result = res
	}()

	select {
	case <-done:
	case <-time.After(8 * time.Second):
		writeJSON(w, http.StatusRequestTimeout, responseBody{Error: "Execution timed out (8s limit)"})
		return
	}

	if execErr != "" {
		writeJSON(w, http.StatusOK, responseBody{Output: execErr, Error: execErr})
		return
	}

	resp := responseBody{
		Output: result.Output,
		AST:    result.AST,
	}

	if result.Error != nil {
		resp.Error = result.Error.Error()
		resp.Output = result.Error.Error()
	}

	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, body responseBody) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
