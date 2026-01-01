<!-- 售后退款管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="aftersales-page art-full-height">
      <!-- 搜索栏 -->
      <AftersalesSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></AftersalesSearch>
  
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
        <AftersalesViewDrawer
          v-model="viewVisible"
          :id="id"
          @submit="refreshData"
        />
      </ElCard>
    </div>
</template>

<script setup lang="ts">
  import { useTable } from '@/hooks/core/useTable'
  import { ElTag, ElImage } from 'element-plus'
  import { useSiteStore } from '@/store/modules/site'
  import { useAuth } from '@/hooks'
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { fetchGetAftersalesList } from '@/api/aftersales'
  import AftersalesSearch from './modules/aftersales-search.vue'
  import { ApplyStatus } from '@/enums/statusEnum'
  import AftersalesViewDrawer from './modules/aftersales-view-drawer.vue'
  
  

  defineOptions({ name: 'Aftersales' })
  
  const {getInfo:site} = useSiteStore()
  // 弹窗相关
  const viewVisible = ref(false)
  const id = ref<number>(0)
  
  // 选中行
  const selectedRows = ref<number[]>([])
  
  // 搜索表单
  const searchForm = ref({
    name: undefined,
    phone: undefined,
    status: undefined
  })
  
  
  
  const AFTERSALES_STATUS = {
    [ApplyStatus.Pending]: { type: 'primary' as const, text: '待审核' },
    [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
    [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝' },
  } as const
  
  const getStatus = (status: number) => {
    return (
      AFTERSALES_STATUS[status as keyof typeof AFTERSALES_STATUS] || {
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
      apiFn: fetchGetAftersalesList,
      apiParams:{
        code: "",
        orderCode: "",
      },
      paginationKey:{
        current: 'page', 
        size: 'limit'
      },
      columnsFactory: () => [
        {
          prop: 'afterInfo',
          label: '售后编号',
          width: 280,
          formatter: (row) => {
            return h('div', {  }, [
                h('p', { }, `售后号:${row.code}`),
                h('p', { }, `订单号:${row.orderCode}`),
              ])
          }
        },
        {
          prop: 'productInfo',
          label: '商品信息',
          width: 220,
          formatter: (row) => {
            return h('div', { class: 'flex-c' }, [
              h(ElImage, {
                class: 'size-12 rounded-md',
                src: row.product.pic,
                previewSrcList: [row.product.pic],
                // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
                previewTeleported: true
              }),
              h('div', { class: 'ml-2 flex-1' }, [
                h('p', { class: 'line-clamp-1' }, row.product.name),
              ])
            ])
          }
        },
        {
          prop: 'user',
          label: '申请人',
          formatter: (row) => {
            return h('p', { }, row.user)
          }
        },
        {
          prop: 'amount',
          label: '退款金额',
          formatter: (row) => {
            return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}` )
          }
        },
        {
          prop: 'status',
          label: '退款状态',
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
            return h('div', { class: 'aftersales flex-c' }, [
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
  
  const handleView = (row:Aftersales.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewVisible.value = true
    })
  }
  
  /**
   * 处理表格行选择变化
   */
  const handleSelectionChange = (selection: Aftersales.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
  }
</script>
  