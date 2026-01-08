<template>
  <ElDrawer
    v-model="visible"
    title="添加威客"
    width="30%"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="威客名" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入威客名" />
      </ElFormItem>
      <ElFormItem label="手机号" prop="phone">
        <ElInput v-model="formData.phone" placeholder="请输入手机号" />
      </ElFormItem>
      <ElFormItem label="密码" prop="password">
        <ElInput v-model="formData.password" placeholder="请输入密码" />
      </ElFormItem>
      <ElFormItem label="所属游戏" prop="gameId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.gameId"
          placeholder="请选择游戏"
          @change="handleChangeGame"
        >
          <el-option
            v-for="item in gameOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem v-if="formData.gameId" label="所属头衔" prop="titleId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.titleId"
          placeholder="请选择头衔"
        >
          <el-option
            v-for="item in titleOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="性别" prop="sex">
        <ElSelect v-model="formData.sex">
          <ElOption label="男" :value="SexType.Male" />
          <ElOption label="女" :value="SexType.Female" />
          <ElOption label="其他" :value="SexType.Other" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="头像" prop="avatar">
        <ArtUpload
        @success="handleAvatar"
        >
          <ElAvatar 
          shape="square"
          v-if="formData.avatar != ''" :size="80" :src="formData.avatar" />
          <ElButton v-if="formData.avatar == ''" type="primary">选择文件</ElButton>
        </ArtUpload>
      </ElFormItem>
      <ElFormItem label="所在地" prop="address">
        <ElCascader
          style="width: 100%;"
          v-model="formData.address"
          :options="region"
          placeholder="请选择省市区"
        />
      </ElFormItem>
      <ElFormItem label="出生日期" prop="birthday">
        <ElDatePicker
          style="width: 100%;"
          v-model="formData.birthday"
          type="date"
          placeholder="请选择出生日期"
          format="YYYY/MM/DD"
          value-format="x"
        />
      </ElFormItem>
      <ElFormItem label="个人介绍" prop="description">
        <ElInput :rows="5" type="textarea" v-model="formData.description" placeholder="请输入个人介绍" />
      </ElFormItem>
      <ElFormItem label="相册" prop="album">
        <ElSpace wrap>
          <div v-for="(item,index) in formData.album" 
              class="w-[100px] h-[100px] relative">
              <ElImage
                class="rounded-lg"
                :src="item"
                :preview-src-list="[item]"
                >
              </ElImage>
              <div @click="deleteAlbum(index)"  class="absolute top-0 right-0 cursor-pointer">
                <ArtSvgIcon icon="solar:close-circle-bold" class="text-red-500 text-2xl"/>
              </div>
          </div>
          <ArtUpload
          @success="handleAlbum"
          >
            <div class="w-[100px] h-[100px] flex items-center justify-center border-1 border-solid rounded-lg cursor-pointer">
              <ArtSvgIcon icon="solar:upload-square-bold" class="text-2xl"/>
            </div>
          </ArtUpload>
        </ElSpace>
      </ElFormItem>
      <ElFormItem label="评分" prop="rate">
        <ElRate 
        allow-half
        v-model="formData.rate"/>
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="Status.Disable">禁用</ElRadio>
          <ElRadio :value="Status.Enable">启用</ElRadio>
        </ElRadioGroup>
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
import { SexType } from '@/enums/typeEnum';
import type { FormInstance, FormRules  } from 'element-plus'
import type { CascaderOption } from 'element-plus';
import { provinceAndCityData } from 'element-china-area-data';
import { fetchGetGameOptionsList, fetchGetTitleOptionsList } from '@/api/site';
import { fetchPostWitkeyCreate} from '@/api/witkey';
import { Status } from '@/enums/statusEnum';

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
const formData = reactive<Witkey.Params.Model>({
  name: null,
  phone: null,
  password: null,
  address: [],
  birthday: 0,
  description: null,
  sex: SexType.Other,
  avatar: "",
  status: Status.Enable,
  titleId: null,
  gameId: null,
  album: [],
  rate: 0,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入威客名称', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式', trigger: 'blur' }
  ],
  gameId: [
    { required: true, message: '请选择所属游戏', trigger: 'blur' },
  ],
  titleId: [
    { required: true, message: '请选择所属头衔', trigger: 'blur' },
  ]
}

const handleAvatar = (e:string) => {
  formData.avatar = e
}
const handleAlbum = (e:string) => {
  formData.album.push(e)
}
const deleteAlbum = (i:number) => {
  formData.album.splice(i,1)
}
interface ListItem {
  value: number
  label: string
}
const gameOptions = ref<ListItem[]>([])
const getGameOptions =  async () => {
  const res = await fetchGetGameOptionsList()
  gameOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}

const handleChangeGame = async (v:any) =>{
  await getTitleOptions(v)
}
const titleOptions = ref<ListItem[]>([])
const getTitleOptions = async (id:number) => {
  const res = await fetchGetTitleOptionsList({gameId:id})
  titleOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}





/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {
  await getGameOptions()
  Object.assign(formData, {
    name: null,
    phone: null,
    password: null,
    address: [],
    birthday: 0,
    description: null,
    sex: SexType.Other,
    avatar: "",
    status: Status.Enable
  })
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
      await fetchPostWitkeyCreate(formData)
      ElMessage.success('添加成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
