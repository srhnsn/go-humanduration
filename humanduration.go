package humanduration

import (
	"github.com/srhnsn/go-humanduration/locale"
	"github.com/srhnsn/go-humanduration/types"
)

func NewEnglishFormatter() types.HumanDurationFormatter {
	return types.HumanDurationFormatter{
		Formatter: locale.EnFormatter,
	}
}

func NewGermanFormatter() types.HumanDurationFormatter {
	return types.HumanDurationFormatter{
		Formatter: locale.DeFormatter,
	}
}
