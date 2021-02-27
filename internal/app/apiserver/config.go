package apiserver

type Config struct{
  //fiedl with something called "tag"
  BindAddr string `toml:"bind_addr"`
}

func NewConfig() *Config{
  return &Config{
    BindAddr: ":8080",
  }
}
