package dasync

import (
	"sync"

	"github.com/diamondburned/arikawa/api"
	"github.com/mavolin/disstate/pkg/state"
)

func Login(s *state.State, email, password string) func() (*api.LoginResponse, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r *api.LoginResponse
	var err error

	go func() {
		r, err = s.Login(email, password)
		wg.Done()
	}()

	return func() (*api.LoginResponse, error) {
		wg.Wait()
		return r, err
	}
}

func TOTP(s *state.State, code, ticket string) func() (*api.LoginResponse, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	var r *api.LoginResponse
	var err error

	go func() {
		r, err = s.TOTP(code, ticket)
		wg.Done()
	}()

	return func() (*api.LoginResponse, error) {
		wg.Wait()
		return r, err
	}
}
