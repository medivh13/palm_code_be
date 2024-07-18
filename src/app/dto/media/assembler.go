package media

func ToMedia(datas []*MediaRespModel) []*MediaRespDTO {
	var resp []*MediaRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnMedia(m))
	}
	return resp
}

func ToReturnMedia(d *MediaRespModel) *MediaRespDTO {
	return &MediaRespDTO{
		ID:        d.ID,
		URL:       d.URL,
		Type:      d.Type,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
