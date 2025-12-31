<!-- 工作台页面 -->
<template>
  <div>
    <ElRow :gutter="20">
      <ElCol :sm="24" :md="12" :lg="16">
        <Banner />
        <CardList />
        <HotProduct/>
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="8">
        <NavBox />
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
import CardList from './modules/card-list.vue'
import ActiveUser from './modules/active-user.vue'

import NewUser from './modules/new-user.vue'
import Dynamic from './modules/dynamic-stats.vue'
import TodoList from './modules/todo-list.vue'
import AboutProject from './modules/about-project.vue'
import SalesOverview from './modules/sales-overview.vue'
import Banner from './modules/banner.vue'
import Tranlist from './modules/tranlist.vue'
import HotProduct from './modules/hot-product.vue'
import NavBox from './modules/nav-box.vue'

defineOptions({ name: 'Console' })
</script>
