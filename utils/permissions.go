package utils

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func CheckUPermission(s *discordgo.Session, userID string, channelID string, permission int64) bool {
	p, err := s.UserChannelPermissions(userID, channelID)
	return (err == nil) && (p&permission == permission)
}

func CheckPermission(s *discordgo.Session, userID string, channelID string, permission int64) bool {
	p, err := s.State.UserChannelPermissions(userID, channelID)
	log.Printf("Result 2: %v, bitset:%v", (p&permission == permission), p)
	return (err == nil) && (p&permission == permission)
}

func IsDiscordInvite(check string) bool {
	return regexp.MustCompile("(gg/invite|gg/)").MatchString(check)
}

// Apparently, this works?
func IsImage(check string) bool {
	regexString := "[^\\s]+(.*?)\\.(jpg|jpeg|png|gif|JPG|JPEG|PNG|GIF)$"
	return regexp.MustCompile(regexString).MatchString(check)
}
