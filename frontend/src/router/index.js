import {
    createRouter,
    createWebHashHistory
} from 'vue-router';
import Global from "../Global/index.js"
const index=()=>import("@/views/index.vue")
const routes=[
    {
        path:"/",
        component:index,
        children:[
            {
                path:"computerInfo",
                component:()=>import("@/views/computer/index.vue")
            },
            {
                path:"PCHardware",
                component:()=>import("@/views/PCHardware/index.vue")
            },
            {
                path:"Manual",
                component:()=>import("@/views/manual/index.vue")
            },
            {
                path:"Settings",
                component:()=>import("@/views/settings/index.vue")
            }
        ]
    }
]
const router = createRouter({
    history: createWebHashHistory(),
    routes
})
router.beforeEach((to,from)=>{
Global.activeUrl=to.fullPath;
})
export default router;