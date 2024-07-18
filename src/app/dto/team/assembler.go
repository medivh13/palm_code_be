package team

func ToTeam(datas []*TeamRespModel) []*TeamRespDTO {
	var resp []*TeamRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnTeam(m))
	}
	return resp
}

func ToReturnTeam(d *TeamRespModel) *TeamRespDTO {
	return &TeamRespDTO{
		ID:             d.ID,
		Name:           d.Name,
		Role:           d.Role,
		Bio:            d.Bio,
		ProfilePicture: d.ProfilePicture,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
}
