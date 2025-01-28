package connections

import (
	"log"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/internal/guild"
	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
	"github.com/gorilla/websocket"
)

type Service struct {
	guild   *guild.Service
	packets *packets.Service
}

func NewService(guild *guild.Service, packets *packets.Service) *Service {
	s := &Service{guild: guild, packets: packets}
	return s
}

func (s *Service) Handler(writer http.ResponseWriter, request *http.Request) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	var userId uint64
	defer func() {
		defer conn.Close()
		if userId != 0 {
			removeConnection(userId)
		}
	}()

	firstPacketRecieved := false

	for {
		messageType, packet, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		if messageType != websocket.TextMessage {
			return
		}

		deserializedPacket, err := s.packets.DeserializePacket(packet)
		if err != nil {
			log.Println(err)
			sendSystemMessageViaConn(conn, types.Danger, "Unrecognized packet:", 0)
			return
		}

		if !firstPacketRecieved {
			handshakePacket, ok := deserializedPacket.(packets.Handshake)
			if !ok {
				conn.Close()
				return
			}

			var err error
			userId, err = s.handshake(handshakePacket)
			if err != nil {
				sendSystemMessageViaConn(conn, types.Danger, err.Error(), 0)
				conn.Close()
				return
			}

			addNewConnection(userId, conn)

			sendSystemMessageViaConn(conn, types.Info, "Connected", 0)

			firstPacketRecieved = true
			continue
		}

		if userId == 0 {
			sendSystemMessageViaConn(conn, types.Danger, "Connection closed due to server error: User Id unknown", 0)
			conn.Close()
			return
		}

		switch p := deserializedPacket.(type) {
		case packets.GuildMessage:
			s.handleGuildMessage(conn, p)
		case packets.Handshake:
			sendSystemMessageViaConn(conn, types.Warning, "Did not expect a handshake packet", 0)
		case packets.SystemMessage:
			sendSystemMessageViaConn(conn, types.Warning, "Did not expect a system message packet", 0)
		default:
			log.Fatal("A deserialized and known packet was not handled")
		}
	}
}
