package entity

type Collection struct {
	ID                  int
	HoldersCount        int
	HoldersDistribution HoldersDistribution
	Holders             []Holder
}

type HoldersDistribution struct {
	ByCommitmentScore map[string]int
	ByPortfolioScore  map[string]int
	ByTradingScore    map[string]int
}

type Holder struct {
	Address         string
	TokensAmount    int
	CommitmentScore float64
	PortfolioScore  float64
	TradingScore    float64
}
