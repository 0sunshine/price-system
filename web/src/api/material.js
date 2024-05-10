import { Console } from 'windicss/utils'
import axios from '~/axios'
import { queryParams } from "~/composables/util"


export function getMaterialList(page, page_size = 10, query = {}) {
    let filter = {}
    filter = query
    filter["page-size"] = page_size
    filter["current-page"] = page

    return axios.post(`/admin/getMaterialList`, {
        filter
    })
}

export function createMaterial(data) {
    return axios.post(`/admin/addMaterial`,data)
}

export function updateMaterial(data){
    return axios.post(`/admin/updateMaterial`,data)
}

export function deleteMaterial(material_id){
    return axios.post(`/admin/deleteMaterial`,{
        material_id
    })
}