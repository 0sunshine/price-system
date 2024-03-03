<template>
  <div>

    <el-card shadow="never" class="border-0">


      <!-- 新增|刷新 -->
      <ListHeader layout="create,refresh" @create="handleCreate" @refresh="getData">

        <el-button type="danger" size="small" @click="handleMultiDelete"
          v-if="searchForm.tab != 'delete'">批量删除</el-button>
      </ListHeader>

      <el-table ref="multipleTableRef" @selection-change="handleSelectionChange" :data="tableData" stripe
        style="width: 100%" row-key="project_kind_id" v-loading="loading" default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
        <el-table-column type="selection" width="55" />
        <el-table-column label="项目类型名称" prop="project_name" width="300" />
        <el-table-column label="备注" prop="notes" />
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <el-button text type="primary" size="small" @click.stop="addChild(scope.row.project_kind_id)">增加</el-button>
            <el-button text type="primary" size="small" @click.stop="handleEdit(scope.row)">修改</el-button>
            <el-popconfirm title="是否要删除该商品？" confirmButtonText="确认" cancelButtonText="取消"
              @confirm="handleDelete([scope.row.project_kind_id])">
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
          <el-form-item label="父项目" prop="parent_project_kind_id">
            <el-cascader v-model="form.parent_project_kind_id" :options="resList"
              :props="{ value: 'project_kind_id', label: 'project_name', children: 'children', checkStrictly: true, emitPath: false }"
              placeholder="空项目" disabled/>
          </el-form-item>
          <el-form-item label="项目名称" prop="project_name">
            <el-input v-model="form.project_name" placeholder="请输入项目名称"></el-input>
          </el-form-item>
          <el-form-item label="备注" prop="notes">
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
  getProjectKindList,
  createProjectKind,
  updateProjectKind,
  deleteProjectKind,
} from "~/api/project_kind"

import { useInitTable, useInitForm } from '~/composables/useCommon.js'

import {
  toast
} from "~/composables/util"

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
  getList: getProjectKindList,
  onGetListSuccess: (res) => {
    tableData.value = res.list
    resList.value = res.list
    total.value = res.totalCount
  },
  delete: deleteProjectKind,
  selectionChange: (e) => e.map(o => o.project_kind_id)
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
    parent_project_kind_id: "",
    project_kind_id: "",
    project_name: "",
    notes: "",
  },
  getData,
  update: updateProjectKind,
  create: createProjectKind
})

// 添加子分类
const addChild = (project_kind_id) => {
  handleCreate()
  form.parent_project_kind_id = project_kind_id
}

</script>