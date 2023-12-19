<template>
  <div>
    概念介绍
  </div>
  <div>

  </div>
  <el-descriptions class="margin-top" title="CPU（只列常见的）" :column="1" size="large" border>
    <el-descriptions-item>
      <template #label>
        <el-tag size="large">Intel桌面端CPU后缀</el-tag>
      </template>
      K: 高性能，未锁频<br>
      F: 没有内置核显,需要独立显卡<br>
      S: 节能版处理器，相比于同型号不带 S 的处理器性能差一些<br>
      T :相比于同型号后面带 S 的更加节能，但是功耗更低，适合长时间开机但是对性能要求不那么高的用户<br>
      X/XE: 最高性能，未锁频<br>
      KS：可以理解为官方超频版，提升了主频的版本
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <el-tag size="large">Intel笔记本端CPU后缀</el-tag>
      </template>
      HX: 最高性能，所有 SKU 未锁频<br>
      HK: 高性能，未锁频<br>
      H: 高电压处理器，功耗高、主频高的处理器<br>
      P: 提供轻薄型设备所需的性能<br>
      U :笔记本低功耗处理器，性能偏弱，续航能力强，多用于轻薄本<br>
      Y: 功耗极低<br>
      G1-G7: 显卡级别（采用较新集成显卡技术的处理器）<br>
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <el-tag size="large">AMD桌面端CPU后缀</el-tag>
      </template>
      X：指支持AMD官方超频XFR技术，也就是自动超频，频率的最大值受到散热器效果的影响，散热器越强，频率越高<br>
      G：代表有核显，也就是常说的APU，集成Vega显卡<br>
      3D：代表使用了3D缓存<br>
      WX：代表最新的超多核心线程撕裂者系列   <br>
      F:7000系全系带核显，f代表没有核显<br>
      无任何后缀：不带核显和不能超频，标准电压   
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <el-tag size="large">AMD笔记本端CPU后缀</el-tag>
      </template>
      H:高性能版本带核心显卡<br>
      HX：一般使用在AMD高端发烧级CPU上，至尊版<br>
      U:低功耗版本带核心显卡<br>
      HS：相当于H功耗略低，通常出现在轻薄全能本，性能较强<br>
      无任何后缀：不带核显和不能超频，标准电压  
    </el-descriptions-item>
  </el-descriptions>
  <el-descriptions class="margin-top" title="显卡" :column="2" size="large" border>
    <el-descriptions-item>
      <template #label>
        <el-tag size="large">AIC厂商</el-tag>
      </template>

      Nvidia将其认可的有权生产采用Nvidia显示芯片显卡的厂商称为Add-In Card厂商，简称AIC。 与之相对的，AMD授权进行显卡生产的厂商称为Add-In Board厂商，简称AIB。
      AIC是一个准入性标准，即只有AIC才具备生产Nvidia显卡的权利。
    </el-descriptions-item>
  </el-descriptions>
</template>
<script setup>

</script>
<style scoped>

</style>