<template>
  <div class="box_one">
    <div class="login_box">
      <span class="login_title">register</span>
      <div class="divider_my"></div>
      <el-form ref="loginForm" :model="userDto" :rules="rules">
        <el-form-item prop="email">
          <el-input
            v-model="userDto.email"
            autocomplete="on"
            type="text"
            placeholder="email"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="username">
          <el-input
            v-model="userDto.username"
            autocomplete="on"
            type="text"
            placeholder="username"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="fisrtName">
          <el-input
            v-model="userDto.firstName"
            autocomplete="on"
            type="text"
            placeholder="firstName"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="LastName">
          <el-input
            v-model="userDto.lastName"
            autocomplete="on"
            type="text"
            placeholder="LastName"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="userDto.password"
            autocomplete="on"
            type="password"
            placeholder="password"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="passwordConf">
          <el-input
            v-model="userDto.passwordConf"
            autocomplete="on"
            type="password"
            placeholder="password Confirmation"
            class="form_input"
          ></el-input>
        </el-form-item>
        <el-form-item prop="rememberPsd">
          <div class="line-box" style="margin-top: 10px">
            <span class="forget_password"></span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            class="commit-button"
            @click="login('loginForm')"
            >register</el-button
          >
        </el-form-item>
                <div class="line-box">
          <span class="no_account"></span
          ><span class="login" @click="tologin()">login</span>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      radio: "1",
      userDto: {
              email: "",
              password: "",
              passwordConf: "",
              username: "",
              firstName: "",
              lastName: "",
      },
      rules: {
        username: [
          {
            required: true,
            message: "require username",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: "require password",
            trigger: "blur",
          },
        ],
        per: ["test", "main"],
      },
    };
  },
  mounted: function () {},
  methods: {
    login(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {

           let data = {
              email: this.userDto.email,
              password: this.userDto.password,
              passwordConf: this.userDto.passwordConf,
              username: this.userDto.username,
              firstName: this.userDto.firstName,
              lastName: this.userDto.lastName,
          };
          this.$http.post(this.$url + "/users", data).then((response) => {
            console.log("返回数据：", response);
              this.$router.push("/home");
          });




          // this.$http({
          //   method: "post",
          //   url: this.$url + "/users", //这里是发送给后台的数据
          //   params: {
          //     email: this.userDto.email,
          //     password: this.userDto.password,
          //     passwordConf: this.userDto.passwordConf,
          //     username: this.userDto.username,
          //     firstName: this.userDto.firstName,
          //     lastName: this.userDto.lastName,
          //   },
          // })
          //   .then((response) => {
          //     sessionStorage.setItem("id", response.data.id);
          //     sessionStorage.setItem(
          //       "shareurl",
          //       "https://api.uwinfotutor.me/tasks/import/" + response.data.id
          //     );

          //     sessionStorage.setItem("username", response.data.username);
          //     sessionStorage.setItem("firstName", response.data.firstName);
          //     sessionStorage.setItem("lastName", response.data.lastName);
          //     this.$router.push("/");
          //   })
          //   .catch(function (error) {
          //     console.log(error);
          //   });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    getCookie(cname) {
      var name = cname + "=";
      var ca = document.cookie.split(";");
      console.log("获取cookie,现在循环");
      console.log(ca);
      for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        console.log(c);
        while (c.charAt(0) == " ") c = c.substring(1);
        if (c.indexOf(name) != -1) {
          return c.substring(name.length, c.length);
        }
      }
      return "";
    },
        tologin(){
                    this.$router.push("/");
    }
  },
};
</script>
<style scoped>
.box_one {
  height: 100%;
  background-image: url(../assets/pic/logo_bg.jpg);
}
.login_box {
  width: 360px;
  height: 480px;
  box-shadow: 0px 5px 25px 0px rgb(202, 120, 26);
  background-color: rgba(255, 255, 255, 0.959);
  margin: auto;
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  padding: 20px;
}
.login_title {
  line-height: 60px;
  font-size: 26px;
  color: rgb(233, 141, 21);
  font-weight: bold;
}
.divider_my {
  height: 2px;
  border-bottom: 2px solid rgba(0, 0, 0, 0.26);
}
.form_input {
  margin: auto;
  margin-top: 20px;
}
.el-form-item {
  margin-bottom: 0px;
}
.commit-button {
  width: 100%;
}
.line-box {
  width: 100%;
  line-height: 30px;
}
.remenberme {
  float: left;
}
.forget_password {
  float: right;
  color: rgb(0, 174, 255);
}
.forget_password:hover {
  cursor: pointer;
}
.register {
  color: rgb(0, 174, 255);
  float: left;
}
.register:hover {
  cursor: pointer;
}
.no_account {
  float: left;
}
</style>
