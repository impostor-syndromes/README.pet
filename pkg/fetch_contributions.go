package pkg

import (
	"context"
	"encoding/json"

	"github.com/machinebox/graphql"

	"README.pet/config"
)

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

func FetchContributions(account string) (string, error) {
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

	req.Var("userName", account)
	req.Var("startDate", "2024-04-01T00:00:00Z")
	req.Var("endDate", "2024-04-07T23:59:59Z")

	token := config.LoadToken()

	req.Header.Set("Authorization", "Bearer "+token)

	ctx := context.Background()

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		return "", err
	}

	response, err := json.Marshal(respData)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
