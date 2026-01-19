<template>
  <div class="art-card h-128 p-5 mb-5 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4>代办工单</h4>
        <p>待处理<span class="text-danger">{{list.length}}</span></p>
      </div>
    </div>
    <div v-if="list.length == 0" class="flex items-center justify-center">
      <el-empty description="无待办" />
    </div>
    <div v-if="list.length > 0" class="h-[calc(100%-40px)] overflow-auto">
      <ElScrollbar>
        <div
          class="flex-cb h-17.5 border-b border-g-300 text-sm last:border-b-0"
          v-for="(item, index) in list"
          :key="index"
        >
          <div>
            <p class="text-sm">{{ item.title }}</p>
            <p class="text-g-500 mt-1">{{ item.createTime }}</p>
          </div>
          <ElTag :type="getPriority(item.priority).type">{{ getPriority(item.priority).text }}</ElTag>
        </div>
      </ElScrollbar>
    </div>
  </div>
</template>

<script setup lang="ts">
import { PriorityType } from '@/enums/typeEnum';

interface Props {
  list: WorkOrder.Response.Info[]
}
const props = withDefaults(defineProps<Props>(), {
  list:() => []
})

const PRIORITY_TYPE = {
  [PriorityType.urgent]: { type: 'danger' as const, text: '紧急' },
  [PriorityType.secondary]: { type: 'danger' as const, text: '高' },
  [PriorityType.tertiary]: { type: 'warning' as const, text: '中' },
  [PriorityType.quaternary]: { type: 'warning' as const, text: '低' },
} as const
const getPriority = (priority: number) => {
  return (
    PRIORITY_TYPE[priority as keyof typeof PRIORITY_TYPE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

</script>
