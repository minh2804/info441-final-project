webpackJsonp([1],{"181v":function(t,e){},"B/WL":function(t,e){},IQXw:function(t,e){},NHnr:function(t,e,s){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=s("7+uW"),a={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"app"}},[e("router-view")],1)},staticRenderFns:[]};var r=s("VU/8")({name:"App"},a,!1,function(t){s("s347")},null,null).exports,i=s("/ocq"),n=s("mvHQ"),l=s.n(n),u=s("bOdI"),c=s.n(u),d={data:function(){var t;return{username:"",tableData:[(t={user:{id:41,username:"username3",firstName:"newFirst",lastName:"newLast"},id:""},c()(t,"user",""),c()(t,"name",""),c()(t,"description",""),c()(t,"isComplete",!1),c()(t,"isHidden",!1),c()(t,"createdAt",""),c()(t,"editedAt",""),t)],shareurl:"",dialogVisible:!1,task:{name:"",description:"",isComplete:!1,isHidden:!1,createdAt:"",editedAt:""}}},methods:{formatter:function(t,e){return 1==t.isComplete&&(t.isComplete="true"),0==t.isComplete&&(t.isComplete="false"),t.isComplete},formatter2:function(t,e){return 1==t.isHidden&&(t.isHidden="true"),0==t.isHidden&&(t.isHidden="false"),t.isHidden},login:function(t){var e=this;this.$refs[t].validate(function(t){if(!t)return console.log("error submit!!"),!1;e.$http({method:"post",url:e.$url+"/authentication/form",params:{username:e.userDto.username,password:e.userDto.password,rememberPsd:e.userDto.rememberPsd}}).then(function(t){sessionStorage.setItem("user",l()(t.data)),sessionStorage.setItem("permissions",t.data.authorities),sessionStorage.setItem("role","admin"),sessionStorage.setItem("login",t.data.authenticated),e.$router.push("/home")}).catch(function(t){console.log(t)})})},toggleSelection:function(t){var e=this;t?(console.log(row),t.forEach(function(t){e.$refs.multipleTable.toggleRowSelection(t)})):this.$refs.multipleTable.clearSelection()},handleSelectionChange:function(t){console.log(t),t.length>1?alert("只能选择一条"):(this.multipleSelectio=t[0],this.task=t[0])},handleClose:function(t){t()},add:function(){var t=this;this.dialogVisible=!0;var e={name:this.task.name,description:this.task.description,isComplete:this.task.isComplete,isHidden:this.task.isHidden};this.$http.post(this.$url+"/tasks",e).then(function(e){console.log("返回数据：",e),t.dialogVisible=!1,t.$http({method:"get",url:t.$url+"/tasks",headers:{"content-type":"application/json",authorization:sessionStorage.getItem("authorization")}}).then(function(e){t.tableData=e.data,console.log(e)}).catch(function(t){console.log(t)})})},del:function(){var t=this;this.$http.delete(this.$url+"/tasks/"+this.task.id,{}).then(function(e){console.log("返回数据：",e),t.$http({method:"get",url:t.$url+"/tasks",headers:{"content-type":"application/json",authorization:sessionStorage.getItem("authorization")}}).then(function(e){t.tableData=e.data,console.log(e)}).catch(function(t){console.log(t)})})},logout:function(){this.$router.push("/")},stats:function(){this.$router.push("/stats")}},mounted:function(){var t=this;this.username=sessionStorage.getItem("username"),null!=sessionStorage.getItem("shareurl")?this.shareurl=sessionStorage.getItem("shareurl"):(sessionStorage.setItem("shareurl","https://api.uwinfotutor.me/tasks/import/7"),this.shareurl=sessionStorage.getItem("shareurl")),this.$http({method:"get",url:this.$url+"/tasks",headers:{"content-type":"application/json",authorization:sessionStorage.getItem("authorization")}}).then(function(e){t.tableData=e.data,console.log(e)}).catch(function(t){console.log(t)})}},m={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("div",{staticClass:"el-header"},[s("div",{staticClass:"header_left"},[s("i",{staticClass:"el-icon-edit",staticStyle:{color:"white"}}),t._v(" "),s("div",{staticClass:"header_left_text"},[t._v("welcome : "+t._s(t.username))])]),t._v(" "),s("div",{staticClass:"header_right"},[s("button",{staticClass:"header_right_text",on:{click:function(e){return t.logout()}}},[t._v("logout")]),t._v(" "),s("button",{staticClass:"header_right_text",on:{click:function(e){return t.stats()}}},[t._v("stats")]),t._v(" "),s("el-tooltip",{staticClass:"item",attrs:{effect:"light",placement:"bottom-end",value:"true",manual:"true"}},[s("div",{attrs:{slot:"content"},slot:"content"},[t._v(t._s(t.shareurl))]),t._v(" "),s("button",{staticClass:"header_right_text"},[t._v("share")])])],1)]),t._v(" "),s("el-table",{ref:"multipleTable",staticStyle:{width:"100%"},attrs:{data:t.tableData},on:{"selection-change":t.handleSelectionChange}},[t.tableData.isHidden?s("el-table-column",{attrs:{prop:"id",label:"id",width:"180"}}):t._e(),t._v(" "),s("el-table-column",{attrs:{prop:"user",type:"selection",label:"user",width:"180"}}),t._v(" "),s("el-table-column",{attrs:{prop:"name",label:"name"}}),t._v(" "),s("el-table-column",{attrs:{prop:"description",label:"description",width:"180"}}),t._v(" "),s("el-table-column",{attrs:{prop:"isComplete",property:"isComplete",formatter:t.formatter,label:"isComplete",width:"180"}}),t._v(" "),s("el-table-column",{attrs:{prop:"isHidden",label:"isHidden",formatter:t.formatter2}}),t._v(" "),s("el-table-column",{attrs:{prop:"createdAt",label:"createdAt",width:"180"}}),t._v(" "),s("el-table-column",{attrs:{prop:"editedAt",label:"editedAt",width:"180"}})],1),t._v(" "),s("el-button",{staticStyle:{"margin-top":"20px"},attrs:{type:"primary"},on:{click:function(e){t.dialogVisible=!0}}},[t._v("add")]),t._v(" "),s("el-button",{staticStyle:{"margin-top":"20px"},attrs:{type:"primary"},on:{click:function(e){return t.del()}}},[t._v("delete")]),t._v(" "),s("el-dialog",{attrs:{title:"create a task",visible:t.dialogVisible,width:"30%","before-close":t.handleClose},on:{"update:visible":function(e){t.dialogVisible=e}}},[s("el-form",{ref:"task",staticClass:"demo-ruleForm",attrs:{model:t.task,"label-width":"100px"}},[s("el-form-item",{attrs:{label:"name",prop:"name"}},[s("el-input",{model:{value:t.task.name,callback:function(e){t.$set(t.task,"name",e)},expression:"task.name"}})],1),t._v(" "),s("el-form-item",{attrs:{label:"description",prop:"description"}},[s("el-input",{model:{value:t.task.description,callback:function(e){t.$set(t.task,"description",e)},expression:"task.description"}})],1),t._v(" "),s("el-form-item",{attrs:{label:"isComplete",prop:"isComplete"}},[s("el-input",{model:{value:t.task.isComplete,callback:function(e){t.$set(t.task,"isComplete",e)},expression:"task.isComplete"}})],1),t._v(" "),s("el-form-item",{attrs:{label:"isHidden",prop:"isHidden"}},[s("el-input",{model:{value:t.task.isHidden,callback:function(e){t.$set(t.task,"isHidden",e)},expression:"task.isHidden"}})],1)],1),t._v(" "),s("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[s("el-button",{on:{click:function(e){t.dialogVisible=!1}}},[t._v("cancel")]),t._v(" "),s("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.add()}}},[t._v("confirm")])],1)],1)],1)},staticRenderFns:[]};var p=s("VU/8")(d,m,!1,function(t){s("IQXw")},null,null).exports,h={data:function(){return{radio:"1",role:"",roles:""}},mounted:function(){this.role=sessionStorage.getItem("role")}},f={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"box_one"},[e("el-container",[e("el-header",{staticStyle:{"background-image":"linear-gradient(to top right, #444, #888, #ddd)"}},[e("div",{staticClass:"header_left"},[e("i",{staticClass:"el-icon-edit",staticStyle:{color:"white"}}),this._v(" "),e("div",{staticClass:"header_left_text"},[this._v("welcome   :  张三")])])]),this._v(" "),e("el-container",[e("el-container",[e("el-aside",[e("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"background-color":"white","text-color":"black","active-text-color":"#409EFF","unique-opened":"",router:""}},[e("el-menu-item",{attrs:{index:"main"}},[e("i",{staticClass:"el-icon-s-home"}),this._v(" "),e("span",{attrs:{slot:"title"},slot:"title"},[this._v("首页")])])],1)],1),this._v(" "),e("el-main",[e("router-view")],1)],1)],1)],1)],1)},staticRenderFns:[]};s("VU/8")(h,f,!1,function(t){s("rjNy")},null,null).exports;var g={data:function(){return{radio:"1",userDto:{username:"",password:"",rememberPsd:""},rules:{username:[{required:!0,message:"require email",trigger:"blur"}],password:[{required:!0,message:"require email",trigger:"blur"}],per:["test","main"]}}},mounted:function(){},methods:{login:function(t){var e=this;this.$refs[t].validate(function(t){if(!t)return console.log("error submit!!"),!1;var s={username:e.userDto.username,password:e.userDto.password};e.$http.post(e.$url+"/sessions",s).then(function(t){console.log("返回数据：",t),sessionStorage.setItem("id",t.data.id),sessionStorage.setItem("shareurl","https://api.thenightbeforeitsdue.de/tasks/import/"+t.data.id),sessionStorage.setItem("username",t.data.username),sessionStorage.setItem("firstName",t.data.firstName),sessionStorage.setItem("lastName",t.data.lastName),console.log(t.headers.authorization),sessionStorage.setItem("authorization",t.headers.authorization),console.log(sessionStorage.getItem("authorization")),console.log(t),e.$router.push("/home")})})},getCookie:function(t){var e=t+"=",s=document.cookie.split(";");console.log("获取cookie,现在循环"),console.log(s);for(var o=0;o<s.length;o++){var a=s[o];for(console.log(a);" "==a.charAt(0);)a=a.substring(1);if(-1!=a.indexOf(e))return a.substring(e.length,a.length)}return""},toRegister:function(){this.$router.push("/register")}}},v={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"box_one"},[s("div",{staticClass:"login_box"},[s("span",{staticClass:"login_title"},[t._v("login")]),t._v(" "),s("div",{staticClass:"divider_my"}),t._v(" "),s("el-form",{ref:"loginForm",attrs:{model:t.userDto,rules:t.rules}},[s("el-form-item",{attrs:{prop:"username"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"text",placeholder:"username"},model:{value:t.userDto.username,callback:function(e){t.$set(t.userDto,"username",e)},expression:"userDto.username"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"password"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"password",placeholder:"password"},model:{value:t.userDto.password,callback:function(e){t.$set(t.userDto,"password",e)},expression:"userDto.password"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"rememberPsd"}},[s("div",{staticClass:"line-box",staticStyle:{"margin-top":"10px"}},[s("span",{staticClass:"forget_password"},[t._v("forget password?")])])]),t._v(" "),s("el-form-item",[s("el-button",{staticClass:"commit-button",attrs:{type:"primary"},on:{click:function(e){return t.login("loginForm")}}},[t._v("login")])],1),t._v(" "),s("div",{staticClass:"line-box"},[s("span",{staticClass:"no_account"}),s("span",{staticClass:"register",on:{click:function(e){return t.toRegister()}}},[t._v("register")])])],1)],1)])},staticRenderFns:[]};var _=s("VU/8")(g,v,!1,function(t){s("181v")},"data-v-8b968360",null).exports,b={data:function(){return{radio:"1",userDto:{email:"",password:"",passwordConf:"",username:"",firstName:"",lastName:""},rules:{username:[{required:!0,message:"require username",trigger:"blur"}],password:[{required:!0,message:"require password",trigger:"blur"}],per:["test","main"]}}},mounted:function(){},methods:{login:function(t){var e=this;this.$refs[t].validate(function(t){if(!t)return console.log("error submit!!"),!1;var s={email:e.userDto.email,password:e.userDto.password,passwordConf:e.userDto.passwordConf,username:e.userDto.username,firstName:e.userDto.firstName,lastName:e.userDto.lastName};e.$http.post(e.$url+"/users",s).then(function(t){console.log("返回数据：",t),e.$router.push("/home")})})},getCookie:function(t){var e=t+"=",s=document.cookie.split(";");console.log("获取cookie,现在循环"),console.log(s);for(var o=0;o<s.length;o++){var a=s[o];for(console.log(a);" "==a.charAt(0);)a=a.substring(1);if(-1!=a.indexOf(e))return a.substring(e.length,a.length)}return""},tologin:function(){this.$router.push("/")}}},C={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"box_one"},[s("div",{staticClass:"login_box"},[s("span",{staticClass:"login_title"},[t._v("register")]),t._v(" "),s("div",{staticClass:"divider_my"}),t._v(" "),s("el-form",{ref:"loginForm",attrs:{model:t.userDto,rules:t.rules}},[s("el-form-item",{attrs:{prop:"email"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"text",placeholder:"email"},model:{value:t.userDto.email,callback:function(e){t.$set(t.userDto,"email",e)},expression:"userDto.email"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"username"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"text",placeholder:"username"},model:{value:t.userDto.username,callback:function(e){t.$set(t.userDto,"username",e)},expression:"userDto.username"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"fisrtName"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"text",placeholder:"firstName"},model:{value:t.userDto.firstName,callback:function(e){t.$set(t.userDto,"firstName",e)},expression:"userDto.firstName"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"LastName"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"text",placeholder:"LastName"},model:{value:t.userDto.lastName,callback:function(e){t.$set(t.userDto,"lastName",e)},expression:"userDto.lastName"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"password"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"password",placeholder:"password"},model:{value:t.userDto.password,callback:function(e){t.$set(t.userDto,"password",e)},expression:"userDto.password"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"passwordConf"}},[s("el-input",{staticClass:"form_input",attrs:{autocomplete:"on",type:"password",placeholder:"password Confirmation"},model:{value:t.userDto.passwordConf,callback:function(e){t.$set(t.userDto,"passwordConf",e)},expression:"userDto.passwordConf"}})],1),t._v(" "),s("el-form-item",{attrs:{prop:"rememberPsd"}},[s("div",{staticClass:"line-box",staticStyle:{"margin-top":"10px"}},[s("span",{staticClass:"forget_password"})])]),t._v(" "),s("el-form-item",[s("el-button",{staticClass:"commit-button",attrs:{type:"primary"},on:{click:function(e){return t.login("loginForm")}}},[t._v("register")])],1),t._v(" "),s("div",{staticClass:"line-box"},[s("span",{staticClass:"no_account"}),s("span",{staticClass:"login",on:{click:function(e){return t.tologin()}}},[t._v("login")])])],1)],1)])},staticRenderFns:[]};var w=s("VU/8")(b,C,!1,function(t){s("B/WL")},"data-v-13ba0cae",null).exports,k={data:function(){return{stats:{create:"",complete:"",comletion:""},username:"",shareurl:""}},methods:{logout:function(){this.$router.push("/")},statss:function(){this.$router.push("/stats")}},mounted:function(){var t=this;this.username=sessionStorage.getItem("username"),sessionStorage.getItem("shareurl"),this.shareurl=sessionStorage.getItem("shareurl"),this.$http({method:"get",url:this.$url+"/stats",headers:{"content-type":"application/json",authorization:sessionStorage.getItem("authorization")}}).then(function(e){t.tableData=e.data,t.stats.complete=e.data.completed.length,t.stats.create=e.data.created.length,t.stats.completion=t.stats.complete/t.stats.create*100,console.log(t.stats)}).catch(function(t){console.log(t)})}},$={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("div",{staticClass:"el-header"},[s("div",{staticClass:"header_left"},[s("i",{staticClass:"el-icon-edit",staticStyle:{color:"white"}}),t._v(" "),s("div",{staticClass:"header_left_text"},[t._v("welcome : "+t._s(t.username))])]),t._v(" "),s("div",{staticClass:"header_right"},[s("button",{staticClass:"header_right_text",on:{click:function(e){return t.logout()}}},[t._v("logout")]),t._v(" "),s("button",{staticClass:"header_right_text",on:{click:function(e){return t.statss()}}},[t._v("stats")]),t._v(" "),s("el-tooltip",{staticClass:"item",attrs:{effect:"light",placement:"bottom-end",value:"true",manual:"true"}},[s("div",{attrs:{slot:"content"},slot:"content"},[t._v(t._s(t.shareurl))]),t._v(" "),s("button",{staticClass:"header_right_text"},[t._v("share")])])],1)]),t._v(" "),s("h1",[t._v("STATS")]),t._v(" "),s("span",[t._v(" TASKS CREATED ALL-TIME:"+t._s(t.stats.create))]),s("br"),t._v(" "),s("span",[t._v(" TASKS COMPLETED ALL-TIME:"+t._s(t.stats.complete))]),s("br"),t._v(" "),s("span",[t._v(" TASKS COMPLETION ALL-TIME:"+t._s(t.stats.completion)+"%")])])},staticRenderFns:[]};var x=s("VU/8")(k,$,!1,function(t){s("Un0H")},null,null).exports;o.default.use(i.a);var S=new i.a({routes:[{path:"/home",name:"HelloWorld",component:p},{path:"/",name:"login",component:_},{path:"/register",name:"register",component:w},{path:"/stats",name:"stats",component:x}]}),D=s("mtWM"),y=s.n(D),I=s("zL8q"),N=s.n(I);s("tvR6");o.default.use(N.a),o.default.config.productionTip=!1,y.a.defaults.headers.post["Content-Type"]="application/json",y.a.interceptors.request.use(function(t){return null!=window.sessionStorage.getItem("authorization")&&(t.headers.contentType="application/json",t.headers.authorization=window.sessionStorage.getItem("authorization")),t}),o.default.prototype.$http=y.a,o.default.prototype.$http=y.a,o.default.prototype.$url="https://api.thenightbeforeitsdue.de",new o.default({el:"#app",router:S,components:{App:r},template:"<App/>"})},Un0H:function(t,e){},rjNy:function(t,e){},s347:function(t,e){},tvR6:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.f8a6566ea99ef2768e7c.js.map