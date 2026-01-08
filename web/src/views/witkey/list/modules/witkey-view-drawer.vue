<template>
    <ElDrawer
      v-model="visible"
      title="威客详情"
      size="30%"
      @close="handleClose"
    >
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <ElAvatar :src="detail?.avatar" shape="square" :size="60" />
                <div class="flex-col flex gap-2">
                    <h2 class="font-bold">
                        {{detail?.name}}
                    </h2>
                    <ElTag size="small" :type="getSex(detail?.sex!).type">{{ getSex(detail?.sex!).text }}</ElTag>
                </div>
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="所在地址">
                {{ address }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="接单状态">
                <ElTag :type="getStatus(detail?.status!).type">{{getStatus(detail?.status!).text}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="所属游戏">
                <ElTag type="primary">{{detail?.game}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="所得头衔">
                <ElTag type="primary">{{detail?.title}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="评分">
                <ElRate showScore
                :scoreTemplate="`${detail?.rate}分`"
                disabled
                :model-value="detail?.rate"
                />
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="入驻时间">
                {{ detail?.createTime }}
            </ElDescriptionsItem>
        </ElDescriptions>
        <template #footer>
            <ElSpace wrap>
                <div v-for="(item) in detail?.album" 
                    class="w-[100px] h-[100px]">
                    <ElImage
                        class="rounded-lg"
                        :src="item"
                        :preview-src-list="[item]"
                        >
                    </ElImage>
                </div>
            </ElSpace>
        </template>
    </ElCard>
    </ElDrawer>
  </template>
  
<script setup lang="ts">
import { fetchGetWitkeyDetail } from '@/api/witkey';
import { Status } from '@/enums/statusEnum';
import { SexType } from '@/enums/typeEnum';
import { codeToText } from 'element-china-area-data';

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

const SEX_CONFIG = {
  [SexType.Female]: { type: 'danger' as const, text: '女' },
  [SexType.Male]: { type: 'primary' as const, text: '男' },
  [SexType.Other]: { type: 'danger' as const, text: '其他' },
} as const

const getSex = (sex: number) => {
  return (
    SEX_CONFIG[sex as keyof typeof SEX_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}
const detail = ref<Witkey.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetWitkeyDetail({id:props.id!})
    detail.value = res
    loading.value = false
}


const address = computed(()=>{
    if (detail.value!.address.length  == 2) {
        return codeToText[detail.value!.address[0]]  + '/' + codeToText[detail.value!.address[1]]
    }
    return codeToText[detail.value!.address[0]]
})


/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const init = async () => {
  getData()
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
}
</script>
  