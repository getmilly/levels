package messages

import (
	"github.com/getmilly/levels/pkg/models"
)

//LevelUpdated ...
type LevelUpdated struct {
	models.CalculatedLevel
}
