package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
	"github.com/mavolin/disstate/pkg/state"
)

func React(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, emoji api.Emoji) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.React(channelID, messageID, emoji)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Unreact(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, emoji api.Emoji) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.Unreact(channelID, messageID, emoji)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func Reactions(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, emoji api.Emoji, limit uint,
) func() ([]discord.User, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var u []discord.User
	var err error

	go func() {
		u, err = s.Reactions(channelID, messageID, emoji, limit)
		wg.Done()
	}()

	return func() ([]discord.User, error) {
		wg.Wait()
		return u, err
	}
}

func ReactionsBefore(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, before discord.UserID, emoji api.Emoji,
	limit uint,
) func() ([]discord.User, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var u []discord.User
	var err error

	go func() {
		u, err = s.ReactionsBefore(channelID, messageID, before, emoji, limit)
		wg.Done()
	}()

	return func() ([]discord.User, error) {
		wg.Wait()
		return u, err
	}
}

func ReactionsAfter(
	s *state.State, channelID discord.ChannelID, messageID discord.MessageID, after discord.UserID, emoji api.Emoji,
	limit uint,
) func() ([]discord.User, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var u []discord.User
	var err error

	go func() {
		u, err = s.ReactionsAfter(channelID, messageID, after, emoji, limit)
		wg.Done()
	}()

	return func() ([]discord.User, error) {
		wg.Wait()
		return u, err
	}
}

func DeleteUserReaction(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, userID discord.UserID, emoji api.Emoji) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteUserReaction(channelID, messageID, userID, emoji)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteReactions(s *state.State, channelID discord.ChannelID, messageID discord.MessageID, emoji api.Emoji) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteReactions(channelID, messageID, emoji)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}

func DeleteAllReactions(s *state.State, channelID discord.ChannelID, messageID discord.MessageID) func() error {
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		err = s.DeleteAllReactions(channelID, messageID)
		wg.Done()
	}()

	return func() error {
		wg.Wait()
		return err
	}
}
