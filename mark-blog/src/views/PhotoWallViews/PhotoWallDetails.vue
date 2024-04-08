<template>
  <div class="container">
    <div class="pictureArray" v-for="(pic,index) in items" :key="index" ref="pictureRefs">
      <img class="picture" :src="pic.url" alt="">
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from "vue";
import { gsap } from "gsap";

const items = reactive([
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6573.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6574.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6575.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6576.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6577.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6578.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6579.jpg?raw=true" },
  { url: "https://github.com/Markydh/MyPicture/blob/img/img/IMG_6580.jpg?raw=true" },
]);

const pictureRefs = ref([]);

onMounted(() => {
  // 使用 $nextTick 确保所有图片元素都被挂载
  pictureRefs.value.forEach((picture, index) => {
    if (index % 2 === 0) {
      gsap.from(picture, { x: -1000, duration: 2 });
    } else {
      gsap.from(picture, { x: 2680, duration: 2 });
    }
    if(index >= 2){
      gsap.from(picture, { y: 2000, duration: 2 });
    }
  });
});

</script>

<style scoped>
.container {
  position: relative;
  display: grid;
  grid-template-columns: 50% 50%;
  column-gap:20px;
  row-gap: 20px;
  left: 20%;
  width: 60%;
  height: 100vh;
}

.pictureArray .picture {
  width: 100%;
  height: 100%;
}
</style>
