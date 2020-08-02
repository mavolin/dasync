package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

func Emojis(s *state.State, guildID discord.GuildID) func() ([]discord.Emoji, error) {
	e, err := s.Store.Emojis(guildID)
	if err == nil {
		return func() ([]discord.Emoji, error) {
			return e, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		e, err = s.Emojis(guildID)
		wg.Done()
	}()

	return func() ([]discord.Emoji, error) {
		wg.Wait()
		return e, err
	}
}

func Emoji(s *state.State, guildID discord.GuildID, emojiID discord.EmojiID) func() (*discord.Emoji, error) {
	e, err := s.Store.Emoji(guildID, emojiID)
	if err == nil {
		return func() (*discord.Emoji, error) {
			return e, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		e, err = s.Emoji(guildID, emojiID)
		wg.Done()
	}()

	return func() (*discord.Emoji, error) {
		wg.Wait()
		return e, err
	}
}

func CreateEmoji(s *state.State, guildID discord.GuildID, data api.CreateEmojiData) func() (*discord.Emoji, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var e *discord.Emoji
	var err error

	go func() {
		e, err = s.CreateEmoji(guildID, data)
		wg.Done()
	}()

	return func() (*discord.Emoji, error) {
		wg.Wait()
		return e, err
	}
}

func ModifyEmoji(s *state.State, guildID discord.GuildID, emojiID discord.EmojiID, data api.ModifyEmojiData) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.ModifyEmoji(guildID, emojiID, data)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteEmoji(s *state.State, guildID discord.GuildID, emojiID discord.EmojiID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteEmoji(guildID, emojiID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
