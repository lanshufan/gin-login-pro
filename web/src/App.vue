<template>
  <div class="common-layout">
    <el-container>
      <!-- <el-header class="header-layout" >Header</el-header> -->
      <el-header class="header-layout">
        <div class="toolbar">
          <el-dropdown>
            <el-icon style="margin-right: 8px; margin-top: 1px; font-size: 20px;"
              ><User
            /></el-icon>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="UserInfo">个人中心</el-dropdown-item>
                <el-dropdown-item @click="Logout">登出</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <span>User</span>
        </div>
      </el-header>
      <el-container>
        <el-aside class="aside-layout">
          <el-scrollbar>
            <el-menu unique-opened default-active="product-list" router>
              <el-sub-menu index="product-manage">
                <template #title>产品管理</template>
                <el-menu-item index="productlists">产品列表</el-menu-item>
                <el-menu-item index="productorder">产品订单</el-menu-item>
              </el-sub-menu>
              <el-sub-menu index="user-manage">
                <template #title>用户管理</template>
                <el-menu-item index="userinfo">个人中心</el-menu-item>
                <el-menu-item index="user-power">权限管理</el-menu-item>
              </el-sub-menu>
            </el-menu>
          </el-scrollbar>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage, ElMessageBox } from 'element-plus'
import type { Action } from 'element-plus'
import 'element-plus/dist/index.css'
import router from '@/router'

const Logout = () => {
  ElMessageBox.alert('您确认要登出吗？', '登出', {
    // if you want to disable its autofocus
    // autofocus: false,
    confirmButtonText: '确认',
    callback: (action: Action) => {
      ElMessage({
        type: 'warning',
        message: '退出登录',
      })
    },
  })
}

const UserInfo = () => {
  router.push("/userinfo")
}
</script>

<style>
.header-layout {
  text-align: right;
  font-size: 16px;
  margin-top: 1%;
  border-bottom: 1px solid #9f9b9b;
}
.aside-layout {
  height: 100%;
  width: 15%;
}
</style>
