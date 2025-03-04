package config

func Parse(cfgDir string) *Config {
	dir := getDirPath(cfgDir)
	dbConfig := parseDBConfig(dir)
	return &Config{
		Service: parseServiceConfig(dir),
		Logs:    parseLoggerConfig(dir),
		DB:      dbConfig,
		Redis:   dbConfig.Redis,
	}
}

func parseServiceConfig(dir string) ServiceConfig {
	cfg := ServiceConfig{}
	parseConfig(dir+"service.yaml", &cfg)
	validateAppConfig(&cfg)
	return cfg
}

func validateAppConfig(cfg *ServiceConfig) {
	if cfg.ServiceName == "" {
		panic("Name is empty in service.yaml")
	}
	if cfg.Port == 0 {
		panic("Port is empty in service.yaml")
	}
	if cfg.RequestPerSecond == 0 {
		cfg.RequestPerSecond = 100
	}
}
func parseLoggerConfig(dir string) LoggerConfig {
	log := LoggerConfig{}
	parseLogger(dir+"logger.yaml", &log)
	return log
}
func parseDBConfig(dir string) DBConfig {
	dbg := DBConfig{}
	parseConfig(dir+"database.yaml", &dbg)
	return dbg
}
