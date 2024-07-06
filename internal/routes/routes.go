package routes
import "github.com/go-chi/chi/v5"

type Router struct{
	config *Config
	router *chi.Mux
}

func NewRouter() *Router{
	return &Router{
		config:  NewConfig().SetTimeOut(serviceConfig.GetConfig().timeout),
		router: chi.NewRouter(),
	}
}

func (r * Router) SetRouters() *chi.Mux{

}

func(r *Router) setConfigsRouter(){

}

func RouterHealth(){

}
func RouterProduct(){

}
func EnableCORS(){

}

func EnableTimeOut(){

}
func EnableRecovery(){

}

func EnableRequestID(){

}

func EnableRealIP(){

}

