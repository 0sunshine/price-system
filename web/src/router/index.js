import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import Admin from "~/layouts/admin.vue";

import ProjectKind from "~/pages/Projectkind/list.vue"
import MaterialKind from "~/pages/MaterialKind/list.vue"
import MaterialAttr from "~/pages/MaterialAttr/list.vue"

const routes = [
    {
        path: "/",
        component: Admin,
        // 子路由
        children: [
            {
                path: "/project_kind/list",
                component: ProjectKind,
            },
            {
                path: "/material_kind/list",
                component: MaterialKind,
            },
            {
                path: "/material_attr/list",
                component: MaterialAttr,
            },
            // {
            //     path: "/material/list",
            //     component: HelloWorld,
            // },
        ]    
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router