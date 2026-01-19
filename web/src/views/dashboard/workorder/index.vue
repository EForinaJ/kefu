<!-- 待办事项管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="workorder-page art-full-height">
      <!-- 搜索栏 -->
      <WorkorderSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></WorkorderSearch>
  
      <ElCard class="art-table-card" shadow="never">
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
          <template #left>
            <ElSpace wrap>
              <ElButton  @click="handleCreate" v-ripple>上报工单</ElButton>
            </ElSpace>
          </template>
        </ArtTableHeader>
  
        <!-- 表格 -->
        <ArtTable
          :loading="loading"
          :data="data"
          :columns="columns"
          :pagination="pagination"
          @selection-change="handleSelectionChange"
          @pagination:size-change="handleSizeChange"
          @pagination:current-change="handleCurrentChange"
        >
        </ArtTable>

        <!-- 弹窗 -->
      <WorkorderDrawer
        v-model="dialogVisible"
        @submit="refreshData"
      />
      </ElCard>
    </div>
  </template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox, ElSpace, ElPopconfirm } from 'element-plus'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'

import { WorkOrderStatus } from '@/enums/statusEnum'
import WorkorderSearch from './modules/workorder-search.vue'
import { fetchGetWorkOrderList, fetchPostWorkOrderCancel } from '@/api/workorder'
import WorkorderDrawer from './modules/workorder-drawer.vue'
import { PriorityType } from '@/enums/typeEnum'

  

defineOptions({ name: 'Workorder' })

// 弹窗相关
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  code: undefined,
  status: undefined
})

const dialogVisible = ref(false)


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

const WORKORDER_STATUS = {
  [WorkOrderStatus.Allocate]: { type: 'primary' as const, text: '待分配' },
  [WorkOrderStatus.Processing]: { type: 'primary' as const, text: '处理中' },
  [WorkOrderStatus.Resolved]: { type: 'success' as const, text: '已解决' },
  [WorkOrderStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
} as const
const getStatus = (status: number) => {
  return (
    WORKORDER_STATUS[status as keyof typeof WORKORDER_STATUS] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}
const {
  columns,
  columnChecks,
  data,
  loading,
  pagination,
  getData,
  searchParams,
  resetSearchParams,
  handleSizeChange,
  handleCurrentChange,
  refreshData
} = useTable({
  // 核心配置
  core: {
    apiFn: fetchGetWorkOrderList,
    apiParams:{
      code: "",
      status: 0,
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      { type: 'selection' }, // 勾选列
      {
        type: 'expand',
        prop: 'content',
        formatter: (row) => {
          return h('p', { }, `${row.content}`)
        }
      }, 
      { prop: 'code', width: 280, label: '工单号' }, // 序号
      {
        prop: 'title',
        label: '工单标题',
        formatter: (row) => {
          return h('p', { }, row.title)
        }
      },
      {
        prop: 'priority',
        label: '优先级',
        formatter: (row) => {
          const priorityConfig = getPriority(row.priority)
          return h(ElTag, { type: priorityConfig.type }, () => priorityConfig.text)
        }
      },
      {
        prop: 'status',
        label: '审核状态',
        formatter: (row) => {
          const statusConfig = getStatus(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '申请时间',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'workorder flex-c' }, [
            (row.status == WorkOrderStatus.Processing && h(ArtButtonTable, {
              type: 'edit',
              icon: 'solar:check-circle-bold',
              onClick: () => handleSolve(row)
            })),
            (row.status == WorkOrderStatus.Processing && h(ArtButtonTable, {
              type: 'delete',
              icon: 'solar:close-circle-bold',
              onClick: () => handleCancel(row)
            })),
          ])
        }
      }
    ],
  },
  // 数据处理
  transform: {
    responseAdapter: (response) => {
      return {
        records: response.list,
        total: response.total,
      };
    },
  },
})


/**
 * 搜索处理
 * @param params 参数
 */
const handleSearch = (params: Record<string, any>) => {
  console.log(params)
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

const handleCreate = (): void => {
  id.value = 0
  nextTick(() => {
    dialogVisible.value = true
  })
}

const handleSolve = async (row: WorkOrder.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`是否完成工单内容`, '确认工单', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
  }).then(async() => {
    ElMessage.success('完成工单内容')
    await fetchPostWorkOrderCancel({id:row.id})
    refreshData()
  })
  .catch(() => {
    // ElMessage.info('已取消处理')
  })
}


const handleCancel = async (row: WorkOrder.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`是否取消处理工单内容`, '取消处理', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    ElMessage.success('确认取消处理')
    await fetchPostWorkOrderCancel({id:row.id})
    refreshData()
  })
  .catch(() => {
    // ElMessage.info('已取消处理')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: WorkOrder.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
  