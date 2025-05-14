package entity

type Solution struct {
	Thoughts        string `json:"thoughts,omitempty"`
	Code            string `json:"code,omitempty"`
	TimeComplexity  string `json:"time_complexity,omitempty"`
	SpaceComplexity string `json:"space_complexity,omitempty"`
}
