<template>
  <div class="box_one">
    <el-container>
      <el-header
        style="
          background-image: linear-gradient(to top right, #444, #888, #ddd);
        "
      >
        <div class="header_left">
          <i class="el-icon-edit" style="color: white"></i>
          <div class="header_left_text">welcome   :  张三</div>
        </div>
      </el-header>
      <el-container>
        <el-container>
          <el-aside>
            <el-menu
              background-color="white"
              text-color="black"
              active-text-color="#409EFF"
              unique-opened=""
              router
              class="el-menu-vertical-demo"
            >
              <el-menu-item index="main">
                <i class="el-icon-s-home"></i>
                <span slot="title">首页</span>
              </el-menu-item>
            </el-menu>
          </el-aside>
          <el-main> <router-view></router-view></el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>
<script>
export default {
  data() {
    return {
      radio: "1",
      role:'',
      roles:'',
    };
  },
  mounted:function(){
    this.role=sessionStorage.getItem('role');
  }
};
</script>
<style >
.el-divider {
  width: 100%;
  margin: auto;
}
.el-divider_box {
  margin: 0 20px;
}
.box_one {
  height: 100%;
}
.el-header {
  background-color: #5f6974;
  color: #333;
  text-align: center;
  height: calc(10vh);
}
.el-header {
  padding: 0px;
  box-shadow: 0px 5px 25px 0px #ddd;
}
.el-aside {
  background-color: #ffffff;
  color: #333;
  text-align: center;
  min-height: calc(100vh - 60px);
  width: 200px !important;
  box-shadow: 0px 5px 25px 0px #ddd;
}
.el-main {
  color: #333;
  text-align: center;
  width: 90% !important;
  padding: 0 !important;
}
body > .el-container {
  margin-bottom: 40px;
}
.header_left {
  line-height: 60px;
  width: 10%;
  height: 100%;
  text-align: left;
  padding: 0 10px;
}

.header_left_text {
  display: inline-block;
  color: white;
}
.menu {
  width: 100%;
  height: 100%;
}
.menu_item {
  font-size: 13px;
  height: 50px;
  line-height: 50px;
  color: black;
}
.el-menu-vertical-demo {
  text-align: left;
  min-height: 100%;
}
.box {
  width: 98%;
  margin: auto;
  margin-top: 15px;
  min-height: calc(100vh - 90px);
  box-shadow: 0px 5px 25px 0px #ddd;
}
</style>
