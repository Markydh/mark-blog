<template>
  <div class="container">
<!--    顶部-->
    <div class="head">
        <div class="logo">
          <img src="/images/logo.png" style="width: 160px;position: relative;left: 30%" alt="大同很不同">
        </div>
        <div class="head-right">
          <Search style="width: 22px;margin-left: 10px"/>
          <input type="search" style="width: 80%;position:relative;top: 12px;border: 0">
        </div>
      <el-button @click="signLogin">登陆｜注册</el-button>
    </div>


<!--    主体-->
    <div class="main">
<!--      侧边栏-->
        <div class="aside">

        </div>
<!--      内容-->
        <div class="context">
          <div  v-if="signUp" class="signDiv">
            <div class="buttonChange">
<!--              如果是注册-->
              <div v-if="func === 'sign'">
                    <button class="changeToLogin" @click="changeToRegister">登陆</button>
              </div>
<!--              如果是登陆-->
              <div v-if="func === 'login'">
                <button class="changeToSign"  @click="changeToLogin">注册</button>
              </div>
            </div>
            <div class="signMsg">
<!--              如果是注册-->
              <div v-if="func === 'sign'">
                <form  ref="loginForm" :model="userInfo" class="registerForm">
                  <h3 style="font-size: 25px;font-weight: normal;color: #595353">注册</h3>
                  <span style="color: #a8a8ab;font-size: small">用户名</span>
                  <div>
                    <input v-model="userInfo.username" type="text" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" />
                  </div>
                  <span style="color: #a8a8ab;font-size: small">邮箱</span>
                  <div>
                    <input v-model="userInfo.email" type="text" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" />
                  </div>
                  <span style="color: #a8a8ab;font-size: small">密码</span>
                  <div>
                    <input v-model="userInfo.password"  type="password" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'"/>
                  </div>
                  <span style="color: #a8a8ab;font-size: small">确认密码</span>
                  <div>
                    <input v-model="userInfo.confirmPwd"  type="password" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'"/>
                  </div>
                  <h3 style="font-size: 13px">Now User </h3>
                  <router-link to="" style="color:rgb(0,0,0);text-decoration: none;font-size: small;margin-right: 120px;">Sign up</router-link>
                  <router-link to="/" style="color:rgb(0,0,0);text-decoration: none;font-size:small">Forget Password?</router-link>
                  <el-button style="margin-top: 48px;background-color: rgb(0,0,0);color: #FFFFFF;width: 165px;height: 40px;position: relative;left: 15%" type="text" @click="doRegister()" round>注册</el-button>
                </form>
              </div>
<!--              如果是登陆-->
              <div v-if="func === 'login'">
                <form  ref="loginForm" :model="userInfo" class="loginForm">
                  <h3 style="font-size: 25px;padding-bottom: 30px;font-weight: normal;color: #595353">Sign in</h3>
                  <span style="color: #a8a8ab;font-size: small">用户名</span>
                    <div>
                    <input v-model="userInfo.username" type="text" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" />
                    </div>
                    <span style="color: #a8a8ab;font-size: small">密码</span>
                    <div>
                    <input v-model="userInfo.password"  type="password" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'"/>
                    </div>
                  <h3 style="font-size: 13px">Now User </h3>
                  <router-link to="" style="color:rgb(0,0,0);text-decoration: none;font-size: small;margin-right: 120px;">Sign Up</router-link>
                  <router-link to="/" style="color:rgb(0,0,0);text-decoration: none;font-size:small">Forget Password?</router-link>
                    <el-button style="margin-top: 48px;background-color: rgb(0,0,0);color: #FFFFFF;width: 165px;height: 40px;position: relative;left: 15%" type="text" @click="doLogin()" round> 登陆 </el-button>
                </form>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import {gsap} from "gsap";
export default {
  name: "community",
  data(){
    return{
      signUp:false,
      func:'sign',
      userInfo:{
        username:'',
        email:'',
        confirmPwd:'',
        password:'',

      }
    }
  },
  methods:{
    async signLogin(){
      console.log("点击登陆注册")
      this.signUp = true
    },
    async doLogin(){
          console.log(this.userInfo)
    },
    async doRegister(){
          console.log(this.userInfo)
    },
    changeToLogin(){
        gsap.to(".buttonChange",{x:0})
        gsap.to(".signMsg",{x:0})
        this.func='sign'
    },
    changeToRegister(){
      gsap.to(".buttonChange",{x:400})
      gsap.to(".signMsg",{x:-400})
      this.func='login'
    }
  }
}
</script>

<style scoped>

.container{
  width: 100%;
  display: flex;
  flex-direction: column
}
.head{
  display: flex;
  width: 100%;
  height: 50px;
  align-items: center;
  justify-content: space-between;
}
.main{
  width: 100%;
  display: flex;
  flex-direction: row;
}
.aside{
  width: 25%;
  background-color: #DD4A68;
}
.context{
  background-color: #63a35c;
  width: 100%;
  height: 100vh;
}

.logo{
  width: 20%;
  margin-right: 40px;
}
.head-right{
  border:1px solid #838383;
  border-radius: 7px;
  display: inline-flex;
  justify-content: flex-start;
  position: relative;
  width: 60%;
  height: 40px;
}
input{
  border: 0;
  font-size: large;
  width: 90%;
}
input:focus {
  outline: none;
}
.signDiv{
  position: absolute;
  width: 800px;
  height: 550px;
  left: 25%;
  top: 25%;
  display: flex;
  flex-direction: row;
}

.buttonChange{
  width: 50%;
  background-color: black;
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url('/images/login.jpg');
  background-size: cover;
  background-position: center;
}
.signMsg{
  width: 50%;
  background-color: #ffffff;
}

.changeToLogin{
  width: 220px;
  height: 55px;
  font-size: 20px;
  color: #656565;
  background-color: rgba(9, 9, 9, 0.8);
  font-family: "Hiragino Sans GB";
  border: 0;
  border-radius: 30px;
}
.changeToSign {
  width: 220px;
  height: 55px;
  font-size: 20px;
  color: #656565;
  background-color: rgba(9, 9, 9, 0.8);
  font-family: "Hiragino Sans GB";
  border: 0;
  border-radius: 30px;
}


.loginForm{
    left: 10%;
    padding-top: 100px;
    position: relative;
}
.registerForm{
  left: 10%;
  padding-top: 60px;
  position: relative;
}

input{
  margin-bottom: 20px;
  border:0;
  border-bottom: 1px solid #999999;
  width:280px;
}
</style>