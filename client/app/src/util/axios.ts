import axios from 'axios'

const http = axios.create({
  baseURL: `${process.env.VUE_APP_SERVER_URL}`,
  headers: {
    'Content-Type': 'application/json'
    //"X-Requested-With": "XMLHttpRequest",
  },
  responseType: 'json'
});

http.interceptors.request.use(
  config => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
    return config
  },
  function(error) {
    return Promise.reject(error)
  }
);
export default http