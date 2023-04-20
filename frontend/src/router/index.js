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
    }
]
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router;