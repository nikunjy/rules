package parser

import (
	"encoding/json"
	"fmt"
)

type ErrVals map[string]interface{}

func (e ErrVals) Dupe() ErrVals {
	ret := make(ErrVals)
	for k, v := range e {
		ret[k] = v
	}
	return ret

}
func (e ErrVals) Merge(vals ErrVals) {
	for k, v := range vals {
		e[k] = v
	}
}

type DebugError struct {
	Err  error
	Msg  string
	Vals ErrVals
}

func (e *DebugError) Original() error {
	switch val := e.Err.(type) {
	case *DebugError:
		return val.Original()
	default:
		return e.Err
	}
}

func (e *DebugError) Error() string {
	if e.Vals == nil {
		e.Vals = make(map[string]interface{})
	}
	e.Vals["err"] = e.Err.Error()
	e.Vals["msg"] = e.Msg
	data, err := json.Marshal(e.Vals.Dupe())
	if err != nil {
		return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
	}
	return string(data)
}

func (e *DebugError) Set(vals ErrVals) *DebugError {
	if e.Vals == nil {
		e.Vals = make(ErrVals)
	}
	e.Vals.Merge(vals.Dupe())
	return e
}

func newDebugError(err error, msg string) *DebugError {
	return &DebugError{
		Err: err,
		Msg: msg,
	}
}
