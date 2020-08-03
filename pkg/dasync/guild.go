package dasync

import (
	"io"
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

func CreateGuild(s *state.State, data api.CreateGuildData) func() (*discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g *discord.Guild
	var err error

	go func() {
		g, err = s.CreateGuild(data)
		wg.Done()
	}()

	return func() (*discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func Guild(s *state.State, id discord.GuildID) func() (*discord.Guild, error) {
	g, err := s.Store.Guild(id)
	if err == nil {
		return func() (*discord.Guild, error) {
			return g, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		g, err = s.Guild(id)
		wg.Done()
	}()

	return func() (*discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func GuildPreview(s *state.State, id discord.GuildID) func() (*discord.GuildPreview, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g *discord.GuildPreview
	var err error

	go func() {
		g, err = s.GuildPreview(id)
		wg.Done()
	}()

	return func() (*discord.GuildPreview, error) {
		wg.Wait()
		return g, err
	}
}

func GuildWithCount(s *state.State, id discord.GuildID) func() (*discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g *discord.Guild
	var err error

	go func() {
		g, err = s.GuildWithCount(id)
		wg.Done()
	}()

	return func() (*discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

// Guilds will not call the state first, as the result, depending on the type
// of caching and the shard configuration, may differ from the API response.
func Guilds(s *state.State, limit uint) func() ([]discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g []discord.Guild
	var err error

	go func() {
		g, err = s.Client.Guilds(limit)
		wg.Done()
	}()

	return func() ([]discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func GuildsBefore(s *state.State, before discord.GuildID, limit uint) func() ([]discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g []discord.Guild
	var err error

	go func() {
		g, err = s.GuildsBefore(before, limit)
		wg.Done()
	}()

	return func() ([]discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func GuildsAfter(s *state.State, after discord.GuildID, limit uint) func() ([]discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g []discord.Guild
	var err error

	go func() {
		g, err = s.GuildsAfter(after, limit)
		wg.Done()
	}()

	return func() ([]discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func LeaveGuild(s *state.State, id discord.GuildID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.LeaveGuild(id)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func ModifyGuild(s *state.State, id discord.GuildID, data api.ModifyGuildData) func() (*discord.Guild, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var g *discord.Guild
	var err error

	go func() {
		g, err = s.ModifyGuild(id, data)
		wg.Done()
	}()

	return func() (*discord.Guild, error) {
		wg.Wait()
		return g, err
	}
}

func DeleteGuild(s *state.State, id discord.GuildID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteGuild(id)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func VoiceRegionsGuild(s *state.State, guildID discord.GuildID) func() ([]discord.VoiceRegion, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r []discord.VoiceRegion
	var err error

	go func() {
		r, err = s.VoiceRegionsGuild(guildID)
		wg.Done()
	}()

	return func() ([]discord.VoiceRegion, error) {
		wg.Wait()
		return r, err
	}
}

func AuditLog(s *state.State, guildID discord.GuildID, data api.AuditLogData) func() (*discord.AuditLog, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var al *discord.AuditLog
	var err error

	go func() {
		al, err = s.AuditLog(guildID, data)
		wg.Done()
	}()

	return func() (*discord.AuditLog, error) {
		wg.Wait()
		return al, err
	}
}

func Integrations(s *state.State, guildID discord.GuildID) func() ([]discord.Integration, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i []discord.Integration
	var err error

	go func() {
		i, err = s.Integrations(guildID)
		wg.Done()
	}()

	return func() ([]discord.Integration, error) {
		wg.Wait()
		return i, err
	}
}

func AttachIntegration(
	s *state.State, guildID discord.GuildID, integrationID discord.IntegrationID, integrationType discord.Service,
) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.AttachIntegration(guildID, integrationID, integrationType)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func ModifyIntegration(
	s *state.State, guildID discord.GuildID, integrationID discord.IntegrationID, data api.ModifyIntegrationData,
) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.ModifyIntegration(guildID, integrationID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func SyncIntegration(s *state.State, guildID discord.GuildID, integrationID discord.IntegrationID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.SyncIntegration(guildID, integrationID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func GuildWidget(s *state.State, guildID discord.GuildID) func() (*discord.GuildWidget, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.GuildWidget
	var err error

	go func() {
		w, err = s.GuildWidget(guildID)
		wg.Done()
	}()

	return func() (*discord.GuildWidget, error) {
		wg.Wait()
		return w, err
	}
}

func ModifyGuildWidget(
	s *state.State, guildID discord.GuildID, data api.ModifyGuildWidgetData,
) func() (*discord.GuildWidget, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.GuildWidget
	var err error

	go func() {
		w, err = s.ModifyGuildWidget(guildID, data)
		wg.Done()
	}()

	return func() (*discord.GuildWidget, error) {
		wg.Wait()
		return w, err
	}
}

func GuildVanityURL(s *state.State, guildID discord.GuildID) func() (*discord.Invite, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var i *discord.Invite
	var err error

	go func() {
		i, err = s.GuildVanityURL(guildID)
		wg.Done()
	}()

	return func() (*discord.Invite, error) {
		wg.Wait()
		return i, err
	}
}

func GuildImage(s *state.State, guildID discord.GuildID, img api.GuildImageStyle) func() (io.ReadCloser, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var rc io.ReadCloser
	var err error

	go func() {
		rc, err = s.GuildImage(guildID, img)
		wg.Done()
	}()

	return func() (io.ReadCloser, error) {
		wg.Wait()
		return rc, err
	}
}
