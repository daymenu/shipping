<template>
  <div class="app-container">

    <div class="filter-container">
      <el-input v-model="listQuery.costumer_id" placeholder="客户编号" style="width: 200px;" class="filter-item" />
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
      <el-table-column label="客户编号" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.customer_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="出发地" align="center">
        <template slot-scope="scope">
          {{ scope.row.origin }}
        </template>
      </el-table-column>
      <el-table-column label="用户编号" align="center">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
      <el-table-column label="长" align="center">
        <template slot-scope="scope">
          {{ scope.row.long }}
        </template>
      </el-table-column>
      <el-table-column label="宽" align="center">
        <template slot-scope="scope">
          {{ scope.row.width }}
        </template>
      </el-table-column>
      <el-table-column label="高" align="center">
        <template slot-scope="scope">
          {{ scope.row.height }}
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
        <el-form-item label="客户编号" prop="customer_id">
          <el-input v-model="editForm.customer_id" type="text" placeholder="请输入客户编号" />
        </el-form-item>
        <el-form-item label="出发地" prop="origin">
          <el-input v-model="editForm.origin" type="text" placeholder="请输入出发地" />
        </el-form-item>
        <el-form-item label="用户编号" prop="user_id">
          <el-input v-model="editForm.user_id" type="text" placeholder="请输入用户编号" />
        </el-form-item>
        <el-form-item label="长" prop="long">
          <el-input v-model="editForm.long" type="text" placeholder="请输入集装箱长度" />
        </el-form-item>
        <el-form-item label="宽" prop="width">
          <el-input v-model="editForm.width" type="text" placeholder="请输入集装箱宽度" />
        </el-form-item>
        <el-form-item label="高" prop="height">
          <el-input v-model="editForm.height" type="text" placeholder="请输入集装箱高度" />
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
import { page, create, update } from '@/api/container'
import { statusConf, buttonNames } from '@/config/common'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

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
    return {
      statusConf: statusConf,

      // 编辑框
      dialogFormVisible: false,
      dialogStatus: '',
      buttonNames: buttonNames,
      editForm: {
        customer_id: '',
        origin: '',
        user_id: '',
        height: '',
        width: '',
        long: '',
        status: ''
      },

      rules: {
        customer_id: [{ required: true, message: '请填写客户编号', trigger: 'blur' }],
        origin: [{ required: true, message: '请填写出发地', trigger: 'blur' }],
        user_id: [{ required: true, message: '请填写用户编号', trigger: 'blur' }],
        long: [{ required: true, message: '请填写长度', trigger: 'blur' }],
        width: [{ required: true, message: '请填写宽度', trigger: 'blur' }],
        height: [{ required: true, message: '请填写高度', trigger: 'blur' }],
        status: [{ required: true, message: '请选择状态', trigger: 'blur' }]
      },

      // 列表
      list: null,
      listQuery: {
        costumer_id: '',
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
        this.list = response.data.containers
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
      form.long = parseInt(form.long)
      form.width = parseInt(form.width)
      form.height = parseInt(form.height)
      return form
    },
    resetEditForm() {
      this.editForm = {
        customer_id: '',
        origin: '',
        user_id: '',
        height: '',
        width: '',
        long: '',
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
