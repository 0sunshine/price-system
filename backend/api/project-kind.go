package api

import (
	"backend/db"
	"backend/protocol"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProjectKindList(c *gin.Context) {
	req := protocol.GetProjectKindListReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	//fmt.Printf("%+v\n", req)
	//tmp, _ := json.Marshal(req)
	//fmt.Println(string(tmp))

	parentProjectKindId := ""
	list, err := db.GetProjectKindList(parentProjectKindId)
	if err != nil {
		c.String(http.StatusInternalServerError, `GetProjectKindList err`)
		return
	}

	res := protocol.GetProjectKindListRes{}

	res.Data.List = list
	res.Data.TotalCount = len(list)

	c.JSON(http.StatusOK, res)
}

func AddProjectKind(c *gin.Context) {
	req := protocol.AddProjectKindListReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.AddProjectKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `AddProjectKind err`, err)
		return
	}

	res := protocol.AddProjectKindListRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func UpdateProjectKind(c *gin.Context) {
	req := protocol.UpdateProjectKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.UpdatedProjectKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `UpdateProjectKind err`, err)
		return
	}

	res := protocol.UpdateProjectKindListRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}

func DeleteProjectKind(c *gin.Context) {
	req := protocol.DeleteProjectKindReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.String(http.StatusBadRequest, `the body err`)
		return
	}

	fmt.Printf("%+v\n", req)

	err = db.DeleteProjectKind(req)
	if err != nil {
		c.String(http.StatusInternalServerError, `DeleteProjectKind err: `, err)
		return
	}

	res := protocol.DeleteProjectKindListRes{}
	res.Code = 0

	c.JSON(http.StatusOK, res)
}
