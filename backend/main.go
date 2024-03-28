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

	r.Run(":8888")
}
