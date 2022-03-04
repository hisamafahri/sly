package helper

import (
	"errors"

	"github.com/bitfield/script"
)

func RunCommand(command string) (string, error) {
	var output string
	for _, c := range []string{command} {
		p := script.Exec(c)
		if err := p.Error(); err != nil {
			p.SetError(nil)
			output, _ := p.String()
			return "", errors.New(output)
		} else {
			result, _ := p.String()
			output = result
		}
	}
	return output, nil
}
