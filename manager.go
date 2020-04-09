package proxymanager

import (
	"os"
	"bufio"
	"errors"
	"math/rand"
	"time"
	"strings"
	"fmt"
)

type ProxyManager struct {
	Proxies []string
	CurrentIndex int
}

func NewManager(filename string) (*ProxyManager, error) {
	file, openErr := os.Open(filename)
	if openErr != nil {
		return nil, openErr
	}
	defer file.Close()
	manager := &ProxyManager{Proxies: []string{}, CurrentIndex: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		manager.Proxies = append(manager.Proxies, scanner.Text())
	}
	return manager, scanner.Err()
}

func (p *ProxyManager) LoadProxies(filename string) error {
	p.Proxies = []string{}
	file, openErr := os.Open(filename)
	if openErr != nil {
		return openErr
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.Proxies = append(p.Proxies, scanner.Text())
	}
	return scanner.Err()
}

func (p *ProxyManager) NextProxy() (string, error) {
	if len(p.Proxies) == 0 {
		return "", errors.New("ProxyManager.Proxies is empty, load proxies")
	}

	p.CurrentIndex++ //increases current index by one
	if p.CurrentIndex > len(p.Proxies) - 1 { 
		p.CurrentIndex = 0 //sets index to 0 if greater than length
	}
	proxy := p.Proxies[p.CurrentIndex]
	if len(strings.Split(proxy, ":")) == 4 {
		splitStr := strings.Split(proxy, ":")
		proxy = fmt.Sprintf("http://%s:%s@%s:%s", splitStr[2], splitStr[3], splitStr[0], splitStr[1])
	}
	return proxy, nil
}

func (p *ProxyManager) RandomProxy() (string, error) {
	if len(p.Proxies) == 0 {
		return "", errors.New("ProxyManager.Proxies is empty, load proxies")
	}

	rand.Seed(time.Now().UnixNano())
	return p.Proxies[rand.Intn(len(p.Proxies))], nil
}