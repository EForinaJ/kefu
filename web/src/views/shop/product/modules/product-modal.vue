<template>
    <ElDialog
      title="商品下单"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="下单用户" prop="userId">
            <ElSelect
                v-model="form.userId"
                filterable
                remote
                reserve-keyword
                placeholder="请输入用户手机号"
                :remote-method="remoteMethod"
                :loading="loading"
          
            >
                <ElOption
                    v-for="item in userOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                />
            </ElSelect>
        </ElFormItem>
        <ElFormItem label="下单数量" prop="quantity">
            <ElInputNumber 
            :precision="0"
            :min="1"
            style="width: 100%;"
            v-model="form.quantity" placeholder="请输入下单数量"
            controls-position="right"> </ElInputNumber>
        </ElFormItem>
        <ElFormItem label="下单备注" prop="requirements">
            <ElInput
            v-model="form.requirements"
            type="textarea"
            :rows="3"
            placeholder="请输入下单备注"
          />
        </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleCancel">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import { fetchGetUserOptionsList } from '@/api/site';
import type { FormInstance, FormRules } from 'element-plus'
import { useDebounce } from '@/hooks/index';
import { fetchGetProductPurchase } from '@/api/product';
import { useRouter } from 'vue-router'
interface Props {
    visible: boolean
    id?: number | null
}

interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
}

const props = withDefaults(defineProps<Props>(), {
    visible: false,
    id: 0,
})

const emit = defineEmits<Emits>()
const router = useRouter()


const formRef = ref<FormInstance>()

/**
 * 弹窗显示状态双向绑定
 */
const visible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
})
/**
 * 表单数据
 */
const form = reactive<Product.Params.Model>({
    id: 0, 
    userId: null,
    quantity: null,
    requirements: null,
})


const rules = reactive<FormRules>({
    quantity: [
        { required: true, message: '请输入下单数量', trigger: 'blur' },
    ],
    userId: [
        { required: true, message: '请输入用户ID', trigger: 'blur' },
    ],
})


// 使用防抖 hook
const { debounce } = useDebounce()
interface ListItem {
  value: number
  label: string,
}

const loading = ref(false)
const userOptions = ref<ListItem[]>([])
const getUserOptions =  async (phone:string) => {
    try {
        const res = await fetchGetUserOptionsList({phone})
        userOptions.value = res.list.map((item) => {
            return {
                value: item.id,
                label: item.name
            }
        })
    } finally {
        loading.value = false
    }
}

// 创建防抖版本
const getUserOptionsDebounced = debounce((query: string) => {
  loading.value = true
  getUserOptions(query)
}, 500)
const remoteMethod = (query: string) => {
  if (query) {
    getUserOptionsDebounced(query)
  } else {
    userOptions.value = []
  }
}


/**
 * 监听弹窗打开，初始化表单数据
 */
watch(
    () => props.visible,
    (newVal) => {
    if (newVal) initForm()
    }
)

/**
 * 初始化表单数据
 * 根据弹窗类型填充表单或重置表单
 */
const initForm = async () => {
    Object.assign(form, {
        id: props.id!,
    })
}

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
    formRef.value?.resetFields()
}

/**
 * 取消操作
 */
const handleCancel = (): void => {
    handleClose()
    emit('update:visible', false)
}

/**
 * 提交表单
 * 验证通过后调用接口保存数据
 */
const handleSubmit = async () => {
    if (!formRef.value) return
    try {
        await formRef.value.validate()
        // TODO: 调用新增/编辑接口
        const res = await fetchGetProductPurchase(form)
        console.log(res)

        ElMessage.success('下单成功')
        handleClose()
        handleCancel()
        router.push({
            name: 'OrderList',
            query: {
                code: res.code,
            }
        })
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>