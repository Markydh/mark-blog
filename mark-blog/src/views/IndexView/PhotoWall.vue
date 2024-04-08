<template>
  <div class="cInnerContent">
    <h1 class="header-section gs_reveal ipsType_center">记录美好生活，mark的照片墙</h1>

    <div class="features">
      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone" v-for="(item , index) in items" :key="index" @click="handleClick()">
        <div class="ipsGrid_span7 ipsType_right" v-if="index % 2 === 1">
          <h2 class="heading_large gs_reveal">{{item.title}}</h2>
        </div>

        <div class="featured-image-container ipsGrid_span5 gs_reveal" :class="{'gs_reveal_fromLeft': index%2 === 0, 'gs_reveal_fromRight':index%2 !== 0}">
          <div class="card">
            <img width="479" :src="item.picUrl" alt="">
          </div>
        </div>

        <div class="ipsGrid_span7 ipsType_left" v-if="index % 2 === 0">
          <h2 class="heading_large gs_reveal">{{item.title}}</h2>
        </div>
      </div>
    </div>

  </div>
</template>


<script>
import gsap from 'gsap';
import ScrollTrigger from 'gsap/ScrollTrigger';
import router from "@/router/index.js";

export default {
  name: "PhotoWall",
  data(){
    return{
      items:[
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6456.jpg",
          title:"眼泪女王💧👩👑",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6588.jpg",
          title:"金智媛❤️❤️❤️",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_5448.jpg",
          title:"最喜欢的男演员 - 方大同",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_6725.jpg",
          title:"最喜欢的女演员 - 邓紫棋",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_5907.jpg",
          title:"最喜欢的歌 - Too Smart",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_2000.jpg",
          title:"杭游-最爱城市🌆",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_5476.JPG",
          title:"上海行",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_1794.JPG",
          title:"武功山一游",
        },
        {
          picUrl:"https://cdn.jsdelivr.net/gh/Markydh/MyPicture@img/img/IMG_1000.JPG",
          title:"未完待续",
        }
      ]
    }
  },
  methods: {
    animateFrom(elem, direction) {
      direction = direction || 1;
      var x = 0,
          y = direction * 100;
      if (elem.classList.contains("gs_reveal_fromLeft")) {
        x = -100;
        y = 0;
      } else if (elem.classList.contains("gs_reveal_fromRight")) {
        x = 100;
        y = 0;
      }
      elem.style.transform = "translate(" + x + "px, " + y + "px)";
      elem.style.opacity = "0";
      gsap.fromTo(elem, {x: x, y: y, autoAlpha: 0}, {
        duration: 1.25,
        x: 0,
        y: 0,
        autoAlpha: 1,
        ease: "expo",
        overwrite: "auto"
      });
    },

    hide(elem) {
      gsap.set(elem, {autoAlpha: 0});
    },

    handleClick(){
      router.push('PhotoWallDetails')
    }
  },

  mounted() {
    gsap.registerPlugin(ScrollTrigger, gsap.TweenMax);

    gsap.utils.toArray(".gs_reveal").forEach((elem) => {
      this.hide(elem); // assure that the element is hidden when scrolled into view

      ScrollTrigger.create({
        trigger: elem,
        // 显示start和end
        markers: false,
        onEnter: () => { this.animateFrom(elem) },
        onEnterBack: () => { this.animateFrom(elem, -1) },
        onLeave: () => { this.hide(elem) } // assure that the element is hidden when scrolled into view
      });
    });
  }
}
</script>



<style scoped>
/* Styles from the GreenSock website */
body {
  font-weight: 300;
}
.ipsType_right {
  text-align: right;
}
.ipsType_center {
  text-align: center;
  font-size: 35px;
}
.cInnerContent {
  max-width: 1240px;
  margin-left: auto;
  margin-right: auto;
}
.ipsSpacer_bottom_double {
  margin-bottom: 30px;
}
.ipsGrid {
  display: inline-block;
  display: -ms-flexbox;
  display: flex;
  flex-wrap: wrap;
}
.ipsGrid::before, .ipsGrid::after {
  display: table;
  content: "";
  line-height: 0;
}
.ipsGrid > [class*="ipsGrid_span"] {
  display: block;
  width: 100%;
  min-height: 30px;
  box-sizing: border-box;
}
.ipsGrid > .ipsGrid_span5 {
  width: 40.42553191489362%;
}
.ipsGrid > .ipsGrid_span7 {
  width: 57.44680851063829%;
}
.ipsGrid > [class*="ipsGrid_span"] {
  float: left;
  margin-left: 2%;
}
.ipsGrid > [class*="ipsGrid_span"]:first-child {
  margin-left: 0;
}
.feature {
  display: flex;
  align-items: center;
}
.card {
  margin-bottom: 30px;
  border: 1px solid #cccccc;
  border-radius: 8px;
  overflow: hidden;
  background: #ffffff;
  box-shadow: 1px 1px 5px 1px #CCCCCC;
  /*   transition: 0.3s; */
}
.featured-image-container .card {
  padding: 10px;
  height: 0;
  padding-bottom: calc(100% - 10px);
}
h2.heading_large {
  font-size: 1.8em;
}
img {
  width: 479px;
  height: 479px;
}

.header-section {
  margin: 200px auto;
}

.gs_reveal {
  opacity: 0;
  visibility: hidden;
  will-change: transform, opacity;
}
</style>
