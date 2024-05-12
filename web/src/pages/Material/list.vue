<template>
  <div>

    <el-card shadow="never" class="border-0">
      <!-- 搜索 -->
      <Search :model="searchForm" @search="getData" @reset="resetSearchForm">
        <SearchItem label="搜索材料">
          <el-input v-model="searchForm.material_name_like" placeholder="模糊查找材料名" clearable></el-input>
        </SearchItem>

        <template #show>
          <SearchItem label="所属项目">
            <el-cascader v-model="searchForm.parent_project_kind_id_like" :options="ProjectList"
              :props="{ value: 'project_kind_id', label: 'project_name', children: 'children', checkStrictly: true, emitPath: false }"
              placeholder="选择所属项目" />
          </SearchItem>
        </template>
      </Search>

      <!-- 新增|刷新 -->
      <ListHeader layout="create,refresh" @create="handleCreateProxy" @refresh="getData">
        <el-popconfirm title="是否批量删除材料？" confirmButtonText="确认" cancelButtonText="取消" @confirm="handleMultiDelete">
          <template #reference>
            <el-button type="danger" size="small" v-if="searchForm.tab != 'delete'">批量删除</el-button>
          </template>
        </el-popconfirm>
      </ListHeader>

      <el-table ref="multipleTableRef" @selection-change="handleSelectionChange" :data="tableData" stripe
        style="width: 100%" row-key="material_id" v-loading="loading" default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
        <el-table-column type="selection" width="55" />
        <el-table-column label="材料名称" prop="material_name" width="300" />
        <el-table-column label="材料价格" prop="price" width="300" />
        <el-table-column label="操作" prop="extra_price" width="300">
          <template #default="scope">
            <el-button text type="primary" size="small" @click.stop="handleEditProxy(scope.row)">修改</el-button>
            <el-popconfirm title="是否要删除该材料？" confirmButtonText="确认" cancelButtonText="取消"
              @confirm="handleDelete([scope.row.material_id])">
              <template #reference>
                <el-button class="px-1" text type="primary" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>

      <div class="flex items-center justify-center mt-5">
        <el-pagination background layout="prev, pager ,next" :total="total" :current-page="currentPage"
          :page-size="limit" @current-change="getData" />
      </div>

      <FormDrawer ref="formDrawerRef" :title="drawerTitle" @submit="handleSubmit">
        <el-form :model="form" ref="formRef" :rules="rules" label-width="80px" :inline="false">
          <el-form-item label="所属材料" prop="material_kind_id">
            <el-cascader v-model="form.material_kind_id" :options="MaterialKindList"
              :props="{ value: 'material_kind_id', label: 'material_name', children: 'children', checkStrictly: true, emitPath: false }"
              placeholder="空项目" @change="handleMaterialKindIdChanged" :disabled=CascaderEnableRef />
          </el-form-item>
          <el-form-item label="规格" prop="attr_descs">
            <el-cascader v-model="MaterialAttrCascader" :options="MaterialKindAttrList"
              :props="{ value: 'attr_id', label: 'attr_desc', checkStrictly: true, emitPath: false, multiple: true }"
              placeholder="规格列表" @change="handleMaterialAttrChanged" :disabled=CascaderEnableRef />
          </el-form-item>
          <el-form-item label="材料价格" prop="price">
            <el-input v-model="form.price" placeholder="选填，备注" style="width: 180px"></el-input>
          </el-form-item>
          <!-- <el-form-item label="材料其他格外价格" prop="extra_price">
            <el-input type="textarea" v-model="form.extra_price" placeholder="选填，备注"></el-input>
          </el-form-item> -->
        </el-form>
      </FormDrawer>

    </el-card>

  </div>
</template>

<script setup>
import { ref } from "vue"
import ListHeader from "~/components/ListHeader.vue";
import FormDrawer from "~/components/FormDrawer.vue";
import Search from "~/components/Search.vue";
import SearchItem from "~/components/SearchItem.vue";

import {
  getMaterialList,
  createMaterial,
  updateMaterial,
  deleteMaterial,
} from "~/api/material"

import {
  getProjectKindList
} from "~/api/project_kind"

import {
  getMaterialKindList
} from "~/api/material_kind"

import {
  getMaterialAttrList,
  createMaterialAttr,
  updateMaterialAttr,
  deleteMaterialAttr,
} from "~/api/material_attr"

import { useInitTable, useInitForm } from '~/composables/useCommon.js'

import {
  toast
} from "~/composables/util"

const MaterialKindList = ref([])
const CascaderEnableRef = ref(false)

const MaterialKindAttrList = ref([])
const MaterialAttrCascader = ref([])

const ProjectList = ref([])

const {
  handleSelectionChange,
  multipleTableRef,
  handleMultiDelete,

  searchForm,
  resetSearchForm,
  tableData,
  loading,
  currentPage,
  total,
  limit,
  getData,
  handleDelete,

  multiSelectionIds
} = useInitTable({
  searchForm: {
    material_name_like: "",
    parent_project_kind_id_like:"",
  },
  getList: getMaterialList,
  onGetListSuccess: (res) => {

    getProjectKindList(1, 10000).then((response) => {
      ProjectList.value = response.list
    });

    getMaterialKindList(1, 10000).then((response) => {
      MaterialKindList.value = response.list
    });

    tableData.value = res.list

    total.value = res.totalCount
  },
  delete: deleteMaterial,
  selectionChange: (e) => e.map(o => o.material_id)
})

const resList = ref([])
const {
  formDrawerRef,
  formRef,
  form,
  rules,
  drawerTitle,
  handleSubmit,
  handleCreate,
  handleEdit
} = useInitForm({
  form: {
    material_id: -1,
    material_kind_id: -1,
    material_name: "",
    price: "",
    extra_price: "",
    attr_descs: []
  },
  getData,
  update: updateMaterial,
  create: createMaterial
})


const handleEditProxy = (row) => {
  handleEdit(row)
  CascaderEnableRef.value = true

  MaterialKindAttrList.value = []
  handleMaterialKindIdChanged()

  MaterialAttrCascader.value = []
  for (var attr of form.attr_descs) {
    MaterialAttrCascader.value.push(attr.attr_id);
  }
}

const handleCreateProxy = (row) => {
  handleCreate(row)
  CascaderEnableRef.value = false

  MaterialKindAttrList.value = []
  MaterialAttrCascader.value = []
}

const handleMaterialKindIdChanged = async () => {

  await getMaterialAttrList(1, 10000, { material_kind_id: form.material_kind_id}).then((response) => {
    MaterialKindAttrList.value = response.list
  });
  
  console.log(MaterialKindAttrList)
}


const handleMaterialAttrChanged = () => {

  form.attr_descs = []

  for (var attr_id of MaterialAttrCascader.value){
    for (var itm of MaterialKindAttrList.value) {
      if (itm.attr_id == attr_id) {
        form.attr_descs.push({
          attr_id: attr_id,
          material_kind_id: form.material_kind_id,
          material_name: form.material_name,
          attr_desc: itm.attr_desc
        })
      }
    }
  }
}

</script>