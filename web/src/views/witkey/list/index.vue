<!-- 威客管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="witkey-page art-full-height">
    <!-- 搜索栏 -->
    <WitkeySearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></WitkeySearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton  @click="handleCreate" v-ripple>添加威客</ElButton>
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
      <WitkeyDialog
        v-model:visible="dialogVisible"
        :id="id"
        @submit="refreshData"
      />
      <!-- 威客弹窗 -->
      <WitkeyViewDrawer
        v-model="viewDrawerVisible"
        :id="id"
        @submit="refreshData"
      />
      <!-- 弹窗 -->
      <WitkeyTitleDialog
        v-model:visible="titleDialogVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import WitkeySearch from './modules/witkey-search.vue'
import { ElTag, ElImage, ElRate } from 'element-plus'
import WitkeyDialog from './modules/witkey-dialog.vue'
import { fetchGetWitkeyList } from '@/api/witkey'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import WitkeyViewDrawer from './modules/witkey-view-drawer.vue'
import { Status } from '@/enums/statusEnum'
import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import WitkeyTitleDialog from './modules/witkey-title-dialog.vue'


defineOptions({ name: 'Witkey' })


// 弹窗相关
const titleDialogVisible = ref(false)
const dialogVisible = ref(false)
const viewDrawerVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
})

const STATUS = {
  [Status.Disable]: { type: 'danger' as const, text: '禁用' },
  [Status.Enable]: { type: 'success' as const, text: '启用' },
} as const

const getStatus = (status: number) => {
  return (
    STATUS[status as keyof typeof STATUS] || {
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
    apiFn: fetchGetWitkeyList,
    apiParams:{
      name: "",
      phone: "",
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      { 
        prop: 'name', 
        label: '威客名称',
      },
      {
        prop: 'userInfo',
        label: '所属用户',
        width: 280,
        formatter: (row) => {
          return h('div', { class: 'user flex-c' }, [
            h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.user.avatar,
              previewSrcList: [row.user.avatar],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2' }, [
              h('p', { class: 'user-name' }, row.user.name),
            ])
          ])
        }
      },
      { 
        prop: 'game', 
        label: '所属游戏',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.game}`)
        }
      },
      { 
        prop: 'title', 
        label: '所属头衔',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.title}`)
        }
      },
      { 
        prop: 'rate', 
        label: '评分',
        width: 260,
        formatter: (row) => {
          return h(ElRate , {
            modelValue: row.rate,
            showScore: true, 
            disabled: true,
            scoreTemplate:`${row.rate}分` 
          })
        }
      },
      {
        prop: 'status',
        label: '结单状态',
        formatter: (row) => {
          const statusConfig = getStatus(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '入驻日期',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'witkey flex-c' }, [
            h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            }),
            h(ArtButtonMore, {
              list: [
                {
                  key: 'changeTitle',
                  label: '变更头衔',
                  icon: 'ep:element-plus',
                },
              ],
              onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
            })
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

const buttonMoreClick = (item: ButtonMoreItem, row: Witkey.Response.Info) => {
  switch (item.key) {
    case 'changeTitle':
      handleChangeTitle(row)
      break
  }
}

/**
 * 搜索处理
 * @param params 参数
 */
const handleSearch = (params: Record<string, any>) => {
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

const handleChangeTitle = (row:Witkey.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    titleDialogVisible.value = true
  })
}
/**
 * 显示威客弹窗
 */
const handleCreate = (): void => {
  id.value = 0
  nextTick(() => {
    dialogVisible.value = true
  })
}

const handleView = (row:Witkey.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewDrawerVisible.value = true
    })
}

/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Witkey.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
