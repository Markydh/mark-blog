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
        <a href="https://mail.google.com" class="iconfont" :id="gmail" @mouseover="hoverEvent('gmail')" @mouseleave="hoverEvent('gmail')">&#xe67c;</a>
        <a href="https://www.instagram.com/" class="iconfont" :id="ins" @mouseover="hoverEvent('ins')" @mouseleave="hoverEvent('ins')">&#xe736;</a>
        <a href="https://discord.com/channels/@me" class="iconfont" :id="discord" @mouseover="hoverEvent('discord')" @mouseleave="hoverEvent('discord')">&#xe618;</a>
      </div>
      <hr style="border: 1px solid rgba(178,178,178,0.31); margin-top: 50px">
    </div>

    <div class="my-project">
      <h1>💻我的项目</h1>
      <div class="project-content">
        <div :class="projectDetail" @mouseenter="hoverEvent('projectDetail')" @mouseleave="hoverEvent('projectDetail')">
            <div class="project-name" >
              <i class="iconfont" style="font-size: 27px;color: #60bddc">&#xe799;</i>
            </div>
          <div class="project-itd">
          </div>
          <div class="project-tags">
          </div>
        </div>
      </div>
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
      projectDetail: 'project-detail-nonActive'
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
  height: 300px;
  display: grid;
  grid-template-columns: repeat(2,50%);
  grid-template-rows: repeat(3,50%);
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
}


.project-detail-Active{
  display: flex;
  justify-content: flex-start;
  background-color: rgba(129, 129, 129, 0.32);
  flex-direction: column;
  border: 2px  solid rgba(103, 103, 103, 0.76);
  border-radius:4px;
}


.project-name{
  padding-left: 20px;
  padding-top: 20px;
}
.project-itd{
  padding-left: 20px;
  padding-top: 20px;
}
.project-tags{
  padding-left: 20px;
  padding-top: 20px;
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
</style>