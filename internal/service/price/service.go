package price

import (
	"fmt"

	"gitlab.tbank.ru/ca-business-common/pkg/pricing"
)

type Service struct {
	pricing.Converter
}

func (s *Service) ProcessPrices() error {
	outPrices, err := s.ConvertPrices(100, "RUB")
	if err != nil {
		return fmt.Errorf("failed to convert prices, %w", err)
	}

	return s.updatePosition(outPrices)
}

func (s *Service) updatePosition(outPrices []string) error {
	return nil
}
