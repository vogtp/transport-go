package utils

var PlatformServerFlagConstants = map[string]map[string]string{
	"Hostname": {
		"FlagName":    "hostname",
		"Description": "Hostname where Plank accepts connections",
	},
	"Port": {
		"FlagName":    "port",
		"Description": "Port where Plank is to be served",
	},
	"RootDir": {
		"FlagName":    "root-dir",
		"Description": "Root directory for the server (default: Current directory)",
	},
	"Cert": {
		"FlagName":    "cert",
		"Description": "X509 Certificate file for TLS",
	},
	"CertKey": {
		"FlagName":    "cert-key",
		"Description": "X509 Certificate private Key file for TLS",
	},
	"Static": {
		"FlagName":    "static",
		"Description": "Path(s) where static files will be served",
	},
	"SpaPath": {
		"FlagName":    "spa-path",
		"Description": "Path to serve Single Page Application (SPA) from. The URI is derived from the leaf directory. A different URI can be specified by providing it following a colon (e.g. --spa-path ./path/to/spa-app:my-spa",
	},
	"NoFabricBroker": {
		"FlagName":    "no-fabric-broker",
		"Description": "Disable Fabric (STOMP) broker",
	},
	"FabricEndpoint": {
		"FlagName":    "fabric-endpoint",
		"Description": "Fabric broker endpoint",
	},
	"TopicPrefix": {
		"FlagName":    "topic-prefix",
		"Description": "Topic prefix for Fabric broker",
	},
	"QueuePrefix": {
		"FlagName":    "query-prefix",
		"Description": "Queue prefix for Fabric broker",
	},
	"RequestPrefix": {
		"FlagName":    "request-prefix",
		"Description": "Application request prefix for Fabric broker",
	},
	"RequestQueuePrefix": {
		"FlagName":    "request-queue-prefix",
		"Description": "Application request queue prefix for Fabric broker",
	},
	"ConfigFile": {
		"FlagName":    "config-file",
		"Description": "Path to the server config JSON file",
	},
	"ShutdownTimeout": {
		"FlagName":    "shutdown-timeout",
		"Description": "Graceful server shutdown timeout in minutes",
	},
	"OutputLog": {
		"FlagName":    "output-log",
		"Description": "Platform log output",
	},
	"AccessLog": {
		"FlagName":    "access-log",
		"Description": "HTTP server access log output",
	},
	"ErrorLog": {
		"FlagName":    "error-log",
		"Description": "HTTP server error log output",
	},
	"Debug": {
		"FlagName":    "debug",
		"Description": "Enable debug logging",
	},
	"NoBanner": {
		"FlagName":    "no-banner",
		"Description": "Do not print Plank banner at startup",
	},
	"Prometheus": {
		"FlagName":    "prometheus",
		"Description": "Enable Prometheus for basic runtime metrics",
	},
}