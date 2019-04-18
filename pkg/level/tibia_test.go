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
	earnedXP := 1

	result, _ := calculator.Calculate(level, totalXP, earnedXP)

	result2, _ := calculator.Calculate(level, totalXP+earnedXP, result.NextLevelExperience)

	assert.True(t, result2.HasUpgraded)
	assert.Equal(t, totalXP+earnedXP+result.NextLevelExperience, result2.TotalExperience)
}

func TestTibia_Calculate_NonUpgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator(10)

	level := 1
	totalXP := 0
	earnedXP := 1

	result, _ := calculator.Calculate(level, totalXP, earnedXP)

	result2, _ := calculator.Calculate(level, totalXP+earnedXP, result.NextLevelExperience-1)

	assert.False(t, result2.HasUpgraded)
}
