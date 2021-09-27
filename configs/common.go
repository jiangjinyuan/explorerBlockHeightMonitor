// Package configs 配置操作公共类
package configs

// MySQLDSN 代表一个 mysql dsn 连接信息
type MySQLDSN struct {
	Name    string
	DSN     string
	Type    string
	SSHName string
}

// MySQLDB 代表一个 mysql 连接信息
type MySQLDB struct {
	Read     MySQLDSN
	Write    MySQLDSN
	Timezone string
	Region   string
	CoinType string
}

// RedisDB 代表一个 redis 连接信息
type RedisDB struct {
	Address  string
	Password string
}

// RedisClusterDB 代表一个 redis 集群连接信息
type RedisClusterDB struct {
	Address  []string
	Password string
}

// RPCClient 代表一个 rpc 连接信息
type RPCClient struct {
	Host string
	Port int64
	User string
	Pass string
}

// Log 代表日志配置
type Log struct {
	Level      string
	ShowSource bool
	Path       string
}

type ConfigURL struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int64  `json:"port"`
}

// AddDatabaseConfig 添加数据库配置
func AddDatabaseConfig(value *MySQLDB, configs map[string]MySQLDSN) {
	if value.Read.DSN != "" && value.Read.Name != "" {
		configs[value.Read.Name] = MySQLDSN{DSN: value.Read.DSN, SSHName: value.Read.SSHName, Type: value.Read.Type}
	}
	if value.Write.DSN != "" && value.Write.Name != "" {
		configs[value.Write.Name] = MySQLDSN{DSN: value.Write.DSN, SSHName: value.Read.SSHName, Type: value.Write.Type}
	}
}
