package main

import (
  "flag"
  "net"
  "net/http"
  "fmt"
)

var port = flag.String("port", "4701", "http service address")

func main() {
  flag.Parse()

  go echo_interfaces();

  err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(".")))
  if err != nil {
    fmt.Printf("error running webserver: %v", err)
  }
}

func echo_interfaces() {
  ifaces, err := net.Interfaces()
  if err != nil {
    fmt.Printf("error fetching interfaces: %v", err)
  }
  fmt.Println("Serving on:")
  for _, i := range ifaces {
    addrs, err := i.Addrs()
    if err != nil {
      fmt.Printf("error fetching ip address from interface: %v", err)
    }
    if len(addrs) == 0 {
      continue;
    }
    fmt.Println()
    for _, addr := range addrs {
      var ip net.IP
      switch v := addr.(type) {
        case *net.IPNet:
          ip = v.IP
        case *net.IPAddr:
          ip = v.IP
			}
      fmt.Printf("      http://%s:%s\n", ip, *port)
    }
  }
}
