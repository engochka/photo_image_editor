package types

import (
	"encoding/json"
	"homework-engochka/repository"
	"net/http"
)

func ProcessError(w http.ResponseWriter, err error, resp any) {
	if err == repository.ErrNotFound {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error())) // что значит?
		return
	}

	if resp != nil {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		} // вроде бы тут он говорил убрать возврат нил если респ нил
	}
}
