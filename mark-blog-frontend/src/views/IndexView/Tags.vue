<template>
  <section>
    <div class="bg"></div>
    <h1>Simple parallax sections</h1>
  </section>
  <section>
    <div class="bg"></div>
    <h1>Hey look, a title</h1>
  </section>
  <section>
    <div class="bg"></div>
    <h1>They just keep coming</h1>
  </section>
  <section>
    <div class="bg"></div>
    <h1>So smooth though</h1>
  </section>
  <section>
    <div class="bg"></div>
    <h1>Nice, right?</h1>
  </section>

</template>

<script>

import gsap from 'gsap';
import ScrollTrigger from 'gsap/ScrollTrigger';

export default {
  name: "Tags",
  mounted() {
    gsap.registerPlugin(ScrollTrigger);

    let getRatio = el => window.innerHeight / (window.innerHeight + el.offsetHeight);

    gsap.utils.toArray("section").forEach((section, i) => {
      section.bg = section.querySelector(".bg");

      // Give the backgrounds some random images
      section.bg.style.backgroundImage = `url(https://picsum.photos/1600/800?random=${i})`;

      // the first image (i === 0) should be handled differently because it should start at the very top.
      // use function-based values in order to keep things responsive
      gsap.fromTo(section.bg, {
        backgroundPosition: () => i ? `50% ${-window.innerHeight * getRatio(section)}px` : "50% 0px"
      }, {
        backgroundPosition: () => `50% ${window.innerHeight * (1 - getRatio(section))}px`,
        ease: "none",
        scrollTrigger: {
          trigger: section,
          start: () => i ? "top bottom" : "top top",
          end: "bottom top",
          scrub: true,
          invalidateOnRefresh: true // to make it responsive
        }
      });

    });
  }
}
</script>

<style scoped>
section {
  position: relative;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
}

h1 {
  color: white;
  text-shadow: 1px 1px 3px black;
  z-index: 1;
  font-size: 3em;
  font-weight: 400;
}

</style>