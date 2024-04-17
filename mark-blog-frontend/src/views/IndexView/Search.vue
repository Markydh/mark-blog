<template>
   <div class="container">
     <div class="searchPart">
       <input class="searchBox" placeholder="Type to search..." >
       <el-icon class="searchIcon" style="vertical-align: bottom;font-size:25px;font-weight:bolder;position: relative;top: 4%;left: 20%" v-show="searchIconVisible">
         <Search />
       </el-icon>
     </div>
   </div>
</template>

<script setup>
import gsap from "gsap";
import {onMounted, ref} from "vue";

const searchIconVisible = ref(false)

onMounted(()=>{
  let tl = gsap.timeline()
  tl.from('.searchBox',{x:1000,duration:1,ease:'elastic'})
  tl.to(searchIconVisible, { value: true}, '>-0.8')
  tl.from('.searchIcon',{y:-20,duration:0.2,ease:'elastic'},">")
})


</script>

<style scoped>
.container{
  width: 100%;
  height: 100vh;
}

.searchPart{
  width: 100%;
}
.searchBox{
  position: relative;
  left: 22%;
  top: 5%;
  width: 55%;
  height: 50px;
  border: 0;
  font-size: 30px; /* 设置字体大小 */
  border-bottom: 4px solid black;
}
.searchBox:focus {
  outline: none; /* 隐藏默认的聚焦样式 */
}

.searchBox::placeholder {
  color: black; /* 修改占位符文本颜色 */
  font-size: 30px;
  font-style: italic; /* 修改占位符文本样式 */
}
</style>