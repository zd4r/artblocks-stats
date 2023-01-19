package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/zd4rova/artblocks-holders/internal/entity"
)

type ArtacleWebAPI struct {
	httpClient *http.Client
}

func New() *ArtacleWebAPI {
	httpClient := &http.Client{}

	return &ArtacleWebAPI{
		httpClient: httpClient,
	}
}

type HoldersCountResp []struct {
	Diversity   int   `json:"diversity"`
	TokenCount  int   `json:"tokenCount"`
	OwnersCount int   `json:"ownersCount"`
	Index       int64 `json:"index"`
}

func (a *ArtacleWebAPI) GetHoldersCount(collection entity.Collection) (entity.Collection, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://artacle.io/api/project/%d/ownersChart", collection.ID), nil)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHoldersCount - http.NewRequest: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHoldersCount - a.httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	var respData HoldersCountResp
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHoldersCount - json.Decode: %w", err)
	}

	collection.HoldersCount = respData[len(respData)-1].OwnersCount

	return collection, nil
}

type GetHoldersResp struct {
	OwnersProfit []struct {
		Owner                 string  `json:"owner"`
		BalanceOut            int     `json:"balanceOut"`
		BalanceIn             float64 `json:"balanceIn"`
		BalanceAll            float64 `json:"balanceAll"`
		TokenDelta            int     `json:"tokenDelta"`
		BalanceDeltaOut       int     `json:"balanceDeltaOut"`
		BalanceDeltaIn        float64 `json:"balanceDeltaIn"`
		BalanceDelta          float64 `json:"balanceDelta"`
		TokensAll             int     `json:"tokensAll"`
		TokenDeltaIn          int     `json:"tokenDeltaIn"`
		TokenDeltaOut         int     `json:"tokenDeltaOut"`
		TransferTokenDeltaIn  int     `json:"transferTokenDeltaIn"`
		TransferTokenDeltaOut int     `json:"transferTokenDeltaOut"`
		LastTransaction       int64   `json:"lastTransaction"`
		OwnerName             string  `json:"ownerName"`
		OwnerOSName           string  `json:"ownerOSName"`
		OwnerLabel            string  `json:"ownerLabel"`
		OwnerAddrType         string  `json:"ownerAddrType"`
		OwnerBalance          string  `json:"ownerBalance"`
	} `json:"ownersProfit"`
}

func (a *ArtacleWebAPI) GetHolders(collection entity.Collection) (entity.Collection, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://artacle.io/api/project/%d", collection.ID), nil)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHolders - http.NewRequest: %w", err)
	}

	q := req.URL.Query()
	q.Add("orderBy", "tokensAll")
	q.Add("sortReverse", "true")
	q.Add("limit", strconv.Itoa(collection.HoldersCount))
	q.Add("address", "")
	q.Add("isHot", "false")
	q.Add("offset", "0")
	req.URL.RawQuery = q.Encode()

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHolders - a.httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	var respData GetHoldersResp
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("ArtacleWebAPI - GetHolders - json.Decode: %w", err)
	}

	var holder entity.Holder
	for _, op := range respData.OwnersProfit {
		holder.Address = op.Owner
		holder.TokensAmount = op.TokensAll
		collection.Holders = append(collection.Holders, holder)
	}

	return collection, nil
}

type GetHolderScoresRest struct {
	Names []struct {
		NameType   string `json:"nameType"`
		WalletName string `json:"walletName"`
	} `json:"names"`
	Balance string `json:"balance"`
	Scores  struct {
		Score                        float64 `json:"score"`
		ScoreConfidence              float64 `json:"scoreConfidence"`
		CommitmentScore              float64 `json:"commitmentScore"`
		CommitmentScoreConfidence    float64 `json:"commitmentScoreConfidence"`
		CommitmentScoreEstimations   int     `json:"commitmentScoreEstimations"`
		TradingScore                 float64 `json:"tradingScore"`
		TradingScoreConfidence       float64 `json:"tradingScoreConfidence"`
		TradingScoreEstimations      int     `json:"tradingScoreEstimations"`
		PortfolioScore               float64 `json:"portfolioScore"`
		PortfolioScoreConfidence     float64 `json:"portfolioScoreConfidence"`
		PortfolioScoreEstimations    int     `json:"portfolioScoreEstimations"`
		PortfolioTags                string  `json:"portfolioTags"`
		LastScoresTs                 int64   `json:"lastScoresTs"`
		CommitmentScoreConfidenceMin float64 `json:"commitmentScoreConfidenceMin"`
		PortfolioScoreConfidenceMin  float64 `json:"portfolioScoreConfidenceMin"`
		TradingScoreConfidenceMin    float64 `json:"tradingScoreConfidenceMin"`
	} `json:"scores"`
}

func (a *ArtacleWebAPI) GetHolderScores(holder entity.Holder) (entity.Holder, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://artacle.io/api/user/%s", strings.ToLower(holder.Address)), nil)
	if err != nil {
		return entity.Holder{}, fmt.Errorf("ArtacleWebAPI - GetHolderScores - http.NewRequest: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return entity.Holder{}, fmt.Errorf("ArtacleWebAPI - GetHolderScores - a.httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	var respData GetHolderScoresRest
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return entity.Holder{}, fmt.Errorf("ArtacleWebAPI - GetHolders - json.Decode: %w", err)
	}

	holder.CommitmentScore = respData.Scores.CommitmentScore
	holder.PortfolioScore = respData.Scores.PortfolioScore
	holder.TradingScore = respData.Scores.TradingScore

	return holder, nil
}