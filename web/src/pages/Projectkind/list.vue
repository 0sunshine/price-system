<template>
  <div>

    <el-card shadow="never" class="border-0">


      <!-- 新增|刷新 -->
      <ListHeader layout="create,refresh" @create="handleCreate" @refresh="getData">

        <el-button type="danger" size="small" @click="handleMultiDelete"
          v-if="searchForm.tab != 'delete'">批量删除</el-button>
        <el-button type="warning" size="small" @click="handleRestoreGoods" v-else>恢复商品</el-button>

        <el-popconfirm v-if="searchForm.tab == 'delete'" title="是否要彻底删除该商品？" confirmButtonText="确认" cancelButtonText="取消"
          @confirm="handleDestroyGoods">
          <template #reference>
            <el-button type="danger" size="small">彻底删除</el-button>
          </template>
        </el-popconfirm>

      </ListHeader>

      <el-table ref="multipleTableRef" @selection-change="handleSelectionChange" :data="tableData" stripe
        style="width: 100%" v-loading="loading">
        <el-table-column type="selection" width="55" />
        <el-table-column label="项目名称" prop="project_name" width="300" />
        <el-table-column label="备注" prop="notes" />
      </el-table>

      <div class="flex items-center justify-center mt-5">
        <el-pagination background layout="prev, pager ,next" :total="total" :current-page="currentPage" :page-size="limit"
          @current-change="getData" />
      </div>

      <FormDrawer ref="formDrawerRef" :title="drawerTitle" @submit="handleSubmit">
        <el-form :model="form" ref="formRef" :rules="rules" label-width="80px" :inline="false">
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
  getGoodsList,
  updateGoodsStatus,
  createGoods,
  updateGoods,
  deleteGoods,
  restoreGoods,
  destroyGoods
} from "~/api/goods"
import {
  getCategoryList
} from "~/api/category"

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
  handleMultiStatusChange,

  multiSelectionIds
} = useInitTable({
  searchForm: {
    title: "",
    tab: "all",
    category_id: null,
  },
  getList: getGoodsList,
  onGetListSuccess: (res) => {
    tableData.value = res.list.map(o => {
      o.bannersLoading = false
      o.contentLoading = false
      o.skusLoading = false
      return o
    })
    total.value = res.totalCount
  },
  delete: deleteGoods,
  updateStatus: updateGoodsStatus
})

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
    project_kind_id: -1,
    project_name: "", 
    parent_project_kind_id: "",
    notes: "",
  },
  getData,
  update: updateGoods,
  create: createGoods
})



const handleRestoreGoods = () => useMultiAction(restoreGoods, "恢复")

const handleDestroyGoods = () => useMultiAction(destroyGoods, "彻底删除")
function useMultiAction(func, msg) {
  loading.value = true
  func(multiSelectionIds.value)
    .then(res => {
      toast(msg + "成功")
      // 清空选中
      if (multipleTableRef.value) {
        multipleTableRef.value.clearSelection()
      }
      getData()
    })
    .finally(() => {
      loading.value = false
    })
}
</script>