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
	calculator := level.NewTibiaCalculator(3)

	level := 1
	totalXP := 350

	result, _ := calculator.Calculate(level, totalXP, 0)

	assert.False(t, result.HasUpgraded)
	assert.Equal(t, result.Level, 3)
	assert.Equal(t, result.TotalExperience, 50)
	assert.Equal(t, result.ExperienceToUpgrade, 350)
}

func TestTibia_Calculate_NonUpgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator(3)

	level := 1
	missingtToUpgrade := 5

	totalXP := calculator.CalculateExperienceByLevel(level+1) - missingtToUpgrade

	result, _ := calculator.Calculate(level, totalXP, 0)

	assert.False(t, result.HasUpgraded)
	assert.Equal(t, result.Level, level)
	assert.Equal(t, result.ExperienceToUpgrade, missingtToUpgrade)
}
