package level_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/getmilly/levels/pkg/level"
)

func TestTibia_Calculate_Upgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator()

	totalXP := 350

	result, _ := calculator.Calculate(totalXP)

	assert.Equal(t, 3, result.Level)
	assert.Equal(t, 50, result.TotalExperience)
	assert.Equal(t, 350, result.ExperienceToUpgrade)
}

func TestTibia_Calculate_NonUpgrade(t *testing.T) {
	calculator := level.NewTibiaCalculator()

	level := 1
	missingtToUpgrade := 5

	totalXP := calculator.CalculateExperienceByLevel(level+1) - missingtToUpgrade

	result, _ := calculator.Calculate(totalXP)

	assert.Equal(t, level, result.Level)
	assert.Equal(t, missingtToUpgrade, result.ExperienceToUpgrade)
}
