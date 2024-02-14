package req

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func NewResponseHandler(r *http.Response, op *ResponseOperationsQueue) (*ResponseHandler, error) {
	res := &ResponseHandler{}
	res.Response = r
	res.performOperations(op)

	return res, nil
}

func NewResponseOperationsQueue() *ResponseOperationsQueue {
	op := &ResponseOperationsQueue{}
	return op
}

func (r *ResponseOperationsQueue) addOperation(op ResponseOperation) {
	r.op = append(r.op, op)
}

func (r *ResponseHandler) performOperations(queue *ResponseOperationsQueue) {
	for _, f := range queue.op {
		f(r)
	}
}

func (r *ResponseOperationsQueue) queueDumpHeader(path string) ResponseOperation {
	fmt.Println("we get here?", len(path))
	return func(rh *ResponseHandler) error {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

		if err != nil {
			panic(err)
		}

		defer f.Close()
		if _, err = f.WriteString(mapToString(rh.Response.Header)); err != nil {
			panic(err)
		}
		return nil
	}
}

func mapToString(m map[string][]string) string {
	sb := strings.Builder{}
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString(":")

		for _, vv := range v {
			sb.WriteString(vv)
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.String()
}
