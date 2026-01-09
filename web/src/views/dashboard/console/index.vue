<!-- 工作台页面 -->
<template>
  <div>
    <ElRow :gutter="20">
      <ElCol :sm="24" :md="12" :lg="16">
        <Banner 
          :today-sales="detail.todaySales"
          :sales-comparison="detail.salesComparison"
        />
        <HotProduct :table-data="detail.hotProductList"/>
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="8">
        <NavBox />
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
import { fetchGetDashboardDetail } from '@/api/dashboard';
import Banner from './modules/banner.vue'
import HotProduct from './modules/hot-product.vue'
import NavBox from './modules/nav-box.vue'

defineOptions({ name: 'Console' })

const detail = ref<Dashboard.Response.Detail>({
  todaySales:0,
  salesComparison:0,
  hotProductList:[]
})
const loading = ref<boolean>(false)
const getData = async () => {
    loading.value = true
    const res = await fetchGetDashboardDetail()
    detail.value = res
    loading.value = false
}
onMounted(()=>{
    getData()
})
</script>
