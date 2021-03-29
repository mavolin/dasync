package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/mavolin/disstate/v3/pkg/state"
)

func SendMessageComplex(s *state.State, channelID discord.ChannelID, data api.SendMessageData) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.SendMessageComplex(channelID, data)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}
