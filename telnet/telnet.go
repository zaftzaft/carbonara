package telnet

import (
	//"fmt"
	"io"
	"net"
	//"os"
	//"time"
)

func Dial(host string, w io.Writer) (conn net.Conn, exit chan error, err error) {
	conn, err = net.Dial("tcp", host)
	if err != nil {
		return conn, exit, err
	}

	exit = make(chan error)

	go func() {
		for {
			var nowcmd TelnetCmd
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				exit <- err
			}

			p := 0

			for p < n {
				c := b[p]

				if TelnetCmd(c) == IAC {
					nowcmd = IAC
				}

				if nowcmd == IAC {
					nowcmd = TelnetCmd(c)
				} else {
					if len(nowcmd.String()) > 0 {
						//fmt.Println(nowcmd, TelnetOp(c).String())
						if nowcmd == DO {
							conn.Write([]byte{byte(IAC), byte(WONT), byte(TelnetOp(c))})
						}
					} else {
						//w.Write(b[p:])
						//break
						// FIXME
						w.Write([]byte{c})
					}
				}

				p++
			}

		}
	}()

	return conn, exit, err
}
