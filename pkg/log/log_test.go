package log

import (
	"testing"
)

func TestFuncName(t *testing.T) {
	// fn := funcName(2)
	// fmt.Println(fn)
}

func TestLog(t *testing.T) {
	// Init(nil)
	// Info("%s", "222322")
	// localUdpAddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:5020")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// svrUdpAddr, err := net.ResolveUDPAddr("udp4", "118.178.140.41:60000")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// conn, _ := net.DialUDP("udp", localUdpAddr, svrUdpAddr)
	// conn.Write([]byte("hhhhhhhhhhhhhhhhhhhhhhhh"))
	Init(nil)
	Info("%s", "222322")
	// h := newNlogHnadle("118.178.140.41", 60000)
	// h.Log(context.Background(), _debugLevel, KVString("test", "test"))
}
