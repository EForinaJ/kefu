<template>
  <ElDrawer
    v-model="visible"
    title="上报工单"
    width="30%"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="工单标题" prop="title">
        <ElInput v-model="formData.title" placeholder="请输入工单标题" />
      </ElFormItem>
      <ElFormItem label="紧急程度" prop="priority">
        <ElSelect v-model="formData.priority">
          <ElOption label="紧急" :value="PriorityType.urgent" />
          <ElOption label="高" :value="PriorityType.secondary" />
          <ElOption label="中" :value="PriorityType.tertiary" />
          <ElOption label="低" :value="PriorityType.quaternary" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="上报内容" prop="content">
        <ElInput v-model="formData.content" 
        :rows="6"
        type="textarea"
        placeholder="请输入上报内容" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="drawer-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDrawer>
</template>

<script setup lang="ts">
import { PriorityType } from '@/enums/typeEnum';
import type { FormInstance, FormRules  } from 'element-plus'
import type { CascaderOption } from 'element-plus';
import { provinceAndCityData } from 'element-china-area-data';
import { fetchPostWorkOrderCreate } from '@/api/workorder';

const region = provinceAndCityData as CascaderOption[]
interface Props {
  modelValue: boolean
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
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



// 表单实例
const formRef = ref<FormInstance>()

// 表单数据
const formData = reactive<WorkOrder.Params.Model>({
  title: null,
  content:  null,
  priority: null,
})

// 表单验证规则
const rules: FormRules = {
  title: [
    { required: true, message: '请输入工单标题', trigger: 'blur' },
  ],
  content: [
    { required: true, message: '请输入上报内容', trigger: 'blur' },
  ],
  priority: [
    { required: true, message: '请选择紧急程度', trigger: 'blur' },
  ]
}


/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {

}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) initFormData()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
  formRef.value?.resetFields()
}
/**
 * 提交表单
 * 验证通过后触发提交事件
 */
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      await fetchPostWorkOrderCreate(formData)
      ElMessage.success('上报成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
