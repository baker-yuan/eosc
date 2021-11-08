package http

type IFilter interface {
	DoFilter(ctx IHttpContext, next IFilterChain) (err error)
}
type IFilterChain interface {
	DoFilter(ctx IHttpContext) error
}
