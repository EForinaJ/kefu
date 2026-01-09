<template>
  <ElRow :gutter="20" class="flex">
    <ElCol  :sm="12" :md="8" :lg="8">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">今日订单</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" :target="12" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">订单总额</span>
          <span class="ml-1 text-xs font-semibold text-success">
            520
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="ri:pie-chart-line" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
    <ElCol  :sm="12" :md="8" :lg="8">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">售后申请</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" :target="12" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">售后总量</span>
          <span class="ml-1 text-xs font-semibold text-success">
            520
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="ri:pie-chart-line" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
    <ElCol  :sm="12" :md="8" :lg="8">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">今日入职</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" :target="12" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">威客总量</span>
          <span class="ml-1 text-xs font-semibold text-success">
            520
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="ri:pie-chart-line" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
  </ElRow>
</template>

<script setup lang="ts">
  interface CardDataItem {
    des: string
    icon: string
    startVal: number
    duration: number
    num: number
    change: string
  }

  /**
   * 卡片统计数据列表
   * 展示总访问次数、在线访客数、点击量和新用户等核心数据指标
   */
  const dataList = reactive<CardDataItem[]>([
    {
      des: '待服务订单',
      icon: 'ri:pie-chart-line',
      startVal: 0,
      duration: 1000,
      num: 9120,
      change: '+20%'
    },
    {
      des: '待处理售后',
      icon: 'ri:group-line',
      startVal: 0,
      duration: 1000,
      num: 182,
      change: '+10%'
    },
    {
      des: '待审核报单',
      icon: 'ri:fire-line',
      startVal: 0,
      duration: 1000,
      num: 9520,
      change: '-12%'
    },
    {
      des: '待审核评论',
      icon: 'ri:fire-line',
      startVal: 0,
      duration: 1000,
      num: 9520,
      change: '-12%'
    },
  ])
</script>
