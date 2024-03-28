package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//type CommonFilter struct {
//	PageSize    int `json:"page-size"`
//	CurrentPage int `json:"current-page"`
//}
//
//type GetProjectKindListReqFilter struct {
//	CommonFilter
//}
//
//type GetProjectKindListReq struct {
//	Filter GetProjectKindListReqFilter `json:"filter"`
//}
//
//type CommonRes struct {
//	Code int `json:"code"`
//	Msg  int `json:"msg"`
//}
//
//type ProjectKindItem struct {
//	ParentProjectKindId string            `json:"parent_project_kind_id"`
//	ProjectKindId       string            `json:"project_kind_id"`
//	ProjectName         string            `json:"project_name"`
//	Notes               string            `json:"notes"`
//	Children            []ProjectKindItem `json:"children"`
//}
//
//type GetProjectKindListResData struct {
//	TotalCount int               `json:"totalCount"`
//	List       []ProjectKindItem `json:"list"`
//}
//
//type GetProjectKindListRes struct {
//	CommonRes
//	Data GetProjectKindListResData `json:"data"`
//}
//
//func getProjectKindList(w http.ResponseWriter, r *http.Request) {
//	data, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		panic(err)
//	}
//
//	req := GetProjectKindListReq{}
//
//	err = json.Unmarshal(data, &req)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(req.Filter.PageSize)
//
//	res := GetProjectKindListRes{}
//
//	res.Data.List = []ProjectKindItem{
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "1",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//			Children: []ProjectKindItem{
//				{
//					ParentProjectKindId: "1",
//					ProjectKindId:       "2",
//					ProjectName:         "测试ttt",
//					Notes:               "tttt",
//				},
//				{
//					ParentProjectKindId: "1",
//					ProjectKindId:       "3",
//					ProjectName:         "测试xxx",
//					Notes:               "tttt",
//				},
//				{
//					ParentProjectKindId: "1",
//					ProjectKindId:       "4",
//					ProjectName:         "测试xxt",
//					Notes:               "tttt",
//				},
//			},
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "5",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "6",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "17",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "8",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "9",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "10",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "11",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "12",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "13",
//			ProjectName:         "测试111",
//			Notes:               "xxxx",
//		},
//		{
//			ParentProjectKindId: "",
//			ProjectKindId:       "14",
//			ProjectName:         "测试222",
//			Notes:               "xxxx",
//		},
//	}
//
//	res.Data.TotalCount = len(res.Data.List)
//
//	resJson, err := json.Marshal(res)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(string(resJson))
//	w.Write(resJson)
//}
//
//func main() {
//	http.HandleFunc("/admin/getProjectKindList", getProjectKindList)
//
//	log.Fatal(http.ListenAndServe(":8888", nil))
//}
