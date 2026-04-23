package common

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// This is not a fatal error, we can rely on env vars
		log.Println("No .env file found")
	}
}

func GetConfig(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Connect() (*nats.Conn, error) {
	url := GetConfig("NATS_SERVERS", "nats://localhost:4222")
	return ConnectWithURL(url)
}

func ConnectWithURL(url string) (*nats.Conn, error) {
	user := GetConfig("NATS_USER", "")
	password := GetConfig("NATS_PASSWORD", "")

	opts := []nats.Option{
		nats.UserInfo(user, password),
		nats.Timeout(10 * time.Second),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(5 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected from NATS: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected to %s", nc.ConnectedUrl())
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

// Message Type Definitions
type DispatchMsg struct {
	Timestamp      time.Time `json:"timestamp"`
	SettlementDate string    `json:"settlementdate"`
	DispatchFlag   bool      `json:"dispatchflag"`
	DispatchCap    float64   `json:"dispatchcap"`
}

type FeedbackMsg struct {
	Timestamp            time.Time `json:"timestamp"`
	ActivePower          float64   `json:"activepower"`
	DispatchCapFeedback  float64   `json:"dispatchcapfeedback"`
	DispatchFlagFeedback bool      `json:"dispatchflagfeedback"`
}

type WatchdogMsg struct {
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
}
