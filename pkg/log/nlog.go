package log

import (
	"context"
	"fmt"
	"net"
)

var arrUDPConn = [6]*net.UDPConn{}

type nlog struct {
	infoUDPCon  *net.UDPConn
	debugUDPCon *net.UDPConn
	fetalUDPCon *net.UDPConn
	errorUDPCon *net.UDPConn
	warnUDPCon  *net.UDPConn
	render      Render
}

func newNlogHnadle(udpAddr, port string) Handle {

	svrUDPAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", udpAddr, port))
	if err != nil {
		panic(err)
	}
	infoUDPAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", "0.0.0.0", 504))
	if err != nil {
		panic(err)
	}
	fetalUDPAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", "0.0.0.0", 501))
	if err != nil {
		panic(err)
	}
	errorUDPAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", "0.0.0.0", 503))
	if err != nil {
		panic(err)
	}
	accessUDPAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", "0.0.0.0", 502))
	if err != nil {
		panic(err)
	}

	infoConn, err := net.DialUDP("udp", infoUDPAddr, svrUDPAddr)
	if err != nil {
		panic(err)
	}
	arrUDPConn[_infoLevel] = infoConn
	fetalConn, err := net.DialUDP("udp", fetalUDPAddr, svrUDPAddr)
	if err != nil {
		panic(err)
	}

	arrUDPConn[_fetalLevel] = fetalConn
	errorConn, err := net.DialUDP("udp", errorUDPAddr, svrUDPAddr)
	if err != nil {
		panic(err)
	}

	arrUDPConn[_errorLevel] = errorConn
	arrUDPConn[_debugLevel] = errorConn
	arrUDPConn[_warnLevel] = errorConn

	accessConn, err := net.DialUDP("udp", accessUDPAddr, svrUDPAddr)
	if err != nil {
		panic(err)
	}
	arrUDPConn[_accessLevel] = accessConn
	return nlog{
		render: newPatternRender(),
	}
}

//Log ...
func (st nlog) Log(ctx context.Context, l Level, d ...D) {
	u := arrUDPConn[l]
	st.render.Render(u, d...)
}

//Close ...
func (st nlog) Close() error {
	return nil
}
