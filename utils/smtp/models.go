package smtp

type Message struct {
	sender     string
	recipients []string
	subject    string
	body       string
}

type SMTPServerConfig struct {
	Address      string
	Port         int
	RequireAuth  bool
	AuthUsername string
	AuthPassword string
}
