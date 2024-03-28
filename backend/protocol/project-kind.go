package protocol

type GetProjectKindListReqFilter struct {
	CommonFilterBase    `json:",inline"`
	ProjectKindId       string `json:"project_kind_id"`
	ParentProjectKindId string `json:"parent_project_kind_id"`
}

type GetProjectKindListReq struct {
	CommonRequestBase[GetProjectKindListReqFilter] `json:",inline"`
}

type ProjectKindItem struct {
	ParentProjectKindId string            `json:"parent_project_kind_id" db:"parent_project_kind_id"`
	ProjectKindId       string            `json:"project_kind_id" db:"project_kind_id"`
	ProjectName         string            `json:"project_name" db:"project_name"`
	Notes               string            `json:"notes" db:"notes"`
	Children            []ProjectKindItem `json:"children"`
}

type getProjectKindListResData struct {
	TotalCount int               `json:"totalCount"`
	List       []ProjectKindItem `json:"list"`
}

type GetProjectKindListRes CommonRespondse[getProjectKindListResData]

type AddProjectKindListReq struct {
	ProjectKindItem `json:",inline"`
}

type AddProjectKindListRes CommonNoDataRespondse

type UpdateProjectKindReq struct {
	ProjectKindItem `json:",inline"`
}

type UpdateProjectKindListRes CommonNoDataRespondse

type DeleteProjectKindReq struct {
	ProjectKindIds []string `json:"project_kind_id"`
}

type DeleteProjectKindListRes CommonNoDataRespondse
