package debian

import (
	"github.com/pritunl/tools/errors"
)

type BuildError struct {
	errors.DropboxError
}
