package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

func AddRole(s *state.State, guildID discord.GuildID, userID discord.UserID, roleID discord.RoleID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.AddRole(guildID, userID, roleID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func RemoveRole(s *state.State, guildID discord.GuildID, userID discord.UserID, roleID discord.RoleID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.RemoveRole(guildID, userID, roleID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Roles(s *state.State, guildID discord.GuildID) func() ([]discord.Role, error) {
	r, err := s.Store.Roles(guildID)
	if err == nil {
		return func() ([]discord.Role, error) {
			return r, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		r, err = s.Roles(guildID)
		wg.Done()
	}()

	return func() ([]discord.Role, error) {
		wg.Wait()
		return r, err
	}
}

func CreateRole(s *state.State, guildID discord.GuildID, data api.CreateRoleData) func() (*discord.Role, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r *discord.Role
	var err error

	go func() {
		r, err = s.CreateRole(guildID, data)
		wg.Done()
	}()

	return func() (*discord.Role, error) {
		wg.Wait()
		return r, err
	}
}

func MoveRole(s *state.State, guildID discord.GuildID, data []api.MoveRoleData) func() ([]discord.Role, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r []discord.Role
	var err error

	go func() {
		r, err = s.MoveRole(guildID, data)
		wg.Done()
	}()

	return func() ([]discord.Role, error) {
		wg.Wait()
		return r, err
	}
}

func ModifyRole(s *state.State, guildID discord.GuildID, roleID discord.RoleID, data api.ModifyRoleData) func() (*discord.Role, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r *discord.Role
	var err error

	go func() {
		r, err = s.ModifyRole(guildID, roleID, data)
		wg.Done()
	}()

	return func() (*discord.Role, error) {
		wg.Wait()
		return r, err
	}
}

func DeleteRole(s *state.State, guildID discord.GuildID, roleID discord.RoleID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteRole(guildID, roleID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
