package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

// Messages will not call the state first, as the result, depending on the type
// of caching and the shard configuration, may differ.
func Messages(s *state.State, channelID discord.ChannelID, limit uint) func() ([]discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Message
	var err error

	go func() {
		m, err = s.Client.Messages(channelID, limit)
		wg.Done()
	}()

	return func() ([]discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func MessagesAround(
	s *state.State, channelID discord.ChannelID, around discord.MessageID, limit uint,
) func() ([]discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Message
	var err error

	go func() {
		m, err = s.MessagesAround(channelID, around, limit)
		wg.Done()
	}()

	return func() ([]discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func MessagesBefore(
	s *state.State, channelID discord.ChannelID, before discord.MessageID, limit uint,
) func() ([]discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Message
	var err error

	go func() {
		m, err = s.MessagesBefore(channelID, before, limit)
		wg.Done()
	}()

	return func() ([]discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func MessagesAfter(
	s *state.State, channelID discord.ChannelID, after discord.MessageID, limit uint,
) func() ([]discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m []discord.Message
	var err error

	go func() {
		m, err = s.MessagesAfter(channelID, after, limit)
		wg.Done()
	}()

	return func() ([]discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func Message(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID,
) func() (*discord.Message, error) {
	m, err := s.Store.Message(channelID, messageID)
	if err == nil {
		return func() (*discord.Message, error) {
			return m, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		m, err = s.Message(channelID, messageID)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func SendText(s *state.State, channelID discord.ChannelID, content string) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.SendText(channelID, content)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func SendEmbed(s *state.State, channelID discord.ChannelID, e discord.Embed) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.SendEmbed(channelID, e)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func SendMessage(
	s *state.State, channelID discord.ChannelID, content string, embed *discord.Embed,
) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.SendMessage(channelID, content, embed)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func EditText(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, content string,
) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.EditText(channelID, messageID, content)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func EditEmbed(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, embed discord.Embed,
) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.EditEmbed(channelID, messageID, embed)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func EditMessage(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, content string, embed *discord.Embed,
	suppressEmbeds bool,
) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.EditMessage(channelID, messageID, content, embed, suppressEmbeds)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func EditMessageComplex(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, data api.EditMessageData) func() (*discord.Message, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var m *discord.Message
	var err error

	go func() {
		m, err = s.EditMessageComplex(channelID, messageID, data)
		wg.Done()
	}()

	return func() (*discord.Message, error) {
		wg.Wait()
		return m, err
	}
}

func DeleteMessage(s *state.State, channelID discord.ChannelID, messageID discord.MessageID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteMessage(channelID, messageID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteMessages(s *state.State, channelID discord.ChannelID, messageIDs []discord.MessageID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteMessages(channelID, messageIDs)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
