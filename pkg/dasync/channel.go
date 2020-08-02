package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

func Channels(s *state.State, guildID discord.GuildID) func() ([]discord.Channel, error) {
	c, err := s.Store.Channels(guildID)
	if err == nil {
		return func() ([]discord.Channel, error) {
			return c, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c, err = s.Channels(guildID)
		wg.Done()
	}()

	return func() ([]discord.Channel, error) {
		wg.Wait()
		return c, err
	}
}

func CreateChannel(
	s *state.State, guildID discord.GuildID, data api.CreateChannelData,
) func() (*discord.Channel, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var c *discord.Channel
	var err error

	go func() {
		c, err = s.CreateChannel(guildID, data)
		wg.Done()
	}()

	return func() (*discord.Channel, error) {
		wg.Wait()
		return c, err
	}
}

func MoveChannel(s *state.State, guildID discord.GuildID, data []api.MoveChannelData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.MoveChannel(guildID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Channel(s *state.State, id discord.ChannelID) func() (*discord.Channel, error) {
	c, err := s.Store.Channel(id)
	if err == nil {
		return func() (*discord.Channel, error) {
			return c, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c, err = s.Channel(id)
		wg.Done()
	}()

	return func() (*discord.Channel, error) {
		wg.Wait()
		return c, err
	}
}

func ModifyChannel(s *state.State, id discord.ChannelID, data api.ModifyChannelData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.ModifyChannel(id, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteChannel(s *state.State, id discord.ChannelID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteChannel(id)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func EditChannelPermission(
	s *state.State, channelID discord.ChannelID, overwriteID discord.Snowflake, data api.EditChannelPermissionData,
) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.EditChannelPermission(channelID, overwriteID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteChannelPermission(s *state.State, channelID discord.ChannelID, overwriteID discord.Snowflake) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteChannelPermission(channelID, overwriteID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Typing(s *state.State, channelID discord.ChannelID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Typing(channelID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func PinnedMessages(s *state.State, channelID discord.ChannelID) func() ([]discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var p []discord.Message
	var err error

	go func() {
		p, err = s.PinnedMessages(channelID)
		wg.Done()
	}()

	return func() ([]discord.Message, error) {
		wg.Wait()
		return p, err
	}
}

func PinMessage(s *state.State, channelID discord.ChannelID, messageID discord.MessageID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.PinMessage(channelID, messageID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func UnpinMessage(s *state.State, channelID discord.ChannelID, messageID discord.MessageID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.UnpinMessage(channelID, messageID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func AddRecipient(
	s *state.State, channelID discord.ChannelID, userID discord.UserID, accessToken, nickname string,
) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.AddRecipient(channelID, userID, accessToken, nickname)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func RemoveRecipient(s *state.State, channelID discord.ChannelID, userID discord.UserID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.RemoveRecipient(channelID, userID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Ack(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, ack *api.Ack) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Ack(channelID, messageID, ack)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
