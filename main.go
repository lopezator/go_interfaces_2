package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Utilizar la funcion Read que implementa el Reader Interface
	//Que a su vez es implementado por el Response Body
	//Para leer el Body Response
	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	//Reader "popula" un []byte
	//Writer, toma ese []byte y lo pinta en terminal, hd o lo que sea

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}