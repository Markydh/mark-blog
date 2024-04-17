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
        <a href="https://github.com/Markydh" class="iconfont" :id="github" @mouseover="hoverEvent('github')" @mouseleave="hoverEvent('github')">&#xe709;</a>
        <a href="https://mail.google.com/mail/u/0/?fs=1&tf=cm&source=mailto&to=dianhong.yang.one@gmail.com" class="iconfont" :id="gmail" @mouseover="hoverEvent('gmail')" @mouseleave="hoverEvent('gmail')">&#xe67c;</a>
        <a href="https://www.instagram.com/" class="iconfont" :id="ins" @mouseover="hoverEvent('ins')" @mouseleave="hoverEvent('ins')">&#xe736;</a>
        <a href="https://discord.com/channels/@me" class="iconfont" :id="discord" @mouseover="hoverEvent('discord')" @mouseleave="hoverEvent('discord')">&#xe618;</a>
      </div>
      <hr style="border: 1px solid rgba(178,178,178,0.31); margin-top: 50px">
    </div>

    <div class="my-project">
      <h1>💻我的项目</h1>
      <div class="project-content">
<!--
            条件渲染 && 列表渲染 子div
            通过isActive的值来决定渲染哪种样式
-->
        <div :class="{ 'project-detail-Active': isActive === index, 'project-detail-nonActive': isActive !== index }" v-for="(item, index) in items" @mouseenter="divHoverEvent(index)" @mouseleave="divHoverEvent(null)"  @click="handleClick(item)" >
          <div class="project-name">
            <i class="iconfont" style="font-size: 27px;color: #3A8ABC">&#xe799;</i><span style="font-size: 18px;color:#3A8ABC;font-family: 'Hiragino Sans GB';padding-left: 5px">{{item.name}}</span>
          </div>
          <div class="project-itd">
            <p>{{item.tags}}</p>
          </div>
          <div class="project-tags">
            <i class="iconfont" style="margin-top: -8px;padding-right: 0px;font-size: 30px;color:#3A8ABC">&#xe604;</i>
            <p style="margin-top: 0px;padding-right: 22px">{{item.language}}</p>
            <i class="iconfont" style="margin-top: -2px;padding-right: 3px;font-size: 18px;">&#xe8b9;</i>
            <p style="margin-top: 0px;padding-right: 20px">3</p>
            <i class="iconfont" style="margin-top: -2px;padding-right: 3px;font-size: 17px;color: rgba(0,0,0,0.68)">&#xe61f;</i>
            <p style="margin-top: 0px;padding-right: 12px">6</p>
          </div>
        </div>
        </div>
       <a  href="https://github.com/Markydh"  class="a-class"  @mouseenter="aHoverEvent(true)" @mouseleave="aHoverEvent(false)">
         <span :style="{ color: linkColor }">查看更多</span>
         <i class="iconfont" id="a-icon" :style="{ color: linkColor }">&#xe600;</i>
       </a>
      <hr style="border: 1px solid rgba(178,178,178,0.31); margin-top: 60px;width: 100%;">
      </div>
    </div>
  </div>
</template>





<script>
export default {
  name: "Index.vue",
  data(){
    return {
      github:'github-nonActive',
      gmail:'gmail-nonActive',
      ins:'ins-nonActive',
      discord:'discord-nonActive',
      isActive: null,
      linkColor:'black',
      items:[
        {
          name:"mark-blog",
          tags:"mark's first blog",
          language:'HTML',
          collect:3,
          fork:0,
          url:"https://github.com/Markydh/mark-blog",
        },
        {
          name:"web-crawler",
          tags:"First attempt at crawling",
          language:'Python',
          collect:2,
          fork:3,
          url:"https://github.com/Markydh/web-crawler"
        },

        {
          name:"Learn_konwledge ",
          tags:"My Learning Document",
          language:'Markdown',
          collect: 2,
          fork:0,
          url:"https://github.com/Markydh/Learn_konwledge"
        },
        {
          name:"Leetcode",
            tags:"冲击算法大佬，从现在开始",
          language:'C++/C',
          collect: 4,
          fork:0,
          url:"https://github.com/Markydh/Leetcode"
        },
        {
          name:"Markydh",
          tags:"My personal repository",
          language:'Markdown',
          collect: 2,
          fork:0,
          url:"https://github.com/Markydh/Markydh"
        },
        {
          name:"MyPicture",
              tags:"My personal PictureBed",
            language:'Image',
            collect: 4,
            fork:0,
            url:"https://github.com/Markydh/MyPicture"
        }
      ,]
    }
  },
  methods:{
    hoverEvent(name){
      if(name=='github') {
        if (this.github == 'github-nonActive')
          this.github = 'github-Active'
        else {
          this.github = 'github-nonActive'
        }
      }
      else if(name=='gmail'){
        if (this.gmail == 'gmail-nonActive')
          this.gmail = 'gmail-Active'
        else {
          this.gmail = 'gmail-nonActive'
        }
      }
      else if(name=='ins'){
        if (this.ins == 'ins-nonActive')
          this.ins = 'ins-Active'
        else {
          this.ins = 'ins-nonActive'
        }
      }
      else if(name=='discord'){
        if (this.discord == 'discord-nonActive')
          this.discord = 'discord-Active'
        else {
          this.discord = 'discord-nonActive'
        }
      }else{
        if (this.projectDetail == 'project-detail-nonActive')
          this.projectDetail = 'project-detail-Active'
        else {
          this.projectDetail = 'project-detail-nonActive'
        }
      }
    },
    // 鼠标悬浮在项目简介时动画效果
    divHoverEvent(index) {
      this.isActive = index;
    },

    aHoverEvent(isHovered){
        this.linkColor = isHovered ? '#3A8ABC':'black'
    },


    /***
     * wind.open() 用法 window.open('url','target')
     * 其中url:想要打开的网址地址，target有三种方法
     * _blank 在新的窗口打开网址
     * _self 在当前窗口下打开网址
     * _top 框架网页中在上部窗口中显示目标网页(不常用)
     * @param item
     */
    handleClick(item){
      const  hrefLink = item.url;
      window.open(hrefLink,'_top');
    }
  }
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

.my-project{
  width: 100%;
  height: auto;
  display: flex;
  flex-direction: column;
}

.project-content{
  width: 100%;
  height: auto;
  display: grid;
  grid-template-columns: repeat(2,50%);
  column-gap: 20px;
  row-gap: 20px;
}


/*项目界面悬浮效果*/
.project-detail-nonActive{
  display: flex;
  justify-content: flex-start;
  flex-direction: column;
  background-color: white;
  border: 2px  solid rgba(103, 103, 103, 0.76);
  border-radius:4px;
  transition: transform 0.3s ease;
}

.project-detail-Active:hover {
  transform: scale(1.01); /* 鼠标悬停时放大1.05倍 */
}

.project-detail-Active{
  display: flex;
  justify-content: flex-start;
  background-color: rgba(129, 129, 129, 0.17);
  flex-direction: column;
  border: 2px  solid rgba(103, 103, 103, 0.76);
  border-radius:4px;
  transition: transform 0.3s ease;
}

.project-name{
  font-family: "Hiragino Sans GB";
  padding-top: 20px;
  padding-left: 15px;
}
.project-itd{
  font-family: "Hiragino Sans GB";
  font-size: 14px;
  color: #3a3a3a;
  padding-top: 5px;
  padding-left: 15px;
}
.project-tags{
  font-family: "Hiragino Sans GB";
  font-size: 14px;
  display: inline-flex;
  justify-content: flex-start;
  padding-left: 9px;
}

.p{
  font-family: "Hiragino Sans GB";
}
a{
  text-decoration: none;
  color: #000;
}

/*鼠标悬浮github图标事件*/
#github-nonActive{
  color: #000;
  font-size: 33px;
  padding-left: 8px
}
#github-Active{
  color: #b22ac0;
  font-size: 33px;
  padding-left: 8px
}

/*鼠标悬浮gmail图标事件*/
#gmail-nonActive{
  color: #000;
  font-size: 35px;
  padding-left: 8px
}
#gmail-Active{
  color: #15d959;
  font-size: 35px;
  padding-left: 8px
}

/*鼠标悬浮ins图标事件*/
#ins-nonActive{
  color: #000;
  font-size: 35px;
  padding-left: 8px
}
#ins-Active{
  color: #ea1186;
  font-size: 35px;
  padding-left: 8px
}


/*鼠标悬浮discord图标事件*/
#discord-nonActive{
  color: #000;
  font-size: 35px;
  padding-left: 8px
}
#discord-Active{
  color: #1b75da;
  font-size: 35px;
  padding-left: 8px
}
.a-class{
  position: relative;
  left: 43%;
  top:8%;
  font-family: 'Hiragino Sans GB';
  font-size: 17px;
}
#a-icon{
  padding-left: 6px;
  font-size: 20px;
}
</style>