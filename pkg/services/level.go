package services

import (
	gnats "github.com/getmilly/grok/nats"
	"github.com/getmilly/levels/pkg/level"
	"github.com/getmilly/levels/pkg/messages"
	"github.com/getmilly/levels/pkg/models"
	"github.com/getmilly/levels/pkg/repositories"
)

//LevelService ...
type LevelService interface {
	HandleReward(event *messages.RewardReached) error
}

type levelService struct {
	levelUpgradedSubject string
	levelUpdatedSubject  string
	producer             *gnats.Producer
	calculator           level.Calculator
	calculatedRepository repositories.CalculatedLevelRepository
}

//NewLevelService ...
func NewLevelService(
	levelUpgradedSubject string,
	levelUpdatedSubject string,
	producer *gnats.Producer,
	calculator level.Calculator,
	calculatedRepository repositories.CalculatedLevelRepository,
) LevelService {
	return &levelService{
		producer:             producer,
		calculator:           calculator,
		levelUpgradedSubject: levelUpgradedSubject,
		levelUpdatedSubject:  levelUpdatedSubject,
		calculatedRepository: calculatedRepository,
	}
}

func (service *levelService) HandleReward(event *messages.RewardReached) error {
	calculated, err := service.calculatedRepository.Get(event.PlayerID)

	if err != nil {
		return err
	}

	result, err := service.calculator.Calculate(calculated.Level, calculated.TotalExperience, event.Data.Experience)

	if err != nil {
		return err
	}

	calculated.Level = result.Level
	calculated.TotalExperience = result.TotalExperience
	calculated.ExperienceToUpgrade = result.ExperienceToUpgrade

	err = service.calculatedRepository.Set(calculated)

	if err != nil {
		return err
	}

	return service.dispatchEvents(result.HasUpgraded, calculated)
}

func (service *levelService) dispatchEvents(hasUpgraded bool, calculated *models.CalculatedLevel) error {
	if hasUpgraded {
		return service.publish(messages.LevelUpgraded{
			CalculatedLevel: calculated,
		})
	}

	return service.publish(messages.LevelUpdated{
		CalculatedLevel: calculated,
	})
}

func (service *levelService) publish(m interface{}) error {
	message, err := gnats.NewMessage(m)

	if err != nil {
		return err
	}

	return service.producer.Publish(service.levelUpgradedSubject, message)
}
