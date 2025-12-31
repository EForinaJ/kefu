<!-- 商品管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
    <div class="product-page art-full-height">
      <!-- 搜索栏 -->
      <ProductSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></ProductSearch>
  
      <ElCard class="art-table-card" shadow="never">
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
          <template #left>
            <ElSpace wrap>
              <!-- <ElButton 
              :disabled="selectedRows.length == 0"
              @click="handleBatchDelete" type="danger" v-ripple>批量删除</ElButton> -->
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
  
        <!-- 商品弹窗 -->
        <ProductView
          v-model="viewDrawerVisible"
          :id="id"
          @submit="refreshData"
        />

        <ProductModal
            v-model:visible="modalVisible"
            :id="id"
            @submit="refreshData"
        />
      </ElCard>
    </div>
  </template>
  
<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
// import ProductSearch from './modules/product-search.vue'
import { ElTag, ElMessageBox, ElImage, ElRate } from 'element-plus'


import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { fetchGetProductList } from '@/api/product';
import { useSiteStore } from '@/store/modules/site';
import ProductSearch from './modules/product-search.vue';
import ProductView from './modules/product-view.vue';
import ProductModal from './modules/product-modal.vue';



const { hasAuth } = useAuth();
defineOptions({ name: 'Product' })

const siteStore = useSiteStore()



// 弹窗相关
const modalVisible = ref(false)
const viewDrawerVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
    name: undefined,
    gameId: undefined,
    categoryId: undefined,
    status: undefined,
})

// 商品状态配置
const PRODUCT_STATUS_CONFIG = {
1: { type: 'info' as const, text: '下架' },
2: { type: 'success' as const, text: '在售' },
} as const

/**
 * 获取商品状态配置
 */
const getProductStatusConfig = (status: number) => {
    return (
        PRODUCT_STATUS_CONFIG[status as keyof typeof PRODUCT_STATUS_CONFIG] || {
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
        apiFn: fetchGetProductList,
        apiParams:{
            name: "",
            status: 0,
            code: ""
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
        { 
            width: 130,
            prop: 'code', 
            label: '商品编号',
        },
        {
            prop: 'productInfo',
            label: '商品信息',
            width: 280,
            formatter: (row) => {
            return h('div', { class: 'product flex-c' }, [
                h(ElImage, {
                    class: 'size-9.5 rounded-md',
                    src: row.pic,
                    previewSrcList: [row.pic],
                    // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
                    previewTeleported: true
                }),
                    h('div', { class: 'ml-2' }, [
                    h('p', { class: 'product-name' }, row.name),
                    h('p', { class: 'email' }, row.category),
                ])
            ])
            }
        },
        { 
            prop: 'game', 
            label: '游戏领域',
            formatter: (row) => {
            return h(ElTag, { type: "primary" }, () => `${row.game}`)
            }
        },
        { 
            prop: 'rate', 
            label: '评分',
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
            prop: 'price', 
            label: '商品售价',
            width: 120,
            formatter: (row) => {
            return h(ElTag, { type: "primary" }, () => `${row.price}${siteStore.getInfo.symbol}`)
            } 
        },
        {
            prop: 'status',
            label: '状态',
            width: 100,
            formatter: (row) => {
            const statusConfig = getProductStatusConfig(row.status)
            return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
            }
        },
        {
            prop: 'createTime',
            label: '创建日期',
            sortable: true
        },
        {
            prop: 'operation',
            label: '操作',
            width: 120,
            fixed: 'right', // 固定列
            formatter: (row) =>{
            return h('div', { class: 'product flex-c' }, [
                h(ArtButtonTable, {
                    type: 'view',
                    onClick: () => handleView(row)
                }),
                h(ArtButtonTable, {
                    type: 'edit',
                    icon: 'solar:lamp-bold',
                    onClick: () => handleBuy(row)
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

const handleView = (row:Product.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        viewDrawerVisible.value = true
    })
}
const handleBuy = (row:Product.Response.Info) => {
    id.value = row.id
    nextTick(() => {
        modalVisible.value = true
    })
}



  
/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Product.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
}
</script>
  