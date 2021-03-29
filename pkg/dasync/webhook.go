package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/mavolin/disstate/v3/pkg/state"
)

func CreateWebhook(s *state.State, channelID discord.ChannelID, data api.CreateWebhookData) func() (*discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.Webhook
	var err error

	go func() {
		w, err = s.CreateWebhook(channelID, data)
		wg.Done()
	}()

	return func() (*discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func ChannelWebhooks(s *state.State, channelID discord.ChannelID) func() ([]discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w []discord.Webhook
	var err error

	go func() {
		w, err = s.ChannelWebhooks(channelID)
		wg.Done()
	}()

	return func() ([]discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func GuildWebhooks(s *state.State, guildID discord.GuildID) func() ([]discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w []discord.Webhook
	var err error

	go func() {
		w, err = s.GuildWebhooks(guildID)
		wg.Done()
	}()

	return func() ([]discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func Webhook(s *state.State, id discord.WebhookID) func() (*discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.Webhook
	var err error

	go func() {
		w, err = s.Webhook(id)
		wg.Done()
	}()

	return func() (*discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func WebhookWithToken(id discord.WebhookID, token string) func() (*discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.Webhook
	var err error

	go func() {
		w, err = webhook.New(id, token).Get()
		wg.Done()
	}()

	return func() (*discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func ModifyWebhook(s *state.State, id discord.WebhookID, data api.ModifyWebhookData) func() (*discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.Webhook
	var err error

	go func() {
		w, err = s.ModifyWebhook(id, data)
		wg.Done()
	}()

	return func() (*discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func ModifyWebhookWithToken(id discord.WebhookID, token string, data api.ModifyWebhookData) func() (*discord.Webhook, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var w *discord.Webhook
	var err error

	go func() {
		w, err = webhook.New(id, token).Modify(data)
		wg.Done()
	}()

	return func() (*discord.Webhook, error) {
		wg.Wait()
		return w, err
	}
}

func DeleteWebhook(s *state.State, id discord.WebhookID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteWebhook(id)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteWebhookWithToken(id discord.WebhookID, token string) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = webhook.New(id, token).Delete()
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Execute(id discord.WebhookID, token string, data webhook.ExecuteData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = webhook.New(id, token).Execute(data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func ExecuteAndWait(id discord.WebhookID, token string, data webhook.ExecuteData) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = webhook.New(id, token).ExecuteAndWait(data)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}
