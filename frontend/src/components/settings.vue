<script setup>
import {reactive,ref} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
const dialogVisible=ref(false)
function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
<el-dialog
v-model="dialogVisible"
    title="设置"
    width="60%"
>
  <el-form :model="form" label-width="120px">
      <el-form-item label="默认导出位置">
          
      </el-form-item>
  </el-form>
</el-dialog>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
