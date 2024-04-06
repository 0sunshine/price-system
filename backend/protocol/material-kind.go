package protocol

type GetMaterialKindListReqFilter struct {
	CommonFilterBase `json:",inline"`
}

type GetMaterialKindReq struct {
	CommonRequestBase[GetMaterialKindListReqFilter] `json:",inline"`
}

type MaterialKindItem struct {
	MaterialKindId int64  `json:"material_kind_id" `
	ProjectKindId  string `json:"project_kind_id" `
	MaterialName   string `json:"material_name" `
	Notes          string `json:"notes" `
	Unit           string `json:"unit" `
}

type getMaterialKindListResData struct {
	TotalCount int                `json:"totalCount"`
	List       []MaterialKindItem `json:"list"`
}

type GetMaterialKindListRes CommonRespondse[getMaterialKindListResData]

type AddMaterialKindReq struct {
	MaterialKindItem `json:",inline"`
}

type AddMaterialKindRes CommonNoDataRespondse

type UpdateMaterialKindReq struct {
	MaterialKindItem `json:",inline"`
}

type UpdateMaterialKindRes CommonNoDataRespondse

type DeleteMaterialKindReq struct {
	MaterialKindIds []int64 `json:"material_kind_id"`
}

type DeleteMaterialKindRes CommonNoDataRespondse
