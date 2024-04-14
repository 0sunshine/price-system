import { Console } from 'windicss/utils'
import axios from '~/axios'
import { queryParams } from "~/composables/util"


export function getMaterialKindList(page, page_size = 10, query = {}) {
    let filter = {}
    filter = query
    filter["page-size"] = page_size
    filter["current-page"] = page

    return axios.post(`/admin/getMaterialKindList`, {
        filter
    })
}

export function createMaterialKind(data) {
    return axios.post(`/admin/addMaterialKind`,data)
}

export function updateMaterialKind(data){
    return axios.post(`/admin/updateMaterialKind`,data)
}

export function deleteMaterialKind(material_kind_id){
    return axios.post(`/admin/deleteMaterialKind`,{
        material_kind_id
    })
}