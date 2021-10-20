import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Homepage',
        component: () => import('@pages/Homepage')
    },
    {
        path: '/server-stream',
        name: 'ServerStream',
        component: () => import('@pages/ServerStream')
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
