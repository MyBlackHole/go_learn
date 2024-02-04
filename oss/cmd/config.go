package oss

type Config struct {
	StoreRoot  string
	Host       string
	Hostname   []string
	ConfigInit bool
	QuietMode  bool
}

var GConfig *Config

func NewConfig() *Config {
	// Log.init
	return &Config{
		StoreRoot:  STORE_ROOT_DIR,
		Host:       HOST,
		Hostname:   HOSTNAMES,
		ConfigInit: true,
		QuietMode:  true,
	}
}

func (config *Config) SetQuietMode(quietMode bool) {
	config.QuietMode = quietMode
	// Log.set_quiet_mode(mode)
}

func (config *Config) SetLogLevel(logLevel string) {
	// Log.set_log_level(level)
}

func (config *Config) SetStore(storeRoot string) {
	config.StoreRoot = storeRoot
}

func (config *Config) GetStore() string {
	return config.StoreRoot
}

func (config *Config) SetHostname(host string) {
	config.Host = host
	config.Hostname = append(config.Hostname, host)
}

func (config *Config) GetHost() string {
	return config.Host
}

func (config *Config) GetHostname() []string {
	return config.Hostname
}
