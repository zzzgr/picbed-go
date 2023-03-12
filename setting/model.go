package setting

type Application struct {
	Server `yaml:"server"`
	Config `yaml:"config"`
}

type Server struct {
	Name   string `yaml:"name"`
	Port   int    `yaml:"port"`
	Notice string `yaml:"notice"`
	Mode   string `yaml:"mode"`
}

type Config struct {
	BliCookie string `yaml:"bliCookie"`
}
