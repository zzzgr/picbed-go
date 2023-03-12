import axios from "axios";
import router from "@/router";

//1. 创建axios对象
const service = axios.create();

//2. 请求拦截器
service.interceptors.request.use(
  (config) => {
    config.headers["Authorization"] = "Bearer " + localStorage.getItem("token");
    return config;
  },
  (error) => {
    Promise.reject(error).then((r) => {
      console.error(r);
    });
  }
);

//3. 响应拦截器
service.interceptors.response.use(
  (response) => {
    let res = response.data;
    //判断code码
    let code = res.code;
    if (code != 200) {
      window.$message.error(res.msg);
      Promise.reject(res.msg).then((r) => {
        console.error(r);
      });
    }
    return res;
  },
  (error) => {
    let httpStatus = error.response.status;
    if (httpStatus == 401) {
      router.push("/");
      return;
    }

    let res = error.response.data;
    //判断code码
    let code = res.code;
    if (code != undefined || code != 200) {
      window.$message.error(res.msg);
      Promise.reject(res.msg).then((r) => {
        console.error(r);
      });
    }
    return Promise.reject(error);
  }
);

export default service;
