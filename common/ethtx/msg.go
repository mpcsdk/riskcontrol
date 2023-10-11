package ethtx

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	vsn = "2.0"
)

type Error interface {
	Error() string
	ErrorCode() int
}
type JsonMessage struct {
	JsonRpc string          `json:"jsonrpc,omitempty"`
	Id      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func NewMessage(method string, paramsIn ...interface{}) (*JsonMessage, error) {
	msg := &JsonMessage{JsonRpc: vsn, Id: GenMsgId(), Method: method}
	if paramsIn != nil { // prevent sending "params":null
		var err error
		if msg.Params, err = json.Marshal(paramsIn); err != nil {
			return nil, err
		}
	}
	return msg, nil
}

func ParseMessage(body []byte) (*JsonMessage, Error) {
	req := &JsonMessage{}
	err := json.Unmarshal(body, req)

	if err != nil {
		return errorMessage(ParseError(err.Error())), ParseError(err.Error())
	}
	return req, nil
}
func (msg *JsonMessage) Hash() string {
	s := msg.Method + string(msg.Params)
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
func (msg *JsonMessage) IsNotification() bool {
	return msg.Id == nil && msg.Method != ""
}

func (msg *JsonMessage) IsCall() bool {
	return msg.HasValidId() && msg.Method != ""
}

func (msg *JsonMessage) IsResponse() bool {
	return msg.HasValidId() && msg.Method == "" && msg.Params == nil && (msg.Result != nil || msg.Error != nil)
}

func (msg *JsonMessage) HasValidId() bool {
	return len(msg.Id) > 0 && msg.Id[0] != '{' && msg.Id[0] != '['
}

func (msg *JsonMessage) IsError() bool {
	return msg.Error != nil
}

func (msg *JsonMessage) String() string {
	b, _ := json.Marshal(msg)
	return string(b)
}

func (msg *JsonMessage) ErrorResponse(err Error) *JsonMessage {

	msg.Params = nil
	msg.Error = &jsonError{
		Code:    err.ErrorCode(),
		Message: err.Error(),
	}

	return msg
}

func (msg *JsonMessage) Response(result interface{}) *JsonMessage {
	enc, err := json.Marshal(result)
	if err != nil {
		return msg.ErrorResponse(ParseError(err.Error()))
	}

	resp := &JsonMessage{JsonRpc: vsn, Id: msg.Id, Result: enc}

	return resp
}

var id int64 = 1
var null = json.RawMessage("null")

func GenMsgId() json.RawMessage {
	// todo: next id
	id++
	rst, _ := json.Marshal(id)
	return rst
}

func errorMessage(err Error) *JsonMessage {
	msg := &JsonMessage{JsonRpc: vsn, Id: null, Error: &jsonError{
		Code:    err.ErrorCode(),
		Message: err.Error(),
	}}
	return msg
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc err %d", err.Code)
	}
	return err.Message
}

func (err *jsonError) ErrorCode() int {
	return err.Code
}

func (err *jsonError) ErrorData() interface{} {
	return err.Data
}

////

type parseError struct {
	Msg string
}

func (e *parseError) Error() string {
	return e.Msg
}
func (e *parseError) ErrorCode() int {
	return -1002
}
func ParseError(err string) *parseError {
	return &parseError{
		Msg: err,
	}
}
