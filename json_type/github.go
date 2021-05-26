package json_type

type ContributionsJson struct {
	Data struct {
		User struct {
			Contributionscollection struct {
				Contributioncalendar struct {
					Totalcontributions int `json:"totalContributions"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}
