<template>
  <a-card :bordered="false">
    <a-row :gutter="8">
      <a-col :span="5">
        <a-tree
          showLine
          :defaultExpandedKeys="['0-0-0']"
          :treeData="treeData"
          @select="onSelect">
        </a-tree>
      </a-col>
      <a-col :span="19" style="border-left: 1px solid #dadada;">
        <div style="text-align: right">
          <a-tag size="small" color="cyan" type="primary">
            <a-icon type="plus"/>
            新增子菜单
          </a-tag>
          <a-tag color="pink" style="margin-left: 8px">
            <a-icon type="delete"/>
            删除
          </a-tag>
        </div>
        <a-form>
          <a-form-item
            label="已选中"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2">
            <b>张三</b>
          </a-form-item>
          <a-form-item
            label="区域名称"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2">
            <a-input placeholder="输入区域名称"/>
          </a-form-item>
          <a-form-item
            label="排序"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2">
            <a-input placeholder="输入排序"/>
          </a-form-item>
          <a-form-item
            label="备注说明"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2">
            <a-textarea placeholder="备注说明" :rows="4" v-decorator="['desc', {rules: [{required: true}]}]"></a-textarea>
          </a-form-item>
          <a-form-item style="text-align: center">
            <a-button style="margin-right: 8px">重置</a-button>
            <a-button type="primary">修改</a-button>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import STree from '@/components/Tree/Tree'
import { STable } from '@/components'
import { getOrgTree, getServiceList } from '@/api/manage'

const treeData = [{
  title: 'parent 1',
  key: '0-0',
  slots: {
    icon: 'smile'
  },
  children: [
    { title: 'leaf', key: '0-0-0' },
    {
      title: 'leaf',
      key: '0-0-1',
      scopedSlots: { icon: 'custom' },
      children: [
        { title: 'leaf', key: '0-0-1-0' },
        { title: 'leaf', key: '0-0-1-1' }
      ]
    }
  ]
}]
export default {
  name: 'SysAreaList',
  components: {
    STable,
    STree
  },
  data () {
    return {
      openKeys: ['key-01'],
      labelCol2: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol2: {
        xs: { span: 24 },
        sm: { span: 15 }
      },
      treeData,
      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: '序',
          dataIndex: 'no'
        },
        {
          title: '姓名',
          dataIndex: 'description'
        },
        {
          title: '登录次数',
          dataIndex: 'callNo',
          sorter: true,
          needTotal: true,
          customRender: (text) => text + ' 次'
        },
        {
          title: '状态',
          dataIndex: 'status',
          needTotal: true
        },
        {
          title: '最近登录时间',
          dataIndex: 'updatedAt',
          sorter: true
        },
        {
          table: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return getServiceList(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.result
          })
      },
      orgTree: [],
      dataList: [],
      selectedRowKeys: [],
      selectedRows: [],

      expandedKeys: [],
      searchValue: '',
      autoExpandParent: true
    }
  },
  created () {
    getOrgTree().then(res => {
      this.orgTree = res.result
      this.dataList = this.orgTree
    })
  },
  methods: {
    onChange (e) {
      const value = e.target.value
      const expandedKeys = this.dataList.map((item) => {
        if (item.key.indexOf(value) > -1) {
          return getParentKey(item.key, gData)
        }
        return null
      }).filter((item, i, self) => item && self.indexOf(item) === i)
      Object.assign(this, {
        expandedKeys,
        searchValue: value,
        autoExpandParent: true
      })
    },
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
      this.$refs.modal.add(item.key)
    },
    handleTitleClick (item) {
      console.log('handleTitleClick', item)
    },
    titleClick (e) {
      console.log('titleClick', e)
    },
    handleSaveOk () {

    },
    handleSaveClose () {

    },
    onSelect (selectedKeys, info) {
      console.log('selected', selectedKeys, info)
    }
  }
}
</script>

<style lang="less">
  .custom-tree {

  /deep/ .ant-menu-item-group-title {
    position: relative;

  &
  :hover {

  .btn {
    display: block;
  }

  }
  }

  /deep/ .ant-menu-item {

  &
  :hover {

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

  &
  :hover {
    transform: scale(1.2);
    transition: 0.5s all;
  }

  }
  }
</style>
