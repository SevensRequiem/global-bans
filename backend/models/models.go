package models

import "time"

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	Groups      Group     `json:"groups" bson:"groups"`
	DateCreated time.Time `json:"date_created" bson:"date_created"`
	DateBanned  time.Time `json:"date_banned" bson:"date_banned"`
	IsBanned    bool      `json:"is_banned" bson:"is_banned"`
	IsPremium   bool      `json:"is_premium" bson:"is_premium"`
	DoesExist   bool      `json:"does_exist" bson:"does_exist"`
}

type LoggedInUser struct {
	ID          string    `json:"id" bson:"_id"`
	Username    string    `json:"username" bson:"username"`
	Groups      string    `json:"groups" bson:"groups"`
	DateCreated time.Time `json:"date_created" bson:"date_created"`
	DateBanned  time.Time `json:"date_banned" bson:"date_banned"`
	IsLoggedIn  bool      `json:"is_logged_in" bson:"is_logged_in"`
	IsPremium   bool      `json:"is_premium" bson:"is_premium"`
	IsBanned    bool      `json:"is_banned" bson:"is_banned"`
}

type Group struct {
	Root    bool `json:"root" bson:"root"`
	Admin   bool `json:"admin" bson:"admin"`
	Mod     bool `json:"mod" bson:"mod"`
	Regular bool `json:"regular" bson:"regular"`
	Banned  bool `json:"banned" bson:"banned"`
}

////////////////////////////////////////

type Config struct {
	Token        string `json:"token" bson:"token"`
	ClientID     string `json:"client_id" bson:"client_id"`
	ClientSecret string `json:"client_secret" bson:"client_secret"`
	GuildID      string `json:"guild_id" bson:"guild_id"`
	ChannelID    string `json:"general_channel_id" bson:"general_channel_id"`
}

////////////////////////////////////////

type Ban struct {
	ID                  string    `json:"id" bson:"_id"`
	IP                  string    `json:"ip" bson:"ip"`
	ServerUUID          string    `json:"server_uuid" bson:"server_uuid"`
	Reason              string    `json:"reason" bson:"reason"`
	Admin               string    `json:"admin" bson:"admin"`
	DateBanned          time.Time `json:"date_banned" bson:"date_banned"`
	Expires             time.Time `json:"expires" bson:"expires"`
	Banned              bool      `json:"banned" bson:"banned,omitempty"`
	Expired             bool      `json:"expired" bson:"expired,omitempty"`
	Unbanned            bool      `json:"unbanned" bson:"unbanned,omitempty"`
	Game                string    `json:"game,omitempty" bson:"game,omitempty"`
	MinecraftPlayerUUID string    `json:"player_uuid,omitempty" bson:"player_uuid,omitempty"`
	SteamID             string    `json:"steam_id,omitempty" bson:"steam_id,omitempty"`
	Identifier          string    `json:"identifier,omitempty" bson:"identifier,omitempty"`
}

////////////////////////////////////////

type FileSync struct {
	ID       string `json:"id" bson:"_id"`
	IP       string `json:"ip" bson:"ip"`
	Port     string `json:"port" bson:"port"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	FromDir  string `json:"from_dir" bson:"from_dir"`
	ToDir    string `json:"to_dir" bson:"to_dir"`
}

////////////////////////////////////////

type Firewall struct {
	ID       string `json:"id" bson:"_id"`
	IP       string `json:"ip" bson:"ip"`
	Port     string `json:"port" bson:"port"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

////////////////////////////////////////

type RCON struct {
	ID       string `json:"id" bson:"_id"`
	IP       string `json:"ip" bson:"ip"`
	Port     string `json:"port" bson:"port"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

////////////////////////////////////////

type Server struct {
	ID          string    `json:"id" bson:"_id"`
	IP          string    `json:"ip" bson:"ip"`
	Port        string    `json:"port" bson:"port"`
	UUID        string    `json:"uuid" bson:"uuid"`
	DateCreated time.Time `json:"date_created" bson:"date_created"`
	Game        string    `json:"game" bson:"game"`
}
