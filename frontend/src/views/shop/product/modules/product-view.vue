<template>
    <ElDrawer
      v-model="visible"
      title="商品详情"
      size="50%"
      @close="handleClose"
    > 
      <div v-loading="loading">
        <div class="flex justify-between items-center mb-2">
          <h2 class="font-bold text-xl">商品编码:{{detail?.code}}</h2>
          <!-- <ElButton type="primary">立即下单</ElButton> -->
        </div>
        <div class="flex gap-6">
          <ElImage class="w-80 h-60 rounded object-cover" 
          :src="detail?.pic"/>
          <div class="flex-1">
            <h2 class="font-bold text-xl mb-3">{{ detail?.name }}</h2>
            <div class="flex justify-between items-center mb-3">
              <ElRate
                :model-value="detail?.rate"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value} 评分"
              />
              <div class="flex items-center gap-2">
                <ElTag type="success">{{getProductStatusConfig(detail?.status!).text}}</ElTag>
                <ElTag type="primary">{{detail?.category}}</ElTag>
                <ElTag type="success">{{detail?.game}}</ElTag>
              </div>
            </div>
            <div class="flex items-center gap-2 mb-3">
              <ElText class="font-bold" size="large" type="danger">{{ detail?.price }}{{ site.symbol }}/{{ detail?.unit }}</ElText>
              <ElDivider  direction="vertical" />
              <ElTag type="primary">限购{{detail?.purchaseLimit != 0 ? detail?.purchaseLimit : 0 }}</ElTag>
            </div>
            <div class="bg-primary/10 p-3 rounded text-xs">
              {{detail?.description}}
            </div>
          </div>
        </div>
        <ElDivider />
        <h2 class="font-bold text-base">商品详情</h2>
        <div v-html="detail?.content"></div>
      </div>
    </ElDrawer>
  </template>
  
<script setup lang="ts">
import { fetchGetProductDetail } from '@/api/product';
import { ProductStatus } from '@/enums/statusEnum';
import { useSiteStore } from '@/store/modules/site';

interface Props {
  modelValue: boolean
  id?: number | null
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
  id: 0
})
interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'submit'): void
}
const emit = defineEmits<Emits>()


// 对话框显示控制
const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
})
// 商品状态配置
const PRODUCT_STATUS_CONFIG = {
  1: { type: 'warning' as const, text: '下架' },
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


const {getInfo:site} = useSiteStore()
const loading = ref(false)
const detail = ref<Product.Response.Detail>()
const getData = async () => {
  loading.value = true 
  try {
    const res = await fetchGetProductDetail({id:props.id!})
    detail.value = res
  } finally {
    loading.value = false 
  }
}

/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initData = async () => {
  getData()
}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) initData()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
}
/**
 * 提交表单
 * 验证通过后触发提交事件
 */
const handleSubmit = async () => {
  
}
</script>
  