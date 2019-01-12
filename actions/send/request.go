package send

import (
    "sync"
    "bytes"
    "net/http"
    "../../utils"
    "../../settings"
)

func Request(client *http.Client, message string) {
    var mutex sync.Mutex
    for _, address := range settings.Addresses {
        send(client, address, message, &mutex)
    }
}

func send(client *http.Client, address, message string, mutex *sync.Mutex) {
    _, err := client.Post(
        "http://" + address + ".onion", 
        settings.CONTENT_TYPE, 
        bytes.NewReader([]byte(message)),
    )

    if err != nil {
        mutex.Lock()
        for index, value := range settings.Addresses {
            if value == address {
                settings.Addresses = utils.Remove(settings.Addresses, index)
            }
        }
        mutex.Unlock()
    }
}
