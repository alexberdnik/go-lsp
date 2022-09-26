package jsonrpc

import (
	"context"
	"sync"
)

type MethodInfo struct {
	Name       string
	NewRequest func() interface{}
	Handler    func(ctx context.Context, req interface{}) (interface{}, error)
}

type Server struct {
	session     *Session
	methods     map[string]MethodInfo
	sessionLock sync.Mutex
}

func NewServer() *Server {
	s := &Server{}
	s.methods = make(map[string]MethodInfo)

	// Register Builtin
	s.RegisterMethod(CancelRequest())

	return s
}

func (s *Server) Notify(msg *NotificationMessage) error {
	if s.session != nil {
		return s.session.Notify(msg)
	}
	return nil
}

func (s *Server) RegisterMethod(m MethodInfo) {
	s.methods[m.Name] = m
}

func (s *Server) Connect(conn ReaderWriter) {
	s.session = s.newSession(conn)
	s.session.Start()
}

func (s *Server) removeSession() {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	s.session = nil
}

func (s *Server) newSession(conn ReaderWriter) *Session {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	session := newSession(s, conn)
	return session
}
