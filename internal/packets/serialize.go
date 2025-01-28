package packets

import (
	"encoding/json"
	"fmt"

	"github.com/TorchofFire/uRelay-guild/internal/types"
)

func (s *Service) DeserializePacket(data []byte) (interface{}, error) {
	var base BasePacket
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base packet: %w", err)
	}

	switch base.Type {
	case types.Handshake:
		var handshake Handshake
		if err := json.Unmarshal(base.Data, &handshake); err != nil {
			return nil, fmt.Errorf("failed to unmarshal handshake packet: %w", err)
		}
		return handshake, nil

	case types.GuildMessage:
		var guildMessage GuildMessage
		if err := json.Unmarshal(base.Data, &guildMessage); err != nil {
			return nil, fmt.Errorf("failed to unmarshal guild message packet: %w", err)
		}
		return guildMessage, nil

	case types.SystemMessage:
		var systemMessage SystemMessage
		if err := json.Unmarshal(base.Data, &systemMessage); err != nil {
			return nil, fmt.Errorf("failed to unmarshal system message packet: %w", err)
		}
		return systemMessage, nil

	default:
		return nil, fmt.Errorf("unknown packet type: %v", base.Type)
	}
}
