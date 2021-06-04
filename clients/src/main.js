// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);
Vue.config.productionTip = false
axios.defaults.headers.post['Content-Type'] = 'application/json';
axios.interceptors.request.use(config => {
	if (window.sessionStorage.getItem("authorization") != null) {
		config.headers['Content-Type'] = 'application/json';
		config.headers['Authorization'] = window.sessionStorage.getItem("authorization")
	}
	return config;
})

Vue.prototype.$http = axios;

//全局引用axios
Vue.prototype.$http = axios;
Vue.prototype.$url = "https://api.thenightbeforeitsdue.de";

/* eslint-disable no-new */
new Vue({
	el: '#app',
	router,
	components: { App },
	template: '<App/>'
})
