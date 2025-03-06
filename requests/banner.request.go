package requests

type BannerRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type BannerIdRequest struct {
	ID int `uri:"id"`
}

type BannerCreateRequest struct {
	Banner string `json:"banner"`
}

type PublishedCreateRequest struct {
	Published string `json:"published"`
}

type BannerUpdateRequest struct {
	ID     int    `json:"id"`
	Banner string `json:"banner"`
}

type PublishedUpdateRequest struct {
	ID        int    `json:"id"`
	Published string `json:"published"`
}
