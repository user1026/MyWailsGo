<template>
  <el-descriptions class="margin-top" title="详细信息" :column="1" :size="size" border>
    <template v-for="(v,i) in List">
      <el-descriptions-item :key="i" :label="v.label" v-if="show(v.val)" :span="2">{{v.val}}
      </el-descriptions-item>
    </template>
  </el-descriptions>
</template>
<script setup>
  import {
    ref,
    reactive,
    onMounted,
  } from "vue"
  const prop = defineProps(["List", "Info"])
  const size = ref("")
  const List = ref({})
  if (prop.Info != null) {
    List.value.forEach(v => {
      v.val = prop.Info[v.key] ? prop.Info[v.key] : "";
    })
  }
  onMounted(() => {

  })
  const show = (str) => {
    return str != "" && str != undefined
  }
</script>
<style scoped>

</style>