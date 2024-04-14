<template>
  <div>

    <el-card shadow="never" class="border-0">


      <!-- 新增|刷新 -->
      <ListHeader layout="create,refresh" @create="handleCreateProxy" @refresh="getData">

        <el-button type="danger" size="small" @click="handleMultiDelete"
          v-if="searchForm.tab != 'delete'">批量删除</el-button>
      </ListHeader>

      <el-table ref="multipleTableRef" @selection-change="handleSelectionChange" :data="tableData" stripe
        style="width: 100%" row-key="material_kind_id" v-loading="loading" default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
        <el-table-column type="selection" width="55" />
        <el-table-column label="所属材料" prop="material_name" width="300" />
        <el-table-column label="规格描述" prop="attr_desc" width="300" />
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <el-button text type="primary" size="small"
              @click.stop="addChild(scope.row.material_kind_id)">增加</el-button>
            <el-button text type="primary" size="small" @click.stop="handleEditProxy(scope.row)">修改</el-button>
            <el-popconfirm title="是否要删除该材料？" confirmButtonText="确认" cancelButtonText="取消"
              @confirm="handleDelete([scope.row.material_kind_id])">
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
            <el-cascader v-model="form.material_kind_id" :options="MaterialList"
              :props="{ value: 'material_kind_id', label: 'material_name', children: 'children', checkStrictly: true, emitPath: false }"
              placeholder="空项目" :disabled=CascaderEnableRef />
          </el-form-item>
          <el-form-item label="规格" prop="attr_desc">
            <el-input type="textarea" v-model="form.notes" placeholder="选填，备注"></el-input>
          </el-form-item>
        </el-form>
      </FormDrawer>

    </el-card>

  </div>
</template>

<script setup>
import { ref } from "vue"
import ListHeader from "~/components/ListHeader.vue";
import FormDrawer from "~/components/FormDrawer.vue";

import {
  getMaterialAttrList,
  createMaterialAttr,
  updateMaterialAttr,
  deleteMaterialAttr,
} from "~/api/material_attr"

import {
  getMaterialKindList
} from "~/api/material_kind"

import { useInitTable, useInitForm } from '~/composables/useCommon.js'

import {
  toast
} from "~/composables/util"

const MaterialList = ref([])
const CascaderEnableRef = ref(false)

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
    title: "",
    tab: "all",
    category_id: null,
  },
  getList: getMaterialAttrList,
  onGetListSuccess: (res) => {

    getMaterialKindList(1, 10000).then((response) => {
      MaterialList.value = response.list
    });

    tableData.value = res.list

    total.value = res.totalCount
  },
  delete: deleteMaterialAttr,
  selectionChange: (e) => e.map(o => o.material_kind_id)
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
    attr_id: -1,
    material_kind_id: -1,
    material_name:"",
    attr_desc: "",
  },
  getData,
  update: updateMaterialAttr,
  create: createMaterialAttr
})

// 添加子分类
const addChild = (project_kind_id) => {
  handleCreate()
  form.project_kind_id = project_kind_id
}

const handleEditProxy = (row) => {
  handleEdit(row)
  CascaderEnableRef.value = true
}

const handleCreateProxy = (row) => {
  handleCreate(row)
  CascaderEnableRef.value = false
}


</script>