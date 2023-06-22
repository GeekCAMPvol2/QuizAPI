package db

import (
	"testing"

	"github.com/GeekCAMPvol2/QuizAPI/util"
	"github.com/stretchr/testify/require"
)

// 必要数取得できているか
func TestGetQuiz(t *testing.T) {
	config, err := util.Loadenv("../")
	require.NoError(t, err)

	client := NewClient(config.MongoTestUri)
	require.NoError(t, err)

	quiz, err := client.GetQuiz(10)

	require.NoError(t, err)

	require.Len(t, quiz, 10)
}
