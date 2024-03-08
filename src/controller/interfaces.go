package controller

type ApiController interface {
	RouteRegisterer()
}

type RouteRegisterer interface {
	RegisterRoutes()
}

type Prefix func(prefix string) string
