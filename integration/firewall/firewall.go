package firewall

import (
	"context"
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

func BanIPTABLES(ip, expires string) {
	servers := database.DB_Main.Collection("firewall_servers")

	cursor, err := servers.Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
	}
	defer cursor.Close(context.TODO())

	bans := database.DB_Main.Collection("firewall_bans")

	for cursor.Next(context.TODO()) {
		var server models.Firewall
		if err := cursor.Decode(&server); err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
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
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		session, err := client.NewSession()
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}
		defer session.Close()

		err = session.Run("iptables -A INPUT -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		err = session.Run("iptables -A OUTPUT -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		err = session.Run("iptables -A FORWARD -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		parsedExpires, err := time.Parse("2006-01-02", expires)
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		_, err = bans.InsertOne(context.TODO(), models.Ban{
			IP:         ip,
			ServerUUID: server.IP,
			Expires:    parsedExpires,
		})

		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		logs.LogInfo("Banned "+ip+" from iptables @"+server.IP, 0, "integrations/firewall/firewall.go")

	}
}

func UnbanIPTABLES(ip string) {
	servers := database.DB_Main.Collection("firewall_servers")

	cursor, err := servers.Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
	}
	defer cursor.Close(context.TODO())

	bans := database.DB_Main.Collection("firewall_bans")

	for cursor.Next(context.TODO()) {
		var server models.Firewall
		if err := cursor.Decode(&server); err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
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
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		session, err := client.NewSession()
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}
		defer session.Close()

		err = session.Run("iptables -D INPUT -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		err = session.Run("iptables -D OUTPUT -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		err = session.Run("iptables -D FORWARD -s " + ip + " -j DROP")
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		_, err = bans.DeleteOne(context.TODO(), bson.M{"ip": ip})
		if err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		logs.LogInfo("Unbanned "+ip+" from iptables @"+server.IP, 0, "integrations/firewall/firewall.go")
	}
}

func ExpireCheck() {
	bans := database.DB_Main.Collection("firewall_bans")

	cursor, err := bans.Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var ban models.Ban
		if err := cursor.Decode(&ban); err != nil {
			logs.LogError(err.Error(), 0, "integrations/firewall/firewall.go")
			return
		}

		if time.Now().After(ban.Expires) {
			UnbanIPTABLES(ban.IP)
		}
	}
}
