<template>
  <div class="box">
  <div class="container">
    <div class="itd-myself">
      <p style="font-size: 45px">👋 Hello!</p>
      <p><span style="font-size: 26px">我是 <strong>杨典宏 (mark)</strong></span>，一名大三学生学生‍👨🏼‍🎓，就读于南昌大学🎓</p>
      <p>业务爱好众多，例如打篮球🏀，踢足球⚽️，喜欢一个人傍晚🌆散步，夜跑等等</p>
      <p>擅长的编程语言：C C++ Java Rust</p>
      <p>也使用过springboot做过几个前后端分离的项目</p>
      <p>除了编码之外，目前做手开发自己的个人博客，你可以在这个网站上浏览我所发布的文章。</p>
      <h1>🎨技术栈</h1>
      <p style="font-size: large"><strong>前端</strong></p>
      <p>Vue3</p>
      <p>element-plus</p>
      <p>Vuetify</p>
      <p style="font-size: large"><strong>后端(还未开发后端，打算使用的技术栈)</strong></p>
      <p>springboot</p>
      <p>lombok</p>
      <div style="display: inline-flex;justify-content: flex-start">
        <a href="https://github.com/Markydh" class="iconfont" style="font-size: 30px;padding-left: 5px">&#xe709;</a>
        <a href="https://mail.google.com" class="iconfont" style="font-size: 35px;padding-left: 8px">&#xe67c;</a>
        <a href="https://www.instagram.com/" class="iconfont" style="font-size: 35px;padding-left: 8px">&#xe736;</a>
        <a href="https://discord.com/channels/@me" class="iconfont" style="font-size: 35px;padding-left: 8px">&#xe618;</a>
      </div>
    </div>
    <div class="test1">

    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: "Index.vue"
}
</script>

<style scoped>
.box{
  display: flex;
  justify-content: center;
}
.container{
  width: 55%;
  height: auto;
  display: flex;
  flex-direction: column;
}
.itd-myself{
  width:100%;
  height: auto;
}
.p{
  font-family: "Hiragino Sans GB";
}
a{
  text-decoration: none;
  color: #000;
}
</style>