package api

import (
	"backend/db"
	"backend/protocol"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMaterialAttrList(c *gin.Context) {
	req := protocol.GetMaterialAttrReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	//fmt.Printf("%+v\n", req)
	//tmp, _ := json.Marshal(req)
	//fmt.Println(string(tmp))

	list, err := db.GetMaterialAttrList(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `GetMaterialAttrList err`)
		return
	}

	res := protocol.GetMaterialAttrListRes{}

	res.Data.List = list
	res.Data.TotalCount = len(list)

	c.JSON(http.StatusOK, res)
}

func AddMaterialAttr(c *gin.Context) {
	req := protocol.AddMaterialAttrReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.AddMaterialAttr(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `AddMaterialAttr err`, err)
		return
	}

	res := protocol.AddMaterialAttrRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func UpdateMaterialAttr(c *gin.Context) {
	req := protocol.UpdateMaterialAttrReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.UpdatedMaterialAttr(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `UpdateMaterialAttr err`, err)
		return
	}

	res := protocol.UpdateMaterialKindRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func DeleteMaterialAttr(c *gin.Context) {
	req := protocol.DeleteMaterialAttrReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.DeleteMaterialAttr(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `DeleteMaterialAttr err: `, err)
		return
	}

	res := protocol.DeleteMaterialAttrRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}
