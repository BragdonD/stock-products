package stockproducts

// Config represents the configuration settings for the application.
// The configuration can be loaded from JSON or YAML files.
type Config struct {
	Http HttpConfig `json:"http" yaml:"http"`
}

// HttpConfig represents the configuration settings for an HTTP server.
type HttpConfig struct {
	Adress string       `json:"address" yaml:"address"`
	Port   int          `json:"port" yaml:"port"` // Need to change that for DNS
	Cors   *CorsConfig  `json:"cors" yaml:"cors"`
	Https  *HttpsConfig `json:"https" yaml:"https"`
}

// CorsConfig represents the configuration settings for
// Cross-Origin Resource Sharing (CORS).
type CorsConfig struct {
	AllowedOrigins []string `json:"allowedOrigins" yaml:"allowedOrigins"`
	AllowedHeaders []string `json:"allowedHeaders" yaml:"allowedHeaders"`
}

// MtlsConfig holds the configuration for MTLS communication.
type MtlsConfig struct {
	ClientsCa []string `json:"clientsCa" yaml:"clientsCa"`
}

// HttpsConfig holds the configuration for HTTPS communication.
type HttpsConfig struct {
	Cert string      `json:"cert" yaml:"cert"`
	Key  string      `json:"key" yaml:"key"`
	Mtls *MtlsConfig `json:"mtls" yaml:"mtls"`
}
