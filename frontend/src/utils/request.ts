import axios from 'axios';
import { store } from '../store';
import { ICustomHeaders } from '../interfaces/headers';

// create an axios instance
const service = axios.create({
    baseURL: process.env.REACT_APP_BASE_API, // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
    config => {
        // do something before request is sent
        (config.headers as ICustomHeaders)['X-Requested-With'] = 'XMLHttpRequest'
        const { userStore } = store;
        if (userStore.token) {
            (config.headers as ICustomHeaders).Authorization = 'Bearer ' + userStore.token
        }
        return config
    },
    error => {
        // do something with request error
        if (process.env.NODE_ENV === 'development') {
            console.log(error) // for debug
        }
        return Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    /**
     * If you want to get http information such as headers or status
     * Please return  response => response
    */
    response => response,
    async error => {
        if (process.env.NODE_ENV === 'development') {
            console.log('err' + error) // for debug
        }

        if (error.response && error.response.status === 401) {
            const { userStore } = store;
            userStore.logout();
            window.location.reload();
        }

        return Promise.reject(error)
    }
)

export default service
