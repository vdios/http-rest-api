package apiserver

import(
  "github.com/sirupsen/logrus"
  "github.com/gorilla/mux"
  "net/http"
  "io"
  "github.com/vdios/http-rest-api/internal/app/store"
)

type APIServer struct{
  config *Config
  logger *logrus.Logger
  router *mux.Router
  store *store.Store
}

func New(config *Config) *APIServer{
  return &APIServer{
    config:config,
    logger:logrus.New(),
    router:mux.NewRouter(),
  }
}

//Start
func (s *APIServer) Start() error{
  if err := s.configureLogger(); err != nil{
    return err
  }

  s.configureRouter()

  if err:=s.configureStore(); err!=nil{
    return err
  }

  s.logger.Info("starting api server")
  return http.ListenAndServe(s.config.BindAddr,s.router)
}

// configure router
func (s *APIServer) configureRouter(){
  s.router.HandleFunc("/hello",s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc{
  return func(w http.ResponseWriter,r *http.Request){
    io.WriteString(w,"hello buddy =)")
  }
}

//configure logger
func (s *APIServer) configureLogger() error{
  level,err := logrus.ParseLevel(s.config.LogLevel)
  if err != nil{
    return err
  }
  s.logger.SetLevel(level)
  return nil
}

//configure store
func (s *APIServer) configureStore() error{
  st := store.New(s.config.Store)
  if err:=st.Open(); err!=nil{
    return err
  }

  s.store = st
  return nil
}
