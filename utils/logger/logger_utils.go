package logger

import (
	"go.uber.org/zap/zapcore"
)

type ColorEncoder struct {
	encoder zapcore.Encoder
}

func NewColorEncoder(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//func (c *ColorEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer, error) {
//	buf := &buffer{}
//
//	buf.AppendString(fmt.Sprintf("\x1b[33m%s\x1b[0m", entry.Time.Format("2006-01-02T15:04:05.999Z07:00")))
//	buf.AppendString(fmt.Sprintf("\x1b[36m%s\x1b[0m", entry.Level.CapitalString()))
//	buf.AppendString(fmt.Sprintf("\x1b[32m%s\x1b[0m", entry.Caller.TrimmedPath()))
//
//	// Include additional fields
//	for _, field := range fields {
//		buf.AppendString(fmt.Sprintf("\x1b[34m%s=%v\x1b[0m", field.Key, field.Interface))
//	}
//
//	if entry.Message != "" {
//		buf.AppendString(fmt.Sprintf("\x1b[36m%s\x1b[0m", entry.Message)) // Cyan color for message
//	}
//
//	if len(fields) > 0 {
//		entryCopy := entry
//		entryCopy.Message = "" // Ensure fields are not duplicated in the message
//		encodeEntry, err := c.encoder.Clone().EncodeEntry(entryCopy, fields)
//		if err != nil {
//			return nil, err
//		}
//		buf.Append(encodeEntry.Bytes())
//	}
//
//	return buf, nil
//}
//
//type buffer struct {
//	bytes.Buffer
//}
//
//func (b *buffer) AppendByte(c byte) error {
//	return b.WriteByte(c)
//}
//
//func (b *buffer) AppendString(s string) error {
//	_, err := b.WriteString(s)
//	return err
//}
//
//func (b *buffer) Append(b2 []byte) error {
//	_, err := b.Write(b2)
//	return err
//}
//
//func (b *buffer) AppendStringByte(s string, c byte) error {
//	_, err := b.WriteString(s)
//	if err == nil {
//		err = b.WriteByte(c)
//	}
//	return err
//}
