<template>
    <div class="f-menu">
        <el-menu :default-openeds="['1']" default-active="2" class="border-0" @select="handleSelect">

            <template v-for="(item,index) in asideMenus" :key="index">
                <el-sub-menu v-if="item.child && item.child.length > 0" :index="item.index">
                    <template #title>
                        <el-icon>
                            <component :is="item.icon"></component>
                        </el-icon>
                        <span>{{ item.name }}</span>
                    </template>
                    <el-menu-item v-for="(item2,index2) in item.child" :key="index2" :index="item2.frontpath">
                        <el-icon>
                            <component :is="item2.icon"></component>
                        </el-icon>
                        <span>{{ item2.name }}</span>
                    </el-menu-item>
                </el-sub-menu>

                <el-menu-item v-else :index="item.frontpath">
                    <el-icon>
                         <component :is="item.icon"></component>
                    </el-icon>
                    <span>{{ item.name }}</span>
                </el-menu-item>
            </template>
        </el-menu>
    </div>
</template>
<script setup>
import { useRouter } from 'vue-router';
const router = useRouter()
const asideMenus = [{
    "name": "系统",
    "icon": "shopping-bag",
    "index": "1",
    "child": [
        {
        "name": "项目类型管理",
        "icon": "shopping-cart-full",
        "frontpath": "/project_kind/list",
        },
        {
            "name": "材料类型管理",
            "icon": "shopping-cart-full",
            "frontpath": "/material_kind/list",
        },
        {
            "name": "材料规格管理",
            "icon": "shopping-cart-full",
            "frontpath": "/material_attr/list",
        },
        {
            "name": "材料管理",
            "icon": "shopping-cart-full",
            "frontpath": "/material/list",
        }
    ]
}]

const handleSelect = (e)=>{
    router.push(e)
}
</script>
<style>
.f-menu {
    width: 250px;
    top: 64px;
    bottom: 0;
    left: 0;
    overflow: auto;
    @apply shadow-md fixed bg-light-50;
}
</style>