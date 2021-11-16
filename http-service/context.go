package http_service

import (
	"context"
	"net/http"
	"net/textproto"
	"net/url"
	"time"
)

type IHttpContext interface {
	RequestId() string

	Context() context.Context
	Value(key interface{}) interface{}
	WithValue(key, val interface{})

	Request() IRequestReader      // 读取原始请求
	Proxy() IRequest              // 读写转发请求
	Response() (IResponse, error) // 处理返回结果，可读可写

	SendTo(address string, timeout time.Duration) error
}

type IHeaderReader interface {
	GetHeader(name string) string
	Headers() http.Header
}

type IHeaderWriter interface {
	SetHeader(key, value string)
	AddHeader(key, value string)
	DelHeader(key string)
}

type IHeader interface {
	IHeaderReader
	IHeaderWriter
}

type IBodyGet interface {
	GetBody() []byte
}

type IBodySet interface {
	SetBody([]byte)
}

type FileHeader struct {
	FileName string
	Header   textproto.MIMEHeader
	Data     []byte
}

type IBodyDataReader interface {
	//protocol() RequestType
	ContentType() string
	//content-Type = application/x-www-form-urlencoded 或 multipart/form-data，与原生request.Form不同，这里不包括 query 参数
	BodyForm() (url.Values, error)
	//content-Type = multipart/form-data 时有效
	Files() (map[string]*FileHeader, error)
	GetForm(key string) string
	GetFile(key string) (file *FileHeader, has bool)
	RawBody() ([]byte, error)
	//encoder()[]byte // 最终数据
}

type IBodyDataWriter interface {
	//设置form数据并将content-type设置 为 application/x-www-form-urlencoded 或 multipart/form-data
	SetForm(values url.Values) error
	SetToForm(key, value string) error
	AddForm(key, value string) error
	// 会替换掉对应掉file信息，并且将content-type 设置为 multipart/form-data
	AddFile(key string, file *FileHeader) error
	//设置 multipartForm 数据并将content-type设置 为 multipart/form-data
	// 重置body，会清除掉未处理掉 form和file
	SetRaw(contentType string, body []byte)
}

type IStatusGet interface {
	StatusCode() int
	Status() string
}

type IStatusSet interface {
	SetStatus(code int, status string)
}

// 原始请求数据的读
type IRequestReader interface {
	IHeaderReader
	IBodyDataReader
	Method() string
	URL() url.URL
	RequestURI() string
	Host() string
	RemoteAddr() string
	Scheme() string
}

// 用于组装转发的request
type IRequest interface {
	IRequestReader
	IHeaderWriter
	IBodyDataWriter
	TargetServer() string
	TargetURL() string
	SetMethod(string)
	SetURL(url.URL)
	SetPath(string)
}

// 返回给client端的
type IResponse interface {
	IHeaderReader
	IBodyGet
	IStatusGet
	IHeaderWriter
	IStatusSet // 设置返回状态
	IBodySet   // 设置返回内容
}
