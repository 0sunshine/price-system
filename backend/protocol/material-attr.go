package protocol

type GetMaterialAttrListReqFilter struct {
	CommonFilterBase `json:",inline"`
	MaterialKindId   int64 `json:"material_kind_id" `
}

type GetMaterialAttrReq struct {
	CommonRequestBase[GetMaterialAttrListReqFilter] `json:",inline"`
}

type MaterialAttrItem struct {
	AttrId         int64  `json:"attr_id" `
	MaterialKindId int64  `json:"material_kind_id" `
	MaterialName   string `json:"material_name,omitempty" `
	AttrDesc       string `json:"attr_desc" `
}

type getMaterialAttrListResData struct {
	TotalCount int                `json:"totalCount"`
	List       []MaterialAttrItem `json:"list"`
}

type GetMaterialAttrListRes CommonRespondse[getMaterialAttrListResData]

type AddMaterialAttrReq struct {
	MaterialAttrItem `json:",inline"`
}

type AddMaterialAttrRes CommonNoDataRespondse

type UpdateMaterialAttrReq struct {
	MaterialAttrItem `json:",inline"`
}

type UpdateMaterialAttrRes CommonNoDataRespondse

type DeleteMaterialAttrReq struct {
	MaterialAttrIds []int64 `json:"attr_id"`
}

type DeleteMaterialAttrRes CommonNoDataRespondse
