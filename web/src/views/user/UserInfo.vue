<template>
    <el-descriptions title="个人中心" column="1">
    <el-descriptions-item label="用户名:">{{ userInfo.name }}</el-descriptions-item>
    <el-descriptions-item label="电话:">{{ userInfo.phone }}</el-descriptions-item>
    <el-descriptions-item label="邮箱:">{{ userInfo.email }}</el-descriptions-item>
    <el-descriptions-item label="QQ:">{{ userInfo.qq }}</el-descriptions-item>
    <el-descriptions-item label="微信:">{{ userInfo.wx }}</el-descriptions-item>
    </el-descriptions>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'

const userInfo = reactive({
    name: "admin",
    phone: "13800000000",
    email: "admin@mail.com",
    qq: "123",
    wx: "13800000000"
})
</script>