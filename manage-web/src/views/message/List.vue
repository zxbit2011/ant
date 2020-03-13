<template>
  <page-view :title="title">
    <a-card :bordered="false">
      <div class="table-page-search-wrapper" v-if="$auth('sys_message_select')">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="8" :sm="24">
              <a-form-item label="信息名称">
                <a-input v-model="queryParam.title" placeholder="输入信息名称"/>
              </a-form-item>
            </a-col>
            <a-col :md="8" :sm="24">
              <a-form-item label="发布状态">
                <a-select v-model="queryParam.status" placeholder="请选择发布状态" default-value="0">
                  <a-select-option value="">全部</a-select-option>
                  <a-select-option value="1">未发布</a-select-option>
                  <a-select-option value="2">已发布</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
                <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleEdit()" v-if="$auth('sys_message_save')" >新建</a-button>
        <a-button type="danger" :disabled="selectedRowKeys.length === 0" @click="delSysNotify" v-if="$auth('sys_message_delete')" style="margin-left: 8px">
          批量删除
        </a-button>
      </div>

      <s-table
        ref="table"
        size="default"
        rowKey="id"
        :columns="columns"
        :data="loadData"
        :rowSelection="rowSelection"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>
        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <span slot="type" slot-scope="text">
          <a-tag v-if="text==='1'" color="green">全部</a-tag>
          <a-tag v-if="text==='2'" color="orange">指定项目</a-tag>
        </span>
        <span slot="created_at" slot-scope="text">
          {{ formatDate(text) }}
        </span>

        <span slot="action" slot-scope="text, record, index">
          <template>
            <a v-if="$auth('sys_message_save')" @click="handleEdit(record)">编辑</a>
            <a-divider type="vertical" />
            <a v-if="$auth('sys_message_release')">
              <a-dropdown>
                <a-menu slot="overlay">
                  <a-menu-item
                    :disabled="record.status==='2'"
                    @click="releaseSysNotify(record,'是否发布消息？')">发布消息</a-menu-item>
                  <a-menu-item
                    :disabled="record.status==='1'"
                    @click="releaseSysNotify(record,'是否取消发布消息？')">取消发布</a-menu-item>
                </a-menu>
                <a color="cyan">发布<a-icon type="down"/></a>
              </a-dropdown>
            </a>
          </template>
        </span>
      </s-table>
      <EditModal ref="editModal" @ok="handleEdit"/>
      <DetailModal ref="detailModal" @ok="handleDetail"/>
    </a-card>
  </page-view>
</template>

<script>
import HeadInfo from '@/components/tools/HeadInfo'
import { PageView } from '@/layouts'
import { STable } from '@/components'
import EditModal from './EditModal'
import DetailModal from './DetailModal'
import { formatDate } from '@/utils/util'
import { GetSysNotifyPage, ReleaseSysNotify, DelSysNotify } from '@/api/notify'
const statusMap = {
  1: {
    status: 'warning',
    text: '未发布'
  },
  2: {
    status: 'success',
    text: '已发布'
  }
}
const typeMap = {
  1: {
    status: 'processing',
    text: '全部'
  },
  2: {
    status: 'success',
    text: '指定项目'
  }
}

let that
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    that.selectedRowKeys = selectedRowKeys
    that.selectedRows = selectedRows
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows)
  },
  onSelect: (record, selected, selectedRows) => {
    console.log(record, selected, selectedRows)
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    console.log(selected, selectedRows, changeRows)
  }
}

export default {
  name: 'MessageList',
  components: {
    HeadInfo,
    PageView,
    STable,
    EditModal,
    DetailModal
  },
  data () {
    return {
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
      rowSelection,
      // 表头
      columns: [
        {
          title: '序',
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '标题',
          dataIndex: 'title',
          scopedSlots: { customRender: 'title' }
        },
        {
          title: '发送对象',
          dataIndex: 'type',
          scopedSlots: { customRender: 'type' }
        },
        {
          title: '发送状态',
          dataIndex: 'status',
          sorter: true,
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '创建时间',
          dataIndex: 'created_at',
          sorter: true,
          scopedSlots: { customRender: 'created_at' }
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return GetSysNotifyPage(Object.assign(parameter, this.queryParam))
          .then(res => {
            if (res.ret === 200) {
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
    statusFilter (type) {
      return statusMap[parseInt(type)].text
    },
    statusTypeFilter (type) {
      return statusMap[parseInt(type)].status
    },
    typeFilter (type) {
      return typeMap[parseInt(type)].text
    },
    typeTypeFilter (type) {
      return typeMap[parseInt(type)].status
    }
  },
  created () {
    that = this
  },
  computed: {
    title () {
      return this.$route.meta.title
    }
  },
  methods: {
    formatDate,
    handleEdit (record) {
      console.log(record)
      this.$refs.editModal.edit(record, function () {
        that.$refs.table.refresh(true)
      })
    },
    handleDetail (record) {
      console.log(record)
      this.$refs.detailModal.detail(record)
    },
    releaseSysNotify (record, msg) {
      this.$confirm({
        title: '消息发布',
        content: msg,
        okText: '确认',
        cancelText: '取消',
        onOk () {
          ReleaseSysNotify({ id: record.id })
            .then(res => {
              if (res.ret === 200) {
                that.$refs.table.refresh()
                that.$notification.success({
                  message: '消息发布',
                  description: res.msg
                })
              }
            })
        }
      })
    },
    delSysNotify () {
      if (this.selectedRowKeys.length === 0) {
        this.$notification.warning({
          message: '消息删除',
          description: '请选择需删除的消息'
        })
      }
      this.$confirm({
        title: '消息删除',
        content: `是否确定删除选择的${this.selectedRowKeys.length}个消息？`,
        okText: '确认',
        cancelText: '取消',
        onOk () {
          DelSysNotify({ ids: that.selectedRowKeys.join(',') }).then(res => {
            if (res.ret === 200) {
              that.selectedRowKeys = []
              that.$refs.table.refresh()
              that.$notification.success({
                message: '消息删除',
                description: res.msg
              })
            }
          })
        }
      })
    },
    handleOk () {
      this.$refs.table.refresh()
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {}
    }
  }
}
</script>
