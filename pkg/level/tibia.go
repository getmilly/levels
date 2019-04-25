package level

import (
	"math"
)

//TibiaCalculator applies Tibia level formula
type TibiaCalculator struct{}

//NewTibiaCalculator creates a new PrtTo TibiaCalculator
func NewTibiaCalculator() *TibiaCalculator {
	return &TibiaCalculator{}
}

//Calculate player level.
func (calc *TibiaCalculator) Calculate(totalXP int) (*CalcResult, error) {
	result := new(CalcResult)

	result.Level = 0
	result.TotalExperience = totalXP

	return calc.calculateRecursive(result), nil
}

//CalculateExperienceByLevel ...
func (calc *TibiaCalculator) CalculateExperienceByLevel(level int) int {
	levelFloat := float64(level)
	return int(50.0 / 3 * (math.Pow(levelFloat, 3) - (6 * math.Pow(levelFloat, 2)) + 17*levelFloat - 12))
}

func (calc *TibiaCalculator) calculateRecursive(result *CalcResult) *CalcResult {
	nextLevelExperience := calc.CalculateExperienceByLevel(result.Level + 1)

	if result.TotalExperience >= nextLevelExperience {
		result.Level = result.Level + 1
		result.TotalExperience = result.TotalExperience - nextLevelExperience
		nextLevelExperience = calc.CalculateExperienceByLevel(result.Level + 1)

		if result.TotalExperience >= nextLevelExperience {
			return calc.calculateRecursive(result)
		}

		result.ExperienceToUpgrade = nextLevelExperience - result.TotalExperience

		return result
	}

	result.ExperienceToUpgrade = nextLevelExperience - result.TotalExperience

	return result
}
