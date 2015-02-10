package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	tftp "github.com/pin/tftp"
)

var localConf *Config

func HandleWrite(filename string, r *io.PipeReader) {
	r.CloseWithError(fmt.Errorf("Server is Read Only"))
}

func HandleRead(filename string, w *io.PipeWriter) {
	fmt.Printf("Filename : %v \n", []byte(filename))
	var exists bool
	d, err := localConf.fs.Get("tftp/" + filename[0:len(filename)-1])
	fmt.Println(d, err)
	if err == nil {
		exists = true
	}
	if exists {
		c, e := io.Copy(w, d)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Can't send %s: %v\n", filename, e)
		} else {
			fmt.Fprintf(os.Stderr, "Sent %s (%d bytes)\n", filename, c)
		}
		w.Close()
	} else {
		w.CloseWithError(fmt.Errorf("File not exists: %s", filename))
	}
}

func tftpServer(conf *Config) {
	fmt.Println("start tftp")
	localConf = conf
	addrStr := flag.String("l", ":69", "Address to listen")
	flag.Parse()
	addr, e := net.ResolveUDPAddr("udp", *addrStr)
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		return
	}
	log := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	s := tftp.Server{addr, HandleWrite, HandleRead, log}
	e = s.Serve()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
}
