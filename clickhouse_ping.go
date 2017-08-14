package clickhouse

import (
	"fmt"

	"github.com/kshvakov/clickhouse/internal/protocol"
)

func (ch *clickhouse) ping() error {
	ch.logf("-> ping")
	if err := ch.encoder.Uvarint(protocol.ClientPing); err != nil {
		return err
	}
	if err := ch.buffer.Flush(); err != nil {
		return err
	}
	packet, err := ch.decoder.Uvarint()
	if err != nil {
		return err
	}
	for {
		switch packet {
		case protocol.ServerException:
			ch.logf("[ping] <- exception")
			return ch.exception()
		case protocol.ServerPong:
			ch.logf("<- pong")
			return nil
		default:
			return fmt.Errorf("unexpected packet [%d] from server", packet)
		}
	}
}
