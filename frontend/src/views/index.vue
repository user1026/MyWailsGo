<template>
  <div id="context">
    <el-container>
      <!-- <el-row :gutter="0">
        <el-col :span="leftSpan"> -->
      <el-aside :width="leftWidth">
        <div id="left">
       
            <el-icon @click="changeCollapse">
              <Menu />
            </el-icon>
        
          <el-menu
          ref="leftElMenu"
            default-active="2"
            class="leftMenu el-menu-vertical-demo"
            :collapse="isCollapse"
          >
            <el-menu-item index="1">
              <el-icon>
                <document />
              </el-icon>
              <template #title>Navigator Three</template>
            </el-menu-item>
            <el-menu-item index="2">
              <el-icon>
                <document />
              </el-icon>
              <template #title>Navigator Three</template>
            </el-menu-item>
          </el-menu>
        </div>
      </el-aside>
      <!-- </el-col>
        <el-col :span="rightSpan"> -->
      <el-container>
        <el-header></el-header>
        <el-main>
          <div id="right">
            <el-button @click="t">+</el-button>
          </div>
        </el-main>
      </el-container>
      <!-- </el-col>
    </el-row> -->
    </el-container>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from "vue";
const leftSpan = ref(4);
const rightSpan = ref(20);
const isCollapse = ref(false);

const leftElMenu=ref(null)
const leftWidth = ref("20vh");
const t = () => {
  leftSpan.value = 8;
  rightSpan.value = 16;
};
onMounted(()=>{
    leftWidth.value=leftElMenu.width+"px"
})
const changeCollapse = () => {
  isCollapse.value = !isCollapse.value;
  // isCollapse.value?leftSpan.value=1:leftSpan.value=4;
   leftWidth.value=leftElMenu.width+"px";
};
</script>
<style lang="scss" scoped>
#context {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

#left {
  
  height: 100vh;
  background-color: #c6cbd6;

  .leftMenu {
    background-color: #e9ecf3;
    border-right: none;
  }
}

#right {
  width: 100%;
  height: 100vh;
  background-color: blue;
}
</style>
