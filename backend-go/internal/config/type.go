package config

const (
	DbIslamy = "islamy"
)

var (
	DbList = []string{DbIslamy}
)

type (
	Config struct {
		Server   ServerConfig
		Database map[string]DatabaseConfig
	}

	ServerConfig struct {
		Ip          string
		Port        int
		CORSOrigins []string
		CORSHeaders []string
		CORSDebug   bool
		StdOut      string
		StdErr      string
	}

	DatabaseConfig struct {
		Driver string
		Debug  bool

		Host         string
		Port         int
		User         string
		Password     string
		DbName       string
		Schema       string
		MaxOpenConns int
		MaxIdleConns int

		SecHost         string
		SecPort         int
		SecUser         string
		SecPassword     string
		SecDbname       string
		SecSchema       string
		SecMaxOpenConns int
		SecMaxIdleConns int
	}
)
