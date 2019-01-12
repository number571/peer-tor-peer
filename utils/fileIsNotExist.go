package utils

import "os"

func FileIsNotExist (name string) bool {
    if _, err := os.Stat(name); os.IsNotExist(err) {
        return true
    }
    return false
}
