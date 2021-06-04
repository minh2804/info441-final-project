<template>
  <div>
    <div class="el-header">
      <div class="header_left">
        <i class="el-icon-edit" style="color: white"></i>
        <div class="header_left_text">welcome : {{ username }}</div>
      </div>

      <div class="header_right">
        <button class="header_right_text" @click="logout()">logout</button>
        <button class="header_right_text" @click="stats()">stats</button>
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
    <el-table
      :data="tableData"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      ref="multipleTable"
    >
      <el-table-column
        prop="id"
        label="id"
        width="180"
        v-if="tableData.isHidden"
      >
      </el-table-column>
      <el-table-column prop="user" type="selection" label="user" width="180">
      </el-table-column>
      <el-table-column prop="name" label="name"> </el-table-column>
      <el-table-column prop="description" label="description" width="180">
      </el-table-column>
      <el-table-column
        prop="isComplete"
        property="isComplete"
        :formatter="formatter"
        label="isComplete"
        width="180"
      >
      </el-table-column>
      <el-table-column prop="isHidden" label="isHidden" :formatter="formatter2">
      </el-table-column>
      <el-table-column prop="createdAt" label="createdAt" width="180">
      </el-table-column>
      <el-table-column prop="editedAt" label="editedAt" width="180">
      </el-table-column>
    </el-table>

    <el-button
      type="primary"
      @click="dialogVisible = true"
      style="margin-top: 20px"
      >add</el-button
    >

    <el-button type="primary" @click="del()" style="margin-top: 20px"
      >delete</el-button
    >

    <el-dialog
      title="create a task"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose"
    >
      <el-form
        :model="task"
        ref="task"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="name" prop="name">
          <el-input v-model="task.name"></el-input>
        </el-form-item>
        <el-form-item label="description" prop="description">
          <el-input v-model="task.description"></el-input>
        </el-form-item>
        <el-form-item label="isComplete" prop="isComplete">
          <el-input v-model="task.isComplete"></el-input>
        </el-form-item>
        <el-form-item label="isHidden" prop="isHidden">
          <el-input v-model="task.isHidden"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">cancel</el-button>
        <el-button type="primary" @click="add()">confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      tableData: [
        {
          user: {
            id: 41,
            username: "username3",
            firstName: "newFirst",
            lastName: "newLast",
          },
          id: "",
          user: "",
          name: "",
          description: "",
          isComplete: false,
          isHidden: false,
          createdAt: "",
          editedAt: "",
        },
      ],
      shareurl: "",
      dialogVisible: false,
      task: {
        name: "",
        description: "",
        isComplete: false,
        isHidden: false,
        createdAt: "",
        editedAt: "",
      },
    };
  },
  methods: {
    formatter(row, index) {
      if (row.isComplete == true) {
        row.isComplete = "true";
      }
      if (row.isComplete == false) {
        row.isComplete = "false";
      }
      return row.isComplete;
    },
    formatter2(row, index) {
      if (row.isHidden == true) {
        row.isHidden = "true";
      }
      if (row.isHidden == false) {
        row.isHidden = "false";
      }
      return row.isHidden;
    },
    login(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$http({
            method: "post",
            url: this.$url + "/authentication/form", //这里是发送给后台的数据
            params: {
              username: this.userDto.username,
              password: this.userDto.password,
              rememberPsd: this.userDto.rememberPsd,
            },
          })
            .then((response) => {
              sessionStorage.setItem("user", JSON.stringify(response.data));
              sessionStorage.setItem("permissions", response.data.authorities);
              sessionStorage.setItem("role", "admin");
              sessionStorage.setItem("login", response.data.authenticated);

              this.$router.push("/home");
            })
            .catch(function (error) {
              console.log(error);
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    toggleSelection(rows) {
      if (rows) {
        console.log(row);
        rows.forEach((row) => {
          this.$refs.multipleTable.toggleRowSelection(row);
        });
      } else {
        this.$refs.multipleTable.clearSelection();
      }
    },
    handleSelectionChange(val) {
      console.log(val);
      if (val.length > 1) {
        alert("只能选择一条");
      } else {
        this.multipleSelectio = val[0];
        this.task = val[0];
      }
    },

    handleClose(done) {

          done();

    },
    add() {
      this.dialogVisible=true;
      let data = {
        name: this.task.name,
        description: this.task.description,
        isComplete: this.task.isComplete,
        isHidden: this.task.isHidden,
      };

      this.$http.post(this.$url + "/tasks", data).then((response) => {
        console.log("返回数据：", response);
        this.dialogVisible = false;
        this.$http({
          method: "get",
          url: this.$url + "/tasks", //这里是发送给后台的数据
          headers: {
            "Content-Type": "application/json",
            "Authorization": sessionStorage.getItem("authorization"),
          },
        })
          .then((response) => {
            this.tableData = response.data;
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
      });

      // this.$http({
      //   method: "post",
      //   url: this.$url + "/tasks", //这里是发送给后台的数据
      //   params: {
      //     name: this.task.name,
      //     description: this.task.description,
      //     isComplete: this.task.isComplete,
      //     isHidden: this.task.isHidden,
      //   },
      // })
      //   .then((response) => {
      //     this.$router.push("/home");
      //   })
      //   .catch(function (error) {
      //     console.log(error);
      //   });

      // alert(this.task.name);
    },
    del() {
      let data = {};

      this.$http
        .delete(this.$url + "/tasks/" + this.task.id, data)
        .then((response) => {
          console.log("返回数据：", response);
          this.$http({
            method: "get",
            url: this.$url + "/tasks", //这里是发送给后台的数据
            headers: {
              "Content-Type": "application/json",
              "Authorization": sessionStorage.getItem("authorization"),
            },
          })
            .then((response) => {
              this.tableData = response.data;
              console.log(response);
            })
            .catch(function (error) {
              console.log(error);
            });
        });
      // this.$http({
      //   method: "delete",
      //   url: this.$url + "/tasks", //这里是发送给后台的数据
      //   params: {
      //     id: this.task.id
      //   },
      // })
      //   .then((response) => {
      //     this.$router.push("/");
      //   })
      //   .catch(function (error) {
      //     console.log(error);
      //   });
    },
    logout() {
      this.$router.push("/");
    },
    stats() {
      this.$router.push("/stats");
    },
  },
  mounted: function () {
    this.username = sessionStorage.getItem("username");
    if (sessionStorage.getItem("shareurl") != null) {
      this.shareurl = sessionStorage.getItem("shareurl");
    } else {
      sessionStorage.setItem(
        "shareurl",
        "https://api.uwinfotutor.me/tasks/import/7"
      );
      this.shareurl = sessionStorage.getItem("shareurl");
    }
    this.$http({
      method: "get",
      url: this.$url + "/tasks", //这里是发送给后台的数据
      headers: {
        "Content-Type": "application/json",
        "Authorization": sessionStorage.getItem("authorization"),
      },
    })
      .then((response) => {
        this.tableData = response.data;
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });

    // let data = {
    //   id: sessionStorage.getItem("id"),
    // };
    // this.$http.post(this.$url + "/tasks", data).then((response) => {
    //   console.log("返回数据：", response);
    // });
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