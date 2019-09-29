package conf

var Cfg = Config{}

type Config struct {
	ListenAddr  int
	ProjectPath string
	Email       *Email
}


type Email struct {
	Sender     string
	UserEmail  string
	Password   string
	Host       string
	Port       int
}
