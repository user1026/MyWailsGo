<template>
  <div>
    概念介绍
  </div>
  <div class="descriptions">

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
          <el-tag size="large">i3,i5,i7,i9</el-tag>
        </template>
        i3:低端<br>
        i5:中端<br>
        i7:高端<br>
        i9:超高端（发烧）。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">AMD桌面端CPU后缀</el-tag>
        </template>
        X：指支持AMD官方超频XFR技术，也就是自动超频，频率的最大值受到散热器效果的影响，散热器越强，频率越高<br>
        G：代表有核显，也就是常说的APU，集成Vega显卡<br>
        3D：代表使用了3D缓存<br>
        WX：代表最新的超多核心线程撕裂者系列 <br>
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
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">R3,R5,R7,R9</el-tag>
        </template>
        R3:低端<br>
        R5:中端<br>
        R7:高端<br>
        R9:超高端（发烧）。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">盒装/散片</el-tag>
        </template>
        盒装CPU：也就是我们所见到的官方正品，包装盒有一颗CPU和CPU散热风扇，而某些盒装CPU还会在出厂时在散热风扇底座涂好导热硅脂，用户可以直接安装使用。 它拥有着质保等正常商品所具备的 一切售后服务 。 散片CPU
        ：一般来说即只有简易包装，没有任何正规包装盒，与正规盒装CPU最大区别在没有CPU原装散热器，无法享受官方正品带来的任何售后以及三包服务，而只有店保。 通常CPU散片主要是指 Intel处理器
        的散片CPU，这类处理器在售价上因为远低于盒装CPU的售价而被众多消费者所青睐。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">ES/QS</el-tag>
        </template>
        CPU的ES全称是Engineering Sample，即工程样品，也有叫ES不显的，在CPU-Z里面是不显示名字的。从开发本代产品开始的每一款处理器样品，会经过若干次修改和优化<br>
        QS全称是Quality Sample，也有叫ES正显的，在CPU-Z里面显示名字，但是末尾会带上“（ES）”的标志,QS是正式版前的最后版本。<br>
        <span style="color:red">小白绝对不要碰</span>
      </el-descriptions-item>
    </el-descriptions>
    <el-descriptions class="margin-top" title="主板" :column="1" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">Intel主板等级</el-tag>
        </template>
        X字母开头：最高级，用来搭配高端CPU，一般CPU型号后缀有“X”字母。比如X299主板，那么就可以搭配i9-7960X或者i7-7800X。<br>
        Z字母开头：次高端，Z字母开头的主板都支持超频，搭配的CPU一般带有“K”字母后缀，比如Z370主板，就能搭配i5-8600K或者i7-8700K。<br>
        B字母开头：中端主流，这种主板不支持超频，B开头的主板性价比最高，主要搭配不带K字母后缀的CPU，比如B360，就能使用i3-8100、i5-8500、i7-8700。<br>
        H字母开头：入门级，不支持超频，价格非常便宜，当然，H字母开头不代表都是低端产品，比如H310和H370，第二个数字越高，规格就越高，H370就相当于不能超频的Z370主板。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">AMD主板等级</el-tag>
        </template>
        X字母开头：最高级，支持自适应动态扩频超频，和Intel一样，也是搭配AMD中代“X”字母后缀的处理器。<br>
        B字母开头：中端主流，可以超频，不支持完整的自适应动态扩频超频，性价比较高。<br>
        A字母开头：入门级，不支持超频，普通办公用户使用，价格非常便宜。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">主板规格</el-tag>
        </template>
        E-ATX型：高性能主板，芯片组都是X字母开头，适合使用带X后缀的处理器，但是价格很高，不推荐普通用户使用。<br>
        ATX型：也就是大板，扩展性好，接口全，一般内存都是四插槽起，2或3个PCIe接口和M.2接口。<br>
        M-ATX型：小板，主流的主板版型，内存插槽一般是两个或者四个，会有一个M.2接口，扩展性虽然不高，但是绝对可以满足大多数用户的需求。<br>
        mini-ITX型：迷你主板，接口数量属于刚好够用，适合ITX迷你机箱，一般用来办公或者家用，不适合做游戏主机。
      </el-descriptions-item>
    </el-descriptions>
    <el-descriptions class="margin-top" title="显卡" :column="1" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">AIC/AIB厂商</el-tag>
        </template>
        Nvidia将其认可的有权生产采用Nvidia显示芯片显卡的厂商称为Add-In Card厂商，简称AIC。 与之相对的，AMD授权进行显卡生产的厂商称为Add-In Board厂商，简称AIB。
        AIC/AIB是一个准入性标准，即只有AIC/AIB才具备生产Nvidia/AMD显卡的权利。
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">Nvidia的AIC厂商</el-tag>
        </template>
        华硕、技嘉、微星、七彩虹、影驰、索泰、映众、耕升、铭瑄（可能会有变化，具体请关注官网）
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">AMD的AIB厂商</el-tag>
        </template>
        华硕 ，迪兰，技嘉，昂达，铭瑄，MSI微星，蓝宝石，盈通，讯景（可能会有变化，具体请关注官网）
      </el-descriptions-item>
    </el-descriptions>
    <el-descriptions class="margin-top" title="硬盘" :column="1" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">颗粒对比</el-tag>
        </template>
        价格：slc＞mlc＞tlc＞qlc<br>
        速度：slc＞mlc＞tlc＞qlc<br>
        寿命：slc＞mlc＞tlc＞qlc
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">m2固态硬盘品牌</el-tag>
        </template>
        系统盘首选：三星，西数，凯侠，致钛，英睿达<br>
        从盘：哪个便宜买哪个
      </el-descriptions-item>
    </el-descriptions>
    <el-descriptions class="margin-top" title="散热器" :column="1" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">风冷</el-tag>
        </template>
        优点：
        <ul>
          <li>1.经久耐用，至少扣具不换就能一直用下去，最多换个风扇</li>
          <li>2.价格相较于水冷便宜很多</li>
        </ul>
       缺点：
       <ul>
        <li>1.体积大（会挡住内存条）</li>
        <li>2.性能上限不如水冷高</li>
        <li>3.相较于水冷，安装比较困难</li>
       </ul>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">水冷</el-tag>
        </template>
          优点：
          <ul>
            <li>1.性能较好</li>
            <li>2.外观漂亮，冷头不占空间</li>
            <li>3.冷头安装拆解简便</li>
            <li>4.短时间使用噪音比风冷小</li>
          </ul>
          缺点：
          <ul>
            <li>1.相较于风冷，使用期限没有风冷长</li>
            <li>2.一旦漏液，会造成其他设备损坏(理赔比较困难)</li>
            <li>3.质保期到了，需要考虑是否更换</li>
            <li>4.长时间高负载使用噪音比风冷大</li>
          </ul>
      </el-descriptions-item>
    </el-descriptions>
    <el-descriptions class="margin-top" title="电源" :column="1" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">品牌</el-tag>
        </template>
      振华（自有工厂），海韵（自有工厂）、长城（自有工厂）、航嘉（自有工厂）、全汉（自有工厂）、台达（自有工厂）、海盗船（侨威、伟训、伟创力代工 ）、安钛克（高端产品海韵代工）、酷冷至尊（高端由益衡代工）、华硕（海韵，长城，益衡代工）、微星（乔威代工）<br>
      <span style="color: red;">首选老牌自有工厂，老牌厂商代工也可</span>
      </el-descriptions-item>
    </el-descriptions>
    <div class="bottom">

    </div>
  </div>
</template>
<script setup>

</script>
<style scoped>
  .descriptions {}
li{
  list-style: none;
}
  .bottom {
    height: 90px;
  }
</style>