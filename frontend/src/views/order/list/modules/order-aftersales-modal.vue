<template>
    <ElDrawer
        v-model="visible"
        title="售后工单"
        size="30%"
        @close="handleClose"
        :close-on-press-escape="false"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="售后原因" prop="reason">
            <ElInput
            v-model="form.reason"
            type="textarea"
            :rows="3"
            placeholder="请输入售后原因"
          />
        </ElFormItem>

        <ElFormItem label="补偿方式" prop="type">
            <ElRadioGroup  v-model="form.type">
                <ElRadio :value="AfterSalesType.Refund">仅退款</ElRadio>
                <ElRadio :value="AfterSalesType.Compensate">仅补偿</ElRadio>
                <ElRadio :value="AfterSalesType.RefundCompensate">退款+补偿</ElRadio>
            </ElRadioGroup>
        </ElFormItem>
        <ElFormItem label="售后金额" prop="amount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.amount"
            placeholder="请输入售后金额" 
            controls-position="right"/>
        </ElFormItem>
        <ElFormItem label="售后详情"  prop="content">
            <ElInput
                v-model="form.content"
                type="textarea"
                :rows="9"
                placeholder="请输入售后方案详情"
            />
        </ElFormItem>
        <ElFormItem label="联系方式" prop="contactMode">
            <ElRadioGroup  v-model="form.contactMode">
                <ElRadio :value="ContactMode.Wechat">微信</ElRadio>
                <ElRadio :value="ContactMode.QQ">QQ</ElRadio>
                <ElRadio :value="ContactMode.Phone">手机</ElRadio>
            </ElRadioGroup>
        </ElFormItem>
        <ElFormItem label="联系号码" prop="contactNumber">
            <ElInput
            v-model="form.contactNumber"
            placeholder="请输入联系号码"
          />
        </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchPostOrderAftersales } from '@/api/order';
import { ContactMode } from '@/enums/modeEnum';
import { AfterSalesType } from '@/enums/typeEnum';
import type { FormInstance, FormRules } from 'element-plus'

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


const formRef = ref<FormInstance>()
/**
 * 表单数据
 */
const form = reactive<Order.Params.Aftersales>({
    id: 0, // 权限ID
    type: AfterSalesType.Refund,
    amount: null,
    reason: null,
    content: "",
    contactMode: ContactMode.Wechat,
    contactNumber: "",
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
    type: [
        { required: true, message: '请选择售后类型', trigger: 'blur' },
    ],
    amount: [
        { required: true, message: '请输入售后金额', trigger: 'blur' },
    ],
    reason: [
        { required: true, message: '请输入售后原因', trigger: 'blur' },
    ],
    content: [
        { required: true, message: '请输入售后内容', trigger: 'blur' },
    ],
    contactMode: [
        { required: true, message: '请选择联系方式', trigger: 'blur' },
    ],
    contactNumber: [
        { required: true, message: '请输入联系号码', trigger: 'blur' },
    ],
})


/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
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
        amount: 0,
    })
}

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
  emit('submit')
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
        await fetchPostOrderAftersales(form)
        ElMessage.success('成功提交售后工单')
        handleClose()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>