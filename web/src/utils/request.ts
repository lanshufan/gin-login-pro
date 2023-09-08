import axios, { AxiosInstance, AxiosRequestConfig } from 'axios'
import { ErrorMsg } from './elmessage';
import router from '@/router'

class HttpRequest {
    private readonly baseUrl: string;
    constructor() {
        this.baseUrl = 'http://127.0.0.1:8888'
    }
    getInsideConfig() {
        const token = localStorage.getItem('token')
        const config = {
            baseURL: this.baseUrl,// 所有的请求地址前缀部分(没有后端请求不用写)  
            timeout: 50000, // 请求超时时间(毫秒)
            withCredentials: true,// 异步请求携带cookie
            headers: {
            // 设置后端需要的传参类型
            'Content-Type': 'application/json',
            'Access-Control-Allow-Credentials': true,
            'Authorization': token
            // 'token': x-auth-token',//一开始就要token
            // 'X-Requested-With': 'XMLHttpRequest',
            },
        }
        return config
    }

    // 拦截器
    interceptors(instance: AxiosInstance, url: string | number | undefined) {
        // 请求拦截
        instance.interceptors.request.use(config => {
            // 添加全局的loading..
            // 请求头携带token
            return config
        }, (error: any) => {
            return Promise.reject(error)
        })

        //响应拦截
        instance.interceptors.response.use(res => {
            //返回数据
            const { data } = res
            if (data.code === 1) {
                data.msg = "后台错误，请联系管理员"
                return data
            }
            console.log('返回数据处理', res)
            return data
        }, (error: any) => {
            if (error.code === 'ERR_NETWORK') {
                console.log('无法连接后台：', error.code)
                ErrorMsg('后台错误，请联系管理员')
                return Promise.reject(error)
            }
            let errMsg = ""
            switch (error.response.status) {
                case 400:
                    errMsg = "请求错误(400)";
                    break;
                case 401:
                    errMsg = "未授权，请重新登录(401)";
                    // 这里可以做清空storage并跳转到登录页的操作
                    localStorage.removeItem('token')
                    ErrorMsg('登录状态失效，请重新登录')
                    router.push('/login')
                    break;
                case 403:
                    errMsg = "拒绝访问(403)";
                    break;
                case 404:
                    errMsg = "请求出错(404)";
                    break;
                case 408:
                    errMsg = "请求超时(408)";
                    break;
                case 500:
                    errMsg = "服务器错误(500)";
                    break;
                case 501:
                    errMsg = "服务未实现(501)";
                    break;
                case 502:
                    errMsg = "网络错误(502)";
                    break;
                case 503:
                    errMsg = "服务不可用(503)";
                    break;
                case 504:
                    errMsg = "网络超时(504)";
                    break;
                case 505:
                    errMsg = "HTTP版本不受支持(505)";
                    break;
                default:
                    errMsg = `连接出错(${error.response.status})!`;
            }

            return Promise.reject(error)
        })
    }

    request(options: AxiosRequestConfig) {
        const instance = axios.create()
        options = Object.assign(this.getInsideConfig(), options)
        this.interceptors(instance, options.url)
        return instance(options)
    }
}

const http = new HttpRequest()
export default http