package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	externalModels "github.com/jetaimejeteveux/e-wallet-ums/external/models"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
)

type External struct {
}

func (e *External) CreateWallet(ctx context.Context, userID int) (*externalModels.Wallet, error) {
	log := helpers.Logger
	wallet := &externalModels.Wallet{
		UserID: userID,
	}

	payload, err := json.Marshal(wallet)
	if err != nil {
		log.Error("failed to marshal JSON: ", err)
		return nil, errors.New(constants.ErrFailedMarshalJson)
	}

	url := fmt.Sprintf("%s/%s", helpers.GetEnv("WALLET_HOST", ""), helpers.GetEnv("WALLET_ENDPOINT_CREATE", ""))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Error("error creating new request: ", err)
		return nil, errors.New(constants.ErrCreateNewRequest)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error executing request: ", err)
		return nil, errors.New(constants.ErrExecuteRequest)
	}

	result := &externalModels.Wallet{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Error("error decoding json response: ", err)
		return nil, errors.New(constants.ErrDecodingJSON)
	}
	defer resp.Body.Close()

	return result, nil
}
