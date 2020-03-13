<template>
  <a-card :bordered="false">
    <a-row :gutter="8">
      <a-col :md="8">
        <a-card title="管理菜单及操作权限" v-if="$auth('sys_menu_select')">
          <a slot="extra" type="primary" @click="handleSaveMenu('0')" v-if="$auth('sys_menu_save')">
            <a-icon type="plus"/> 新增顶级菜单</a>
          <a-spin :spinning="menuLoading">
            <a-tree
              showIcon
              ref="menuTree"
              @select="onSelectMenu"
              :expandedKeys="menuExpandedKeys"
              :treeData="menuTreeData"
            >
              <template slot="custom" slot-scope="item">
                <a-icon :type="item.dataRef.icon" />
              </template>
            </a-tree>
          </a-spin>
        </a-card>
      </a-col>
      <a-col :span="16" >
        <a-card :title="handleMenuTitle">
          <a slot="extra" type="primary" v-if="selectMenuId!=='' && $auth('sys_menu_save')" @click="handleSaveMenu(selectMenuId)" >
            <a-icon type="plus"/> 新增子菜单</a>
          <a
            type="primary"
            slot="extra"
            v-if="selectMenuId!==''"
            @click="handleDelMenu"
            style="color: red;margin-left: 10px"><a-icon type="delete" v-if="$auth('sys_menu_delete')"/> 删除选中菜单</a>

          <a-spin :spinning="menuEditLoading">
            <a-form
              :form="form">
              <a-form-item
                label="菜单名称"
                :labelCol="labelCol2"
                :wrapperCol="wrapperCol2">
                <a-input placeholder="输入菜单名称" v-decorator="['name', {rules: [{required: true}]}]" />
              </a-form-item>
              <a-form-item
                label="权限标识"
                :labelCol="labelCol2"
                :wrapperCol="wrapperCol2">
                <a-input placeholder="输入权限标识" v-decorator="['permission', {rules: [{required: true}]}]" />
              </a-form-item>
              <a-form-item
                label="菜单图标"
                :labelCol="labelCol2"
                :wrapperCol="wrapperCol2">
                <a-input placeholder="输入菜单图标" v-decorator="['icon']" />
              </a-form-item>
              <a-form-item
                label="菜单描述"
                :labelCol="labelCol2"
                :wrapperCol="wrapperCol2">
                <a-textarea placeholder="输入菜单描述" :rows="4" v-decorator="['remarks']"></a-textarea>
              </a-form-item>
              <a-form-item style="text-align: center">
                <a-button style="margin-right: 8px" @click="resetMenuForm">重置</a-button>
                <a-button type="primary" @click="saveMenu" v-if="$auth('sys_menu_save')">保存</a-button>
              </a-form-item>
            </a-form>
          </a-spin>
        </a-card>
        <a-card
          title="操作权限"
          v-if="selectMenuId!==''">
          <a slot="extra" type="primary" @click="handleMenuBtnEdit()" v-if="$auth('sys_menu_btn_save')"><a-icon type="plus"/> 新增操作权限</a>
          <a-table
            :dataSource="menuBtnSource"
            :columns="columns"
            rowKey="id"
            :loading="menuBtnLoading">
            <span slot="serial" slot-scope="text, record, index">
              {{ index + 1 }}
            </span>
            <template slot="action" slot-scope="text, record">
              <a @click="handleMenuBtnEdit(record)"   v-if="$auth('sys_menu_btn_save')">编辑</a>
              <a-divider type="vertical" />
              <a-popconfirm
                v-if="menuBtnSource.length && $auth('sys_menu_btn_delete')"
                title="是否删除此操作权限?"
                @confirm="() => onMenuBtnDelete(record.id)">
                <a href="javascript:;" style="color: red">删除</a>
              </a-popconfirm>
            </template>
          </a-table>
        </a-card>
      </a-col>
    </a-row>
    <MenuBtnEdit ref="MenuBtnEdit" @ok="handleMenuBtnEdit"/>
  </a-card>
</template>

<script>
import pick from 'lodash.pick'
import {
  GetMenuList,
  GetBtnMenuList,
  SaveMenu,
  GetMenu,
  DelMenu,
  DelMenuBtn
} from '@/api/sys'
import MenuBtnEdit from './MenuBtnEdit'

let that
export default {
  name: 'SysMenuList',
  components: {
    MenuBtnEdit
  },
  data () {
    return {
      id: '',
      form: this.$form.createForm(this),
      labelCol2: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol2: {
        xs: { span: 24 },
        sm: { span: 15 }
      },
      menuExpandedKeys: [],
      menuTreeData: [],
      menuLoading: true,
      menuEditLoading: false,
      menuBtnLoading: false,
      menuBtnSource: [],
      selectMenuId: '',
      parentMenuId: '0',
      handleMenuTitle: '新增顶级菜单',
      // 表头
      columns: [
        {
          title: '序',
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '名称',
          dataIndex: 'name'
        },
        {
          title: '请求方法',
          dataIndex: 'method'
        },
        {
          title: '权限标识',
          dataIndex: 'permission'
        },
        {
          title: '描述',
          dataIndex: 'remarks'
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
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
    that = this
    that.loadMenuList()
  },
  methods: {
    loadMenuList () {
      that.menuLoading = true
      GetMenuList({ type: 'antTree' }).then(res => {
        that.menuLoading = false
        if (res.ret === 200) {
          that.menuTreeData = res.data
          that.setMenuExpandedKeys(that.menuTreeData)
        }
      })
    },
    loadMenuBtnList (menuId) {
      that.menuBtnLoading = true
      GetBtnMenuList({ menuId: menuId }).then(res => {
        that.menuBtnLoading = false
        if (res.ret === 200) {
          that.menuBtnSource = res.data
        }
      })
    },
    setMenuExpandedKeys (list) {
      if (list && list.length > 0) {
        for (let i = 0; i < list.length; i++) {
          that.menuExpandedKeys.push(list[i].key)
          if (list[i].children && list[i].children.length > 0) {
            that.setMenuExpandedKeys(list[i].children)
          }
        }
      }
    },
    onSelectMenu (selectedKeys, info, e) {
      that.handleMenuTitle = '编辑菜单'
      that.selectMenuId = info.node.dataRef.key
      that.parentMenuId = info.node.dataRef.parentId
      that.id = that.selectMenuId
      const { form: { setFieldsValue } } = that
      that.menuEditLoading = true
      GetMenu({ id: that.selectMenuId }).then(res => {
        that.menuEditLoading = false
        if (res.ret === 200) {
          that.$nextTick(() => {
            setFieldsValue(pick(res.data, 'name', 'icon', 'permission', 'remarks'))
          })
        }
      })
      // 加载菜单按钮
      that.loadMenuBtnList(selectedKeys[0])
    },
    handleSaveMenu (menuId) {
      that.id = ''
      this.form.resetFields()
      if (menuId !== '0') {
        that.handleMenuTitle = '编辑菜单'
        this.parentMenuId = menuId
      } else {
        that.handleMenuTitle = '新增顶级菜单'
        that.selectMenuId = ''
        that.parentMenuId = '0'
      }
    },
    handleDelMenu () {
      that.$confirm({
        title: '是否删除此菜单?',
        okText: '删除',
        okType: 'danger',
        cancelText: '取消',
        onOk () {
          that.menuEditLoading = true
          that.menuLoading = true
          that.menuBtnLoading = true
          DelMenu({ id: that.selectMenuId }).then(res => {
            that.menuEditLoading = false
            that.menuLoading = false
            that.menuBtnLoading = false
            if (res.ret === 200) {
              that.$notification.success({
                message: '系统提示',
                description: res.msg
              })
              that.loadMenuList()
              that.resetMenuForm()
              that.selectMenuId = ''
            }
          })
        }
      })
    },
    onMenuBtnDelete (id) {
      that.menuBtnLoading = true
      DelMenuBtn({ id: id }).then(res => {
        that.menuBtnLoading = false
        if (res.ret === 200) {
          this.$notification.success({
            message: '系统提示',
            description: res.msg
          })
          that.loadMenuBtnList(that.selectMenuId)
        }
      })
    },
    saveMenu () {
      const { form: { validateFields } } = this
      validateFields((errors, values) => {
        if (!errors) {
          values.id = that.id
          values.parent_id = that.parentMenuId
          that.confirmLoading = true
          that.menuEditLoading = true
          SaveMenu(values)
            .then(res => {
              values.parent_id = that.parentMenuId
              that.menuEditLoading = false
              if (res.ret === 200) {
                this.$notification.success({
                  message: '系统提示',
                  description: res.msg
                })
                that.loadMenuList()
              }
            })
        }
      })
    },
    resetMenuForm () {
      this.form.resetFields()
    },
    handleMenuBtnEdit (record) {
      this.$refs.MenuBtnEdit.edit(record, that.selectMenuId, function () {
        that.loadMenuBtnList(that.selectMenuId)
      })
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
