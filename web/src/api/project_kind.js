import { Console } from 'windicss/utils'
import axios from '~/axios'
import { queryParams } from "~/composables/util"


export function getProjectKindList(page, query = {}) {
    let filter = {}
    filter = query
    filter["page-size"] = 10
    filter["current-page"] = page

    return axios.post(`/admin/getProjectKindList`, {
        filter
    })
}

export function createProjectKind(data) {
    return axios.post(`/admin/addProjectKind`,data)
}

export function updateProjectKind(data){
    return axios.post(`/admin/updateProjectKind`,data)
}

export function deleteProjectKind(project_kind_id){
    return axios.post(`/admin/deleteProjectKind`,{
        project_kind_id
    })
}