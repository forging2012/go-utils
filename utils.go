package utils

import "os"
import "github.com/Sirupsen/logrus"
import "net"

func GetHostName() string {
  hostname, err := os.Hostname()
  CheckError(err)
  return hostname
}


func CheckError(err error) {
    if err != nil {
        logrus.Errorf("Fatal Error: %v",err)
        os.Exit(1)
    }
}

func GetIP() string {
    addresses, err := net.InterfaceAddrs()

    CheckError(err)

    for _, address := range addresses {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

