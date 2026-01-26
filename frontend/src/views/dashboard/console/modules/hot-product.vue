<template>
    <div class="art-card p-5 h-[27.8rem] mb-5 overflow-hidden max-sm:mb-4">
      <div class="art-card-header">
        <div class="title">
          <h4>热销产品</h4>
          <p>近期销售情况</p>
        </div>
      </div>
  
      <ElScrollbar style="height: 21.55rem" class="w-full">
        <ArtTable
          :data="tableData"
          style="margin-top: 0 !important"
          :border="false"
          :stripe="false"
          :header-cell-style="{ background: 'transparent' }"
        >
          <template #default>
            <ElTableColumn label="商品编号" prop="code">
              <template #default="scope">
                <span class="font-semibold">{{ scope.row.code }}</span>
              </template>
            </ElTableColumn>
            <ElTableColumn label="产品" prop="product" width="320px">
              <template #default="scope">
                <div class="flex-c">
                  <img class="size-12.5 object-cover rounded-md" :src="scope.row.pic" />
                  <div class="flex flex-col ml-3">
                    <div class="font-medium">{{ scope.row.name }}</div>
                    <div class="text-xs text-slate-500">{{ scope.row.category }}</div>
                  </div>
                </div>
              </template>
            </ElTableColumn>
            <ElTableColumn label="游戏" prop="rate">
              <template #default="scope">
                <ElTag type="primary">{{ scope.row.game }}</ElTag>
              </template>
            </ElTableColumn>
            <ElTableColumn label="价格" prop="price">
              <template #default="scope">
                <span class="font-semibold">{{ scope.row.price }}{{ site.symbol }}</span>
              </template>
            </ElTableColumn>
            <ElTableColumn label="评分" prop="rate">
              <template #default="scope">
                <ElRate 
                  disabled
                  :model-value="scope.row.rate"
                />
              </template>
            </ElTableColumn>
            <!-- <ElTableColumn label="销量" prop="sales" />
            <ElTableColumn label="销售趋势" width="240">
              <template #default="scope">
                <ElProgress :percentage="scope.row.pro" :color="scope.row.color" :stroke-width="4" />
              </template>
            </ElTableColumn> -->
          </template>
        </ArtTable>
      </ElScrollbar>
    </div>
  </template>
  
<script setup lang="ts">
import { useSiteStore } from '@/store/modules/site';

interface Props {
  tableData: Product.Response.Info[]
}
const props = withDefaults(defineProps<Props>(), {
})
const { getInfo:site } = useSiteStore()

</script>