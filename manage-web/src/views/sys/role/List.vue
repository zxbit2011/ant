<template>
  <a-card :bordered="false" :style="{ height: '100%' }">
    <a-row :gutter="24">
      <a-col :md="6" style="padding: 0">
        <a-card title="选择角色">
          <a slot="extra" type="primary" @click="handleRoleEdit()" v-if="$auth('sys_role_save')"><a-icon type="plus"/> 新增角色</a>
          <a-list itemLayout="horizontal" :dataSource="roles" :loading="roleLoading" v-if="$auth('sys_role_select')">
            <a-list-item
              slot="renderItem"
              slot-scope="item, index"
              :key="index"
              :class="item.id===checkRoleId?'hover':''"
              @click="checkRole(item)"
              :style="{cursor:item.role_type!=='2'?'cursor':'not-allowed'}" >
              <a slot="actions" v-if="item.role_type!=='2' && $auth('sys_role_delete')"><a-icon type="edit" @click="handleRoleEdit(item)" /></a>
              <a slot="actions" v-if="item.role_type!=='2' && $auth('sys_role_save')"><a-icon type="delete" @click="delRole(item)" style="color: #ff0000" /></a>
              <a-list-item-meta :style="{ marginBottom: '0' }">
                <span slot="title" style="color: #9e9e9e" v-if="item.role_type==='2'"><a-icon type="sliders" /> {{ item.name }}（所有权限）</span>
                <a slot="title" v-else><a-icon type="sliders" /> {{ item.name }}</a>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </a-card>
      </a-col>
      <a-col :md="9">
        <a-card title="选择菜单权限">
          <a-spin :spinning="menuLoading">
            <a-tree
              :checkable="$auth('sys_role_menu_save')"
              ref="menuTree"
              @select="onSelectMenu"
              @check="onCheckMenu"
              :checkedKeys="menuCheckedKeys"
              :expandedKeys="menuExpandedKeys"
              :treeData="menuTreeData"
            />
          </a-spin>
        </a-card>
      </a-col>
      <a-col :md="9">
        <a-card title="选择操作权限" >
          <!--<div slot="extra">
            <a-checkbox
              :indeterminate="indeterminate"
              @change="onCheckAllChange"
              :checked="checkAll"
            >全选
            </a-checkbox>
          </div>-->
          <a-spin :spinning="menuBtnLoading">
            <!--<a-checkbox-group :options="menuBtnOptions" v-model="checkedMenuBtnOptions" @change="onChange" />-->
            <a-checkbox
              class="ant-checkbox-group-item"
              :checked="checkedMenuBtnOptions.indexOf(item.value)>-1"
              v-for="item in menuBtnOptions"
              @change="onChange"
              :value="item.value"
              :key="item.value"
              :disabled="!$auth('sys_role_menu_save')"
            >{{ item.label }}</a-checkbox>
          </a-spin>
        </a-card>
      </a-col>
    </a-row>
    <EditModal ref="editModal" @ok="handleRoleEdit"/>
  </a-card>
</template>

<script>
import {
  DelRole,
  GetRoleList,
  GetMenuList,
  GetCheckedMenuLimit,
  UpdateRoleMenuByIds,
  GetBtnMenuList,
  GetCheckedMenuBtnLimit,
  UpdateRoleMenuBtnById } from '@/api/sys'
import { mixinDevice } from '@/utils/mixin'
import EditModal from './EditModal'

var that
export default {
  name: 'SysRoleList',
  mixins: [mixinDevice],
  components: {
    EditModal
  },
  data () {
    return {
      menuExpandedKeys: [],
      menuCheckedKeys: [],
      menuCheckedKeysCache: [],
      checkedMenuBtnOptions: [],
      checkedMenuBtnValues: [],
      selectMenu: null,
      menuBtnOptions: [],
      menuTreeData: [],
      roleTreeData: [],
      checkRoleId: 0,
      checkMenuId: 0,
      form: this.$form.createForm(this),
      mdl: {},
      roles: [],
      indeterminate: true,
      checkAll: false,
      roleLoading: true,
      menuLoading: true,
      menuBtnLoading: false,
      permissions: [],
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      }
    }
  },
  created () {
    that = this
    that.loadRoleList()
    that.loadMenuList()
  },
  methods: {
    checkRole (item) {
      if (item.role_type !== '2') {
        that.loadRoleMenuList(item.id)
      }
    },
    onChange (e) {
      console.log('checked = ', e)
      if (e.target.checked) {
        that.checkedMenuBtnOptions.push(e.target.value)
      } else {
        that.checkedMenuBtnOptions = that.checkedMenuBtnOptions.filter(item => item !== e.target.value)
      }
      UpdateRoleMenuBtnById({
        id: e.target.value,
        roleId: this.checkRoleId,
        menuId: this.checkMenuId,
        add: e.target.checked
      }).then((res) => {
        if (res.ret === 200) {
          this.$message.success('保存成功')
        }
      })
    },
    onCheckAllChange (e) {
      that.checkedMenuBtnOptions = e.target.checked ? that.menuBtnOptions.map(item => item.value) : []
      that.indeterminate = false
      that.checkAll = e.target.checked
    },
    handleRoleEdit (record) {
      console.log(record)
      this.$refs.editModal.edit(record, function () {
        that.loadRoleList()
      })
    },
    delRole (record) {
      this.$confirm({
        title: '禁用系统用户',
        content: `是否确定删除【${record.name}】角色？`,
        okText: '确认',
        cancelText: '取消',
        onOk () {
          DelRole({ id: record.id }).then((res) => {
            if (res.ret === 200) {
              that.$message.success('删除成功')
              that.loadRoleList()
            }
          })
        }
      })
    },
    loadRoleList () {
      GetRoleList().then((res) => {
        that.roleLoading = false
        if (res.ret === 200) {
          that.roles = res.data
          if (that.roles && that.roles.length > 0) {
            that.checkedDefaultRole(that.roles, 0)
          }
        }
      })
    },
    checkedDefaultRole (roles, i) {
      if (roles.length - 1 > i) {
        if (that.roles[i].role_type === '2') {
          i++
          that.checkedDefaultRole(roles, i)
        } else {
          that.loadRoleMenuList(roles[i].id)
        }
      }
    },
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
    loadRoleMenuList (roleId) {
      that.checkRoleId = roleId
      that.menuCheckedKeys = []
      that.menuCheckedKeysCache = []
      that.menuLoading = true
      GetCheckedMenuLimit({ id: roleId }).then(res => {
        if (res.ret === 200) {
          that.menuLoading = false
          const list = res.data
          if (list && list.length > 0) {
            for (let i = 0; i < list.length; i++) {
              if (list.filter(item => item.relation_ids !== list[i].relation_ids && item.relation_ids.indexOf(list[i].relation_ids) > -1).length === 0) {
                that.menuCheckedKeys.push(list[i].id)
              }
              that.menuCheckedKeysCache.push(list[i].id)
            }
          }
        }
      })
    },
    onCheckMenu (checkedKeys, e) {
      // 当前资源
      const ids = that.menuCheckedKeysCache
      // 最新资源
      const halfIds = e.halfCheckedKeys
      let newIds = checkedKeys
      if (halfIds && halfIds.length > 0) {
        newIds = newIds.concat(halfIds)
      }
      let changeIds
      if (e.checked) {
        // 新增
        changeIds = newIds.filter(item => ids.filter(item2 => item2 === item).length === 0)
        that.menuCheckedKeys.push(...checkedKeys)
      } else {
        // 删除
        changeIds = ids.filter(item => newIds.filter(item2 => item2 === item).length === 0)
        // 排除半选状态
        changeIds = changeIds.filter(item => halfIds.filter(item2 => item2 === item).length === 0)
        that.menuCheckedKeys = that.menuCheckedKeys.filter(item => changeIds.indexOf(item) === -1)
      }
      that.menuCheckedKeysCache = newIds
      const param = {
        ids: JSON.stringify(changeIds),
        roleId: that.checkRoleId,
        add: e.checked
      }
      UpdateRoleMenuByIds(param).then((res) => {
        if (res.ret === 200) {
          this.$message.success('保存成功')
        }
      })
    },
    onSelectMenu (selectedKeys, info) {
      if (that.selectMenu === info.selectedNodes[0].key) {
        return
      }
      that.selectMenu = info.selectedNodes[0].key
      that.getBtnResourceList(that.selectMenu)
    },
    // 根据当前选中菜单资源获取按钮资源
    getBtnResourceList (menuId) {
      that.checkMenuId = menuId
      that.menuBtnOptions = []
      that.checkedMenuBtnOptions = []
      this.menuBtnLoading = true
      // 获取已拥有的菜单按钮资源
      GetBtnMenuList({
        menuId: menuId
      }).then((res) => {
        if (res.ret === 200) {
          const arr = []
          if (res.data && res.data.length > 0) {
            for (let i = 0; i < res.data.length; i++) {
              arr.push({ label: res.data[i].name, value: res.data[i].id })
            }
          }
          that.menuBtnOptions = arr
          GetCheckedMenuBtnLimit({ menuId: menuId,
            roleId: that.checkRoleId }).then((res) => {
            this.menuBtnLoading = false
            if (res.ret === 200) {
              const list = res.data
              if (list && list.length > 0) {
                for (let i = 0; i < list.length; i++) {
                  that.checkedMenuBtnOptions.push(list[i].id)
                }
                that.checkedMenuBtnValues = that.checkedMenuBtnOptions
              }
            }
          })
        }
      })
    }
  }
}
</script>

<style scoped>
  .ant-list-split .ant-list-item{    padding: 10px;}

  .ant-list-split .ant-list-item.hover{background: #e7e7e7}

</style>
