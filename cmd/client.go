package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect to the cmpr server",
	Long:  "Connect to the cmpr server",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := net.Dial("unix", SockFilePath)
		if err != nil {
			panic(err)
		}
		defer c.Close()

		ch := make(chan int)
		go reader(c, ch)

		_, err = c.Write([]byte("hello"))
		if err != nil {
			log.Fatal("net.Conn.Write() error: ", err)
		}

		<-ch
	},
}

func reader(c net.Conn, ch chan int) {
	buf := make([]byte, 4092)
	println("reader")
	for {
		println("read: start")
		n, err := c.Read(buf)
		println("read: end")
		if err != nil {
			// TODO: Check EOF
			log.Fatal("net.Conn.Read() error: ", err)
			break
		}
		println("read: ", string(buf[0:n]))
	}
	ch <- 1
}
