package portfolio_test

import (
	"testing"

	"github.com/gandarez/semver-demo/pkg/portfolio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPortfolio(t *testing.T) {
	portfolio, err := portfolio.GetPortfolio(1)
	require.NoError(t, err)

	assert.Equal(t, "My Portfolio", portfolio)
}

func TestGetPortfolioErr(t *testing.T) {
	_, err := portfolio.GetPortfolio(99)

	assert.Error(t, err)
}
