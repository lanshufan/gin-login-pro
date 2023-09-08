import { ElMessage } from "element-plus";
import 'element-plus/dist/index.css'

export const SuccessMsg = (data: any) => {
    ElMessage({
        showClose: true,
        message: data,
        type: 'success'
    })
}

export const WarnMsg = (data: any) => {
    ElMessage({
        showClose: true,
        message: data,
        type: 'warning'
    })
}

export const ErrorMsg = (data: any) => {
    ElMessage({
        showClose: true,
        message: data,
        type: 'error'
    })
}