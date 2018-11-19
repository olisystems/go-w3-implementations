package whisper

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

func Whisper() {

	shhClient, err := shhclient.Dial("ws://127.0.0.1:8546")

	if err != nil {
		log.Fatal(err)
	}

	_ = shhClient

	fmt.Println("we have a whisper connection")

	keyID, err := shhClient.NewKeyPair(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(keyID)

	publicKey, err := shhClient.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(hexutil.Encode(publicKey))

	message := whisperv6.NewMessage{
		Payload:   []byte("Hello"),
		PublicKey: publicKey,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Topic:     whisperv6.BytesToTopic([]byte("TEST")),
	}

	messageHash, err := shhClient.Post(context.Background(), message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messageHash)

}
