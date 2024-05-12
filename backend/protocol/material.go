package protocol

type GetMaterialListReqFilter struct {
	CommonFilterBase        `json:",inline"`
	MaterialNameLike        string `json:"material_name_like" `
	ParentProjectKindIdLike string `json:"parent_project_kind_id_like" `
}

type GetMaterialReq struct {
	CommonRequestBase[GetMaterialListReqFilter] `json:",inline"`
}

type MaterialItem struct {
	MaterialId     int64              `json:"material_id" `
	MaterialKindId int64              `json:"material_kind_id" `
	MaterialName   string             `json:"material_name" `
	Price          string             `json:"price" `
	ExtraPrice     string             `json:"extra_price" `
	AttrDescs      []MaterialAttrItem `json:"attr_descs" `
}

type getMaterialListResData struct {
	TotalCount int            `json:"totalCount"`
	List       []MaterialItem `json:"list"`
}

type GetMaterialListRes CommonRespondse[getMaterialListResData]

type AddMaterialReq struct {
	MaterialItem `json:",inline"`
}

type AddMaterialRes CommonNoDataRespondse

type UpdateMaterialReq struct {
	MaterialItem `json:",inline"`
}

type UpdateMaterialRes CommonNoDataRespondse

type DeleteMaterialReq struct {
	MaterialIds []int64 `json:"material_id"`
}

type DeleteMaterialRes CommonNoDataRespondse
