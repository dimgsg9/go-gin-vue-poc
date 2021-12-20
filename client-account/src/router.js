import { createRouter, createWebHistory } from 'vue-router';

import Home from './views/Home.vue'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import NotFound from './views/NotFound.vue'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
        // beforeEnter: requireAuth,
    },
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/:catchAll(.*)*',
        name: 'notFound',
        component: NotFound,
    },
];

const router = createRouter({
    // env variable provided base "base" key of vite config
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;