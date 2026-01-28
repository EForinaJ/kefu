<template>
    <ElDrawer
      v-model="visible"
      title="订单详情"
      size="50%"
      @close="handleClose"
      :close-on-press-escape="false"
    >
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <div class="font-bold">
                    订单编号: {{ detail?.code }}
                </div>
                <el-divider direction="vertical" />
                <div class="font-bold">
                    下单时间: {{ detail?.createTime }}
                </div>
            </div>
        </template>
        <template #footer>
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                    <ElAvatar :src="detail?.user.avatar" shape="square" :size="28" />
                    <div class="flex-col flex gap-2">
                        <h2 class="font-bold">
                            {{detail?.user.name}}
                        </h2>
                    </div>
                </div>
                <ElSpace>
                    <!-- 未支付 -->
                    <ElButton v-if="detail?.status == OrderStatus.PendingPayment"
                    @click="handleAddDiscount" size="small" type="primary">添加优惠</ElButton>
                    <ElButton  v-if="detail?.status == OrderStatus.PendingPayment" 
                    @click="handlePaid" size="small" type="success">确认收款</ElButton>
                    <ElButton v-if="detail?.status == OrderStatus.PendingPayment" 
                    @click="handleCancel" size="small" type="danger">关闭订单</ElButton>

                    <!-- 已支付 -->
                    <ElButton v-if="detail?.status == OrderStatus.PendingService"
                    @click="handleAftersales" size="small" type="warning">售后工单</ElButton>
                    <ElButton v-if="detail?.status == OrderStatus.PendingService"
                    @click="handleDistribute" size="small" type="primary">立即派单</ElButton>
                    <ElButton v-if="detail?.status == OrderStatus.PendingService"
                    @click="handleStart" size="small" type="primary">服务开始</ElButton>

                    <!-- 进行中 -->
                    <ElButton v-if="detail?.status == OrderStatus.InProgress"
                    @click="handleAftersales" size="small" type="warning">售后退款</ElButton>
                    <ElButton v-if="detail?.status == OrderStatus.InProgress"
                    @click="handleComplete" size="small" type="primary">结束完成</ElButton>
                </ElSpace>
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="1"  label="订单状态">
                <ElTag :type="getOrderStatusConfig(detail?.status!).type">{{ getOrderStatusConfig(detail?.status!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="1"  label="支付状态">
                <ElTag :type="getOrderPayStatus(detail?.payStatus!).type">{{ getOrderPayStatus(detail?.payStatus!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="1"  label="支付方式">
                <ElTag :type="getOrderPayMode(detail?.payMode!).type">{{ getOrderPayMode(detail?.payMode!).text }}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="1"  label="支付时间">
                {{ detail?.payTime  ? detail?.payTime :  '未支付' }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="1"  label="开始时间">
                {{ detail?.startTime  ? detail?.startTime : "未开始" }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="1"  label="结束时间">
                {{ detail?.finishTime  ? detail?.finishTime :  "未结束" }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="用户要求">
                {{ detail?.requirements == "" ? '无备注' : detail?.requirements }}
            </ElDescriptionsItem>
        </ElDescriptions>
    </ElCard>

    <ElCard v-loading="loading" shadow="never" class="mt-[20px]">
        <template #header>
            <div class="font-bold">
                商品详情
            </div>
        </template>
        <ElTable :data="productData" style="width: 100%">
            <ElTableColumn prop="productInfo" label="商品信息" width="280">
                <template #default="scope">
                    <div class="flex gap-3">
                        <ElImage class="size-12 rounded-md" :src="scope.row.pic"/>
                        <h2 class="text-base line-clamp-1">{{ scope.row.name }}</h2>
                    </div>
                </template>
            </ElTableColumn>
            <ElTableColumn prop="game" label="所属游戏" width="180" >
                <template #default="scope">
                    <ElTag type="primary">{{ scope.row.game }}</ElTag>
                </template>
            </ElTableColumn>
            <ElTableColumn prop="category" label="所属分类" width="180" >
                <template #default="scope">
                    <ElTag type="primary">{{ scope.row.category }}</ElTag>
                </template>
            </ElTableColumn>
            <ElTableColumn prop="unitPrice" label="单价">
                <template #default="scope">
                    {{ `${scope.row.unitPrice}${site.symbol}` }}
                </template>
            </ElTableColumn>
            <ElTableColumn prop="quantity" label="数量">
                <template #default="scope">
                    {{ `${scope.row.quantity}${scope.row.unit}` }}
                </template>
            </ElTableColumn>
        </ElTable>
        <template #footer>
            <div class="flex justify-end">
                <div class="flex items-center gap-2">
                    <div class="flex items-center">
                    订单总额：<span class="text-[red] font-bold">{{ detail?.totalAmount }} {{ site.symbol }}</span>
                    </div>
                    <div class="font-bold">
                    -
                    </div>
                    <div class="flex items-center">
                    优惠金额：<span class="text-[red] font-bold">{{ detail?.discountAmount }}  {{ site.symbol }}</span>
                    </div>
                </div>
            </div>
            <div class="flex justify-end">
                <div class="flex items-center">
                    需付金额：<span class="text-[red] font-bold">{{ detail?.actualAmount }} {{ site.symbol }}</span>
                </div>
            </div>
            <div class="flex justify-end">
                <div class="flex items-center">
                    已支付：<span class="text-[red] font-bold">{{ detail?.payAmount ? detail?.payAmount :  0 }} {{ site.symbol }}</span>
                </div>
            </div>
        </template>
    </ElCard>

    <OrderAddDiscountModal
        v-model:visible="addDiscountModalVisible"
        :id="id"
        @submit="getData"
    />
    
    <OrderAftersalesModal
        v-model="aftersalesModalVisible"
        :id="id"
        @submit="getData"
    />
    <OrderPaidModal
        v-model:visible="paidModalVisible"
        :id="id"
        @submit="getData"
    />
    <OrderDistributeModal
        v-model:visible="distributeModalVisible"
        :id="id"
        @submit="getData"
    />
  </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetOrderDetail, fetchPostOrderCancel, fetchPostOrderComplete, fetchPostOrderStart } from '@/api/order';
import { PayMode } from '@/enums/modeEnum';
import { OrderStatus, PayStatus } from '@/enums/statusEnum';
import { useSiteStore } from '@/store/modules/site';
import { ElMessageBox } from 'element-plus';
import OrderDistributeModal from './order-distribute-modal.vue';
import OrderPaidModal from './order-paid-modal.vue';
import OrderAftersalesModal from './order-aftersales-modal.vue';
import OrderAddDiscountModal from './order-add-discount-modal.vue';

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

// 订单状态配置
const ORDER_STATUS_CONFIG = {
  [OrderStatus.PendingPayment]: { type: 'info' as const, text: '待支付' },
  [OrderStatus.PendingService]: { type: 'primary' as const, text: '待服务' },
  [OrderStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
  [OrderStatus.Completed]: { type: 'success' as const, text: '已完成' },
  [OrderStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
  [OrderStatus.AfterSales]: { type: 'warning' as const, text: '已售后' },
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

// 支付方式配置
const ORDER_PAY_MODE = {
  [PayMode.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [PayMode.Wechat]: { type: 'success' as const, text: '微信' },
  [PayMode.Balance]: { type: 'warning' as const, text: '余额' },
  [PayMode.PersonalTransfer]: { type: 'primary' as const, text: '个人转账' },
} as const

/**
 * 获取支付方式配置
 */
const getOrderPayMode = (mode: number) => {
  return (
    ORDER_PAY_MODE[mode as keyof typeof ORDER_PAY_MODE] || {
      type: 'info' as const,
      text: '未支付'
    }
  )
}

// 支付状态配置
const ORDER_PAY_STATUS = {
  [PayStatus.Success]: { type: 'success' as const, text: '支付成功' },
  [PayStatus.Pending]: { type: 'primary' as const, text: '待支付' },
  [PayStatus.Fail]: { type: 'warning' as const, text: '支付失败' },
} as const

/**
 * 获取支付状态配置
 */
const getOrderPayStatus = (status: number) => {
  return (
    ORDER_PAY_STATUS[status as keyof typeof ORDER_PAY_STATUS] || {
      type: 'info' as const,
      text: '未支付'
    }
  )
}

const {getInfo:site} = useSiteStore()
const detail = ref<Order.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetOrderDetail({id:props.id!})
    detail.value = res
    loading.value = false
}
const productData = computed(()=>{
    return detail.value ? [detail.value?.product] :[]
})

const addDiscountModalVisible = ref(false)
const handleAddDiscount = (): void => {
    nextTick(() => {
        addDiscountModalVisible.value = true
    })
}

const aftersalesModalVisible = ref(false)
const handleAftersales = () => {
    nextTick(() => {
        aftersalesModalVisible.value = true
    })
}


const distributeModalVisible = ref(false)
const handleDistribute = () => {
    nextTick(() => {
        distributeModalVisible.value = true
    })
}

const paidModalVisible = ref(false)
const handlePaid = (): void => {
    nextTick(() => {
        paidModalVisible.value = true
    })
}
const handleCancel = (): void => {
ElMessageBox.confirm(`确定要关闭订单吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
}).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderCancel({id:props.id!})
    getData()
})
.catch(() => {
    ElMessage.info('已取消')
})
}

const handleStart = (): void => {
ElMessageBox.confirm(`确定要开始服务吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
}).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderStart({id:props.id!})
    getData()
})
.catch(() => {
    ElMessage.info('已取消')
})
}

const handleComplete = (): void => {
ElMessageBox.confirm(`确定服务完成吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
}).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderComplete({id:props.id!})
    getData()
})
.catch(() => {
    ElMessage.info('已取消')
})
}



/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const init = async () => {
  await getData()
}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) init()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
  emit('submit')
}
</script>
  