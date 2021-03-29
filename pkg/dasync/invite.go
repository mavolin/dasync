package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/mavolin/disstate/v3/pkg/state"
)

func Invite(s *state.State, code string) func() (*discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i *discord.Invite
	var err error

	go func() {
		i, err = s.Invite(code)
		wg.Done()
	}()

	return func() (*discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func InviteWithCounts(s *state.State, code string) func() (*discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i *discord.Invite
	var err error

	go func() {
		i, err = s.InviteWithCounts(code)
		wg.Done()
	}()

	return func() (*discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func ChannelInvites(s *state.State, channelID discord.ChannelID) func() ([]discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i []discord.Invite
	var err error

	go func() {
		i, err = s.ChannelInvites(channelID)
		wg.Done()
	}()

	return func() ([]discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func GuildInvites(s *state.State, guildID discord.GuildID) func() ([]discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i []discord.Invite
	var err error

	go func() {
		i, err = s.GuildInvites(guildID)
		wg.Done()
	}()

	return func() ([]discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func CreateInvite(s *state.State, channelID discord.ChannelID, data api.CreateInviteData) func() (*discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i *discord.Invite
	var err error

	go func() {
		i, err = s.CreateInvite(channelID, data)
		wg.Done()
	}()

	return func() (*discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func DeleteInvite(s *state.State, code string) func() (*discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i *discord.Invite
	var err error

	go func() {
		i, err = s.DeleteInvite(code)
		wg.Done()
	}()

	return func() (*discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}
