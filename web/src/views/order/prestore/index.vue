<!-- 预存管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="prestore-page art-full-height">
    <!-- 搜索栏 -->
    <PrestoreSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></PrestoreSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton  @click="handleCreate" v-ripple>新增预存</ElButton>
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

      <PrestoreModal
        v-model:visible="modalVisible"
        @submit="refreshData"
      />
      <PrestoreViewDrawer
        v-model="viewVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'

import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { ApplyStatus } from '@/enums/statusEnum'
import { fetchGetPrestoreList } from '@/api/prestore'
import PrestoreViewDrawer from './modules/prestore-view-drawer.vue'
import PrestoreModal from './modules/prestore-modal.vue'
import PrestoreSearch from './modules/prestore-search.vue'



defineOptions({ name: 'Prestore' })

const {getInfo:site} = useSiteStore()
// 弹窗相关
const modalVisible = ref(false)
const viewVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  code: undefined,
  name: undefined,
  status: undefined
})



const PRESTORE_STATUS = {
  [ApplyStatus.Pending]: { type: 'primary' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝' },
} as const

const getStatus = (status: number) => {
  return (
    PRESTORE_STATUS[status as keyof typeof PRESTORE_STATUS] || {
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
    apiFn: fetchGetPrestoreList,
    apiParams:{
      code: "",
      name: "",
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      {
        prop: 'user',
        label: '预存用户',
        formatter: (row) => {
          return h('p', { }, row.user)
        }
      },
      {
        prop: 'amount',
        label: '预存金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}` )
        }
      },
      {
        prop: 'bonusAmount',
        label: '赠送金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.bonusAmount}${site.symbol}` )
        }
      },
      {
        prop: 'status',
        label: '预存状态',
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
          return h('div', { class: 'prestore flex-c' }, [
            h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
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

const handleView = (row:Prestore.Response.Info) => {
  id.value = row.id
  nextTick(() => {
    viewVisible.value = true
  })
}

const handleCreate = () => {
  nextTick(() => {
    modalVisible.value = true
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Prestore.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
