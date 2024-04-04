<template>
  <div class="container">
  <div class="aside">
    <div style="position:relative;left: 40%;top:5%;display: flex;flex-direction: column">
    <img src="/images/markAvatar.png" alt="" style="border-radius: 10%;width: 100px;height: 100px;">
    <h><span style="font-family: 'Hiragino Sans GB';font-size: 22px;font-weight: bold;position:relative;top:12px;right: 15px">Mark 的博客</span></h>
      <div class="subscribe">
        Subscribe
      </div>
      <div v-for="(item,index) in type" :key="index" style="margin-bottom: -22px;position:relative;top: 8%;width: 40%;">
        <p><router-link :to="item.link">{{item.name}}</router-link></p>
      </div>
    </div>
  </div>
    <div class="content">
      <RouterView/>
    </div>
  </div>
</template>

<script>
import { RouterView } from 'vue-router'
export default {
  name: "Article",
  components:{
    RouterView
  },
  data(){
    return{
      type:[
        {
          name: "📰 月刊",
          link:'/monthArticle',
        },
        {
          name: "🏆年报",
          link:"/yearArticle",
        },
        {
          name: "✏️ 随笔",
          link:"/freeArticle",
        },
        {
          name: "💻 编程",
          link:"/codeArticle",
        },
        {
          name: "🏀生活",
          link:"/lifeArticle",
        }
      ]
    }
  }
}
</script>

<style scoped>
.container{
  width: 100%;
  height: auto;
  display: flex;
  justify-content: flex-start;
}

.aside{
  width: 22%;
  height: 800px;
  display: flex;
  flex-direction: column;
}
.content{
  width: 57%;
  height: auto;
}
/*订阅按钮样式*/
.subscribe{
  width: 100px;
  height: 20px;
  position: relative;
  top:20px;
  right: 16px;
  border: 2px solid black;
  border-radius: 45px;
  text-align: center;
  font-family: 'Hiragino Sans GB';
  font-size: small;
  font-weight:bold;
}
a{
  color:black;
  text-decoration:none;
}
</style>