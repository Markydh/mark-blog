<template>
  <div class="cInnerContent">
    <h1 class="header-section gs_reveal ipsType_center">Scroll down and up to see different reveal animations</h1>

    <div class="features">

      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromLeft">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=1" alt="">
          </div>
        </div>

        <div class="ipsGrid_span7 ipsType_left">
          <h2 class="heading_large gs_reveal">Create amazing <strong>SVG</strong> animations</h2>
        </div>
      </div>

      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="ipsGrid_span7 ipsType_right">
          <h2 class="heading_large gs_reveal">Supercharge immersive <strong>WebGL</strong> experiences</h2>
        </div>

        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromRight">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=2" alt="">
          </div>
        </div>
      </div>

      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromLeft">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=3" alt="">
          </div>
        </div>

        <div class="ipsGrid_span7 ipsType_left">
          <h2 class="heading_large gs_reveal">Control performant <strong>Canvas</strong> animations</h2>
        </div>
      </div>

      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="ipsGrid_span7 ipsType_right">
          <h2 class="heading_large gs_reveal"><strong>Award winning</strong> websites</h2>
        </div>

        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromRight">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=4" alt="">
          </div>
        </div>
      </div>
      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromLeft">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=3" alt="">
          </div>
        </div>

        <div class="ipsGrid_span7 ipsType_left">
          <h2 class="heading_large gs_reveal">Control performant <strong>Canvas</strong> animations</h2>
        </div>
      </div>

      <div class="feature ipsSpacer_bottom_double ipsGrid ipsGrid_collapsePhone">
        <div class="ipsGrid_span7 ipsType_right">
          <h2 class="heading_large gs_reveal"><strong>Award winning</strong> websites</h2>
        </div>

        <div class="featured-image-container ipsGrid_span5 gs_reveal gs_reveal_fromRight">
          <div class="card">
            <img width="479" src="https://picsum.photos/479/479?index=4" alt="">
          </div>
        </div>
      </div>

    </div>

  </div>
</template>


<script>
import gsap from 'gsap';
import ScrollTrigger from 'gsap/ScrollTrigger';

export default {
  name: "PhotoWall",
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
  max-width: 100%;
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
