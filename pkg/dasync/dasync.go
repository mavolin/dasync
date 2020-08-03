// Package dasync provides functions that allow asynchronous access to the
// discord api.
// If a resource isn't state cached, then the respective function will start
// a goroutine and fetch it from Discord.
package dasync
