import axios from "axios";
const http=axios.create({
  timeout:50000
})

http.interceptors.response.use(res=>{
  return res;
})
http.interceptors.request.use(req=>{
  return req;
})

export let get=({url,param})=>{
return http.get(url,param);
}
export let post=({url,data})=>{
return http.post(url,data)
}