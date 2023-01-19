package entity

type Collection struct {
	ID                  int                 `json:"id"`
	HoldersCount        int                 `json:"holders_count"`
	HoldersDistribution HoldersDistribution `json:"holders_distribution"`
	Holders             []Holder            `json:"holders"`
}

type HoldersDistribution struct {
	ByCommitmentScore map[string]int `json:"by_commitment_score"`
	ByPortfolioScore  map[string]int `json:"by_portfolio_score"`
	ByTradingScore    map[string]int `json:"by_trading_score"`
}

type Holder struct {
	Address         string  `json:"address"`
	TokensAmount    int     `json:"tokens_amount"`
	CommitmentScore float64 `json:"commitment_score"`
	PortfolioScore  float64 `json:"portfolio_score"`
	TradingScore    float64 `json:"trading_score"`
}
