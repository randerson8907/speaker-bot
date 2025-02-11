package commands

import (
	"fmt"

	"github.com/aadithpm/speaker-bot/internal/data"
	"github.com/aadithpm/speaker-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type DungeonCommand struct {
	Name string
}

func NewDungeonCommand() (c SpeakerCommand) {
	return DungeonCommand{
		Name: Dungeon,
	}
}

func (c DungeonCommand) GetCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        Dungeon,
		Type:        discordgo.ChatApplicationCommand,
		Description: "Featured Dungeon for the week",
	}
}

func (c DungeonCommand) GetName() string {
	return c.Name
}

func (c DungeonCommand) Handler(s *discordgo.Session, d *discordgo.ApplicationCommandInteractionData) (res string, err error) {
	log.Infof("got command %v from handler", d.Name)

	dungeons := data.ReadRotationData("./data/dungeons.json")
	current_week := utils.GetTimeDifferenceInWeeks(dungeons.StartDate)
	dungeon := dungeons.ContentRotation[current_week % len(dungeons.ContentRotation)]

	str := "Featured Dungeon for this week is **%v** at %v."
	if dungeon.MasterAvailable {
		str += " Master difficulty is available!"
	}

	msg := fmt.Sprintf(
		str,
		dungeon.Name,
		dungeons.LocationList[dungeon.Location],
	)

	return msg, nil
}
