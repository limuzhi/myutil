/*
 * @PackageName: errors
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 15:11
 */

package errors

import (
	"fmt"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
)

func NewHTTPError(code int, reason string, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Reason:  reason,
		Message: message,
		Ts:      time.Now().UTC().Unix(),
	}
}

type HTTPError struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Ts      int64  `json:"ts"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		se.Ts = time.Now().UTC().Unix()
		return se
	}
	if se := new(errors.Error); errors.As(err, &se) {
		return NewHTTPError(int(se.Code), se.Reason, se.Message)
	}
	gs, ok := status.FromError(err)
	if !ok {
		NewHTTPError(errors.UnknownCode, errors.UnknownReason, err.Error())
	}
	ret := NewHTTPError(
		httpstatus.FromGRPCCode(gs.Code()),
		errors.UnknownReason,
		gs.Message(),
	)
	return ret
}
