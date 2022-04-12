package trie

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type RouteGroup struct {
	prefix     string
	middleware []HandleFunc
	parent     *RouteGroup
	engine     *Engine
}

type Engine struct {
	*RouteGroup
	router        *router
	groups        []*RouteGroup
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (s *RouteGroup) Use(middleware ...HandleFunc) {
	s.middleware = append(s.middleware, middleware...)
}

func (s *RouteGroup) Group(prefix string) *RouteGroup {
	engine := s.engine
	newGroup := &RouteGroup{
		prefix: s.prefix + prefix,
		parent: s,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (s *RouteGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := s.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	s.engine.router.addRoute(method, pattern, handler)
}

func (s *RouteGroup) GET(pattern string, handler HandleFunc) {
	s.addRoute("GET", pattern, handler)
}

func (s *RouteGroup) POST(pattern string, handler HandleFunc) {
	s.addRoute("POST", pattern, handler)
}

func (s *RouteGroup) Static(relativePath string, root string) {
	handler := s.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	s.GET(urlPattern, handler)
}

func (s *RouteGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(s.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		file := ctx.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(ctx.Writer, ctx.Req)
	}
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// 处理分组的中间件
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = e
	e.router.handle(c)
}
