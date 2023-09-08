<template>
  <div class="login-content">
    <div class="login-box-style">
      <el-form
        ref="ruleFormRef"
        :model="ruleForm"
        status-icon
        :rules="rules"
        label-width="120px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="ruleForm.username"
            type="text"
            autocomplete="off"
            @keydown.enter="submitForm(ruleFormRef)"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="ruleForm.password" type="password" autocomplete="off" @keydown.enter="submitForm(ruleFormRef)"/>
        </el-form-item>
        <el-form-item>
          <el-button :plain="true" type="primary" @click="submitForm(ruleFormRef)"
            >登录</el-button> 
          <el-button @click="submitRegister">注册</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
  <div>
    <el-dialog v-model="dialogFormVisible" title="用户注册" width="40%">
      <!-- <p>用户注册</p>
      <el-button type="primary" @click="dialogFormVisible = false">
        取消
      </el-button> -->
      <UserRegister @submit="closeRegister"></UserRegister>
    </el-dialog>
  </div>
</template>
  
  <script lang="ts" setup>
  import { reactive, ref } from 'vue'
  import { FormInstance, FormRules, ElMessage } from 'element-plus'
  import router from '@/router'
  import { SuccessMsg, WarnMsg, ErrorMsg} from '@/utils/elmessage'
  import UserRegister from '../user/UserRegister'
  
  const ruleFormRef = ref<FormInstance>()
  const dialogFormVisible = ref(false)
  
  const ruleForm = reactive({
    username: '',
    password: '',
  })
  
  const rules = reactive<FormRules<typeof ruleForm>>({
    // password: [{ required: true, validator: validatePass, trigger: 'blur' }],
    // username: [{ required: true, validator: validUserName, trigger: 'blur' }]
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }]
  })
  
  const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
      if (valid) {
        login(ruleForm)
      } else {
        console.log('error submit!')
        return false
      }
    })
  }

  // 注册
  const submitRegister = () => {
    dialogFormVisible.value=true
  }
  // 关闭注册窗口
  const closeRegister = () => {
    // console.log('关闭窗口')
    dialogFormVisible.value=false
  }

  import { LoginReq } from '../../api/index'
  // login
  const login = (data: any) => {
    LoginReq(data).then((res) => {
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
      // 保存token
      const accessToken = JSON.stringify(resultData.data.LToken).replace(/\"/g, '')
      localStorage.setItem('token', accessToken)
      router.push("/productlists")
    })
  }
  </script>
  
<style scoped>
form {
  width: 350px;
  margin-top: 10%;
  margin-left: 5%;
}
@import '@/style/login.css'
</style>