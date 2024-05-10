package api

import (
	"backend/db"
	"backend/protocol"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMaterialList(c *gin.Context) {
	req := protocol.GetMaterialReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	//fmt.Printf("%+v\n", req)
	//tmp, _ := json.Marshal(req)
	//fmt.Println(string(tmp))

	list, err := db.GetMaterialList(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `GetMaterialAttrList err`)
		return
	}

	res := protocol.GetMaterialListRes{}

	res.Data.List = list
	res.Data.TotalCount = len(list)

	c.JSON(http.StatusOK, res)
}

func AddMaterial(c *gin.Context) {
	req := protocol.AddMaterialReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.AddMaterial(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `AddMaterial err`, err)
		return
	}

	res := protocol.AddMaterialRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func UpdateMaterial(c *gin.Context) {
	req := protocol.UpdateMaterialReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.UpdatedMaterial(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `UpdateMaterial err`, err)
		return
	}

	res := protocol.UpdateMaterialRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func DeleteMaterial(c *gin.Context) {
	req := protocol.DeleteMaterialReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.DeleteMaterial(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `DeleteMaterial err: `, err)
		return
	}

	res := protocol.DeleteMaterialRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}
