package connections

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/nacl/sign"
)

func unlockAndVerifySignedMessage(publicKey, encryptedMessage string, timestampTolerance int) (string, error) {
	pubKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode publicKey: %w", err)
	}
	encMessage, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", fmt.Errorf("failed to decode encryptedMessage: %w", err)
	}

	openedMessage, ok := sign.Open(nil, encMessage, (*[32]byte)(pubKey))
	if !ok {
		return "", errors.New("failed to verify signed message")
	}
	unlockedMessage := string(openedMessage)
	parts := strings.SplitN(unlockedMessage, "|", 2)
	if len(parts) < 2 {
		return "", errors.New("invalid message format: expected 'timestamp|message'")
	}

	stringTimestamp, payload := parts[0], parts[1]

	timestamp, err := strconv.ParseInt(stringTimestamp, 10, 64)
	if err != nil {
		return "", errors.New("invalid timestamp: expected a unix timestamp")
	}

	messageTime := time.Unix(timestamp, 0)
	if time.Since(messageTime).Seconds() > float64(timestampTolerance) {
		return "", errors.New("timestamp falls outside the expected range; check your device clock")
	}

	return payload, nil
}
