import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
      //主页
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
          //文章
        {
          path:'/article',
          name:'Article',
          component: () => import('../views/IndexView/Article.vue'),
          redirect:'monthArticle',
          children:[
              // 月刊
            {
              path:'/monthArticle',
              name:'MonthArticle',
              component: () => import('../views/ArticleViews/MonthArticle.vue')
            },
              //年刊
            {
              path:'/yearArticle',
              name:'YearArticle',
              component: () => import('../views/ArticleViews/YearArticle.vue')
            },
              //随笔
            {
              path:'/freeArticle',
              name:'FreeArticle',
              component: () => import('../views/ArticleViews/FreeArticle.vue')
            },
              //编程
            {
              path:'/codeArticle',
              name:'CodeArticle',
              component: () => import('../views/ArticleViews/CodeArticle.vue')
            },
              //生活
            {
              path:'/lifeArticle',
              name:'LifeArticle',
              component: () => import('../views/ArticleViews/LifeArticle.vue')
            },
              //文章详情页
            {
              path:'/articleDetail',
              name:'ArticleDetail',
              component:() => import('../views/ArticleViews/ArticleDetail.vue')
            }

          ]
        },
          //标签
        {
          path:'tags',
          name:'Tags',
          component: () => import('../views/IndexView/Tags.vue')
        },
          //时间轴
        {
          path:'timeLine',
          name:'TimeLine',
          component: () => import('../views/IndexView/TimeLine.vue')
        },
          //照片墙
        {
          path:'photoWall',
          name:'PhotoWall',
          component: () => import('../views/IndexView/PhotoWall.vue')
        },
        //照片墙内容
        {
          path:'/photoWallDetails',
          name:'PhotoWallDetails',
          component:()=>import('../views/PhotoWallViews/PhotoWallDetails.vue')
        },
          //主页搜索
        {
          path:'search',
          name:'Search',
          component: () => import('../views/IndexView/Search.vue')
        }
      ]
    },
    //社区
    {
      path:'/community',
      name:'Community',
      component: () => import('../views/CommunityView/Community.vue')
    },
  ]
})

export default router
