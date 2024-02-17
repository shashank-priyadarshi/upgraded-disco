package smtp

import "github.com/shashank-priyadarshi/upgraded-disco/utils/logger"

type Client interface {
	Send(*Message) error
	SetMessageTemplate(string) error
	ConfigureServer(SMTPServerConfig) error
}

type SMTPClient struct {
	defaultConfig, customConfig SMTPServerConfig
	log                         logger.Logger
}

func NewSMTPClient(config SMTPServerConfig, log logger.Logger) (*SMTPClient, error) {
	return &SMTPClient{
		defaultConfig: config,
		log:           log,
	}, nil
}

func (S *SMTPClient) Send(message *Message) error {
	//TODO implement me
	return nil
}

func (S *SMTPClient) SetMessageTemplate(s string) error {
	//TODO implement me
	return nil
}

func (S *SMTPClient) ConfigureServer(config SMTPServerConfig) error {
	//TODO implement me
	return nil
}
