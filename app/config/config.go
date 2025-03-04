package config

type Config struct {
	Service ServiceConfig
	Logs    LoggerConfig
	DB      DBConfig
}
type LogConfig struct {
	Level    string      `yaml:"level"`
	File     LogFile     `yaml:"file"`
	Rotation LogRotation `yaml:"rotation"`
}
type ServiceConfig struct {
	ServiceName      string `yaml:"name"`
	BaseURL          string `yaml:"base_url"`
	TracingExporter  string `yaml:"tracing_exporter"`
	AppEnv           string `yaml:"app_env"`
	Port             int    `yaml:"port"`
	RequestPerSecond int    `yaml:"rps"`
	Timeout          int    `yaml:"timeout"`
	EncrytionKey     string `yaml:"encryption_key"`
	JWTSecret        string `yaml:"jwt_secret"`
}
type DBConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Database        string `yaml:"database"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	PoolSize        int    `yaml:"pool_size"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}
type LogRotation struct {
	MaxSize    int  `yaml:"max_size"`
	MaxBackups int  `yaml:"max_backups"`
	MaxAge     int  `yaml:"max_age"`
	Compress   bool `yaml:"compress"`
}

type LogFile struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type LoggerConfig struct {
	Logs []LogConfig `yaml:"logs"`
}
