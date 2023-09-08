<template>
 <div class="register-box-style">
    <el-form 
        :model="registerForm"
        ref="ruleFormRef"
        :rules="rules"
        :label-position="labelPosition"
    >
        <el-form-item label="用户名" :label-width="formLabelWidth" prop="username">
            <el-input v-model="registerForm.username"  type="text" />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth" prop="password">
            <el-input v-model="registerForm.password"  type="password" />
        </el-form-item>
    </el-form>
 </div>
 <div class="register-button-style">
    <el-button @click="cancelRegister">
        取消
    </el-button>
    <el-button type="primary" @click="submitToRegister(ruleFormRef)">
        注册
    </el-button>
 </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import { FormInstance, FormRules, ElMessage, ElMessageBox } from 'element-plus'

// 对齐
const labelPosition = ref('left')
const ruleFormRef = ref<FormInstance>()
const formLabelWidth = '80px'
const registerForm = reactive({
    username: '',
    password: ''
})
const emit = defineEmits(['submit'])

// 账号密码必填
const rules = reactive<FormRules<typeof registerForm>>({
    // password: [{ required: true, validator: validatePass, trigger: 'blur' }],
    // username: [{ required: true, validator: validUserName, trigger: 'blur' }]
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }]
  })

// 取消注册
const cancelRegister = () => {
    emit('submit')
}

// 注册
const submitToRegister = (formEl: FormInstance | undefined) => {
  ElMessageBox.confirm(
    '是否确认注册用户：'+registerForm.username+' ?' ,
    'Info',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'info',
    }
  ).then(() => {
    if (!formEl) return
    formEl.validate((valid) => {
      if (valid) {
        register(registerForm)
      } else {
        console.log('error submit!')
        return false
      }
    })
  }).catch(() => {
    ElMessage({
        type: 'info',
        message: '取消注册',
    })
    cancelRegister()
  })
}

// 调用注册接口
import { RegisterReq } from '@/api'
const register = (data: any) => {
    RegisterReq(data).then((res) => {
        let resData = JSON.stringify(res)
        let resultData = JSON.parse(resData)
        if (resultData.code !== 0) {
            ElMessage({
                showClose: true,
                message: resultData.msg,
                type: 'error',
            })
            return
        }
        // 注册成功，不登陆
        ElMessage({
            showClose: true,
            message: resultData.msg,
            type: 'success'
        })
        // 打印token
        const accessToken = JSON.stringify(resultData.data.LToken).replace(/\"/g, '')
        console.log(accessToken)
        // 关闭dialog
        cancelRegister()
    })
}

</script>

<style scoped>
@import '@/style/register.css'
 
</style>
