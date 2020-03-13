<template>
  <a-card :bordered="false">
    <div class="table-page-search-wrapper" v-if="$auth('sys_role_save')">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="操作内容">
              <a-input  v-model="queryParam.title" placeholder="请输入操作内容"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="请求路径">
              <a-input  v-model="queryParam.url" placeholder="请输入请求路径"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="IP">
              <a-input  v-model="queryParam.ip" placeholder="请输入IP"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="操作类型">
              <a-select v-model="queryParam.type" placeholder="请选择操作类型" default-value="0">
                <a-select-option value="">全部</a-select-option>
                <a-select-option value="1">管理后台</a-select-option>
                <a-select-option value="2">用户前台</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="操作时间">
              <a-range-picker
                v-model="queryTime"
                :showTime="{
                    hideDisabledOptions: true,
                    defaultValue: [moment('00:00:00', 'HH:mm:ss'), moment('11:59:59', 'HH:mm:ss')]
                  }"
                format="YYYY-MM-DD HH:mm:ss"
                @change="onChangeTime"
              />
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <span class="table-page-search-submitButtons">
              <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
              <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>
    <s-table
      ref="table"
      size="default"
      :columns="columns"
      :data="loadData"
      rowKey="id"
    >
      <span slot="serial" slot-scope="text, record, index">
        {{ index + 1 }}
      </span>
      <span slot="created_at" slot-scope="text">
        {{ formatDate(text) }}
      </span>
      <span slot="type" slot-scope="text">
        <a-tag color="purple" v-if="text==='1'">管理后台</a-tag>
        <a-tag color="green" v-if="text==='2'">用户前台</a-tag>
      </span>
      <div
        slot="expandedRowRender"
        slot-scope="record"
        style="margin: 0">
        <a-row
          :gutter="24"
          :style="{ marginBottom: '12px' }">
          <a-col :span="24" style="margin-bottom: 12px;">
            <a-col :lg="4" :md="24">
              <span>请求路径：</span>
            </a-col>
            <a-col :lg="20" :md="24" >
              <b>{{ record.request_url }}</b>
            </a-col>
          </a-col>
          <a-col :span="24" style="margin-bottom: 12px;">
            <a-col :lg="4" :md="24">
              <span>参数：</span>
            </a-col>
            <a-col :lg="20" :md="24" >
              <b>{{ record.params }}</b>
            </a-col>
          </a-col>
          <a-col :span="24" style="margin-bottom: 12px;">
            <a-col :lg="4" :md="24">
              <span>IP：</span>
            </a-col>
            <a-col :lg="20" :md="24" >
              <b>{{ record.remote_addr }}</b>
            </a-col>
          </a-col>
          <a-col :span="24" style="margin-bottom: 12px;">
            <a-col :lg="4" :md="24">
              <span>User-Agent：</span>
            </a-col>
            <a-col :lg="20" :md="24" >
              <b>{{ record.user_agent }}</b>
            </a-col>
          </a-col>
        </a-row>
      </div>
    </s-table>
  </a-card>
</template>

<script>
import { STable } from '@/components'
import { GetManageActionLog } from '@/api/sys'
import { formatDate } from '@/utils/util'
import moment from 'moment'
var that
export default {
  name: 'SysLogList',
  components: {
    STable,
    moment
  },
  data () {
    return {
      queryTime: [],
      visible: false,
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      },
      form: null,
      mdl: {},

      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: '序',
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '操作内容',
          dataIndex: 'title'
        },
        {
          title: '操作人员',
          dataIndex: 'name'
        },
        {
          title: '操作类型',
          dataIndex: 'type',
          scopedSlots: { customRender: 'type' }
        },
        {
          title: '操作时间',
          dataIndex: 'created_at',
          sorter: true,
          scopedSlots: { customRender: 'created_at' }
        }
      ],
      // 向后端拉取可以用的操作列表
      permissionList: null,
      AccountNames: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return GetManageActionLog(Object.assign(parameter, this.queryParam))
          .then(res => {
            if (res.ret === 200) {
              if (res.data.data != null && res.data.data.length > 0 && res.data.extendData != null && res.data.extendData.length > 0) {
                for (let i = 0; i < res.data.data.length; i++) {
                  if (res.data.data[i].api_id !== '') {
                    const an = res.data.extendData.filter(item => item.id === res.data.data[i].api_id)
                    if (an.length > 0) {
                      res.data.data[i].name = an[0].name
                    }
                  }
                }
              }
              return res.data
            } else {
              return []
            }
          })
      },

      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    statusFilter (status) {
      const statusMap = {
        1: '正常',
        2: '禁用'
      }
      return statusMap[status]
    }
  },
  created () {
    that = this
    this.loadPermissionList()
  },
  methods: {
    moment,
    formatDate,
    onChangeTime (value, dateString) {
      this.queryParam.startTime = dateString[0]
      this.queryParam.endTime = dateString[1]
    },
    loadPermissionList () {
      // permissionList
      new Promise(resolve => {
        const data = [
          { label: '新增', value: 'add', defaultChecked: false },
          { label: '查询', value: 'get', defaultChecked: false },
          { label: '修改', value: 'update', defaultChecked: false },
          { label: '列表', value: 'query', defaultChecked: false },
          { label: '删除', value: 'delete', defaultChecked: false },
          { label: '导入', value: 'import', defaultChecked: false },
          { label: '导出', value: 'export', defaultChecked: false }
        ]
        setTimeout(resolve(data), 1500)
      }).then(res => {
        this.permissionList = res
      })
    },
    handleEdit (record) {
      this.mdl = Object.assign({}, record)
      console.log(this.mdl)
      this.visible = true
    },
    handleOk () {

    },
    onChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    }
  },
  watch: {
    /*
      'selectedRows': function (selectedRows) {
        this.needTotalList = this.needTotalList.map(item => {
          return {
            ...item,
            total: selectedRows.reduce( (sum, val) => {
              return sum + val[item.dataIndex]
            }, 0)
          }
        })
      }
      */
  }
}
</script>
