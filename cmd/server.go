package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the cmpr server",
	Long:  "Start the cmpr server",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := net.Listen("unix", SockFilePath)
		if err != nil {
			log.Fatal("net.Listen() error: ", err)
		}
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal("net.Conn.Accept() error: ", err)
			}
			go processConnection(c)
		}
	},
}

func processConnection(c net.Conn) {
	// TODO: Implement this method. Currently, this method acts as echo server.
	buf := make([]byte, 4092)
	for {
		n, err := c.Read(buf)
		if err != nil {
			// TODO: Check EOF
			return
		}
		data := buf[0:n]
		println("read: ", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("net.Conn.Write() err: ", err)
		}
	}
}
