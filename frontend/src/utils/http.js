import axios from "axios";
const http=axios.create({
  timeout:50000
})

http.interceptors.response.use(res=>{
  return res.data;
})
http.interceptors.request.use(req=>{
  return req;
})

export let get=(url,params)=>{
return http.get(url,{params});
}
export let post=(url,data)=>{
return http.post(url,data)
}