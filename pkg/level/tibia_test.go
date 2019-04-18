package level_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/getmilly/levels/pkg/level"
)

func TestTibia_Calculate_Panics(t *testing.T) {
	assert.Panics(t, func() {
		level.NewTibiaCalculator(-1)
	})
}

func TestTibia_Calculate_Upgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator(10)

	level := 1
	totalXP := 0
	earnedXP := 200

	result, _ := calculator.Calculate(level, totalXP, earnedXP)

	assert.Equal(t, result.Level, 3)
	assert.Equal(t, result.TotalExperience, 40)
	assert.Equal(t, result.NextLevelExperience, 185)
}

func TestTibia_Calculate_NonUpgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator(10)

	level := 1
	totalXP := 0
	earnedXP := 1

	result, _ := calculator.Calculate(level, totalXP, earnedXP)

	assert.Equal(t, result.Level, level)
	assert.Equal(t, result.TotalExperience, totalXP+earnedXP)
	assert.Equal(t, result.NextLevelExperience, 65)
}
