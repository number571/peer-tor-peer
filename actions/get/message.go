package get

import (
    "fmt"
    "net/http"
    "../../settings"
)

func Message(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "hello, world")
		return
	}

	var (
        buffer []byte = make([]byte, settings.BUFF_SIZE)
        message string
    )

    for {
        length, err := r.Body.Read(buffer)
        message += string(buffer[:length])
        if length == 0 || err != nil { break }
    }

    fmt.Println(message)
}