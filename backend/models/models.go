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
	ID          string `json:"id" bson:"_id"`
	IP          string `json:"ip" bson:"ip"`
	SteamID     string `json:"steamid" bson:"steamid"`
	DiscordID   string `json:"discordid" bson:"discordid"`
	MinecraftID string `json:"minecraftid" bson:"minecraftid"`
	MiscID      string `json:"miscid" bson:"miscid"`
	Username    string `json:"username" bson:"username"`
	Reason      string `json:"reason" bson:"reason"`
	Admin       string `json:"admin" bson:"admin"`
	Game        string `json:"game" bson:"game"`
	DateBanned  string `json:"date_banned" bson:"date_banned"`
	Expires     string `json:"expires" bson:"expires"`
	ServerIP    string `json:"server_ip" bson:"server_ip"`
	ServerPort  string `json:"server_port" bson:"server_port"`
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
