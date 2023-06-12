import {
    createRouter,
    createWebHashHistory
} from 'vue-router';
const login =()=>import("@/views/login.vue")
const index=()=>import("@/views/index.vue")
const routes=[
    {
        path:"/",
        component:login,
    },
    {
        path:"/index",
        component:index,
        children:[
            {
                path:"echarts",
                component:()=>import("@/views/computer/index.vue")
            },
            {
                path:"info",
                component:()=>import("@/views/personInfo/index.vue")
            }
        ]
    }
]
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router;