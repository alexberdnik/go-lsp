package lsp

import (
	"github.com/alexberdnik/go-lsp/jsonrpc"
	"github.com/alexberdnik/go-lsp/lsp/defines"
	jsoniter "github.com/json-iterator/go"
)

func (s *Server) PublishDiagnostics(params defines.PublishDiagnosticsParams) error {
	bytes, err := jsoniter.Marshal(params)
	if err != nil {
		return err
	}

	msg := jsonrpc.NotificationMessage{
		BaseMessage: jsonrpc.BaseMessage{},
		Method:      "textDocument/publishDiagnostics",
		Params:      bytes,
	}
	return s.rpcServer.Notify(&msg)
}
