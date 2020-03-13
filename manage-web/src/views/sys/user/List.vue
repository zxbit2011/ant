<template>
  <a-card :bordered="false">
    <a-row :gutter="8">
      <a-col :span="5" >
        <a-card title="组织架构" v-if="$auth('sys_menu_select')">
          <a slot="extra" type="primary" @click="handleOfficeMenu('0')" v-if="$auth('sys_menu_save')">
            <a-icon type="plus"/> 新增子组织架构</a>
          <a-spin :spinning="officeLoading">
            <a-tree
              showLine
              ref="officeTree"
              @select="onOfficeMenu"
              :treeData="officeTree"
            >
            </a-tree>
          </a-spin>
        </a-card>
      </a-col>
      <a-col :span="19">

        <a-card :title="checkedOfficeTitle">
          <div class="table-page-search-wrapper" v-if="$auth('sys_message_select')">
            <a-form layout="inline">
              <a-row :gutter="48">
                <a-col :md="8" :sm="24">
                  <a-form-item label="用户类型">
                    <a-select v-model="queryParam.usertype" placeholder="请选择用户类型" default-value="0">
                      <a-select-option value="">全部</a-select-option>
                      <a-select-option value="sec">教委</a-select-option>
                      <a-select-option value="teacher">教师</a-select-option>
                      <a-select-option value="student">学生</a-select-option>
                    </a-select>
                  </a-form-item>
                </a-col>
                <a-col :md="8" :sm="24">
                  <span class="table-page-search-submitButtons" >
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
            :alert="false"
            rowKey="id"
            v-if="$auth('sys_user_select')"
          >
            <span slot="serial" slot-scope="text, record, index">
              {{ index + 1 }}
            </span>
            <span slot="user_type" slot-scope="text">
              <span v-if="text==='sec'">教委</span>
              <span v-else-if="text==='teacher'">教师</span>
              <span v-else-if="text==='student'">学生</span>
              <span v-else>text</span>
            </span>
            <span slot="login_date" slot-scope="text">
              {{ formatDate(text) }}
            </span>
            <span slot="role" slot-scope="text, record">
              <template>
                <span v-if="$auth('sys_user_disable') && roleList[record.id]" >
                  <a >
                    <a-tag color="purple" v-for="role in roleList[record.id]" :key="role">{{ role }}</a-tag>
                  </a>
                </span>
              </template>
            </span>
            <span slot="action" slot-scope="text, record">
              <template>
                <span v-if="$auth('sys_user_disable')" >
                  <a v-if="record.login_flag==='1'" @click="handleDisable(record)" style="color: red">禁用</a>
                  <a v-else @click="handleDisable(record)">取消禁用</a>
                </span>
              </template>
            </span>
          </s-table>
        </a-card>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import STree from '@/components/Tree/Tree'
import { STable } from '@/components'
import EditModal from './EditModal'
import { formatDate } from '@/utils/util'
import { GetSysOfficeAll, GetSysUserList, DisableSysUser } from '@/api/sys'
var that = this
export default {
  name: 'SysUserList',
  components: {
    STable,
    STree,
    EditModal
  },
  data () {
    return {
      openKeys: ['key-01'],
      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: '序',
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '姓名',
          dataIndex: 'name'
        },
        {
          title: '登录名',
          dataIndex: 'username'
        },
        {
          title: '用户类型',
          dataIndex: 'user_type',
          scopedSlots: { customRender: 'user_type' }
        },
        {
          title: '用户角色',
          dataIndex: 'role',
          width: '150px',
          scopedSlots: { customRender: 'role' }
        },
        {
          title: '最近登录时间',
          dataIndex: 'login_date',
          sorter: true,
          scopedSlots: { customRender: 'login_date' }
        },
        {
          title: '最近登录IP',
          dataIndex: 'login_ip',
          sorter: true
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      checkedOfficeTitle: '系统用户 - 全部',
      officeLoading: false,
      roleList: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return GetSysUserList(Object.assign(parameter, this.queryParam))
          .then(res => {
            if (res.ret === 200) {
              that.roleList = res.data.extendData.roleList
              if (res.data.data != null && res.data.data.length > 0 && res.data.extendData.accountNameList != null && res.data.extendData.accountNameList.length > 0) {
                for (let i = 0; i < res.data.data.length; i++) {
                  if (res.data.data[i].api_id !== '') {
                    const an = res.data.extendData.accountNameList.filter(item => item.id === res.data.data[i].api_id)
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
      officeTree: [],
      props: {
        label: 'name',
        children: 'children'
      },
      onSearch: ['sdf'],
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  created () {
    that = this
    GetSysOfficeAll().then(res => {
      if (res.ret === 200) {
        that.officeLoading = false
        that.officeTree = res.data
      }
    })
  },
  methods: {
    formatDate,
    handleClick (e) {
      console.log('handleClick', e)
      this.queryParam = {
        key: e.key
      }
      this.$refs.table.refresh(true)
    },
    handleAdd (item) {
      console.log('add button, item', item)
      this.$message.info(`提示：你点了 ${item.key} - ${item.title} `)
      this.$refs.editModal.add(item.key)
    },
    handleTitleClick (item) {
      console.log('handleTitleClick', item)
    },
    titleClick (e) {
      console.log('titleClick', e)
    },
    handleDisable (record) {
      this.$confirm({
        title: '禁用系统用户',
        content: `是否确定禁用【${record.name}】系统用户？`,
        okText: '确认',
        cancelText: '取消',
        onOk () {
          DisableSysUser({ id: record.id }).then(res => {
            if (res.ret === 200) {
              that.$refs.table.refresh(true)
              that.$notification.success({
                message: '禁用系统用户',
                description: res.msg
              })
            }
          })
        }
      })
    },
    onOfficeMenu (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    }
  }
}
</script>

<style lang="less">
  .custom-tree {

  /deep/ .ant-menu-item-group-title {
    position: relative;
  &:hover {
  .btn {
    display: block;
  }
  }
  }

  /deep/ .ant-menu-item {
  &:hover {
  .btn {
    display: block;
  }
  }
  }

  /deep/ .btn {
    display: none;
    position: absolute;
    top: 0;
    right: 10px;
    width: 20px;
    height: 40px;
    line-height: 40px;
    z-index: 1050;

  &:hover {
     transform: scale(1.2);
     transition: 0.5s all;
   }
  }
  }
</style>
