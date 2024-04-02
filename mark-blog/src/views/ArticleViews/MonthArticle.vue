<template>
  <div class="container">
<!--    左侧修饰-->
    <div style="width: 10%">
        <div style="width: 5px;height: 90px;background-color: #da5827"></div>
    </div>
    <div class="content">
      <div style="width: 100%;height: 200px">
        <p style="font-size: 35px;margin-bottom: 5px;">月刊</p>
        <p style="font-size: 20px;color: #818181;margin-top: 10px;">个人月刊，分享近期的所见所思所感，每月更新一期。若喜欢可以点击左侧 Subscribe 输入您的邮箱地址订阅（移动端点击上方「Subscribe」按钮订阅）</p>
      </div>
     <div v-for="(item , index) in articles" :key="index" class="article">
       <hr style="border: 1px solid rgba(204,204,204,0.31); width: 100%">
<!--       如果此文章有展示图片-->
          <div v-if="item.url !== ''" class="article-with-image" :style="{backgroundImage: 'url('+ item.url +')' }">
<!--            加一层阴影-->
            <div class="image-overlay">
              <div style="position: relative;top:53%;left: 4%">
              <p style="color: #ffffff;font-size: 15px;margin-bottom:10px">{{item.date}}</p>
              <p style="color:#ffffff;font-size: 30px;font-weight: 500;margin-top:25px;padding-bottom: 20px">{{item.title}}</p>
              </div>
            </div>
          </div>
<!--       如果此文章没有展示图片-->
          <div v-else>
              <p style="color: #b0b0b0;font-size: 14px;margin-bottom:10px">{{item.date}}</p>
              <p style="font-size: 25px;font-weight: 500;margin-top:10px;padding-bottom: 20px">{{item.title}}</p>
          </div>
     </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "MonthArticle",
  data(){
    return{
      imageUrl:'/images/LoveMan.jpg',
      articles:[
        {
          date:'April 02, 2024',
          title:'月刊（第6期）：月色血风暴',
          url:'/images/storm.jpg'
        },
        {
          date:'April 01, 2024',
          title:'月刊（第5期）：梦的开始',
          url:''
        },
        {
          date:'March 28 , 2024',
          title:'月刊（第4期）：爱具体的人',
          url:'/images/LoveMan.jpg',
        },
        {
          date:'March 14, 2024',
          title:'月刊（第3期）：当下的快乐',
          url:''
        },
        {
          date:'March 12, 2024',
          title:' 月刊（第2期）：眼泪女王',
          url:'/images/tearQueen.jpg'
        },
        {
          date:'February 04, 2024',
          title:'月刊（第1期）：编程之路',
          url:''
        }
      ]
    }
  }
}
</script>

<style scoped>
.container{
  display: flex;
  justify-content: flex-start;
}
.content{
  width: 100%;
  display: flex;
  flex-direction: column;
}
p{
  font-family: "Hiragino Sans GB";
}
.article{
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
}
.article-with-image {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 400px;
  top:5px;
  background-size: cover;
  background-position: top;
  position: relative;
}
.image-overlay{
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.4); /* 阴影颜色和透明度 */
  padding: 20px;
  transition: background-color 0.3s ease; /* 过渡效果 */
}
.article-with-image:hover .image-overlay {
  background-color: rgba(0, 0, 0, 0.6); /* 悬停时阴影加深 */
}
</style>