import { createRouter, createWebHistory } from 'vue-router';

import Home from './views/Home.vue'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import NotFound from './views/NotFound.vue'
import { createApp } from 'vue';

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/register',
        name: 'Register',
        component: Register,
        // beforeEnter: requireAuth,
    },
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/:catchAll(.*)*',
        name: 'NotFound',
        component: NotFound,
    },
];

const router = createRouter({
    // env variable provided base "base" key of vite config
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;