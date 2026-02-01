package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yuriy-kazanov/arch-go-example/internal/model"
	"io"
	"net/http"
)

type Instruments struct {
	client *http.Client
}

type Instrument struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	FaceValue int    `json:"faceValue"`
}

func (i *Instrument) Validate() error {
	if i.ID == "" {
		return errors.New("empty instrument id")
	}
	if i.Type == "bond" && i.FaceValue == 0 {
		return errors.New("empty faceValue for bond")
	}

	return nil
}

func (i *Instruments) GetInstruments(ctx context.Context) (model.Instrument, error) {
	resp, _ := i.client.Get("https://some-url-instrument-service.ru")

	out := &Instrument{}

	body, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(body, out)

	if err := out.Validate(); err != nil {
		return model.Instrument{}, fmt.Errorf("validation failed, %w", err)
	}

	return model.Instrument{
		ID:        out.ID,
		Type:      out.Type,
		FaceValue: out.FaceValue,
	}, nil
}
