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

func (c *Collection) CountHoldersDistribution() error {
	for _, h := range c.Holders {
		switch {
		case h.CommitmentScore < 3.5:
			c.HoldersDistribution.ByCommitmentScore["[3 - 3.5)"] += 1
		case h.CommitmentScore >= 3.5 && h.CommitmentScore < 4:
			c.HoldersDistribution.ByCommitmentScore["[3.5 - 4)"] += 1
		case h.CommitmentScore >= 4 && h.CommitmentScore < 4.5:
			c.HoldersDistribution.ByCommitmentScore["[4 - 4.5)"] += 1
		case h.CommitmentScore >= 4.5 && h.CommitmentScore <= 5:
			c.HoldersDistribution.ByCommitmentScore["[4 - 4.5]"] += 1
		}

		switch {
		case h.PortfolioScore < 3.5:
			c.HoldersDistribution.ByPortfolioScore["[3 - 3.5)"] += 1
		case h.PortfolioScore >= 3.5 && h.PortfolioScore < 4:
			c.HoldersDistribution.ByPortfolioScore["[3.5 - 4)"] += 1
		case h.PortfolioScore >= 4 && h.PortfolioScore < 4.5:
			c.HoldersDistribution.ByPortfolioScore["[4 - 4.5)"] += 1
		case h.PortfolioScore >= 4.5 && h.PortfolioScore <= 5:
			c.HoldersDistribution.ByPortfolioScore["[4 - 4.5]"] += 1
		}

		switch {
		case h.TradingScore < 3.5:
			c.HoldersDistribution.ByTradingScore["[3 - 3.5)"] += 1
		case h.TradingScore >= 3.5 && h.TradingScore < 4:
			c.HoldersDistribution.ByTradingScore["[3.5 - 4)"] += 1
		case h.TradingScore >= 4 && h.TradingScore < 4.5:
			c.HoldersDistribution.ByTradingScore["[4 - 4.5)"] += 1
		case h.TradingScore >= 4.5 && h.TradingScore <= 5:
			c.HoldersDistribution.ByTradingScore["[4 - 4.5]"] += 1
		}
	}

	return nil
}
