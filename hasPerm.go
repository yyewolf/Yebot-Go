package main

import (
	"github.com/andersfylling/disgord"
)

var PERM_CREATE_INSTANT_INVITE = 0x00000001
var PERM_KICK_MEMBERS = 0x00000002
var PERM_BAN_MEMBERS = 0x00000004
var PERM_ADMINISTRATOR = 0x00000008
var PERM_MANAGE_CHANNELS = 0x00000010
var PERM_MANAGE_GUILD = 0x00000020
var PERM_ADD_REACTIONS = 0x00000040
var PERM_VIEW_AUDIT_LOG = 0x00000080
var PERM_VIEW_CHANNEL = 0x00000400
var PERM_SEND_MESSAGES = 0x00000800
var PERM_SEND_TTS_MESSAGES = 0x00001000
var PERM_MANAGE_MESSAGES = 0x00002000
var PERM_EMBED_LINKS = 0x00004000
var PERM_ATTACH_FILES = 0x00008000
var PERM_READ_MESSAGE_HISTORY = 0x00010000
var PERM_MENTION_EVERYONE = 0x00020000
var PERM_USE_EXTERNAL_EMOJIS = 0x00040000
var PERM_CONNECT = 0x00100000
var PERM_SPEAK = 0x00200000
var PERM_MUTE_MEMBERS = 0x00400000
var PERM_DEAFEN_MEMBERS = 0x00800000
var PERM_MOVE_MEMBERS = 0x01000000
var PERM_USE_VAD = 0x02000000
var PERM_PRIORITY_SPEAKER = 0x00000100
var PERM_CHANGE_NICKNAME = 0x04000000
var PERM_MANAGE_NICKNAMES = 0x08000000
var PERM_MANAGE_ROLES = 0x10000000
var PERM_MANAGE_WEBHOOKS = 0x20000000
var PERM_MANAGE_EMOJIS = 0x40000000

func hasPermission(User *disgord.Member, s disgord.Session, GuildID disgord.Snowflake, Perm int) (isHe bool) {
	for _, roleID := range User.Roles {
		g, err := s.GetGuild(GuildID)
		if err != nil {
			return false
		}
		role, err := g.Role(roleID)
		if err != nil {
			return false
		}
		if role.Permissions&0x8 == 0x8 {
			return true
		}
	}
	return false
}