package dial_tcp

import (
"log"
"strconv"
"time"

"github.com/StabbyCutyou/buffstreams"
	"fmt"
)

// Test client to send a sample payload of data endlessly
// By default it points locally, but it can point to any network address
// TODO Make that externally configurable to make automating the test easier
func Connect() {
	cfg := &buffstreams.TCPConnConfig{
		MaxMessageSize: 2048,
		Address:        buffstreams.FormatAddress("127.0.0.1", strconv.Itoa(6000)),
	}
	/*name := "Client"
	date := time.Now().UnixNano()
	data := "This is an intenntionally long and rambling sentence to pad out the size of the message."
	msg := &message.Note{Name: &name, Date: &date, Comment: &data}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Print(err)
	}*/
	btw, err := buffstreams.DialTCP(cfg)
	if err != nil {
		log.Fatal(err)
	}
	var msgBytes []byte
	//0x0006
	//msgBytes = []byte{0x7e,0x00,0x2a,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x36,0x30,0x31,0x32,0x32,0x31,0x38,0x35,0x30,0x33,0x34,0xf3,0x00,0x00,0x00,0x00,0x06,0x05,0x01,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00}
	//0x0003
	//msgBytes =[]byte{0x7e,0x00,0x37,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x36,0x30,0x31,0x32,0x32,0x31,0x38,0x35,0x32,0x34,0x35,0xfa,0x00,0x00,0x00,0x00,0x03,0x00,0x03,0x00,0x13,0xa2,0x00,0x40,0xe9,0x2a,0xae,0x00,0x00,0x13,0xa2,0x00,0x40,0x8e,0x05,0x8c,0x00,0x00,0x13,0xa2,0x00,0x40,0xc7,0x6c,0x05,0x00}
	//0x3001
	//msgBytes =[]byte{0x07e,0x00,0x29,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x36,0x30,0x31,0x32,0x32,0x31,0x38,0x35,0x39,0x32,0x33,0x14,0x01,0x00,0x00,0x30,0x01,0x30,0x00,0x13,0xa2,0x00,0x40,0xc7,0x6b,0x8c,0x00,0xff,0x03,0x00,0xff,0x03}

	//e000
	//msgBytes =[]byte{0x7e,0x00,0x64,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x36,0x30,0x31,0x32,0x32,0x31,0x38,0x35,0x32,0x32,0x32,0xf8,0x00,0x00,0x00,0xe0,0x00,0x00,0x03,0x00,0x13,0xa2,0x00,0x40,0xe9,0x2a,0xae,0x00,0x01,0x00,0x00,0x00,0x07,0xd4,0x06,0x6d,0x00,0x00,0x00,0x00,0x00,0x49,0x47,0x00,0x13,0xa2,0x00,0x40,0x8e,0x05,0x8c,0x00,0x01,0x01,0x01,0x01,0x00,0x00,0x00,0x00,0x02,0x00,0x14,0x02,0x00,0x49,0x00,0x00,0x13,0xa2,0x00,0x40,0xc7,0x6c,0x05,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x00,0x14,0x02,0x00,0x49,0x00}
	
	//msgBytes =[]byte{0x7e,0x00,0x7c,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x35,0x31,0x32,0x31,0x39,0x31,0x32,0x32,0x36,0x30,0x30,0x83,0x00,0x00,0x00,0xe0,0x00,0x00,0x04,0x00,0x13,0xa2,0x00,0x40,0xe9,0x2a,0xae,0x00,0x01,0x00,0x00,0x00,0x07,0xd7,0x06,0x82,0x00,0x00,0x00,0x00,0x00,0x41,0x48,0x00,0x13,0xa2,0x00,0x40,0xc7,0x6c,0x05,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x00,0x14,0x02,0x00,0x41,0x00,0x00,0x13,0xa2,0x00,0x40,0x8e,0x05,0x8c,0x00,0x01,0x01,0x01,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x41,0x00,0x00,0x13,0xa2,0x00,0x40,0xc7,0x6b,0x8c,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x00,0x14,0x02,0x00,0x41,0x00}

	//0002
	msgBytes = []byte{0x7e,0x00,0x1a,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x35,0x31,0x32,0x31,0x39,0x31,0x32,0x32,0x38,0x33,0x33,0x8c,0x00,0x00,0x00,0x00,0x02}
	fmt.Print(msgBytes," len=>",len(msgBytes))
	_, err = btw.Write(msgBytes)
	if err != nil {
		log.Print("There was an error")
		log.Print(err)
	}
	time.Sleep(100 * time.Millisecond)
	readBytes := make([]byte, 4096)
	_, err = btw.Read(readBytes)
	if err != nil {
		log.Printf("Address %s: Failure to read from connection", err)
		btw.Close()
		return
	}
	fmt.Println("dsds")
	fmt.Print(readBytes)



}