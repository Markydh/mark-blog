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
                <form  ref="loginForm" :model="userInfo" class="registerForm" >
                  <h3 style="font-size: 25px;font-weight: normal;color: #595353">注册</h3>
                  <span style="color: #a8a8ab;font-size: small">用户名</span>
                  <div>
                    <input v-model="userInfo.username" type="text" @input="checkUsername" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" />
                    <p v-if="usernameExist" style="color: red; font-size: small;">用户名已存在</p>
                  </div>
                  <span style="color: #a8a8ab;font-size: small">密码</span>
                  <div>
                    <input v-model="userInfo.password"  type="password" @input="checkPassword" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'"/>
                    <p v-if="passwordError" style="color: red; font-size: small;">至少包含一个大写字母、一个数字，且长度为8到20位</p>
                  </div>
                  <span style="color: #a8a8ab;font-size: small">确认密码</span>
                  <div>
                    <input v-model="userInfo.confirmPwd"  type="password" @input="checkConfirmPassword" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'"/>
                    <p v-if="confirmPasswordError" style="color: red; font-size: small;">密码不一致</p>
                  </div>
                  <span style="color: #a8a8ab;font-size: small">邮箱</span>
                  <div>
                    <input v-model="userInfo.email" type="text"  @input="checkEmail" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" />
                    <p v-if="emailError" style="color: red; font-size: small;">邮箱格式错误</p>
                  </div>
                  <div style="display: flex;flex-direction: row;padding-top: 2px">
                    <input v-model="userInfo.checkCode" type="text" onMouseOver="style.background='#eee'" onMouseOut="style.background='#fff'" style="  border:0;border-bottom: 1px solid #999999;width: 50%;margin-right: 5px" />
                    <a href="javascript: void(0)" class="btn" @click="SendMsgToUser()">
                      验证码
                    </a>
                  </div>
                  <h3 style="font-size: 13px">Now User </h3>
                  <router-link to="" style="color:rgb(0,0,0);text-decoration: none;font-size: small;margin-right: 120px;">Sign up</router-link>
                  <router-link to="/" style="color:rgb(0,0,0);text-decoration: none;font-size:small">Forget Password?</router-link>
                  <el-button style="margin-top: 10px;background-color: rgb(0,0,0);color: #FFFFFF;width: 165px;height: 40px;position: relative;left: 15%" type="text" @click="doRegister()" round>注册</el-button>
                </form>
              </div>
<!--              如果是登陆-->
              <div v-if="func === 'login'">
                <form  ref="loginForm" :model="userInfo" class="loginForm">
                  <h3 style="font-size: 25px;padding-bottom: 30px;font-weight: normal;color: #595353">登陆</h3>
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
import axios from "axios";
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
        checkCode:'',
      },
      usernameExist:false,
      passwordError:false,
      confirmPasswordError:false,
      emailError:false
    }
  },
  methods:{
    //判断用户名是否存在
    checkUsername() {
      this.request.get("/user/getUserInfo", {
        params: {
          username: this.userInfo.username
        }
      }).then(res => {
        if (res.data.code === 400) {
          this.usernameExist = true;
        } else {
          this.usernameExist = false;
        }
      }).catch(error => {
        console.error("Error occurred while checking username:", error);
        // 处理错误情况，例如显示错误提示或者其他逻辑
      });
    },

    //检查密码格式是否正确
    checkPassword() {
      // 密码格式要求：至少包含一个大写字母、一个小写字母、一个数字，且长度为8到20位
      const passwordRegex = /^(?=.*[A-Z])(?=.*\d)[A-Za-z\d@$!%*?&]{8,20}$/;
      if (!passwordRegex.test(this.userInfo.password)) {
       this.passwordError=true
      }else{
        this.passwordError=false
      }
    },
    //两次密码是否一致
    checkConfirmPassword() {
      if (this.userInfo.password !== this.userInfo.confirmPwd) {
      this.confirmPasswordError=true
      }else{
        this.confirmPasswordError=false
      }
    },
    //邮件格式是否正确
    checkEmail() {
      // 邮箱格式要求：符合常见的邮箱格式要求
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailRegex.test(this.userInfo.email)) {
        this.emailError=true
      }else{
        this.emailError=false
      }
    },

    async signLogin(){
      console.log("点击登陆注册")
      this.signUp = true
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
    },
    async SendMsgToUser(){
      console.log(this.userInfo)
     this.request.post("/user/sendCodeToUser",this.userInfo).then(res=>{
           if(!res){
              this.$message.warning("验证码发送失败")
            }
            else{
              this.$message.info("验证码已发送")
            }
        })
    },
    async doLogin(){
      console.log(this.userInfo)
    },
    async doRegister(){
      this.request.post("/user/userRegister",this.userInfo).then(response=>{
        if(!response){
          this.$message.warning("注册失败")
        }else{
          this.$message.success("注册成功")
        }
      })
    },
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
  padding-top: 50px;
  position: relative;
}

input{
  margin-bottom: 20px;
  border:0;
  border-bottom: 1px solid #999999;
  width:280px;
  font-size:large;
}
input:focus {
  outline: none;
}

/*验证码样式*/
.btn,
.btn:focus {
  position: relative;
  z-index: 0;
  width: 20%;
  height: 30px;
  border: 2px solid currentColor;
  border-radius: 20px;
  color: #000000;
  font-size: 1rem;
  text-align: center;
  text-decoration: none;
  overflow: hidden;
  transition: 0.2s transform ease-in-out;
  will-change: transform;
  bottom: 10px;
}
.btn:after {
  content: '';
  display: block;
  position: absolute;
  height: 100%;
  width: 100%;
  left: 0;
  top: 0;
  z-index: -1;
  background-color: #000000;
  border-radius: 3rem;
  transform: translate(-100%, 0) rotate(10deg);
  transform-origin: top left;
  transition: 0.2s transform ease-out;
  will-change: transform;
}
.btn:hover:after {
  transform: translate(0, 0);
}
.btn:hover {
  border: 2px solid transparent;
  color: #FFFFFF;
  transform: scale(1.05);
  will-change: transform;
}
p{
  font-size: smaller;
  position: absolute;
  margin-top: -20px;
}
</style>


