import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Homepage',
        component: () => import('@pages/Homepage.tsx')
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
