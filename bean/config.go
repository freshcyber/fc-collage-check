package bean

// ServerConfig ServerConfig
type ServerConfig struct {
	Exec     bool
	Host     string
	Decimal  int64
	Limit    int
	Spec     string
	Urlgo    string
	Recharge string
	Callback string
	Salt     string
	APIAppendKey string
	APIMd5Key    string
}

// DBConfig DBConfig Struct
type DBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}
