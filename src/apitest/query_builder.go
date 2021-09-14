package apitest

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//)
//
//func NewQueryBuilder(cl *Client) *QueryBuilder {
//	return &QueryBuilder{
//		client: cl,
//	}
//}
//
//type QueryBuilder struct {
//	client *Client
//
//	url          string
//	requestBody  *interface{}
//	responseBody *interface{}
//}
//
//func (qb *QueryBuilder) Pathf(path string, args ...interface{}) *QueryBuilder {
//	qb.url = fmt.Sprintf(path, args...)
//	return qb
//}
//
//func (qb *QueryBuilder) WithRequestBody(b interface{}) *QueryBuilder {
//	qb.requestBody = &b
//	return qb
//}
//
//func (qb *QueryBuilder) WithResponseBody(b interface{}) *QueryBuilder {
//	qb.responseBody = &b
//	return qb
//}
//
//func (qb *QueryBuilder) Post() *Response {
//	requestBody := qb.getRequestBody()
//
//	resp := qb.client.Post(qb.url, bytes.NewReader(requestBody))
//
//	if qb.responseBody != nil {
//		if err := json.Unmarshal(resp.Body, qb.responseBody); err != nil {
//			panic(err)
//		}
//	}
//
//	return resp
//}
//
//func (qb *QueryBuilder) getRequestBody() []byte {
//	if qb.requestBody == nil {
//		return []byte("{}")
//	}
//
//	b, err := json.Marshal(*qb.requestBody)
//	if err != nil {
//		panic(err)
//	}
//
//	return b
//}
