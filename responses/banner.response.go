package response

type BannerReponses struct {
	ID     int    `json:"id"`
	Banner string `json:"banner"`
}

type PublishedResponses struct {
	ID        int    `json:"id"`
	Published string `json:"published"`
}
