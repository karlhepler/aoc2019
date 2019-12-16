package terminator

import "golang.org/x/crypto/ssh/terminal"

// Raw mode puts the terminal into raw mode and returns a function to
// restore it.
// Example Usage:
// func main() {
//		defer term.RawMode()
// }
func RawMode() (func() error, error) {
	oldState, err := terminal.MakeRaw(0)
	return func() error {
		return terminal.Restore(0, oldState)
	}, err
}
