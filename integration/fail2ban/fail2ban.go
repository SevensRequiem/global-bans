package fail2ban

import (
	"context"
	"encoding/json"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/ssh"
)

var secret string

func init() {
	secret = os.Getenv("SECRET")
}

func Ban(ip string, admin string) error {
	bandb := database.DB_Main.Collection("fail2ban_bans")
	_, err := bandb.InsertOne(context.TODO(), bson.M{
		"ip":          ip,
		"admin":       admin,
		"reason":      "Banned by Fail2Ban",
		"date_banned": time.Now(),
		"expires":     time.Now().Add(time.Hour * 72),
		"identifier":  "fail2ban",
	})
	if err != nil {
		logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
		return err
	}
	return nil
}

// un tested
func IngestFail2ban() {
	servers := database.DB_Main.Collection("fail2ban_servers")

	cursor, err := servers.Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var server models.Fail2BanServer
		if err := cursor.Decode(&server); err != nil {
			logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
			continue
		}

		config := &ssh.ClientConfig{
			User: server.Username,
			Auth: []ssh.AuthMethod{
				ssh.Password(server.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		client, err := ssh.Dial("tcp", server.IP+":"+server.Port, config)
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
			continue
		}

		session, err := client.NewSession()
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
			continue
		}
		defer session.Close()

		bannedOutput, err := session.Output("fail2ban-client banned")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
			continue
		}

		var data []map[string][]string
		err = json.Unmarshal(bannedOutput, &data)
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/fail2ban/fail2ban.go")
			continue
		}

		for _, entry := range data {
			for _, ip := range entry["sshd"] {
				Ban(ip, "fail2ban")
			}
		}
	}
}
