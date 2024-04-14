package main

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)
import "backend/db"

func main() {

	err := db.InitDB()
	if err != nil {
		return
	}

	r := gin.Default()

	r.POST("/admin/getProjectKindList", api.GetProjectKindList)
	r.POST("/admin/addProjectKind", api.AddProjectKind)
	r.POST("/admin/updateProjectKind", api.UpdateProjectKind)
	r.POST("/admin/deleteProjectKind", api.DeleteProjectKind)

	r.POST("/admin/getMaterialKindList", api.GetMaterialKindList)
	r.POST("/admin/addMaterialKind", api.AddMaterialKind)
	r.POST("/admin/updateMaterialKind", api.UpdateMaterialKind)
	r.POST("/admin/deleteMaterialKind", api.DeleteMaterialKind)

	r.POST("/admin/getMaterialAttrList", api.GetMaterialAttrList)
	r.POST("/admin/addMaterialAttr", api.AddMaterialAttr)
	r.POST("/admin/updateMaterialAttr", api.UpdateMaterialAttr)
	r.POST("/admin/deleteMaterialAttr", api.DeleteMaterialAttr)

	r.Run(":8888")
}
