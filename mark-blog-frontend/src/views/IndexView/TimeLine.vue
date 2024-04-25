<template>
<div class="box">
  <div class="container">

    <div class="top">
      <div style="display: flex;justify-content: flex-start;align-items: center;">
        <a href="Index">首页 </a>
        <p style="margin-left: 5px;color: #989898;font-family: Hiragino Sans GB"> > 时间轴 </p>
      </div>
      <h1 style="font-family: 'Hiragino Sans GB';font-size: 33px;margin-top: -19px">时间轴</h1>
      <p style="font-size:16px;font-family:'Hiragino Sans GB';margin-top: -19px;">我的故事 from 2002 - 2024</p>
    </div>

    <div v-for="(item,index) in timeLine" :key="index">
      <h><span style="font-family: 'Hiragino Sans GB';font-size: 15px">{{item.year}}</span></h>
      <ul>
        <li v-for="(task, taskIndex) in item.tasks" :key="taskIndex">
          <i class="iconfont" style="color: black;font-size: 20px;font-weight: normal">&#xe682;</i>
          <span style="padding-left: 5px; position: relative; bottom: 1px;font-size: 16px;font-family:'Hiragino Sans GB'"><strong>{{ task.description }}</strong></span>
          <p style="font-family: 'Hiragino Sans GB'">{{ task.case }}</p>
        </li>
      </ul>
      <hr style="border: 1px solid rgba(178,178,178,0.31); margin-top: 30px">
    </div>



  </div>
</div>
</template>

<script>
export default {
  name: "TimeLine",
  data(){
    return{
      timeLine:[
        {
          year:'2024',
          tasks:[
            {
              description:'着手个人Blog',
              case:'部署阿里云服务器 结合Vue Gin NGINX等，开发个人博客网站'
            }
          ]
        },
        {
          year:'2023',
          tasks:[
            {
              description:'独立开发网页，未部署服务器',
              case:'初习Vue2 + Springboot，开发了智享餐饮的前后端分离项目'
            }
          ]
        },
        {
          year:'2022',
          tasks:[
            {
              description:'初识web前端',
              case:'感受到了web前端的魅力'
            }
          ]
        },
        {
          year:'2021',
          tasks:[
            {
              description:'✌️胜利',
              case:'获得南昌大学录取通知书📖'
            },
            {
              description:'程序员🐒',
              case:'选择软件工程🧑‍💻'
            },
          ]
        },
        {
          year:'2020',
          tasks:[
            {
              description:'疫情起始',
              case:'新冠肆虐，居家学习，埋下伏笔'
            },
            {
              description:'人生至暗时刻plus',
              case:'🏳️高考失利'
            },
            {
              description:'独自一人前往另一个城市🌆',
              case:'这次复读，可以说是影响我一生，改正了以前的许多坏习惯'
            }
          ]
        },
        {
          year:'2019',
          tasks:[
            {
              description:'语文老师的赌约2',
              case:'语文老师貌似还不满意考到班上第一的我，又一次在语文课上，对着全班人说：杨典宏这个小朋友呀，不够持久，上次考第一时运气，下次肯定考不到，要是下次考到，' +
                  '我就请他吃饭。戏剧的是，我又一次蝉联班级第一☝️，而且在我印象中，好像是连续三次还是四次，考班级第一勒。'
            },
            {
              description:'神之领悟',
              case:'突然领悟到老师的良苦用心😁😁'
            }
          ]
        },
        {
          year:'2018',
          tasks:[
            {
              description:'18年的上半年，人生第二至暗时刻',
              case:'这一学期，有了手机的我本就饿狼扑食，加之一次又一次的活动，放了许多的假，如虎添翼，我玩疯了，man😭 期末考试下滑到班上21名'
            },
            {
              description:'18年的下半年，曙光就在前方',
              case:'带着21321的学号，开启了我大二的学习生活。痛定思痛，痛改前非，开始拿起书本，投入到学习的海洋。随后的考试考到了班上第7，小win！！！🥳🥳'
            },
            {
              description:'语文老师的赌约1',
              case:'班上第七的我，幸运的在下次考试斩获第四的名次，深知此数字些许不吉利，所以我下次以及以后都是考了第二名🥈(开玩笑😝),有次语文老师的心血来潮，' +
                  '在班上同学面前，狠狠批斗我。说：杨典宏，你这家伙还是不行，不拔尖，肯定考不到班上第一。如果你考上班上第一我就奖励你一瓶阿萨姆（南昌人就系爱恰阿萨姆😝）'+
                  '没想到的是，在下次考试中，我就斩获第一名☝️（老师到现在还没奖我呢）'
            }
          ]
        },
        {
          year:'2017',
          tasks:[
            {
              description:'顺利熬过中考，进入省重点高中',
              case:'17年的我还没有对知识充满强烈的渴望，导致我并没有通过零班考试'
            },
            {
              description:'高中唯一一次分班，学号11307',
              case:'一年时间，我的变化并不多，所以还是错过了进入零班的机会(也是最后一次)'
            },
            {
              description:'进入到高一十三班的班集体',
              case:'起初沉浸在考试失利的低气压下，但是遇到了几个有意思的兄弟👬，让我的枯燥乏味学习生活有了一丝趣味'
            }
          ]
        },
        {
          year:'2008',
          tasks:[
            {
              description:'🏫步入小学学堂',
              case:'正式踏入求知之旅'
            }
          ]
        },
        {
          year:'2002',
          tasks:[
            {
              description:'🏥',
              case:'👶🍼 江西南昌'
            },
            {
              description: '被赋予与众不同的姓名',
              case:'😎杨典宏出现了🌍'
            }
          ]
        }
      ]
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
.top{
  display: flex;
  flex-direction: column;
}
a{
  font-family: "Hiragino Sans GB";
  color: #989898;
  text-decoration: none;
}
a:hover{
  color: black; ;
}
ul{
  list-style-type: none;
}
</style>