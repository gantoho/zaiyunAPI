package config

import "time"

type RedisConfig struct {
	Address  string        // Redis server address (e.g., "localhost:6379" or "redis.example.com:6380")
	Password string        // Optional password for authenticating with Redis
	DB       int           // Database number to use (default is 0)
	TLS      bool          // Enable SSL/TLS connection (default is false)
	TLSCert  string        // Path to the client certificate file (optional)
	TLSKey   string        // Path to the client private key file (optional)
	TLSCA    string        // Path to the CA certificate file or directory (optional)
	PoolSize int           // Connection pool size (default is 10)
	Timeout  time.Duration // Connection and command timeout (default is 5 seconds)

	// Other options can be added here as needed
}
