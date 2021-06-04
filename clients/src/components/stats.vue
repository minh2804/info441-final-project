<template>
  <div>
    <div class="el-header">
      <div class="header_left">
        <i class="el-icon-edit" style="color: white"></i>
        <div class="header_left_text">welcome : {{ username }}</div>
      </div>

      <div class="header_right">
        <button class="header_right_text" @click="logout()">logout</button>
        <button class="header_right_text" @click="statss()">stats</button>
        <el-tooltip
          class="item"
          effect="light"
          placement="bottom-end"
          value="true"
          manual="true"
        >
          <div slot="content">{{ shareurl }}</div>
          <button class="header_right_text">share</button>
        </el-tooltip>
      </div>
    </div>

    <h1>STATS</h1>
    <span> TASKS CREATED ALL-TIME:{{ stats.create }}</span><br>
    <span> TASKS COMPLETED ALL-TIME:{{ stats.complete }}</span><br>
    <span> TASKS COMPLETION ALL-TIME:{{ stats.completion }}%</span>
  </div>
</template>

<script>
export default {
  data() {
    return {
      stats: {
        create: "",
        complete: "",
        comletion: "",
      },
      username: "",
      shareurl: "",
    };
  },
  methods: {
    logout() {
      this.$router.push("/");
    },
    statss() {
      this.$router.push("/stats");
    },
  },
  mounted: function () {
    this.username = sessionStorage.getItem("username");
    if (sessionStorage.getItem("shareurl") != null) {
      this.shareurl = sessionStorage.getItem("shareurl");
    } else {
      this.shareurl = sessionStorage.getItem("shareurl");
    }
    this.$http({
      method: "get",
      url: this.$url + "/stats", //这里是发送给后台的数据
      headers: {
        "content-type": "application/json",
        authorization: sessionStorage.getItem("authorization"),
      },
    })
      .then((response) => {
        this.tableData = response.data;
        this.stats.complete = response.data.completed.length;
        this.stats.create = response.data.created.length;
        this.stats.completion = (this.stats.complete / this.stats.create) * 100;
        console.log(this.stats);
      })
      .catch(function (error) {
        console.log(error);
      });
  },
};
</script>
<style >
.el-header {
  background-image: linear-gradient(to top right, #444, #888, #ddd);
  z-index: 9;
  line-height: 60px;
  height: 60px;
}
.header_left {
  line-height: 60px;
  width: 10%;
  height: 100%;
  text-align: left;
  float: left;
  z-index: 99;
  padding: 0 10px;
}
.header_left_text {
  display: inline-block;
  color: white;
}
.header_right {
  line-height: 60px;
  width: 10%;
  height: 100%;
  text-align: right;
  float: right;
  padding: 0 10px;
}
.header_right_text {
  display: inline-block;
  color: black;
}
</style>