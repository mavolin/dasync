package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/mavolin/disstate/v3/pkg/state"
)

func User(s *state.State, id discord.UserID) func() (*discord.User, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var u *discord.User
	var err error

	go func() {
		u, err = s.User(id)
		wg.Done()
	}()

	return func() (*discord.User, error) {
		wg.Wait()
		return u, err
	}
}

func Me(s *state.State) func() (*discord.User, error) {
	me, err := s.Cabinet.Me()
	if err == nil {
		return func() (*discord.User, error) {
			return me, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		me, err = s.Me()
		wg.Done()
	}()

	return func() (*discord.User, error) {
		wg.Wait()
		return me, err
	}
}

func ModifyMe(s *state.State, data api.ModifySelfData) func() (*discord.User, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var u *discord.User
	var err error

	go func() {
		u, err = s.ModifyMe(data)
		wg.Done()
	}()

	return func() (*discord.User, error) {
		wg.Wait()
		return u, err
	}
}

func ChangeOwnNickname(s *state.State, guildID discord.GuildID, nick string) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.ChangeOwnNickname(guildID, nick)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func PrivateChannels(s *state.State) func() ([]discord.Channel, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var c []discord.Channel
	var err error

	go func() {
		c, err = s.PrivateChannels()
		wg.Done()
	}()

	return func() ([]discord.Channel, error) {
		wg.Wait()
		return c, err
	}
}

func CreatePrivateChannel(s *state.State, recipientID discord.UserID) func() (*discord.Channel, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var c *discord.Channel
	var err error

	go func() {
		c, err = s.CreatePrivateChannel(recipientID)
		wg.Done()
	}()

	return func() (*discord.Channel, error) {
		wg.Wait()
		return c, err
	}
}

func UserConnections(s *state.State) func() ([]discord.Connection, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var c []discord.Connection
	var err error

	go func() {
		c, err = s.UserConnections()
		wg.Done()
	}()

	return func() ([]discord.Connection, error) {
		wg.Wait()
		return c, err
	}
}

func SetNote(s *state.State, userID discord.UserID, note string) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.SetNote(userID, note)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func SetRelationship(s *state.State, userID discord.UserID, t discord.RelationshipType) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.SetRelationship(userID, t)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteRelationship(s *state.State, userID discord.UserID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteRelationship(userID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
