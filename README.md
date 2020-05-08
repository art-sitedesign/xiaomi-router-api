# xiaomi-router-api
Клиент для API роутера Xiaomi Mi Router 4

### Пример использования:
```go
package main

import (
    "fmt"
    xiaomirouterapi "github.com/art-sitedesign/xiaomi-router-api"
)

func main() {
    api := xiaomirouterapi.NewMiWifiApi("192.168.31.1")
    
    err := api.Auth("admin", "12345678")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    defer func() {
        _ = api.Logout()
    }()
    
    status, err := api.Status()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(status)
}
```