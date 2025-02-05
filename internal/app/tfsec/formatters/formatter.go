package formatters

import (
	"io"

	"github.com/liamg/tfsec/internal/app/tfsec/scanner"
)

// Formatter formats scan results into a specific format
type Formatter func(w io.Writer, results []scanner.Result) error
