package pkg

import (
	"context"
	"errors"
	"time"

	"github.com/machinebox/graphql"

	"README.pet/config"
)

// APIから取得する日数を設定
const validDays = 10

// APIから取得するデータの構造体
type ResponseStruct struct {
	User struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				Weeks []struct {
					ContributionDays []struct {
						ContributionCount int    `json:"contributionCount"`
						Date              string `json:"date"`
					} `json:"contributionDays"`
				} `json:"weeks"`
			} `json:"contributionCalendar"`
		} `json:"contributionsCollection"`
	} `json:"user"`
}

func FetchContributions(account string) ([]int, error) {
	client := graphql.NewClient("https://api.github.com/graphql")

	req := graphql.NewRequest(`
		query($userName:String!, $startDate:DateTime!, $endDate:DateTime!) {
			user(login: $userName){
			contributionsCollection(from: $startDate, to: $endDate) {
				contributionCalendar {
				weeks {
					contributionDays {
					contributionCount
					date
					}
				}
				}
			}
			}
		}
	`)

	// 取得開始日の日付と現在の時刻を取得
	startDate, endDate := getValidDates()

	req.Var("userName", account)
	req.Var("startDate", startDate)
	req.Var("endDate", endDate)

	// トークンが設定されている場合はリクエストヘッダに追加
	token := config.LoadToken()
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	ctx := context.Background()

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, err
	}

	// レスポンスのコントリビューション数を日ごとに配列に格納
	contributions := []int{}
	for _, week := range respData.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			contributions = append(contributions, day.ContributionCount)
		}
	}

	// 取得した値の数が合っているか検証
	if len(contributions) != validDays {
		return nil, errors.New("contributions count do not match")
	}

	return contributions, nil
}

func getValidDates() (string, string) {
	now := time.Now()

	// Goでは時刻のフォーマットをこのように特定の時刻で指定して行う
	startDate := now.AddDate(0, 0, -validDays+1).Format("2006-01-02T") + "00:00:00Z"
	endDate := now.Format("2006-01-02T15:04:05Z")

	return startDate, endDate
}
