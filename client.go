package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	address  string
	password string
	commands string
	timeout  int
)

func initPropertiesFromEnv() {
	address = os.Getenv("PLUGIN_ADDRESS")
	password = os.Getenv("PLUGIN_PASSWORD")
	commands = os.Getenv("PLUGIN_COMMANDS")
	t, err := strconv.Atoi(os.Getenv("PLUGIN_TIMEOUT"))
	if err != nil {
		fmt.Println("PLUGIN_TIMEOUT convert err:", err)
	}
	timeout = t
}

func main() {
	initPropertiesFromEnv()
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	if password != "" {
		result := login(password, conn)
		if !result {
			return
		}
	}
	commandList := strings.Split(commands, ",")
	for _, cmd := range commandList {
		sendCommand(cmd, conn)
	}
}

//Packet data struct about rcon
type Packet struct {
	size        uint32
	id          int32
	_type       uint32
	body        []byte
	emptyString []byte
}

var tail = []byte{0x00, 0x00}

func login(password string, conn net.Conn) bool {
	passwordBytes := []byte(password)
	packet := Packet{uint32(len(passwordBytes) + 10), 0, 3, passwordBytes, tail}
	_, err := conn.Write(encode(packet))
	if err != nil {
		fmt.Println("Login err:", err)
	}
	result := decode(conn)
	if result.id == -1 {
		fmt.Println("login failed , check password")
		return false
	}
	fmt.Println("login success")
	return true

}

func sendCommand(command string, conn net.Conn) {
	fmt.Println("sendCommand :", command)
	commandBytes := []byte(command)
	packet := Packet{uint32(len(commandBytes) + 10), 1, 2, commandBytes, tail}
	_, err := conn.Write(encode(packet))
	if err != nil {
		fmt.Println("SendCommand err:", err)
	}
	result := decode(conn)
	if len(result.body) > 0 {
		fmt.Println(string(result.body))
	} else {
		fmt.Println("empty response")
	}
}

func encode(packet Packet) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, packet.size)
	binary.Write(buf, binary.LittleEndian, packet.id)
	binary.Write(buf, binary.LittleEndian, packet._type)
	binary.Write(buf, binary.BigEndian, packet.body)
	binary.Write(buf, binary.BigEndian, packet.emptyString)
	return buf.Bytes()
}

func decode(conn net.Conn) Packet {
	var size uint32
	var id int32
	var _type uint32
	var body []byte
	var emptyString []byte
	s := make([]byte, 4)
	conn.Read(s)
	size = binary.LittleEndian.Uint32(s)
	conn.Read(s)
	id = int32(binary.LittleEndian.Uint32(s))
	conn.Read(s)
	_type = binary.LittleEndian.Uint32(s)
	b := make([]byte, size-10)
	conn.Read(b)
	body = b
	e := make([]byte, 2)
	conn.Read(e)
	emptyString = e
	return Packet{size, id, _type, body, emptyString}
}
