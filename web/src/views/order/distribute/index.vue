<!-- 售后退款管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="distribute-page art-full-height">
      <!-- 搜索栏 -->
      <DistributeSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></DistributeSearch>
  
      <ElCard class="art-table-card" shadow="never">
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
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

        <DistributeViewDrawer
          v-model="viewVisible"
          :id="id"
          @submit="refreshData"
        />
        <DistributeCancelModal
          v-model:visible="cancelVisible"
          :id="id"
          @submit="getData"
      />
      </ElCard>
    </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag } from 'element-plus'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import DistributeSearch from './modules/distribute-search.vue'
import {  DistributeStatus } from '@/enums/statusEnum'
import { DistributeType } from '@/enums/typeEnum'
import { fetchGetDistributeList } from '@/api/distribute'
import DistributeCancelModal from './modules/distribute-cancel-modal.vue'
import DistributeViewDrawer from './modules/distribute-view-drawer.vue'
  
  

defineOptions({ name: 'Distribute' })

// 弹窗相关
const viewVisible = ref(false)
const cancelVisible = ref(false)
const id = ref<number>(0)
  
// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
  status: undefined
})

  
  

const STATUS = {
  [DistributeStatus.Pending]: { type: 'primary' as const, text: '待服务' },
  [DistributeStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
  [DistributeStatus.Completed]: { type: 'success' as const, text: '已完成' },
  [DistributeStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
} as const

const getStatus = (isCancel: number) => {
  return (
    STATUS[isCancel as keyof typeof STATUS] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const TYPE = {
  [DistributeType.Self]: { type: 'primary' as const, text: '个人服务' },
  [DistributeType.Team]: { type: 'primary' as const, text: '自带队伍' },
} as const

const getType = (type: number) => {
  return (
    TYPE[type as keyof typeof TYPE] || {
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
    apiFn: fetchGetDistributeList,
    apiParams:{
      code: "",
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      {
          prop: 'order',
          label: '订单号',
          width: 240,
          formatter: (row) => {
              return h('p', { }, row.order)
          }
      },
      {
          prop: 'manage',
          label: '派单客服',
          width: 160,
          formatter: (row) => {
              return h('p', { }, row.manage)
          }
      },
      {
          prop: 'witkey',
          label: '接单者',
          width: 160,
          formatter: (row) => {
              return h('p', { }, row.witkey)
          }
      },
      {
          prop: 'game',
          label: '游戏领域',
          formatter: (row) => {
              return h(ElTag, { type:"primary" }, () => row.game )
          }
      },
      {
          prop: 'title',
          label: '所属头衔',
          formatter: (row) => {
              return h(ElTag, { type:"primary" }, () => row.title )
          }
      },
      {
          prop: 'type',
          label: '派单类型',
          formatter: (row) => {
          const typeConfig = getType(row.type)
          return h(ElTag, { type: typeConfig.type }, () => typeConfig.text)
          }
      },
      {
          prop: 'status',
          label: '服务状态',
          formatter: (row) => {
          const statusConfig = getStatus(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
          }
      },
      {
          prop: 'createTime',
          label: '派单时间',
          width: 200,
          sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'distribute flex-c' }, [
            h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            }),
            h(ArtButtonTable, {
              type: 'delete',
              icon: 'solar:close-square-bold',
              onClick: () => handleCancel(row)
            }),
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

const handleView = (row:Distribute.Response.Info) => {
  id.value = row.id
  nextTick(() => {
    viewVisible.value = true
  })
}


const handleCancel = (row:Distribute.Response.Info) => {
  id.value = row.id
  nextTick(() => {
    cancelVisible.value = true
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Distribute.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
  