# proxymanager

A simple GoLang implementation of a proxy manager package. Can be initialized with a txt file filled with proxies.
From the ProxyManager object, you can grab either a random proxy, or sequentially go through the proxies in the object.

### Installation
```
go get github.com/aiomonitors/goproxymanager
```

### Usage

Initialize new object
```golang
package main 

import "github.com/aiomonitors/goproxymanager"

func main() {
    manager, err := proxymanager.NewManager("proxies.txt")
    if err != nil{
        panic(err)
    }
    
    //Gets next proxy in the manager
    next, err := manager.NextProxy()
    if err != nil{
        panic(err)
    }
   
    //Gets random proxy in the manager
    randomProxy, err := manager.RandomProxy()
    if err != nil{
        panic(err)
    }
}
```
