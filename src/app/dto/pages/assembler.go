package pages

func ToPage(datas []*PageRespModel) []*PageRespDTO {
	var resp []*PageRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnPage(m))
	}
	return resp
}

func ToReturnPage(d *PageRespModel) *PageRespDTO {
	return &PageRespDTO{
		ID:              d.ID,
		Title:           d.Title,
		Slug:            d.Slug,
		BannerMedia:     d.BannerMedia,
		Content:         d.Content,
		PublicationDate: d.PublicationDate,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}
