package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	example := "protocol://user:pass@host.com:port/path?query=param#queryfragment"

	u, err := url.Parse(example)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) //prints: protocol

	fmt.Println(u.User)            //prints: user:pass
	fmt.Println(u.User.Username()) //prints: user
	pass, _ := u.User.Password()
	fmt.Println(pass) //prints: pass

	fmt.Println(u.Host) //prints: host.com:port
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) //prints: host
	fmt.Println(port) //prints: port

	fmt.Println(u.Path) //prints: /path

	fmt.Println(u.RawQuery) //prints: query=param
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)             //prints: map[query:[param]]
	fmt.Println(m["query"][0]) //prints: param

	fmt.Println(u.Fragment) //prints: queryfragment
}
