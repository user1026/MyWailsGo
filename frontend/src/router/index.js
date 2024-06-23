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
                //介绍
                path:"Introduction",
                component:()=>import("@/views/Introduction/index.vue")
            },
            {
                //挑选
                path:"PCHardware",
                component:()=>import("@/views/PCHardware/index.vue")
            },
            {   
                //概念介绍
                path:"Manual",
                component:()=>import("@/views/manual/index.vue")
            },
            {
                //设置
                path:"Settings",
                component:()=>import("@/views/settings/index.vue")
            },
            {
                //装机大法
                path:"Warning",
                component:()=>import("@/views/Warning/index.vue")
            },
            {
                //购买
                path:"Pay",
                component:()=>import("@/views/Pay/index.vue")
            },
            {
                //到货
                path:"Arrival",
                component:()=>import("@/views/Arrival/index.vue")
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