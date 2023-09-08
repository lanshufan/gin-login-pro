<template>
    <!-- <ul>
        <li v-for="(items, index) in newData" :key="index">
            <router-link :to="'/blog/'+items['id']">{{ items["name"] }}</router-link>
        </li>
    </ul> -->
    <el-table :data="tableData" border style="width: 810px;">
        <el-table-column prop="name" label="名称" width="250" />
        <el-table-column prop="author" label="作者" width="180" />
        <el-table-column prop="date" label="日期" width="200" />
        <el-table-column prop="sale" label="销量" width="180" />
    </el-table>
    <!-- <ul>
        <li v-for="(item, index) in Lists" :key="index">
            <router-link :to="'/blog/'+index">{{ item }}</router-link>
        </li>
    </ul> -->
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive } from 'vue';
import { FormInstance, FormRules, ElMessage } from 'element-plus'

const tableData = [
    {
        name: '精通Java',
        date: '2023-01-01',
        author: 'xxk',
        sale: 100
    },
    {
        name: '深入浅出Docker',
        date: '2023-02-01',
        author: 'xxk',
        sale: 200
    },
    {
        name: 'Python从入门到放弃',
        date: '2023-03-01',
        author: 'xxk',
        sale: 350
    },
    {
        name: 'Golang入门',
        date: '2023-02-01',
        author: 'xxk',
        sale: 250
    },
]

// 本地测试
const Lists = reactive({
    1: "精通Java",
    2: "Python入门到放弃",
    3: "深入浅出Docker"
})

// 调用后端
// let newData = ref([])
// onMounted(() => {
//     ListReq().then((res) => {
//         const { data } = res
//         for (let i=0; i<data.length; i++) {
//             // console.log(data[i])

//             newData.value = data
//         }

//         console.log(newData)

//     })
// }) 

</script>