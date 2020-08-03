# dasync

Dasync is a utility library, used to perform asynchronous calls to the Discord API.
It uses [disstate](https://github.com/mavolin/disstate) to access state and, if necessary the Discord API.

## Performance and Usage

To keep the performance and memory overhead as small as possible, a goroutine will only be started if an element is not available in the state.
Values can be accessed by calling the returned futures, which in turn are functions that block until the resource becomes available.

All functions are static, so that you won't need an instance of a dasync type, and can just pass in your `State`.

## Example

Assume you want to get a guild and a user simultaneously:

```go
var (
    s = state.New(myToken)
    guildID discord.GuildID = 123
    userID discord.UserID = 465
)

gf := dasync.Guild(s, guildID) // each function
uf := dasync.User(s, userID) // returns a future

g, err := gf() // blocks until the
u, err := uf() // resources become available
```