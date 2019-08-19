<template>
  <div class="app-container">

    <div class="filter-container">
      <el-input v-model="listQuery.search" placeholder="用户名/邮件" style="width: 200px;" class="filter-item" />
      <el-button class="filter-item" type="primary" icon="el-icon-search">搜索</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="create">添加</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="公司名称" align="center">
        <template slot-scope="scope">
          {{ scope.row.company }}
        </template>
      </el-table-column>
      <el-table-column label="邮件" align="center">
        <template slot-scope="scope">
          {{ scope.row.email }}
        </template>
      </el-table-column>
      <el-table-column label="电话" align="center">
        <template slot-scope="scope">
          {{ scope.row.mobile }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ statusConf[scope.row.status] }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="edit(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="listQuery.total>0" :total="listQuery.total" :page.sync="listQuery.page" :limit.sync="listQuery.pageSize" @pagination="getList" />
    <el-dialog :title="buttonNames[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="editForm" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="editForm.name" type="text" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="公司名称" prop="company">
          <el-input v-model="editForm.company" type="text" placeholder="请输入公司名称" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="editForm.email" type="text" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="电话" prop="mobile">
          <el-input v-model="editForm.mobile" type="text" placeholder="请输入电话" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="editForm.password" type="password" placeholder="请输入密码" />
        </el-form-item>
        <el-form-item label="确认密码" prop="repassword">
          <el-input v-model="editForm.repassword" type="password" placeholder="请再次输入密码" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="editForm.status" placeholder="请选择">
            <el-option
              v-for="(value, key) in statusConf"
              :key="key"
              :label="value"
              :value="key"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogStatus==='add'?store():update()">{{ buttonNames[dialogStatus] }}</el-button>
        <el-button @click="dialogFormVisible = false">{{ buttonNames['cancel'] }}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { page, create, update } from '@/api/user'
import { statusConf, buttonNames } from '@/config/common'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import { validatPassword } from '@/utils/validate'

export default {
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: 'success',
        2: 'gray',
        3: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    var vPassword = (rule, value, callback) => {
      if (this.dialogStatus === "add" && !validatPassword(value)) {
        callback(new Error('请输入数字或字母并且在6~20个字符的密码'))
      } else {
        if (this.dialogStatus === "add" && this.editForm.repassword !== '') {
          this.$refs.dataForm.validateField('repassword')
        }
        callback()
      }
    }
    var vRePassword = (rule, value, callback) => {
      if (this.dialogStatus === "add" && !validatPassword(value)) {
        callback(new Error('请再次输入密码'))
      } else if (this.dialogStatus === "add" && value !== this.editForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      statusConf: statusConf,

      // 编辑框
      dialogFormVisible: false,
      dialogStatus: '',
      buttonNames: buttonNames,
      editForm: {
        name: '',
        email: '',
        company: '',
        mobile: '',
        password: '',
        repassword: '',
        status: ''
      },

      rules: {
        name: [{ required: true, message: '请填写用户名', trigger: 'blur' }],
        email: [{ required: true, message: '请填写邮箱', trigger: 'blur' }],
        company: [{ required: true, message: '请填写公司名称', trigger: 'blur' }],
        mobile: [{ required: true, message: '请填写手机号', trigger: 'blur' }],
        password: [{ required: true, validator: vPassword, trigger: 'blur' }],
        repassword: [{ required: true, validator: vRePassword, trigger: 'blur' }],
        status: [{ required: true, message: '请选择状态', trigger: 'blur' }]
      },

      // 列表
      list: null,
      listQuery: {
        search: '',
        username: '',
        page: 1,
        pageSize: 10,
        total: 0
      },
      listLoading: true
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      page(this.listQuery).then(response => {
        this.list = response.data.users
        this.listQuery.total = response.data.count
        this.listLoading = false
      })
    },
    create() {
      this.dialogStatus = 'add'
      this.resetEditForm()
      this.showDialog()
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    store() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          create(this.formatForm(this.editForm)).then((res) => {
            this.list.unshift(res.data.container)
            this.hideDialog()
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          }).catch(error => {
            console.log(error)
          })
        }
      })
    },
    edit(row) {
      this.editForm = Object.assign({}, row)
      this.editForm.password = ''
      this.editForm.status = '' + this.editForm.status
      this.dialogStatus = 'edit'
      this.showDialog()
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    update() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          update(this.formatForm(this.editForm)).then((res) => {
            for (const v of this.list) {
              if (v.id === this.editForm.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.editForm)
                break
              }
            }
            this.hideDialog()
            this.$notify({
              title: '成功',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
          }).catch(error => {
            console.log(error)
          })
        }
      })
    },
    search() {

    },
    formatForm(form) {
      form.status = parseInt(form.status)
      return form
    },
    resetEditForm() {
      this.editForm = {
        name: '',
        email: '',
        company: '',
        mobile: '',
        password: '',
        repassword: '',
        status: ''
      }
    },
    showDialog() {
      this.dialogFormVisible = true
    },
    hideDialog() {
      this.dialogFormVisible = false
    }
  }
}
</script>
