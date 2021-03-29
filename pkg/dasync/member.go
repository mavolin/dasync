package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/mavolin/disstate/v3/pkg/state"
)

func Member(s *state.State, guildID discord.GuildID, userID discord.UserID) func() (*discord.Member, error) {
	m, err := s.Cabinet.Member(guildID, userID)
	if err == nil {
		return func() (*discord.Member, error) {
			return m, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		m, err = s.Member(guildID, userID)
		wg.Done()
	}()

	return func() (*discord.Member, error) {
		wg.Wait()
		return m, err
	}
}

// Members will not call the state first, as the result, depending on the type
// of caching and the shard configuration, may differ from the API response.
func Members(s *state.State, guildID discord.GuildID, limit uint) func() ([]discord.Member, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Member
	var err error

	go func() {
		m, err = s.Client.Members(guildID, limit)
		wg.Done()
	}()

	return func() ([]discord.Member, error) {
		wg.Wait()
		return m, err
	}
}

func MembersAfter(s *state.State, guildID discord.GuildID, after discord.UserID, limit uint) func() ([]discord.Member, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Member
	var err error

	go func() {
		m, err = s.MembersAfter(guildID, after, limit)
		wg.Done()
	}()

	return func() ([]discord.Member, error) {
		wg.Wait()
		return m, err
	}
}

func AddMember(
	s *state.State, guildID discord.GuildID, userID discord.UserID, data api.AddMemberData,
) func() (*discord.Member, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Member
	var err error

	go func() {
		m, err = s.AddMember(guildID, userID, data)
		wg.Done()
	}()

	return func() (*discord.Member, error) {
		wg.Wait()
		return m, err
	}
}

func ModifyMember(s *state.State, guildID discord.GuildID, userID discord.UserID, data api.ModifyMemberData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.ModifyMember(guildID, userID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func PruneCount(s *state.State, guildID discord.GuildID, data api.PruneCountData) func() (uint, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var count uint
	var err error

	go func() {
		count, err = s.PruneCount(guildID, data)
		wg.Done()
	}()

	return func() (uint, error) {
		wg.Wait()
		return count, err
	}
}

func Prune(s *state.State, guildID discord.GuildID, data api.PruneData) func() (uint, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var count uint
	var err error

	go func() {
		count, err = s.Prune(guildID, data)
		wg.Done()
	}()

	return func() (uint, error) {
		wg.Wait()
		return count, err
	}
}

func Kick(s *state.State, guildID discord.GuildID, userID discord.UserID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Kick(guildID, userID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Bans(s *state.State, guildID discord.GuildID) func() ([]discord.Ban, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var b []discord.Ban
	var err error

	go func() {
		b, err = s.Bans(guildID)
		wg.Done()
	}()

	return func() ([]discord.Ban, error) {
		wg.Wait()
		return b, err
	}
}

func GetBan(s *state.State, guildID discord.GuildID, userID discord.UserID) func() (*discord.Ban, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var b *discord.Ban
	var err error

	go func() {
		b, err = s.GetBan(guildID, userID)
		wg.Done()
	}()

	return func() (*discord.Ban, error) {
		wg.Wait()
		return b, err
	}
}

func Ban(s *state.State, guildID discord.GuildID, userID discord.UserID, data api.BanData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Ban(guildID, userID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Unban(s *state.State, guildID discord.GuildID, userID discord.UserID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Unban(guildID, userID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
