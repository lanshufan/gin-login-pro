<template>
    <el-table :data="productData" border style="width: 810px;">
        <el-table-column prop="name" label="名称" width="250" />
        <el-table-column prop="author" label="作者" width="180" />
        <el-table-column prop="date" label="日期" width="200" />
        <el-table-column label="操作" width="180">
                <el-button type="primary" size="small">
                    编辑
                </el-button>
        </el-table-column>
    </el-table>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { ListReq } from '@/api'

const productData = ref([])


// 调用后端
onMounted(() => {
    ListReq().then((res) => {
        let resultDataStr = JSON.stringify(res)
        let resultData = JSON.parse(resultDataStr)
        if (resultData.code !== 0) {
            ElMessage({
                showClose: true,
                message: resultData.msg,
                type: 'error',
        })
        return
        }

        productData.value = resultData.data
        console.log(productData.value)

    })
}) 

</script>