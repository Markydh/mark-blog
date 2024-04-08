<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <div class="container">
          <div class="nameIndex">
            <router-link to="/">Mark's Blog</router-link>
          </div>
          <div class="tags">
            <p><router-link to="/article">文章</router-link></p>
            <p><router-link to="/tags">标签</router-link></p>
            <p><router-link to="/timeLine">时间轴</router-link></p>
            <p><router-link to="/photoWall">照片墙</router-link></p>
            <p><router-link to="/search">搜索</router-link></p>
          </div>
        </div>
      </el-header>
      <hr style="border: 1px solid rgba(178,178,178,0.31);width: 55%">
      <el-main>
        <RouterView />
      </el-main>
      <el-footer style="display:flex;justify-content: space-between;width: 100%">
          <div><span class="copyright"><a href="https://beian.miit.gov.cn/" target="_blank">赣ICP备2024029658号-1</a></span></div>
          <div><span style="padding-right: 20px">Copyright©markydh 2024-2025</span></div>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { RouterView } from 'vue-router'
export  default {
  components:{
    RouterView
  }
}

</script>
<style scoped>
.el-header{
  display: flex;
  justify-content: center;
}
.el-main {
  height: auto;
}
.el-footer{
  display: flex;
  justify-content: center;
}
.el-footer div{
  display: flex;
  flex-direction: column-reverse
}

/*导航条内容*/
.container{
  width: 55%;
  height: auto;
  display: flex;
  justify-content: space-between;
}
.nameIndex{
  display: inline-flex;
  justify-content: center;
}
.nameIndex a{
  position: relative;
  top: 20%;
  font-family: "Hiragino Sans GB";
  font-weight: bold;
  font-size: 25px;
}

.tags{
  display: inline-flex;
  justify-content: flex-end;
}

.tags p{
  font-family: "Hiragino Sans GB";
  font-size: 18px;
  font-weight: inherit;
  padding-left: 20px;
}

a{
  text-decoration: none;
  color: #000;
}
</style>
