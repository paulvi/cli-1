// This file was generated by counterfeiter
package fake_ssh

import (
	"net"
	"sync"

	"golang.org/x/crypto/ssh"
)

type FakeConn struct {
	UserStub        func() string
	userMutex       sync.RWMutex
	userArgsForCall []struct{}
	userReturns     struct {
		result1 string
	}
	SessionIDStub        func() []byte
	sessionIDMutex       sync.RWMutex
	sessionIDArgsForCall []struct{}
	sessionIDReturns     struct {
		result1 []byte
	}
	ClientVersionStub        func() []byte
	clientVersionMutex       sync.RWMutex
	clientVersionArgsForCall []struct{}
	clientVersionReturns     struct {
		result1 []byte
	}
	ServerVersionStub        func() []byte
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct{}
	serverVersionReturns     struct {
		result1 []byte
	}
	RemoteAddrStub        func() net.Addr
	remoteAddrMutex       sync.RWMutex
	remoteAddrArgsForCall []struct{}
	remoteAddrReturns     struct {
		result1 net.Addr
	}
	LocalAddrStub        func() net.Addr
	localAddrMutex       sync.RWMutex
	localAddrArgsForCall []struct{}
	localAddrReturns     struct {
		result1 net.Addr
	}
	SendRequestStub        func(name string, wantReply bool, payload []byte) (bool, []byte, error)
	sendRequestMutex       sync.RWMutex
	sendRequestArgsForCall []struct {
		name      string
		wantReply bool
		payload   []byte
	}
	sendRequestReturns struct {
		result1 bool
		result2 []byte
		result3 error
	}
	OpenChannelStub        func(name string, data []byte) (ssh.Channel, <-chan *ssh.Request, error)
	openChannelMutex       sync.RWMutex
	openChannelArgsForCall []struct {
		name string
		data []byte
	}
	openChannelReturns struct {
		result1 ssh.Channel
		result2 <-chan *ssh.Request
		result3 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
}

func (fake *FakeConn) User() string {
	fake.userMutex.Lock()
	fake.userArgsForCall = append(fake.userArgsForCall, struct{}{})
	fake.userMutex.Unlock()
	if fake.UserStub != nil {
		return fake.UserStub()
	} else {
		return fake.userReturns.result1
	}
}

func (fake *FakeConn) UserCallCount() int {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	return len(fake.userArgsForCall)
}

func (fake *FakeConn) UserReturns(result1 string) {
	fake.UserStub = nil
	fake.userReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConn) SessionID() []byte {
	fake.sessionIDMutex.Lock()
	fake.sessionIDArgsForCall = append(fake.sessionIDArgsForCall, struct{}{})
	fake.sessionIDMutex.Unlock()
	if fake.SessionIDStub != nil {
		return fake.SessionIDStub()
	} else {
		return fake.sessionIDReturns.result1
	}
}

func (fake *FakeConn) SessionIDCallCount() int {
	fake.sessionIDMutex.RLock()
	defer fake.sessionIDMutex.RUnlock()
	return len(fake.sessionIDArgsForCall)
}

func (fake *FakeConn) SessionIDReturns(result1 []byte) {
	fake.SessionIDStub = nil
	fake.sessionIDReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeConn) ClientVersion() []byte {
	fake.clientVersionMutex.Lock()
	fake.clientVersionArgsForCall = append(fake.clientVersionArgsForCall, struct{}{})
	fake.clientVersionMutex.Unlock()
	if fake.ClientVersionStub != nil {
		return fake.ClientVersionStub()
	} else {
		return fake.clientVersionReturns.result1
	}
}

func (fake *FakeConn) ClientVersionCallCount() int {
	fake.clientVersionMutex.RLock()
	defer fake.clientVersionMutex.RUnlock()
	return len(fake.clientVersionArgsForCall)
}

func (fake *FakeConn) ClientVersionReturns(result1 []byte) {
	fake.ClientVersionStub = nil
	fake.clientVersionReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeConn) ServerVersion() []byte {
	fake.serverVersionMutex.Lock()
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct{}{})
	fake.serverVersionMutex.Unlock()
	if fake.ServerVersionStub != nil {
		return fake.ServerVersionStub()
	} else {
		return fake.serverVersionReturns.result1
	}
}

func (fake *FakeConn) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeConn) ServerVersionReturns(result1 []byte) {
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeConn) RemoteAddr() net.Addr {
	fake.remoteAddrMutex.Lock()
	fake.remoteAddrArgsForCall = append(fake.remoteAddrArgsForCall, struct{}{})
	fake.remoteAddrMutex.Unlock()
	if fake.RemoteAddrStub != nil {
		return fake.RemoteAddrStub()
	} else {
		return fake.remoteAddrReturns.result1
	}
}

func (fake *FakeConn) RemoteAddrCallCount() int {
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	return len(fake.remoteAddrArgsForCall)
}

func (fake *FakeConn) RemoteAddrReturns(result1 net.Addr) {
	fake.RemoteAddrStub = nil
	fake.remoteAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeConn) LocalAddr() net.Addr {
	fake.localAddrMutex.Lock()
	fake.localAddrArgsForCall = append(fake.localAddrArgsForCall, struct{}{})
	fake.localAddrMutex.Unlock()
	if fake.LocalAddrStub != nil {
		return fake.LocalAddrStub()
	} else {
		return fake.localAddrReturns.result1
	}
}

func (fake *FakeConn) LocalAddrCallCount() int {
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	return len(fake.localAddrArgsForCall)
}

func (fake *FakeConn) LocalAddrReturns(result1 net.Addr) {
	fake.LocalAddrStub = nil
	fake.localAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeConn) SendRequest(name string, wantReply bool, payload []byte) (bool, []byte, error) {
	fake.sendRequestMutex.Lock()
	fake.sendRequestArgsForCall = append(fake.sendRequestArgsForCall, struct {
		name      string
		wantReply bool
		payload   []byte
	}{name, wantReply, payload})
	fake.sendRequestMutex.Unlock()
	if fake.SendRequestStub != nil {
		return fake.SendRequestStub(name, wantReply, payload)
	} else {
		return fake.sendRequestReturns.result1, fake.sendRequestReturns.result2, fake.sendRequestReturns.result3
	}
}

func (fake *FakeConn) SendRequestCallCount() int {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return len(fake.sendRequestArgsForCall)
}

func (fake *FakeConn) SendRequestArgsForCall(i int) (string, bool, []byte) {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return fake.sendRequestArgsForCall[i].name, fake.sendRequestArgsForCall[i].wantReply, fake.sendRequestArgsForCall[i].payload
}

func (fake *FakeConn) SendRequestReturns(result1 bool, result2 []byte, result3 error) {
	fake.SendRequestStub = nil
	fake.sendRequestReturns = struct {
		result1 bool
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConn) OpenChannel(name string, data []byte) (ssh.Channel, <-chan *ssh.Request, error) {
	fake.openChannelMutex.Lock()
	fake.openChannelArgsForCall = append(fake.openChannelArgsForCall, struct {
		name string
		data []byte
	}{name, data})
	fake.openChannelMutex.Unlock()
	if fake.OpenChannelStub != nil {
		return fake.OpenChannelStub(name, data)
	} else {
		return fake.openChannelReturns.result1, fake.openChannelReturns.result2, fake.openChannelReturns.result3
	}
}

func (fake *FakeConn) OpenChannelCallCount() int {
	fake.openChannelMutex.RLock()
	defer fake.openChannelMutex.RUnlock()
	return len(fake.openChannelArgsForCall)
}

func (fake *FakeConn) OpenChannelArgsForCall(i int) (string, []byte) {
	fake.openChannelMutex.RLock()
	defer fake.openChannelMutex.RUnlock()
	return fake.openChannelArgsForCall[i].name, fake.openChannelArgsForCall[i].data
}

func (fake *FakeConn) OpenChannelReturns(result1 ssh.Channel, result2 <-chan *ssh.Request, result3 error) {
	fake.OpenChannelStub = nil
	fake.openChannelReturns = struct {
		result1 ssh.Channel
		result2 <-chan *ssh.Request
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeConn) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeConn) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

var _ ssh.Conn = new(FakeConn)
