<!-- 订单管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="order-page art-full-height">
      <!-- 搜索栏 -->
      <OrderSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></OrderSearch>
  
      <ElCard class="art-table-card" shadow="never">
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData"></ArtTableHeader>
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
  
        <!-- 订单弹窗 -->
        <OrderView
          v-model="viewDrawerVisible"
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
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { fetchGetOrderList, } from '@/api/order'
import { OrderStatus} from '@/enums/statusEnum'
import OrderSearch from './modules/order-search.vue'
import OrderView from './modules/order-view.vue'

const route = useRoute()
defineOptions({ name: 'Order' })
const siteStore = useSiteStore()


// 弹窗相关

const viewDrawerVisible = ref(false)

const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])
// 搜索表单
const searchForm = ref({
    code: undefined,
    status: undefined
})

// 订单状态配置
const ORDER_STATUS_CONFIG = {
    [OrderStatus.PendingPayment]: { type: 'info' as const, text: '待支付' },
    [OrderStatus.PendingService]: { type: 'primary' as const, text: '待服务' },
    [OrderStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
    [OrderStatus.Completed]: { type: 'success' as const, text: '已完成' },
    [OrderStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
    [OrderStatus.Refund]: { type: 'warning' as const, text: '已退款' },
} as const

/**
 * 获取订单状态配置
 */
const getOrderStatusConfig = (status: number) => {
    return (
        ORDER_STATUS_CONFIG[status as keyof typeof ORDER_STATUS_CONFIG] || {
            type: 'info' as const,
            text: '未知'
        }
    )
}

// 监听查询参数变化
watch(
  () => route.query,
  (newQuery) => {
    handleSearch({code:route.query.code})
  }
)

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
        apiFn: fetchGetOrderList,
        apiParams:{
            code: "",
            status: 0,
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            {
                prop: 'code',
                label: '订单号',
                width: 280,
            },
            {
                prop: 'productInfo',
                label: '商品信息',
                width: 400,
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
                        h('div', { class: 'flex-1 flex gap-2' }, [
                            h(ElTag, { type:"primary",size:"small" }, () => row.product.category),
                            h(ElTag, { type:"primary",size:"small" }, () => row.product.game)
                        ])
                    ])
                ])
                }
            },
            {
                prop: 'orderInfo',
                label: '用户信息',
                width: 200,
                formatter: (row) => {
                    return h('p', { }, row.user)
                }
            },
            {
                prop: 'totalAmount',
                label: '订单总额',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => `${row.totalAmount}${siteStore.getInfo.symbol}`)
                }
            },
            {
                prop: 'status',
                label: '订单状态',
                formatter: (row) => {
                const statusConfig = getOrderStatusConfig(row.status)
                return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
                }
            },
            {
                prop: 'createTime',
                label: '下单时间',
                width: 180,
                sortable: true
            },
            {
                prop: 'operation',
                label: '操作',
                width: 120,
                fixed: 'right', // 固定列
                formatter: (row) =>{
             
                return h('div', { class: 'order flex-c' }, [
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


const handleView = (row:Order.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        viewDrawerVisible.value = true
    })
}

/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Order.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
}
</script>
