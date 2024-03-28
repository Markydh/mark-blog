import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
      redirect:'/index',
      children:[
        {
          path:'index',
          name:'Index',
          component: () => import('../views/IndexView/Index.vue')
        },
        {
          path:'article',
          name:'Article',
          component: () => import('../views/IndexView/Article.vue')
        },
        {
          path:'tags',
          name:'Tags',
          component: () => import('../views/IndexView/Tags.vue')
        },
        {
          path:'timeLine',
          name:'TimeLine',
          component: () => import('../views/IndexView/TimeLine.vue')
        },
        {
          path:'photoWall',
          name:'PhotoWall',
          component: () => import('../views/IndexView/PhotoWall.vue')
        },
        {
          path:'search',
          name:'Search',
          component: () => import('../views/IndexView/Search.vue')
        }
      ]
    }
  ]
})

export default router
