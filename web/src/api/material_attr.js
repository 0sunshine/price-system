import { Console } from 'windicss/utils'
import axios from '~/axios'
import { queryParams } from "~/composables/util"


export function getMaterialAttrList(page, page_size = 10, query = {}) {
    let filter = {}
    filter = query
    filter["page-size"] = page_size
    filter["current-page"] = page

    return axios.post(`/admin/getMaterialAttrList`, {
        filter
    })
}

export function createMaterialAttr(data) {
    return axios.post(`/admin/addMaterialAttr`,data)
}

export function updateMaterialAttr(data){
    return axios.post(`/admin/updateMaterialAttr`,data)
}

export function deleteMaterialAttr(material_attr_id){
    return axios.post(`/admin/deleteMaterialAttr`,{
        material_attr_id
    })
}