package api

import (
	"backend/db"
	"backend/protocol"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMaterialKindList(c *gin.Context) {
	req := protocol.GetMaterialKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	//fmt.Printf("%+v\n", req)
	//tmp, _ := json.Marshal(req)
	//fmt.Println(string(tmp))

	list, err := db.GetMaterialKindList(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `GetMaterialKindList err`)
		return
	}

	res := protocol.GetMaterialKindListRes{}

	res.Data.List = list
	res.Data.TotalCount = len(list)

	c.JSON(http.StatusOK, res)
}

func AddMaterialKind(c *gin.Context) {
	req := protocol.AddMaterialKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.AddMaterialKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `AddMaterialKind err`, err)
		return
	}

	res := protocol.AddMaterialKindRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func UpdateMaterialKind(c *gin.Context) {
	req := protocol.UpdateMaterialKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.UpdatedMaterialKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `UpdatedMaterialKind err`, err)
		return
	}

	res := protocol.UpdateMaterialKindRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func DeleteMaterialKind(c *gin.Context) {
	req := protocol.DeleteMaterialKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.DeleteMaterialKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `DeleteMaterialKind err: `, err)
		return
	}

	res := protocol.DeleteMaterialKindRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}
