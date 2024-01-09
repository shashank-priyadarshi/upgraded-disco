package smtp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"io"
	"net"
)

type Server interface {
	Start() error
	Stop() error
}

type SMTPServer struct {
	config     SMTPServerConfig
	log        logger.Logger
	onNewEmail func(*Message) error
}

func NewSMTPServer(config SMTPServerConfig, onNewEmail func(message *Message) error, log logger.Logger) (*SMTPServer, error) {
	return &SMTPServer{
		config:     config,
		onNewEmail: onNewEmail,
		log:        log,
	}, nil
}

func (S *SMTPServer) Start() error {
	//TODO implement me
	return nil
}

func (S *SMTPServer) Stop() error {
	//TODO implement me
	return nil
}

// handleSession handles an SMTP session for an incoming connection.
func (s *SMTPServer) handleSession(conn net.Conn) {
	// TODO: Implement handling of SMTP commands during the session.
}

// handleDataCommand handles the DATA command during an SMTP session.
// https://www.rfc-editor.org/rfc/rfc5321.html
func (s *SMTPServer) handleDataCommand(conn net.Conn) {
	// TODO: Implement receiving and processing email content during the DATA command.
}

// parseEmailContent parses the received email content.
func (s *SMTPServer) parseEmailContent(reader io.Reader) (*Message, error) {
	// TODO: Implement parsing of email content.
	return nil, nil
}

// storeEmail stores the received email.
func (s *SMTPServer) storeEmail(email *Message) {
	// TODO: Implement storing or forwarding the received email.
	// You can use the onNewEmail callback or any other mechanism to handle incoming emails.
	if s.onNewEmail != nil {
		s.onNewEmail(email)
	}
}
