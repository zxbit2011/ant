<template>
  <a-modal
    title="角色编辑"
    :width="1000"
    :visible="visible"
    :keyboard="false"
    :maskClosable="false"
    :destroyOnClose="true"
    :confirmLoading="confirmLoading"
    @cancel="handleCancel"
  >
    <a-spin :spinning="confirmLoading">
      <a-form :form="form">
        <!-- step1 -->
        <div>
          <a-form-item
            label="角色名称"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-decorator="['name', {rules: [{required: true}]}]" placeholder="请输入角色名称" />
          </a-form-item>
          <a-form-item
            label="角色类型"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-radio-group
              @change="onChangeType"
              v-decorator="['role_type', {initialValue: 1, rules: [{required: true}]}]">
              <a-radio :value="1">管理员
                <a-tooltip>
                  <template slot="title">
                    <div><b>可管理子级组织的数据权限：</b></div>
                  </template>
                  <a-icon type="question-circle-o"/>
                </a-tooltip></a-radio>
              <a-radio :value="0">普通账户
                <a-tooltip>
                  <template slot="title">
                    <div>仅管理自己创建的项目数据</div>
                  </template>
                  <a-icon type="question-circle-o"/>
                </a-tooltip></a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item
            label="备注"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-textarea :rows="4" v-decorator="['remarks', {rules: [{required: true}]}]"></a-textarea>
          </a-form-item>
        </div>
      </a-form>
    </a-spin>
    <template slot="footer">
      <a-button key="cancel" @click="handleCancel">取消</a-button>
      <a-button key="forward" :loading="confirmLoading" type="primary" @click="handleNext()">完成</a-button>
    </template>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'
import Vue from 'vue'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import { AddRole, GetRoleInfo } from '@/api/sys'
let that, backFunc
export default {
  name: 'QuotaEdit',
  data () {
    return {
      id: null,
      years: [],
      projectList: [],
      projectListCache: [],
      projectTypeList: [],
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this),
      typeOptions: [
        { label: '全部', value: '1' },
        { label: '指定项目', value: '2' }
      ],
      typeCheckedValue: 0,
      fileList: [],
      headers: {
        'X-Requested-With': 'XMLHttpRequest',
        'Access-Token': Vue.ls.get(ACCESS_TOKEN)
      }
    }
  },
  methods: {
    edit (record, backFunc) {
      that = this
      this.backFunc = backFunc
      this.visible = true
      // 初始化
      this.typeCheckedValue = 0
      if (record) {
        that.id = record.id
        GetRoleInfo({ id: record.id })
          .then(res => {
            that.loading = false
            if (res.ret === 200) {
              record = res.data
              record.role_type = parseInt(record.role_type)
              that.typeCheckedValue = record.role_type
              that.Form = res.data
              const { form: { setFieldsValue } } = that
              that.$nextTick(() => {
                setFieldsValue(pick(record, 'name', 'role_type', 'remarks'))
              })
            }
          })
      } else {
        that.id = null
        that.record = null
        that.loading = false
      }
    },
    onChangeType (checkedValue) {
      console.log('checked = ', checkedValue)
      console.log('value = ', checkedValue.target.value)
      this.typeCheckedValue = checkedValue.target.value
    },
    filterOption (input, option) {
      return option.componentOptions.children[0].text.toUpperCase().indexOf(input.toUpperCase()) >= 0
    },
    filter (inputValue, path) {
      return (path.some(option => (option.label).toLowerCase().indexOf(inputValue.toLowerCase()) > -1))
    },
    handleNext () {
      const { form: { validateFields } } = this
      // last step
      this.confirmLoading = true
      validateFields((errors, values) => {
        console.log('errors:', errors, 'val:', values)
        if (!errors) {
          // 保存消息
          values.id = that.id
          AddRole(values)
            .then(res => {
              that.confirmLoading = false
              if (res.ret === 200) {
                that.$notification.success({
                  message: '系统提示',
                  description: res.msg
                })
                this.visible = false
                if (that.backFunc) {
                  that.backFunc()
                }
              }
            })
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    handleCancel () {
      // clear form
      this.visible = false
    }
  }
}
</script>
